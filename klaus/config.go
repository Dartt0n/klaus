package klaus

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token  string  `yaml:"token"`
	Admins []int64 `yaml:"admins"`
}

func LoadConfig() (*Config, error) {
	c := Config{}

	f, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
