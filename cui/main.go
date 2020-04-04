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

// Controller can access Feeds and Config
var Controller *controller.Controller

var (
	viewArr     = []string{"feeds", "Items"}
	active      = 0
	currentFeed = 0 // TODO: move to Controller
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			v.SetCursor(cx, cy-1)
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// openFeed opens all items in the feed
func openFeed(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	currentFeed = cy
	feed := Controller.Rss.Feeds[currentFeed]
	ov, _ := g.View("Items")

	ov.Clear()
	for _, item := range feed.Items {
		fmt.Fprintln(ov, "-", item.Title)
	}
	nextView(g, ov)

	return nil
}

// openItem opens the feed in an external browser
func openItem(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	item := Controller.Rss.Feeds[currentFeed].Items[cy]
	viewer := Controller.Config.ExternalViewer
	err := exec.Command(viewer, item.Link).Start()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func showLoading(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("loading", maxX/2-4, maxY/2-1, maxX/2+4, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Loading")
	}
	return nil
}

func hideLoading(g *gocui.Gui) error {
	if err := g.DeleteView("loading"); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func refreshFeeds(g *gocui.Gui, v *gocui.View) error {
	showLoading(g)
	Controller.Rss.Update(Controller.Config.Feeds)
	//hideLoading(g)
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
		v.Title = "Feeds"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

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
		v.Title = "Items"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
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
