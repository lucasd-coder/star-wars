package repository

import (
	"context"

	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetRepository struct {
	Config     *config.Config
	Connection *mongo.Client
}

func NewPlanetRepository(cfg *config.Config, connection *mongo.Client) *PlanetRepository {
	return &PlanetRepository{
		Config:     cfg,
		Connection: connection,
	}
}

func (repo *PlanetRepository) Save(planet *models.Planet) error {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection("planets")
	_, err := collection.InsertOne(context.TODO(), planet)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (repo *PlanetRepository) FindByName(name string) (*models.Planet, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection("planets")
	filter := bson.M{"name": bson.M{"$eq": name}}

	var planet *models.Planet
	if err := collection.FindOne(context.TODO(), filter).Decode(&planet); err != nil {
		return &models.Planet{}, err
	}

	return planet, nil
}

func (repo *PlanetRepository) FindById(id string) (*models.Planet, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection("planets")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.Planet{}, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var planet *models.Planet
	if err := collection.FindOne(context.TODO(), filter).Decode(&planet); err != nil {
		return &models.Planet{}, err
	}
	return planet, nil
}

func (repo *PlanetRepository) FindAll() ([]*models.Planet, error) {
	return nil, nil
}

func (repo *PlanetRepository) Delete(id string) error {
	return nil
}
