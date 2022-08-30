package repository

import (
	"context"

	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "planets"
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
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)
	repo.createIndex("name", true)
	_, err := collection.InsertOne(context.TODO(), planet)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (repo *PlanetRepository) FindByName(name string) (*models.Planet, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)
	filter := bson.M{"name": bson.M{"$eq": name}}

	var planet *models.Planet
	if err := collection.FindOne(context.TODO(), filter).Decode(&planet); err != nil {
		return &models.Planet{}, err
	}

	return planet, nil
}

func (repo *PlanetRepository) FindById(id string) (*models.Planet, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)
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

func (repo *PlanetRepository) FindAll() ([]models.Planet, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)
	var planets []models.Planet
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return []models.Planet{}, err
	}

	if err := cursor.All(context.TODO(), &planets); err != nil {
		return []models.Planet{}, err
	}

	return planets, nil
}

func (repo *PlanetRepository) Delete(id string) (*mongo.DeleteResult, error) {
	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}

	return result, err
}

func (repo *PlanetRepository) createIndex(field string, unique bool) {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	collection := repo.Connection.Database(repo.Config.MongoDbDabase).Collection(collection)

	_, err := collection.Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		logger.Log.Error(err.Error())
	}
}
