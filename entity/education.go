package entity

type Education struct {
	Institution string `json:"institution"`
	URL         string `json:"url"`
	Area        string `json:"area"`
	StudyType   string `json:"studyType"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Score       string `json:"score"`
	Courses     []any  `json:"courses"`
}
