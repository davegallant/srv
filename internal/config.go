package internal

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Configuration stores the global config
type Configuration struct {
	Feeds          []string `yaml:"feeds"`
	ExternalViewer string   `yaml:"externalViewer,omitempty"`
}

// DefaultConfiguration can be used if a config is missing
var DefaultConfiguration = Configuration{
	Feeds: []string{
		"https://news.ycombinator.com/rss",
		"https://www.reddit.com/r/golang/.rss",
		"https://www.zdnet.com/topic/security/rss.xml",
	},
}

// Determines the default viewer
func DetermineExternalViewer() (string, error) {
	switch os := runtime.GOOS; os {
	case "linux":
		return "xdg-open", nil
	case "darwin":
		return "open", nil
	}

	return "", errors.New("Unable to determine a default external viewer")
}

// LoadConfiguration takes a filename (configuration) and loads it.
func LoadConfiguration(file string) Configuration {
	var config Configuration

	// If the configuration file does not exist,
	// write a default config
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		WriteConfig(DefaultConfiguration, file)
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if config.ExternalViewer == "" {
		config.ExternalViewer, err = DetermineExternalViewer()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicln(err)
	}
	return config
}

// WriteConfig writes a config to disk
func WriteConfig(config Configuration, file string) error {
	c, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Unable to marshal default config: %v", err)
	}

	err = ioutil.WriteFile(file, c, 0644)
	if err != nil {
		log.Fatalf("Unable to write default config: %v", err)
	}
	return nil
}
