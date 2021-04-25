package models

import "github.com/mitchellh/mapstructure"

type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work"`
	Volunteer    []Work        `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []Award       `json:"awards"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Skill       `json:"interests"`
	References   []Reference   `json:"references"`
}

var MyResume Resume = Resume{}

func (r *Resume) Patch(resume map[string]interface{}) {
	for k, v := range resume {
		switch k {
		case "basics":
			r.Basics = v.(Basics)
		case "work":
			work := v.([]interface{})
			list := make([]Work, 0, len(work))
			for _, v := range work {
				var result Work
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Work = list
		case "volunteer":
			volunteer := v.([]interface{})
			list := make([]Work, 0, len(volunteer))
			for _, v := range volunteer {
				var result Work
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Volunteer = list
		case "education":
			education := v.([]interface{})
			list := make([]Education, 0, len(education))
			for _, v := range education {
				var result Education
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Education = list
		case "awards":
			award := v.([]interface{})
			list := make([]Award, 0, len(award))
			for _, v := range award {
				var result Award
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Awards = list
		case "publications":
			publication := v.([]interface{})
			list := make([]Publication, 0, len(publication))
			for _, v := range publication {
				var result Publication
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Publications = list
		case "skills":
			skill := v.([]interface{})
			list := make([]Skill, 0, len(skill))
			for _, v := range skill {
				var result Skill
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Skills = list
		case "languages":
			lang := v.([]interface{})
			list := make([]Language, 0, len(lang))
			for _, v := range lang {
				var result Language
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Languages = list
		case "interests":
			interest := v.([]interface{})
			list := make([]Skill, 0, len(interest))
			for _, v := range interest {
				var result Skill
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.Interests = list
		case "references":
			reference := v.([]interface{})
			list := make([]Reference, 0, len(reference))
			for _, v := range reference {
				var result Reference
				mapstructure.Decode(v, &result)
				list = append(list, result)
			}
			r.References = list
		}
	}
}
