package models

type Work struct {
	Organization string   `json:"organization,omitempty"`
	Company      string   `json:"company,omitempty"`
	Position     string   `json:"position"`
	Website      string   `json:"website"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

func (w *Work) Patch(work map[string]interface{}) {
	for k, v := range work {
		switch k {
		case "company":
			w.Company = v.(string)
		case "organization":
			w.Organization = v.(string)
		case "position":
			w.Position = v.(string)
		case "website":
			w.Website = v.(string)
		case "startDate":
			w.StartDate = v.(string)
		case "endDate":
			w.EndDate = v.(string)
		case "summary":
			w.Summary = v.(string)
		case "highlights":
			w.Highlights = v.([]string)
		}
	}
}

func (w *Work) IfCompanyExists(works []Work) bool {
	for _, v := range works {
		if v.Company == w.Company {
			return true
		}
	}
	return false
}

func (w *Work) IfOrganizationExists(volunteers []Work) bool {
	for _, v := range volunteers {
		if v.Organization == w.Organization {
			return true
		}
	}
	return false
}

func FindWork(works []Work, company string) (*Work, int) {
	for k, v := range works {
		if v.Company == company {
			return &v, k
		}
	}
	return &Work{}, -1
}
func FindVolunteer(works []Work, organization string) (*Work, int) {
	for k, v := range works {
		if v.Organization == organization {
			return &v, k
		}
	}
	return &Work{}, -1
}
