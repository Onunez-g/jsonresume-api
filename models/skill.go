package models

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level,omitempty"`
	Keywords []string `json:"keywords"`
}

func (s *Skill) Patch(skill map[string]interface{}) {
	for k, v := range skill {
		switch k {
		case "name":
			s.Name = v.(string)
		case "level":
			s.Level = v.(string)
		case "keywords":
			s.Keywords = v.([]string)
		}
	}
}

func (s *Skill) IfNameExists(skills []Skill) bool {
	for _, v := range skills {
		if v.Name == s.Name {
			return true
		}
	}
	return false
}

func FindSkill(skills []Skill, name string) (*Skill, int) {
	for k, v := range skills {
		if v.Name == name {
			return &v, k
		}
	}
	return &Skill{}, -1
}
