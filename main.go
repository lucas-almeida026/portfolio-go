package main

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"text/template"

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
	e.Use(middleware.Recover())
	e.Static("/assets", "/assets")
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

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
			"Greeting":   textMap[lang]["greeting"],
			"Title":      textMap[lang]["title"],
			"Roles": []string{
				textMap[lang]["role0"],
				textMap[lang]["role1"],
				textMap[lang]["role2"],
				textMap[lang]["role3"],
			},
		}
		return c.Render(http.StatusOK, "index", data)
	})

	e.Logger.Fatal(e.Start(":4488"))
}
