package config

var cfg *Config

type (
	Config struct {
		App         `yaml:"app"`
		HTTP        `yaml:"http"`
		Log         `yaml:"logger"`
		Integration `yaml:"integration"`
		MongoDb     `yaml:"mongodb"`
		Cache       `yaml:"cache"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Integration struct {
		SwapiApi `yaml:"swapi-api"`
	}

	SwapiApi struct {
		SwapiUrl            string `env-required:"true" yaml:"url"  env:"SWAPI_URL"`
		SwapiApiMaxConn     int    `yaml:"maxConn"`
		SwapiApiMaxRoutes   int    `yaml:"maxRoutes"`
		SwapiApiReadTimeout string `yaml:"ReadTimeout"`
		SwapiApiConnTimeout string `yaml:"connTimeout"`
	}

	MongoDb struct {
		MongoDbDabase string `env-required:"true" yaml:"database" env:"DATABASE_MONGODB"`
		MongoDbHost   string `env-required:"true" yaml:"host" env:"HOST_MONGODB"`
		MongoDbPort   string `env-required:"true" yaml:"port" env:"PORT_MONGODB"`
	}

	Cache struct {
		RedisUrl      string `yaml:"url" env:"REDIS_URL"`
		RedisPort     int    `env:"REDIS_PORT"`
		RedisDb       int    `env:"REDIS_DB"`
		RedisPassword string `yaml:"password" env:"REDIS_PASSWORD"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
