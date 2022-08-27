package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Climate string             `bson:"climate"`
	Terrain string             `bson:"terrain"`
}

type PlanetDTO struct {
	Name    string `json:"name" binding:"required"`
	Climate string `json:"climate" binding:"required"`
	Terrain string `json:"terrain" binding:"required"`
}

func NewPlanet(dto PlanetDTO) *Planet {
	return &Planet{
		ID:      primitive.NewObjectID(),
		Name:    dto.Name,
		Climate: dto.Climate,
		Terrain: dto.Terrain,
	}
}
