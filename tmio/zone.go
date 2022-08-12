package tmio

type Zone struct {
	Name   string `json:"name"`
	Flag   string `json:"flag"`
	Parent *Zone  `json:"parent"`
}
