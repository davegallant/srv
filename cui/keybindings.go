package cui

import "github.com/jroimartin/gocui"

// Map keys to actions
func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("feeds", gocui.KeyEnter, gocui.ModNone, openFeed); err != nil {
		return err
	}
	if err := g.SetKeybinding("Items", gocui.KeyEnter, gocui.ModNone, openItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, refreshFeeds); err != nil {
		return err
	}
	return nil
}
