package models

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	Website     string `json:"website"`
	Summary     string `json:"summary"`
}

func (p *Publication) Patch(award map[string]interface{}) {
	for k, v := range award {
		switch k {
		case "name":
			p.Name = v.(string)
		case "releaseDate":
			p.ReleaseDate = v.(string)
		case "publisher":
			p.Publisher = v.(string)
		case "website":
			p.Website = v.(string)
		case "summary":
			p.Summary = v.(string)
		}
	}
}

func (p *Publication) IfNameExists(publications []Publication) bool {
	for _, v := range publications {
		if v.Name == p.Name {
			return true
		}
	}
	return false
}

func FindPublication(publications []Publication, name string) (*Publication, int) {
	for k, v := range publications {
		if v.Name == name {
			return &v, k
		}
	}
	return &Publication{}, -1
}
