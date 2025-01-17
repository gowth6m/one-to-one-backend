package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host        string `envconfig:"APP_HOST" default:"localhost"`
		Port        string `envconfig:"APP_PORT" default:"8080"`
		ApiVersion  string `envconfig:"API_VERSION" default:"v0"`
		AppVersion  string `envconfig:"APP_VERSION" default:"v0.0.1"`
		Environment string `envconfig:"ENVIRONMENT" default:"local"`
	}
	Database struct {
		MongoURI      string `envconfig:"DATABASE_URL" default:"mongodb://localhost:27017"`
		AccessTimeout int    `envconfig:"MONGODB_ACCESS_TIMEOUT" default:"5"`
		MongoDBName   string `envconfig:"MONGODB_DB_NAME" default:"one-to-one"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		JWTExpireInHours int    `envconfig:"JWT_EXPIRE" default:"24"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
		JWTIssuer        string `envconfig:"JWT_ISSUER" default:"one-to-one.vercel.app"`
	}
	Pusher struct {
		AppID   string `envconfig:"PUSHER_APP_ID"`
		Key     string `envconfig:"PUSHER_KEY"`
		Secret  string `envconfig:"PUSHER_SECRET"`
		Cluster string `envconfig:"PUSHER_CLUSTER"`
	}
	Cors struct {
		AllowOrigins     []string `envconfig:"CORS_ALLOW_ORIGINS" default:"*"`
		AllowMethods     []string `envconfig:"CORS_ALLOW_METHODS" default:"GET, POST, PUT, DELETE, OPTIONS"`
		AllowHeaders     []string `envconfig:"CORS_ALLOW_HEADERS" default:"Origin, Content-Length, Content-Type, Authorization, Tenant"`
		AllowCredentials bool     `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	}
	Vercel struct {
		CronSecret  string `envconfig:"CRON_SECRET" default:""`
		DeployedURL string `envconfig:"DEPLOYED_URL" default:"https://one-to-one.vercel.app"`
	}
}

// It is initialized once when the application starts.
var appConfig = &Config{}

// This function provides access to the appConfig variable.
func AppConfig() *Config {
	return appConfig
}

// LoadConfig loads environment variables and populates appConfig.
// It first attempts to load variables from a .env file and then
// processes environment variables according to the Config struct tags.
func LoadConfig() error {
	godotenv.Load()
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
