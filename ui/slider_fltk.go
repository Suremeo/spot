//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeSlider = *goFltk.Slider

func (c *Slider) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Slider)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	if next.Min != c.Min {
		c.Min = next.Min
		c.ref.SetMinimum(c.Min)
	}

	if next.Max != c.Max {
		c.Max = next.Max
		c.ref.SetMaximum(c.Max)
	}

	if next.Value != c.Value {
		c.Value = next.Value
		c.ref.SetValue(c.Value)
	}

	c.OnValueChanged = next.OnValueChanged

	return true
}

func (c *Slider) callback() {
	if c.OnValueChanged != nil {
		c.OnValueChanged(c.ref.Value())
	}
}

func (c *Slider) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	c.ref = goFltk.NewSlider(c.X, c.Y, c.Width, c.Height)
	c.ref.SetMaximum(c.Max)
	c.ref.SetMinimum(c.Min)
	c.ref.SetValue(c.Value)
	// b.ref.SetType(b.Type)
	// b.ref.SetType(goFltk.HOR_SLIDER)
	c.ref.SetType(goFltk.HOR_NICE_SLIDER)
	c.ref.SetBox(goFltk.FLAT_BOX)
	c.ref.SetCallback(c.callback)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *Slider) Unmount() {
	if c.ref != nil {
		c.ref.Destroy()
		c.ref = nil
	}
}
