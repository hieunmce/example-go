package config

import (
	"os"

	"github.com/spf13/viper"
)

// FileReader read config from file
type FileReader struct {
	filename string
	dirname  string
}

// NewFileReader create new file reader with filename and dirname
func NewFileReader(filename, dirname string) *FileReader {
	return &FileReader{filename, dirname}
}

func (r *FileReader) Read() (*Config, error) {
	v := viper.New()
	if os.Getenv("PRODUCTION") != "true" {
		v.SetConfigName(r.filename)
		v.AddConfigPath(r.dirname)
	}

	// priority for env varirable
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBType:          v.GetString("DB_TYPE"),
		DBUserName:      v.GetString("DB_USERNAME"),
		DBPassword:      v.GetString("DB_PASSWORD"),
		DBName:          v.GetString("DB_NAME"),
		DBSSLModeOption: v.GetString("DB_SSLMODE_OPTION"),
		DBHostname:      v.GetString("DB_HOSTNAME"),
		DBPort:          v.GetString("DB_PORT"),
		DBEnvironment:   v.GetString("DB_ENVIRONMENT"),
	}, nil
}
