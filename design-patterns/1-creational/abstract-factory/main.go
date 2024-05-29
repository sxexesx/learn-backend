package abstractfactory

type IAbstractFactory interface {
	CreateButton() IButton
}

func getUI(osType string) IAbstractFactory {
	if osType == "win" {
		return &WinFactory{}
	}
	if osType == "mac" {
		return &MacFactory{}
	}
	return nil
}

// WinFactory concrete fabric
type WinFactory struct {
}

func (w *WinFactory) CreateButton() IButton {
	return &WinButton{
		Button: Button{
			height: 100,
			width:  100,
		},
	}
}

// MacFactory concrete fabric
type MacFactory struct {
}

func (m *MacFactory) CreateButton() IButton {
	return &MacButton{
		Button: Button{
			height: 80,
			width:  80,
		},
	}
}

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

func (b *Button) getHeight() int       { return b.height }
func (b *Button) setHeight(height int) { b.height = height }
func (b *Button) getWidth() int        { return b.width }
func (b *Button) setWidth(width int)   { b.width = width }

type WinButton struct {
	Button
}

type MacButton struct {
	Button
}

func main() {
	ui := getUI("win")
	button := ui.CreateButton()
	println(button.getHeight())
}
