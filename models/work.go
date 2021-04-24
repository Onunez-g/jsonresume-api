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
