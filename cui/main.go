package cui

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/davegallant/srv/controller"
	"github.com/davegallant/srv/utils"
	"github.com/jroimartin/gocui"
	"github.com/mmcdole/gofeed"
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

	if err := ov.SetOrigin(0, 0); err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Items {
		fmt.Fprintln(ov, "-", item.Title)
	}
	err := nextView(g, ov)
	if err != nil {
		log.Printf("Unable to get next view: %s", err)
	}
	displayDescription(g, ov)

	return nil
}

func getCurrentFeedItem(v *gocui.View) *gofeed.Item {
	_, oy := v.Origin()
	return Controller.Rss.Feeds[Controller.CurrentFeed].Items[oy]
}

// displayDescription displays feed description if it exists
func displayDescription(g *gocui.Gui, v *gocui.View) {

	ov, _ := g.View("Description")
	ov.Clear()

	item := getCurrentFeedItem(v)
	description := utils.StripHTMLTags(item.Description)
	fmt.Fprintln(ov, description)
}

// openItem opens the feed in an external browser
func openItem(g *gocui.Gui, v *gocui.View) error {

	item := getCurrentFeedItem(v)
	err := exec.Command(
		Controller.Config.ExternalViewer,
		append(Controller.Config.ExternalViewerArgs, item.Link)...).Start()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func refreshFeeds(g *gocui.Gui, v *gocui.View) error {
	g.Close()
	Start()
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
	if v, err := g.SetView("feeds", 0, 0, maxX-1, maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = feedNameSelectionBgColor
		v.SelFgColor = feedNameSelectionFgColor
		v.Title = "Feeds"

		if _, err = setCurrentViewOnTop(g, "feeds"); err != nil {
			return err
		}
		for _, f := range Controller.Rss.Feeds {
			fmt.Fprintln(v, "-", f.Title)
		}
	}
	if v, err := g.SetView("Items", 0, maxY/3, maxX-1, maxY-(maxY/5)-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = feedItemSelectionBgColor
		v.SelFgColor = feedItemSelectionFgColor
		v.Title = "Items"
	}
	if v, err := g.SetView("Description", 0, maxY-maxY/5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.SelBgColor = feedItemSelectionBgColor
		v.SelFgColor = feedItemSelectionFgColor
		v.Title = "Description"
		v.FgColor = descriptionFgColor
		v.Wrap = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Start initializes the application
func Start() {
	Controller = &controller.Controller{}
	Controller.Init()

	g, err := gocui.NewGui(gocui.Output256)
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
