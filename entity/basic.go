package entity

type Basic struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    string    `json:"image"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	URL      string    `json:"url"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}
