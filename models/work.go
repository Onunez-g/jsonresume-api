package models

type Work struct {
	Organization string   `json:"organization,omitempty"`
	Position     string   `json:"position"`
	Website      string   `json:"website"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
	Company      string   `json:"company,omitempty"`
}

func (w *Work) IfCompanyExists(works []Work) bool {
	for _, v := range works {
		if v.Company == w.Company {
			return true
		}
	}
	return false
}

func FindWork(works []Work, company string) *Work {
	for _, v := range works {
		if v.Company == company {
			return &v
		}
	}
	return &Work{}
}
func FindVolunteer(works []Work, organization string) *Work {
	for _, v := range works {
		if v.Organization == organization {
			return &v
		}
	}
	return &Work{}
}
