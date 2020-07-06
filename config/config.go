package config

import (
	"github.com/juju/errors"

	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/davegallant/srv/file"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// ConfigPath defines where the configuration is stored
const ConfigPath = ".config/srv/config.yml"

// Configuration stores the global config
type Configuration struct {
	Feeds              []string `yaml:"feeds"`
	ExternalViewer     string   `yaml:"externalViewer,omitempty"`
	ExternalViewerArgs []string `yaml:"externalViewerArgs,omitempty"`
	Path               string
}

// DefaultConfiguration can be used if a config is missing
var DefaultConfiguration = Configuration{
	Feeds: []string{
		"https://aws.amazon.com/blogs/security/feed/",
		"https://www.phoronix.com/rss.php",
		"https://www.zdnet.com/topic/security/rss.xml",
	},
	Path: ConfigPath,
}

// GetUGetUGetUserConfigPath returns the full configuration path for the current user
func GetUserConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	createConfigPath := []string{usr.HomeDir}
	createConfigPath = append(createConfigPath, strings.Split(ConfigPath, "/")...)
	return path.Join(createConfigPath...), nil
}

// DetermineExternalViewer checks the OS to decide the default viewer
func DetermineExternalViewer() (string, error) {
	switch os := runtime.GOOS; os {
	case "linux":
		return "xdg-open", nil
	case "darwin":
		return "open", nil
	}

	return "", errors.New("Unable to determine a default external viewer")
}

// EnsureConfigDirExists ensures directory exists with correct permissions
func EnsureConfigDirExists(d string) error {
	var AppFs = afero.NewOsFs()
	return AppFs.MkdirAll(d, 0700)
}

// LoadConfiguration loads a configuration from a file
func LoadConfiguration(f string) (Configuration, error) {
	var config Configuration

	if !file.Exists(f) {
		err := WriteConfig(DefaultConfiguration, f)
		return DefaultConfiguration, errors.Annotate(err, "failed to load configuration")
	}

	data, err := ioutil.ReadFile(f)

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
	return config, nil
}

// WriteConfig writes a config to disk
func WriteConfig(config Configuration, f string) error {

	d := filepath.Dir(f)
	err := EnsureConfigDirExists(d)
	if err != nil {
		return errors.Annotatef(err, "Unable to to create config directory '%s'", d)
	}

	c, err := yaml.Marshal(&config)
	if err != nil {
		return errors.Annotatef(err, "Unable to marshal config '%s'", f)
	}

	err = ioutil.WriteFile(f, c, 0600)
	if err != nil {
		return errors.Annotatef(err, "Unable to write default config: '%s'", f)
	}
	return nil
}
