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
	textMap["en"]["role0"] = "Software Developer"
	textMap["en"]["role1"] = "Software Engineer"
	textMap["en"]["role2"] = "Software Architect"
	textMap["en"]["role3"] = "Full Stack Developer"
	// English END

	// Portuguese START
	// header
	textMap["pt"]["title"] = "Desenvolvedor de Software"
	// hero
	textMap["pt"]["greeting"] = "Olá, meu nome é"
	textMap["pt"]["role0"] = "Desenvolvedor de Software"
	textMap["pt"]["role1"] = "Engenheiro de Software"
	textMap["pt"]["role2"] = "Arquiteto de Software"
	textMap["pt"]["role3"] = "Desenvolvedor Full Stack"
	// Portuguese END
	return textMap
}
