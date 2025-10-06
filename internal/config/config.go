// sentiric-vertical-legal-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Hukuk servisi bağımlılıkları (Placeholder)
	CaseManagementSystem string // CMS API
	CMSAPIKey            string
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Dikey Servisler, 204XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("VERTICAL_LEGAL_SERVICE_GRPC_PORT", "20411"),
		HttpPort: GetEnv("VERTICAL_LEGAL_SERVICE_HTTP_PORT", "20410"),

		CertPath: GetEnvOrFail("VERTICAL_LEGAL_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("VERTICAL_LEGAL_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		CaseManagementSystem: GetEnv("CMS_SYSTEM", "mock_cms"),
		CMSAPIKey:            GetEnv("CMS_API_KEY", ""),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
