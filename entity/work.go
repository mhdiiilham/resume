package entity

type Work struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Position    string `json:"position"`
	URL         string `json:"url"`
	StartDate   string `json:"startDate"`
	Summary     string `json:"summary"`
	Highlights  []any  `json:"highlights"`
	EndDate     string `json:"endDate,omitempty"`
}
