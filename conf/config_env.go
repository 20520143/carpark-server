package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// AppConfig presents app conf
type AppConfig struct {
	AppEnv           string `envconfig:"APP_ENV" envDefault:"dev"`
	Port             string `envconfig:"PORT" envDefault:"8088"`
	LogFormat        string `envconfig:"LOG_FORMAT" envDefault:"text"`
	DBHost           string `envconfig:"DB_HOST" envDefault:"localhost"`
	DBPort           string `envconfig:"DB_PORT" envDefault:"5432"`
	DBUser           string `envconfig:"DB_USER" envDefault:"postgres"`
	DBPass           string `envconfig:"DB_PASS" envDefault:"1"`
	DBName           string `envconfig:"DB_NAME" envDefault:"postgres"`
	EnableDB         string `envconfig:"ENABLE_DB" envDefault:"true"`
	TwilioAccountSID string `envconfig:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken  string `envconfig:"TWILIO_AUTH_TOKEN"`
	TwilioServiceSID string `envconfig:"VERIFY_SERVICE_SID"`
}

var config *AppConfig

func init() {
	config = &AppConfig{}

	_ = godotenv.Load()

	err := envconfig.Process("", config)
	if err != nil {
		err = errors.Wrap(err, "Failed to decode config env")
		logrus.Fatal(err)
	}
}

func GetConfig() *AppConfig {
	return config
}
