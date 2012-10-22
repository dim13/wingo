package xclient

import (
	"github.com/BurntSushi/wingo/layout"
	"github.com/BurntSushi/wingo/stack"
	"github.com/BurntSushi/wingo/wm"
	"github.com/BurntSushi/wingo/workspace"
)

func (c *Client) FloatingToggle() {
	if c.floating {
		c.Unfloat()
	} else {
		c.Float()
	}
}

func (c *Client) Float() {
	if c.floating {
		return
	}
	if wrk, ok := c.Workspace().(*workspace.Workspace); ok {
		c.floating = true
		wrk.CheckFloatingStatus(c)
	}
}

func (c *Client) Unfloat() {
	if !c.floating {
		return
	}
	if wrk, ok := c.Workspace().(*workspace.Workspace); ok {
		c.floating = false
		wrk.CheckFloatingStatus(c)
	}
}

func (c *Client) StackAboveToggle() {
	if c.layer == stack.LayerAbove {
		c.unstackAbove()
	} else {
		c.stackAbove()
	}
}

func (c *Client) stackAbove() {
	c.layer = stack.LayerAbove
	c.Raise()

	c.removeState("_NET_WM_STATE_BELOW")
	c.addState("_NET_WM_STATE_ABOVE")
}

func (c *Client) unstackAbove() {
	c.layer = stack.LayerDefault
	c.Raise()

	c.removeState("_NET_WM_STATE_ABOVE")
}

func (c *Client) StackBelowToggle() {
	if c.layer == stack.LayerBelow {
		c.unstackBelow()
	} else {
		c.stackBelow()
	}
}

func (c *Client) stackBelow() {
	c.layer = stack.LayerBelow
	c.Raise()

	c.removeState("_NET_WM_STATE_ABOVE")
	c.addState("_NET_WM_STATE_BELOW")
}

func (c *Client) unstackBelow() {
	c.layer = stack.LayerDefault
	c.Raise()

	c.removeState("_NET_WM_STATE_BELOW")
}

func (c *Client) StickyToggle() {
	if c.sticky {
		c.unstick()
	} else {
		c.stick()
	}
}

func (c *Client) unstick() {
	c.sticky = false
	c.workspace = nil
	wm.Workspace().Add(c)

	c.removeState("_NET_WM_STATE_STICKY")
}

func (c *Client) stick() {
	if c.sticky {
		return
	}

	c.sticky = true
	if c.workspace != nil {
		c.workspace.(*workspace.Workspace).CheckFloatingStatus(c)
		c.workspace.Remove(c)
	}
	c.WorkspaceSet(wm.StickyWrk)

	c.addState("_NET_WM_STATE_STICKY")
}

func (c *Client) FullscreenToggle() {
	if c.fullscreen {
		c.Fullscreened()
	} else {
		c.Unfullscreened()
	}
}

func (c *Client) Fullscreened() {
	if c.workspace == nil || !c.workspace.IsVisible() {
		return
	}
	if c.fullscreen {
		return
	}
	if _, ok := c.Layout().(layout.Floater); ok {
		c.SaveState("last-floating")
	}
	c.fullscreen = true

	// Make sure the window has been forced into a floating layout.
	if wrk, ok := c.Workspace().(*workspace.Workspace); ok {
		wrk.CheckFloatingStatus(c)
	}

	// Resize outside of the constraints of a layout.
	g := c.Workspace().HeadGeom()
	c.FrameNada()
	c.MoveResize(g.X(), g.Y(), g.Width(), g.Height())

	c.addState("_NET_WM_STATE_FULLSCREEN")
}

func (c *Client) Unfullscreened() {
	if !c.fullscreen {
		return
	}
	c.fullscreen = false

	// Make sure the window is no longer forced into a floating layout just
	// because of its fullscreen status.
	if wrk, ok := c.Workspace().(*workspace.Workspace); ok {
		wrk.CheckFloatingStatus(c)
	}

	// If the window's layout is now floating, restore geometry.
	if _, ok := c.Layout().(layout.Floater); ok {
		c.LoadState("last-floating")
	}

	c.removeState("_NET_WM_STATE_FULLSCREEN")
}

func (c *Client) MaximizeToggle() {
	if c.IsMaximized() {
		c.Unmaximize()
	} else {
		c.Maximize()
	}
}

func (c *Client) Maximize() {
	if !c.canMaxUnmax() {
		return
	}
	if !c.IsMaximized() {
		c.SaveState("before-maximize")
		c.maximize()
	}
}

func (c *Client) Unmaximize() {
	if !c.canMaxUnmax() {
		return
	}
	if c.IsMaximized() {
		c.unmaximize()
		c.LoadState("before-maximize")
	}
}

func (c *Client) Remaximize() {
	if !c.IsMaximized() {
		return
	}
	c.maximize()
}

func (c *Client) maximize() {
	if !c.canMaxUnmax() {
		return
	}

	c.maximized = true
	c.addState("_NET_WM_STATE_MAXIMIZE_HORZ")
	c.addState("_NET_WM_STATE_MAXIMIZE_VERT")

	c.frames.maximize()

	g := c.Workspace().Geom()
	c.LayoutMoveResize(g.X(), g.Y(), g.Width(), g.Height())
}

func (c *Client) unmaximize() {
	if c.Workspace() == nil || !c.Workspace().IsVisible() {
		return
	}
	if c.maximized {
		c.maximized = false
		c.removeState("_NET_WM_STATE_MAXIMIZE_HORZ")
		c.removeState("_NET_WM_STATE_MAXIMIZE_VERT")
		c.frames.unmaximize()
	}
}

func (c *Client) canMaxUnmax() bool {
	if c.Workspace() == nil || !c.Workspace().IsVisible() {
		return false
	}
	if _, ok := c.Layout().(layout.Floater); !ok {
		return false
	}
	if c.fullscreen {
		return false
	}
	return true
}