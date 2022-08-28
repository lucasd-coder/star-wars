package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Climate string             `json:"climate" bson:"climate"`
	Terrain string             `json:"terrain" bson:"terrain"`
}

type PlanetDTO struct {
	Name    string `binding:"required"`
	Climate string `binding:"required"`
	Terrain string `binding:"required"`
}

type PlanetResponse struct {
	Planet
	MovieAppearances int `json:"movie_appearances"`
}

func NewPlanet(dto PlanetDTO) *Planet {
	return &Planet{
		ID:      primitive.NewObjectID(),
		Name:    dto.Name,
		Climate: dto.Climate,
		Terrain: dto.Terrain,
	}
}

func NewPlanetResponse(planet Planet, movieAppearances int) *PlanetResponse {
	return &PlanetResponse{
		Planet:           planet,
		MovieAppearances: movieAppearances,
	}
}
