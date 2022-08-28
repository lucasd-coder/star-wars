package external

type swapiPlanet struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

type ResultsSwaApi struct {
	Results []*swapiPlanet `json:"results"`
}
