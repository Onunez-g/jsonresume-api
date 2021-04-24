package models

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

func (b *Basics) Patch(basics map[string]interface{}) {
	for k, v := range basics {
		switch k {
		case "name":
			b.Name = v.(string)
		case "label":
			b.Label = v.(string)
		case "picture":
			b.Picture = v.(string)
		case "email":
			b.Email = v.(string)
		case "phone":
			b.Phone = v.(string)
		case "website":
			b.Website = v.(string)
		case "summary":
			b.Summary = v.(string)
		case "location":
		case "profiles":
		}
	}
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

func (p *Profile) Patch(profile map[string]interface{}) {
	for k, v := range profile {
		switch k {
		case "network":
			p.Network = v.(string)
		case "username":
			p.Username = v.(string)
		case "url":
			p.URL = v.(string)
		}
	}
}

func (p *Profile) IfNetworkExists(profiles []Profile) bool {
	for _, v := range profiles {
		if v.Network == p.Network {
			return true
		}
	}
	return false
}

func FindProfile(profiles []Profile, network string) (*Profile, int) {
	for k, v := range profiles {
		if v.Network == network {
			return &v, k
		}
	}
	return &Profile{}, -1
}
