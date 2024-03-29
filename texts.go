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
	textMap["en"]["role0"] = "Web Development"
	textMap["en"]["role1"] = "Performance optimization"
	textMap["en"]["role2"] = "Mentorship & Leadership"
	textMap["en"]["role3"] = "Creator of libraries, CLIs and parsers"
	// English END

	// Portuguese START
	// header
	textMap["pt"]["title"] = "Desenvolvedor de Software"
	// hero
	textMap["pt"]["greeting"] = "Olá, meu nome é"
	textMap["pt"]["role0"] = "Desenvolvimento Web"
	textMap["pt"]["role1"] = "Otimização de Performance"
	textMap["pt"]["role2"] = "Mentoria e liderança"
	textMap["pt"]["role3"] = "Criador de bibliotecas, CLIs e parsers"
	// Portuguese END
	return textMap
}
