package cui

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"path"

	"github.com/davegallant/srv/controller"
	"github.com/jroimartin/gocui"
)

// Controller can access internal state
var Controller *controller.Controller

var (
	viewArr = []string{"feeds", "Items"}
	active  = 0
)

// openFeed opens all items in the feed
func openFeed(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	feed := Controller.Rss.Feeds[oy]
	Controller.CurrentFeed = oy
	ov, _ := g.View("Items")

	ov.Clear()

	if err := v.SetOrigin(0, 0); err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Items {
		fmt.Fprintln(ov, "-", item.Title)
	}
	nextView(g, ov)

	return nil
}

// openItem opens the feed in an external browser
func openItem(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	item := Controller.Rss.Feeds[Controller.CurrentFeed].Items[oy]
	err := exec.Command(
		Controller.Config.ExternalViewer,
		append(Controller.Config.ExternalViewerArgs, item.Link)...).Start()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func refreshFeeds(g *gocui.Gui, v *gocui.View) error {
	Controller.Rss.Update(Controller.Config.Feeds)
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	_, err := g.View("Items")
	if err != nil {
		return err
	}

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("feeds", 0, 0, maxX-1, maxY/4-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = selectionBgColor
		v.SelFgColor = selectionFgColor
		v.Title = "Feeds"

		if _, err = setCurrentViewOnTop(g, "feeds"); err != nil {
			return err
		}
		for _, f := range Controller.Rss.Feeds {
			fmt.Fprintln(v, "-", f.Title)
		}
	}
	if v, err := g.SetView("Items", 0, maxY/4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = selectionBgColor
		v.SelFgColor = selectionFgColor
		v.Title = "Items"
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Start initializes the application
func Start() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configPath := path.Join(usr.HomeDir, ".config", "srv", "config.yaml")

	Controller = &controller.Controller{}
	Controller.Init(configPath)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
