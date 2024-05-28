package main

func main() {

}

type Button interface {
	render()
	onClick()
}

type WindowsButton struct {
}

func (w *WindowsButton) render()  {}
func (w *WindowsButton) onClick() {}

type MacButton struct {
}

func (m *MacButton) render()  {}
func (m *MacButton) onClick() {}

type Dialog struct {
}

func (d *Dialog) render() {

}
