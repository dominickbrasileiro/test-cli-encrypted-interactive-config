package config

type Config struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	URL      string `yaml:"url"`
}
