package cui

import (
	"strings"

	"github.com/jroimartin/gocui"
)

func scroll(g *gocui.Gui, v *gocui.View, direction int) error {
	if v != nil {
		_, y := v.Size()
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
		// If we're nearing a boundary
		if oy+direction > strings.Count(v.ViewBuffer(), "\n")+y-direction {
			v.Autoscroll = true
		} else {
			v.Autoscroll = false
			v.SetOrigin(ox, oy+direction)
			return nil
		}
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	return scroll(g, v, 1)
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	return scroll(g, v, -1)
}
