package entity

type Resume struct {
	Basics       Basic         `json:"basics"`
	Work         []Work        `json:"work"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []Award       `json:"awards"`
	Certificates []Certificate `json:"certificates"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
	Projects     []Project     `json:"projects"`
}
