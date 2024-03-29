package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, paths ...string) {
	tmpl := &template.Template{}
	for i := range paths {
		template.Must(tmpl.ParseGlob(paths[i]))
	}
	t := newTemplate(tmpl)
	e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template{
		Templates: templates,
	}
}

func main() {

	textMap := getTextMap()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variable
	ENV := os.Getenv("ENV")
	PORT := os.Getenv("PORT")
	PARSE_IMG := os.Getenv("PARSE_IMG")

	if PARSE_IMG == "1" {
		bytes, err := os.ReadFile("assets/favicon.ico")
		if err != nil {
			fmt.Println(err)
		}
		os.WriteFile("_imgs/_img_favicon-ico.txt", []byte(base64.StdEncoding.EncodeToString(bytes)), 0644)
	} else {
		fmt.Println("skipped parse img")
	}

	if PORT == "" {
		PORT = "4488"
	}
	if ENV == "" {
		ENV = "prod"
	}
	fmt.Println("ENV =", ENV)

	htmxFile, err := fs.ReadFile(os.DirFS("local"), "htmx.js")
	if err != nil {
		fmt.Println(err)
	}
	stylesFile, err := fs.ReadFile(os.DirFS("local"), "styles.css")
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	if ENV == "prod" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://almeidadeveloper.com", "http://almeidadeveloper.com"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
		e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
			rate.Limit(20),
		)))
	} else {
		e.Use(middleware.CORS())
	}
	e.Use(middleware.Recover())

	NewTemplateRenderer(e, "templates/*.html")
	e.GET("/", func(c echo.Context) error {
		mode := c.QueryParam("visualmode")
		lang := c.QueryParam("lang")

		if mode != "dark" && mode != "light" {
			return c.Redirect(http.StatusFound, "/?visualmode=dark&lang="+lang)
		}

		if lang != "en" && lang != "pt" {
			return c.Redirect(http.StatusFound, "/?lang=en&visualmode="+mode)
		}

		data := map[string]interface{}{
			"HTMX":       string(htmxFile),
			"Styles":     string(stylesFile),
			"IsDarkmode": mode == "dark",
			"Lang":       lang,
			"Greeting":   textMap[lang]["greeting"],
			"Title":      textMap[lang]["title"],
			"Roles": []string{
				textMap[lang]["role0"],
				textMap[lang]["role1"],
				textMap[lang]["role2"],
				textMap[lang]["role3"],
			},
			"IMG_en":   getENIcon(),
			"IMG_pt":   getPTIcon(),
			"IMG_logo": getOrangeLogo(),
			"IMG_sun":  getSunIcon(),
			"IMG_moon": getMoonIcon(),
		}
		return c.Render(http.StatusOK, "index", data)
	})

	e.GET("/favicon.ico", func(c echo.Context) error {
		decoded, error := base64.StdEncoding.DecodeString(getFavicon())
		if error != nil {
			return c.String(http.StatusInternalServerError, "Error decoding base64")
		}
		return c.Blob(http.StatusOK, "image/x-icon", decoded)
	})

	e.Logger.Fatal(e.Start("0.0.0.0:" + PORT))
}
