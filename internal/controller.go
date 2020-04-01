package internal

import "time"

// Controller keeps everything together
type Controller struct {
	Config     Configuration
	lastUpdate time.Time
	Rss        *RSS
}

// Init initiates the controller with config
func (c *Controller) Init(config string) {
	c.Config = LoadConfiguration(config)
	c.Rss = &RSS{}
	c.Rss.New(c)
	c.Rss.Update()
}
