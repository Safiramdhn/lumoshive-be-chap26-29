package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitViper initializes the Viper configuration
func InitViper() *viper.Viper {
	start := time.Now()

	// Initialize logger
	log := IntiLogger()

	// Create a new Viper instance
	configViper := viper.New()
	configViper.SetConfigFile(".env") // Specify the .env file
	configViper.AutomaticEnv()        // Automatically override values with environment variables

	// Read the configuration file
	err := configViper.ReadInConfig()
	if err != nil {
		log.Panic("Failed to read configuration file",
			zap.String("config_type", "viper"),
			zap.Error(err),
			zap.Duration("duration", time.Since(start)),
		)
	}

	// Log success
	log.Info("Configuration successfully loaded",
		zap.String("config_type", "viper"),
		zap.Duration("duration", time.Since(start)),
	)

	return configViper
}
