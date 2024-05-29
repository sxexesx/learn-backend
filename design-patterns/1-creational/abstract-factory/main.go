package abstractfactory

import "fmt"

type IButton interface {
	getHeight() int
	setHeight(height int)

	getWidth() int
	setWidth(width int)
}

type Button struct {
	height int
	width  int
}

func (m *Button) getHeight() int       { return m.height }
func (m *Button) setHeight(height int) { m.height = height }
func (m *Button) getWidth() int        { return m.width }
func (m *Button) setWidth(width int)   { m.width = width }

type ICheckbox interface {
	getHeight() int
	setHeight(height int)

	getWidth() int
	setWidth(width int)
}

type Checkbox struct {
	height int
	width  int
}

func (c *Checkbox) getHeight() int       { return c.height }
func (c *Checkbox) setHeight(height int) { c.height = height }
func (c *Checkbox) getWidth() int        { return c.width }
func (c *Checkbox) setWidth(width int)   { c.width = width }

type MacButton struct{}

func (m *MacButton) createButton() IButton {
	return &Button{
		height: 100,
		width:  150,
	}
}
func (m *MacButton) createCheckbox() ICheckbox {
	return &Checkbox{
		height: 20,
		width:  20,
	}
}

type WinButton struct{}

func (m *WinButton) createButton() IButton {
	return &Button{
		height: 80,
		width:  100,
	}
}
func (m *WinButton) createCheckbox() ICheckbox {
	return &Checkbox{
		height: 25,
		width:  25,
	}
}

// func (m *MacButton) getHeight() int       { return m.height }
// func (m *MacButton) setHeight(height int) { m.height = height }
// func (m *MacButton) getWidth() int        { return m.width }
// func (m *MacButton) setWidth(width int)   { m.width = width }

// type WinButton struct {
// 	height int
// 	width  int
// }

// func (m *MacButton) getHeight() int       { return m.height }
// func (m *MacButton) setHeight(height int) { m.height = height }
// func (m *MacButton) getWidth() int        { return m.width }
// func (m *MacButton) setWidth(width int)   { m.width = width }

type IAbstractFactory interface {
	createButton() IButton
	createCheckbox() ICheckbox
}

type WinFactory struct{}

func (w *WinFactory) createButton() {
}

func (w *WinFactory) createCheckbox() {
}

type MacFactory struct{}

func (m *MacFactory) createButton() IButton {
}

func (m *MacFactory) createCheckbox() ICheckbox {
}

func getUI(osType string) (*IAbstractFactory, error) {
	if osType == "win" {
		return &WinFactory{}
	}
	if osType == "mac" {

	}
}

func main() {
}
