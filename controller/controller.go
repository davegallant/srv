package controller

import (
	config "github.com/davegallant/srv/config"
	feeds "github.com/davegallant/srv/feeds"
)

// Controller keeps everything together
type Controller struct {
	Config      config.Configuration
	Rss         *feeds.RSS
	CurrentFeed int
}

// Init initiates the controller with config
func (c *Controller) Init(conf string) {
	c.Config = config.LoadConfiguration(conf)
	c.Rss = &feeds.RSS{}
	c.Rss.Update(c.Config.Feeds)
	c.CurrentFeed = 0
}
