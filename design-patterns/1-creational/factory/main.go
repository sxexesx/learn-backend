package main

import "fmt"

type IButton interface {
	setName(name string)
	getName() string
	setAction(action string)
	getAction() string
}

type Button struct {
	name   string
	action string
}

func (b *Button) setName(name string) {
	b.name = name
}

func (b *Button) getName() string {
	return b.name
}

func (b *Button) setAction(action string) {
	b.action = action
}

func (b *Button) getAction() string {
	return b.action
}

type WindowsButton struct {
	Button
}

func NewWindowsButton() IButton {
	return &WindowsButton{
		Button: Button{
			name:   "windows-button",
			action: "windows-action",
		},
	}
}

type MacButton struct {
	Button
}

func NewMacButton() IButton {
	return &MacButton{
		Button: Button{
			name:   "mac-button",
			action: "mac-action",
		},
	}
}

// fabric
func getButton(buttonType string) IButton {
	if buttonType == "windows-button" {
		return NewWindowsButton()
	}
	if buttonType == "mac-button" {
		return NewMacButton()
	}
	return nil
}

// client
func main() {
	winButton := getButton("windows-button")
	fmt.Printf("Button: %s; %s\n", winButton.getName(), winButton.getAction())

	macButton := getButton("mac-button")
	fmt.Printf("Button: %s; %s\n", macButton.getName(), macButton.getAction())
}
