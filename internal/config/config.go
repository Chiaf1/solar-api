package config

import (
	"fmt"
	"os"

	"github.com/chiaf1/solar-api/internal/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InfluxDB InfluxConfig `yaml:"InfluxDB"`
}

type InfluxConfig struct {
	Url    string `yaml:"url"`
	Token  string `yaml:"token"`
	Org    string `yaml:"org"`
	Bucket string `yaml:"bucket"`
}

// Load loads the values frrom the file "path" to the struct c, if the file is not present:
// the default values are loaded and the file is created.
func (c *Config) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			c.SetDefault()
			c.Save(path)
			return nil
		}
		return fmt.Errorf("Error while reading the config file: %w", err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("Error during parsing of YAML config file: %w", err)
	}
	return nil
}

// SetDefault sets the config default values
func (c *Config) SetDefault() {
	c.InfluxDB.Url = "http://localhost:8086"
	c.InfluxDB.Token = ""
	c.InfluxDB.Org = ""
	c.InfluxDB.Bucket = ""
}

// Save saves the configs to the "path" using the WriteFileAtomic function
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error while parsing to YAML: %w", err)
	}
	return utils.WriteFileAtomic(path, data, 0644)
}

// Data validation after loading
func (c *Config) Validate() error {
	if c.InfluxDB.Token == "" {
		return fmt.Errorf("[config.influx] no token defined")
	}
	if c.InfluxDB.Org == "" {
		return fmt.Errorf("[config.influx] no org defined")
	}
	if c.InfluxDB.Bucket == "" {
		return fmt.Errorf("[config.influx] no bucket defined")
	}
	return nil
}
