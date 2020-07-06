package controller

import (
	"log"

	config "github.com/davegallant/srv/config"
	feeds "github.com/davegallant/srv/feeds"
)

// Controller keeps everything together
type Controller struct {
	Config      config.Configuration
	Rss         *feeds.RSS
	CurrentFeed int
}

// Init initiates the controller
func (c *Controller) Init() {

	configPath, err := config.GetUserConfigPath()

	if err != nil {
		log.Fatal("Unable to locate user's config path")
	}

	c.Config, err = config.LoadConfiguration(configPath)
	if err != nil {
		log.Fatalf("Unable to load configuration: %s", err)
	}
	c.Rss = &feeds.RSS{}
	c.Rss.Update(c.Config.Feeds)
	c.CurrentFeed = 0
}
