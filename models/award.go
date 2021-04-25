package models

type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}

func (a *Award) Patch(award map[string]interface{}) {
	for k, v := range award {
		switch k {
		case "title":
			a.Title = v.(string)
		case "date":
			a.Date = v.(string)
		case "awarder":
			a.Awarder = v.(string)
		case "summary":
			a.Summary = v.(string)
		}
	}
}

func (a *Award) IfTitleExists(awards []Award) bool {
	for _, v := range awards {
		if v.Title == a.Title {
			return true
		}
	}
	return false
}

func FindAward(awards []Award, title string) (*Award, int) {
	for k, v := range awards {
		if v.Title == title {
			return &v, k
		}
	}
	return &Award{}, -1
}
