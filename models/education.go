package models

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

func (e *Education) Patch(education map[string]interface{}) {
	for k, v := range education {
		switch k {
		case "institution":
			e.Institution = v.(string)
		case "area":
			e.Area = v.(string)
		case "studyType":
			e.StudyType = v.(string)
		case "startDate":
			e.StartDate = v.(string)
		case "endDate":
			e.EndDate = v.(string)
		case "gpa":
			e.Gpa = v.(string)
		case "courses":
			e.Courses = v.([]string)
		}
	}
}

func (e *Education) IfInstitutionExists(educations []Education) bool {
	for _, v := range educations {
		if v.Institution == e.Institution {
			return true
		}
	}
	return false
}

func FindEducation(educations []Education, institution string) (*Education, int) {
	for k, v := range educations {
		if v.Institution == institution {
			return &v, k
		}
	}
	return &Education{}, -1
}
