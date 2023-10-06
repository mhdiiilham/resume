package entity

type Project struct {
	Name        string   `json:"name"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	URL         string   `json:"url"`
}
