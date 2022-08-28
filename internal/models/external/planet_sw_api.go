package external

type SwapiPlanet struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

type ResultsSwaApi struct {
	Results []*SwapiPlanet `json:"results"`
}
