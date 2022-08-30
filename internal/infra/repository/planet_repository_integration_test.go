package repository_test

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/infra/repository"
	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDocumentation = "documentation_examples"
	testCollection    = "planets"
	testId            = "630b7bcb419f837457644cbc"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Print("skipping mtest integration test in short mode")
		return
	}

	if err := mtest.Setup(); err != nil {
		log.Fatal(err)
	}
	defer os.Exit(m.Run())
	if err := mtest.Teardown(); err != nil {
		log.Fatal(err)
	}
}

func SetUpConfig() *config.Config {
	err := setEnvValues()
	if err != nil {
		panic(err)
	}
	var cfg config.Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func insertData(collation *mongo.Collection) {
	plant := mock.FactoryPlanet()
	_, err := collation.InsertOne(context.TODO(), plant)
	if err != nil {
		panic(err)
	}
}

func cleanData(collation *mongo.Collection) {
	_, err := collation.DeleteMany(context.TODO(), options.Delete())
	if err != nil {
		panic(err)
	}
}

func TestSaveData(t *testing.T) {
	cfg := SetUpConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))

	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)

	cleanData(collection)

	require.NoError(t, err)

	planet := mock.FactoryPlanet()

	testRepository := repository.PlanetRepository{cfg, client}

	errRepo := testRepository.Save(planet)

	assert.Nil(t, errRepo)
	cleanData(collection)
}

func TestFindByName(t *testing.T) {
	cfg := SetUpConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)

	insertData(collection)

	planet := mock.FactoryPlanet()

	testRepository := repository.PlanetRepository{cfg, client}

	result, errRepo := testRepository.FindByName(planet.Name)

	assert.Nil(t, errRepo)
	assert.Equal(t, planet, result)
	cleanData(collection)
}

func TestFindById(t *testing.T) {
	cfg := SetUpConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)
	cleanData(collection)
	insertData(collection)

	planet := mock.FactoryPlanet()

	testRepository := repository.PlanetRepository{cfg, client}

	result, errRepo := testRepository.FindById(testId)

	assert.Nil(t, errRepo)
	assert.Equal(t, planet, result)

	cleanData(collection)
}

func TestFindAll(t *testing.T) {
	cfg := SetUpConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)
	cleanData(collection)
	insertData(collection)

	testRepository := repository.PlanetRepository{cfg, client}

	result, errRepo := testRepository.FindAll()

	assert.Nil(t, errRepo)
	assert.NotEmpty(t, result)

	cleanData(collection)
}

func TestDeleteData(t *testing.T) {
	cfg := SetUpConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mtest.ClusterURI()))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	collection := client.Database(testDocumentation).Collection(testCollection)
	cleanData(collection)
	insertData(collection)

	testRepository := repository.PlanetRepository{cfg, client}

	result, errRepo := testRepository.Delete(testId)

	assert.Nil(t, errRepo)
	assert.NotEmpty(t, result)

	cleanData(collection)
}

func setEnvValues() error {
	err := os.Setenv("DATABASE_MONGODB", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting DATABASE_MONGODB, err = %v", err)
	}

	err = os.Setenv("APP_NAME", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting APP_NAME, err = %v", err)
	}

	err = os.Setenv("APP_VERSION", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting APP_VERSION, err = %v", err)
	}

	err = os.Setenv("HTTP_PORT", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting HTTP_PORT, err = %v", err)
	}

	err = os.Setenv("LOG_LEVEL", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting LOG_LEVEL, err = %v", err)
	}
	err = os.Setenv("SWAPI_URL", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting SWAPI_URL, err = %v", err)
	}
	err = os.Setenv("HOST_MONGODB", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting HOST_MONGODB, err = %v", err)
	}

	err = os.Setenv("PORT_MONGODB", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting PORT_MONGODB, err = %v", err)
	}
	err = os.Setenv("REDIS_URL", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting REDIS_URL, err = %v", err)
	}

	err = os.Setenv("REDIS_PASSWORD", testDocumentation)
	if err != nil {
		return fmt.Errorf("Error setting REDIS_PASSWORD, err = %v", err)
	}

	return nil
}
