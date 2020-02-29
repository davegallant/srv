package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/mmcdole/gofeed"
)

var Configuration struct {
	Feeds []string
}

var filename = flag.String("config", "config.json", "Location of the config file.")

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Panicln(err)
	}
	if err != nil {
		log.Panicln(err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		log.Panicln(err)
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("feeds", maxX/2-26, maxY/5, maxX/2+8, maxY/2+6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fp := gofeed.NewParser()
		for _, f := range Configuration.Feeds {
			feed, err := fp.ParseURL(f)
			if err != nil {
				fmt.Println(err)
			}
			if feed != nil {
				fmt.Fprintln(v, feed.Title)
			}
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
