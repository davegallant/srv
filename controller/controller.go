package controller

import (
	"time"

	config "github.com/davegallant/srv/config"
	feeds "github.com/davegallant/srv/feeds"
)

// Controller keeps everything together
type Controller struct {
	Config     config.Configuration
	lastUpdate time.Time
	Rss        *feeds.RSS
}

// Init initiates the controller with config
func (c *Controller) Init(conf string) {
	c.Config = config.LoadConfiguration(conf)
	c.Rss = &feeds.RSS{}
	c.Rss.Update(c.Config.Feeds)
}
