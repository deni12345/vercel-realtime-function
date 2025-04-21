package configs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	Local      = "local"
	Production = "production"
)

type config struct {
	Credentials             []byte `envconfig:"GOOGLE_CREDENTIALS"`
	ProjectID               string `json:"project_id"`
	Type                    string `json:"type"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

type ConfigLoader struct {
	config  *config
	loadFns []func(*config) error
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{
		loadFns: []func(*config) error{
			loadEnvFile,
			loadCredentials,
			decodeCredentials,
		},
	}
}

func (l *ConfigLoader) load() (*config, error) {
	if l.config != nil {
		return l.config, nil
	}

	l.config = &config{}
	for _, fn := range l.loadFns {
		if err := fn(l.config); err != nil {
			return nil, err
		}
	}
	return l.config, nil
}

func loadEnvFile(value *config) error {
	if os.Getenv("ENVIRONMENT") == "" {
		return godotenv.Load()
	}
	return nil
}

func loadCredentials(value *config) error {
	return envconfig.Process("", value)
}

func decodeCredentials(value *config) error {
	if len(value.Credentials) == 0 {
		return fmt.Errorf("empty credentials")
	}

	decodedValue := make([]byte, base64.RawStdEncoding.DecodedLen(len(value.Credentials)))
	n, err := base64.RawStdEncoding.Decode(decodedValue, value.Credentials)
	if err != nil {
		return err
	}
	return json.Unmarshal(decodedValue[:n], value)
}

var value *config

func GetConfig() (*config, error) {
	if value != nil {
		return value, nil
	}

	loader := NewConfigLoader()
	config, err := loader.load()
	if err != nil {
		return nil, err
	}

	value = config
	return value, nil
}
