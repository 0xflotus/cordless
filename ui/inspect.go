package ui

import "github.com/Bios-Marcel/tview"

type InspectionDialog struct {
	view tview.Primitive
}

const imagePlaceHolder = `[red]█[green]█[red]█[green]█[red]█[green]█[red]█[green]█
[green]█[red]█[green]█[red]█[green]█[red]█[green]█[red]█
[red]█[green]█[red]█[green]█[red]█[green]█[red]█[green]█
[green]█[red]█[green]█[red]█[green]█[red]█[green]█[red]█
[red]█[green]█[red]█[green]█[red]█[green]█[red]█[green]█
[green]█[red]█[green]█[red]█[green]█[red]█[green]█[red]█
[red]█[green]█[red]█[green]█[red]█[green]█[red]█[green]█
[green]█[red]█[green]█[red]█[green]█[red]█[green]█[red]█`

func NewInspectionDialog() *InspectionDialog {
	dialog := &InspectionDialog{}
	// The header contains a pixelated version of the users avatar and some
	// textual information.
	header := tview.NewFlex().SetDirection(tview.FlexColumn)

	image := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetText(imagePlaceHolder)
	header.AddItem(image, 8, 0, false)

	infoText := tview.NewTextView().SetText("(Nickname) User#1234\nNitro: Yes\nUser-ID: 8912347815879125687235\nRelationship: Friend")
	header.AddItem(infoText, 0, 1, false)

	leftArea := tview.NewFlex()
	leftArea.SetDirection(tview.FlexRow)

	serverList := tview.NewTreeView()
	leftArea.AddItem(serverList, 0, 1, false)

	userList := tview.NewTreeView()
	leftArea.AddItem(userList, 0, 1, false)

	rightArea := tview.NewFlex()
	rightArea.SetDirection(tview.FlexRow)

	message := tview.NewTextView().SetText("12:04 - 19.04.2019 (Last edited 13-04 - 20.04.2019)\n\nExample message")
	message.SetDynamicColors(true).SetRegions(true)
	rightArea.AddItem(message, 0, 1, false)

	actionArea := tview.NewFlex().SetDirection(tview.FlexColumn)
	rightArea.AddItem(actionArea, 3, 0, false)

	root := tview.NewFlex().SetDirection(tview.FlexRow)
	root.AddItem(header, 8, 0, false)

	body := tview.NewFlex().SetDirection(tview.FlexColumn)
	body.AddItem(leftArea, 0, 5, false)
	body.AddItem(rightArea, 0, 20, false)

	root.AddItem(body, 0, 1, false)

	dialog.view = root

	return dialog
}

func (i *InspectionDialog) GetPrimitive() tview.Primitive {
	return i.view
}
