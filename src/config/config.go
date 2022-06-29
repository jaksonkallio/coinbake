package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database    Database   `yaml:"database"`
	MarketData  MarketData `yaml:"market_data"`
	Oauth       Oauth      `yaml:"oauth"`
	Environment string     `yaml:"environment"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type MarketData struct {
	ApiKey  string `yaml:"api_key"`
	BaseUrl string `yaml:"base_url"`
}

type Oauth struct {
	Google OauthProvider `yaml:"google"`
}

type OauthProvider struct {
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

var CurrentConfig Config

func LoadConfig(path string) error {
	// Reset/blank the current config
	CurrentConfig = Config{}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var config Config

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return err
	}

	// Update the current config with the configuration we loaded.
	CurrentConfig = config

	InitOauthProviderConfigs()

	return nil
}

func IsDev() bool {
	return CurrentConfig.Environment == "dev"
}
