package models

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

func (a *Language) Patch(language map[string]interface{}) {
	for k, v := range language {
		switch k {
		case "language":
			a.Language = v.(string)
		case "fluency":
			a.Fluency = v.(string)
		}
	}
}

func (a *Language) IfLanguageExists(languages []Language) bool {
	for _, v := range languages {
		if v.Language == a.Language {
			return true
		}
	}
	return false
}

func FindLanguage(languages []Language, language string) (*Language, int) {
	for k, v := range languages {
		if v.Language == language {
			return &v, k
		}
	}
	return &Language{}, -1
}
