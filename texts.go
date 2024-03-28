package main

func getTextMap() map[string]map[string]string {
	textMap := make(map[string]map[string]string)
	textMap["en"] = make(map[string]string)
	textMap["pt"] = make(map[string]string)

	// English START
	// header
	textMap["en"]["title"] = "Software Developer"
	// hero
	textMap["en"]["greeting"] = "Hello, I'm"
	textMap["en"]["aboutme"] = "!About Me!"
	// English END

	// Portuguese START
	// header
	textMap["pt"]["title"] = "Desenvolvedor de Software"
	// hero
	textMap["pt"]["greeting"] = "Olá, meu nome é"
	// Portuguese END
	return textMap
}
