package internal

import (
	"github.com/floriansw/go-hll-rcon/rconv2/api"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"slices"
)

type Fence struct {
	X       *string `yaml:"X,omitempty"`
	Y       *int    `yaml:"Y,omitempty"`
	Numpads []int   `yaml:"Numpad,omitempty"`
}

func (f Fence) Includes(w api.Grid) bool {
	if f.X != nil && w.X != *f.X {
		return false
	}
	if f.Y != nil && w.Y != *f.Y {
		return false
	}
	if len(f.Numpads) == 0 {
		return true
	}
	return slices.Contains(f.Numpads, w.Numpad)
}

type Server struct {
	Host               string  `yaml:"Host"`
	Port               int     `yaml:"Port"`
	Password           string  `yaml:"Password"`
	PunishAfterSeconds *int    `yaml:"PunishAfterSeconds,omitempty"`
	AxisFence          []Fence `yaml:"AxisFence"`
	AlliesFence        []Fence `yaml:"AlliesFence"`
}

type Config struct {
	Servers []Server `yaml:"Servers"`
	path    string
}

func (c *Config) Save() error {
	config, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, config, 0655)
}

func NewConfig(path string, logger *slog.Logger) (*Config, error) {
	config, err := readConfig(path, logger)
	if err != nil {
		return config, err
	}

	return config, config.Save()
}

func readConfig(path string, logger *slog.Logger) (*Config, error) {
	var config Config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Info("create-config")
		config = Config{}
	} else {
		logger.Info("read-existing-config")
		c, err := os.ReadFile(path)
		if err != nil {
			return &Config{}, err
		}
		err = yaml.Unmarshal(c, &config)
		if err != nil {
			return &Config{}, err
		}
	}
	config.path = path
	return &config, nil
}
