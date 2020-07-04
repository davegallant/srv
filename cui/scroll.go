package cui

import (
	"github.com/jroimartin/gocui"
)

func scroll(g *gocui.Gui, v *gocui.View, direction int) error {
	if v != nil {
		ox, oy := v.Origin()
		if oy+direction >= len(v.BufferLines())-1 {
			return nil
		}
		v.SetOrigin(ox, oy+direction)
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	err := scroll(g, v, 1)
	if g.CurrentView().Title == "Items" {
		displayDescription(g, v)
	}
	return err
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	err := scroll(g, v, -1)
	if g.CurrentView().Title == "Items" {
		displayDescription(g, v)
	}
	return err
}
