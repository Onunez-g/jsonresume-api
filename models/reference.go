package models

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

func (r *Reference) Patch(reference map[string]interface{}) {
	for k, v := range reference {
		switch k {
		case "name":
			r.Name = v.(string)
		case "reference":
			r.Reference = v.(string)
		}
	}
}

func (r *Reference) IfNameExists(references []Reference) bool {
	for _, v := range references {
		if v.Name == r.Name {
			return true
		}
	}
	return false
}

func FindReference(references []Reference, name string) (*Reference, int) {
	for k, v := range references {
		if v.Name == name {
			return &v, k
		}
	}
	return &Reference{}, -1
}
