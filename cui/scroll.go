package cui

import (
	"github.com/jroimartin/gocui"
)

func scroll(v *gocui.View, direction int) error {
	if v != nil {
		ox, oy := v.Origin()
		if oy+direction >= len(v.BufferLines())-1 {
			// hit bottom
			return nil
		}
		if oy+direction < 0 {
			// hit top
			return nil
		}
		if err := v.SetOrigin(ox, oy+direction); err != nil {
			return err
		}
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	err := scroll(v, 1)
	if g.CurrentView().Title == "Items" {
		displayDescription(g, v)
	}
	return err
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	err := scroll(v, -1)
	if g.CurrentView().Title == "Items" {
		displayDescription(g, v)
	}
	return err
}
