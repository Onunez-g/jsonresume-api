package models

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
