package entity

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	URL         string `json:"url"`
	Summary     string `json:"summary"`
}
