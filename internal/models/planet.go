package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Climate string             `bson:"climate"`
	Terrain string             `bson:"terrain"`
}

type PlanetDTO struct {
	Name    string `binding:"required"`
	Climate string `binding:"required"`
	Terrain string `binding:"required"`
}

type PlanetResponse struct {
	Planet
	MovieAppearances int
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
