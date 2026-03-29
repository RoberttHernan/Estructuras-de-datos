package main

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("VLangCherry IDE")

	// === Editor ===
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Escribe tu código VLangCherry aquí...")

	// === Consola ===
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Salida del intérprete")
	output.Disable()

	// === Botones ===
	btnOpen := widget.NewButton("Abrir", func() {
		dialog.ShowFileOpen(func(r fyne.URIReadCloser, _ error) {
			if r != nil {
				data, _ := ioutil.ReadAll(r)
				editor.SetText(string(data))
				r.Close()
			}
		}, w)
	})

	btnSave := widget.NewButton("Guardar", func() {
		dialog.ShowFileSave(func(wc fyne.URIWriteCloser, _ error) {
			if wc != nil {
				wc.Write([]byte(editor.Text))
				wc.Close()
			}
		}, w)
	})

	btnRun := widget.NewButton("Ejecutar", func() {
		salida := EjecutarCodigo(editor.Text)
		output.SetText(salida)
	})

	// === Panel horizontal: editor | consola ===
	hSplit := container.NewHSplit(
		container.NewMax(editor),
		container.NewMax(output),
	)
	hSplit.Offset = 0.5 // 50/50

	// === Panel vertical: (botones arriba) | (hSplit abajo) ===
	vSplit := container.NewVSplit(
		container.NewVBox(btnOpen, btnSave, btnRun),
		hSplit,
	)
	vSplit.Offset = 0.1 // botones ocupan 10% vertical

	// === Final ===
	w.SetContent(vSplit)
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}
