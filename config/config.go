package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config contains all the configuration for harpoonee
type Config struct {
	Port    int      `yaml:"port" json:"port"`
	TLS     TLS      `yaml:"tls" json:"tls"`
	Logging Logging  `yaml:"log" json:"log"`
	Plugins []Plugin `yaml:"plugins" json:"plugins"`
}

// Plugin contains the config for one plugin
type Plugin struct {
}

// Logging contains all logging related config
type Logging struct {
	Level    string `yaml:"level" json:"level"`
	Filepath string `yaml:"file" json:"file"`
}

// TLS contains all config for tls
type TLS struct {
	Key  string `yaml:"key" json:"key"`
	Cert string `yaml:"certificate" json:"certification"`
}

// Default contains the default configuration for harpoonee
// This is assumed, if no configuration was provided, or an error occured while loading the config
var defaultConfig Config = Config{
	Port: 6678,
	Logging: Logging{
		Level:    "info",
		Filepath: "harpooneer.log",
	},
}

// Default returns the default config
func Default() Config {
	return defaultConfig
}

// FromYAML loads the configuration from yaml data
func fromYAML(src io.Reader) (Config, error) {
	c := Config{}
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return defaultConfig, err
	}
	if err = yaml.Unmarshal(data, &c); err != nil {
		return defaultConfig, err
	}
	return c, nil
}

// FromJSON loads the configuration from json data
func fromJSON(src io.Reader) (Config, error) {
	c := Config{}
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return defaultConfig, err
	}
	if err = json.Unmarshal(data, &c); err != nil {
		return defaultConfig, err
	}
	return c, nil
}

// ToYAML converts the given configuration to a yaml string
func (c Config) ToYAML() (string, error) {
	out, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// ToJSON converts the given configuration to a json string
func (c Config) ToJSON() (string, error) {
	out, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Load attemptes to load the config from the given file
// In case of an error, the default config is returned
func Load(filename string) Config {
	filename, err := filepath.Abs(filename)
	if err != nil {
		log.Printf("Failed to load config from %s. Starting with default config. Error was %v", filename, err)
		return defaultConfig
	}
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Printf("Failed to load config from %s. Starting with default config. Error was %v", filename, err)
		return defaultConfig
	}
	var config Config
	switch filepath.Ext(file.Name()) {
	case ".yaml":
		config, err = fromYAML(file)
	case ".yml":
		config, err = fromYAML(file)
	case ".json":
		config, err = fromJSON(file)
	default:
		config = defaultConfig
	}
	if err != nil {
		log.Println("Unable to unmarshal input. Starting with default config. Error was ", err)
		return defaultConfig
	}
	file.Close()
	return config
}
