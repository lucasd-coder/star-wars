package config

var cfg *Config

type (
	Config struct {
		App         `yaml:"app"`
		HTTP        `yaml:"http"`
		Log         `yaml:"logger"`
		Integration `yaml:"integration"`
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
		SwapiUrl            string `yaml:"url"`
		SwapiApiMaxConn     int    `yaml:"maxConn"`
		SwapiApiMaxRoutes   int    `yaml:"maxRoutes"`
		SwapiApiReadTimeout string `yaml:"ReadTimeout"`
		SwapiApiConnTimeout string `yaml:"connTimeout"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
