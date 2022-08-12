package tmio

type Player struct {
	Name string         `json:"name"`
	Tag  string         `json:"tag"`
	ID   string         `json:"id"`
	Zone *Zone          `json:"zone"`
	Meta map[string]any `json:"meta"`
}
