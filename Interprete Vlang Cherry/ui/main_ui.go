package main

/*
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setupUI(window fyne.Window) {
	// Componentes principales
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Escribe tu código VLangCherry aquí...")
	editor.Wrapping = fyne.TextWrapWord

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Salida del intérprete")
	output.Disable()

	// Tabs para los reportes
	tabs := container.NewAppTabs(
		container.NewTabItem("Editor", editor),
		container.NewTabItem("Consola", output),
		container.NewTabItem("Errores", createErrorReportTable()),
		container.NewTabItem("Tabla de Símbolos", createSymbolTable()),
		container.NewTabItem("AST", widget.NewMultiLineEntry()),
	)

	// Botones de la interfaz
	btnOpen := widget.NewButtonWithIcon("Abrir", theme.FolderOpenIcon(), func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			editor.SetText(string(data))
		}, window)

		fileDialog.SetFilter(fyne.NewExtensionFileFilter([]string{".vch"}))
		fileDialog.Show()
	})

	btnSave := widget.NewButtonWithIcon("Guardar", theme.DocumentSaveIcon(), func() {
		fileDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			_, err = writer.Write([]byte(editor.Text))
			if err != nil {
				dialog.ShowError(err, window)
			}
		}, window)

		fileDialog.SetFilter(fyne.NewExtensionFileFilter([]string{".vch"}))
		fileDialog.Show()
	})

	btnRun := widget.NewButtonWithIcon("Ejecutar", theme.MediaPlayIcon(), func() {
		code := editor.Text
		if code == "" {
			dialog.ShowInformation("Aviso", "No hay código para ejecutar", window)
			return
		}

		progress := dialog.NewCustom("Ejecutando", "Cancelar", widget.NewProgressBarInfinite(), window)
		progress.Show()

		go func() {
			defer progress.Hide()

			response, err := sendToBackend(code)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Actualizar salida
			output.SetText(response.Output)

			// Actualizar pestañas de reportes
			updateErrorReportTable(tabs.Items[2].Content.(*widget.Table), response.Errors)
			updateSymbolTable(tabs.Items[3].Content.(*widget.Table), response.SymbolTable)
			tabs.Items[4].Content.(*widget.Entry).SetText(response.AST)
		}()
	})

	btnViewAST := widget.NewButtonWithIcon("Ver AST", theme.VisibilityIcon(), func() {
		exec.Command("xdg-open", "http://localhost:8080/ast-visual").Start()
	})

	btnExportPDF := widget.NewButtonWithIcon("Exportar PDF", theme.DocumentPrintIcon(), func() {
		resp, err := http.Get("http://localhost:8080/export/pdf")
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer resp.Body.Close()

		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			defer writer.Close()
			io.Copy(writer, resp.Body)
		}, window)
		saveDialog.SetFileName("report.pdf")
		saveDialog.Show()
	})

	btnExportHTML := widget.NewButtonWithIcon("Exportar HTML", theme.FileTextIcon(), func() {
		resp, err := http.Get("http://localhost:8080/export/html")
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer resp.Body.Close()

		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			defer writer.Close()
			io.Copy(writer, resp.Body)
		}, window)
		saveDialog.SetFileName("report.html")
		saveDialog.Show()
	})

	// Barra de herramientas
	toolbar := container.NewHBox(
		btnOpen,
		btnSave,
		btnRun,
		widget.NewSeparator(),
		btnViewAST,
		btnExportPDF,
		btnExportHTML,
		layout.NewSpacer(),
		widget.NewLabel("VLangCherry IDE"),
	)

	// Layout principal
	content := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		tabs,
	)

	window.SetContent(content)
}

// Funciones para manejar tablas de reportes
func createErrorReportTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return 0, 4 },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText("")
		},
	)
	table.SetColumnWidth(0, 50)  // Línea
	table.SetColumnWidth(1, 50)  // Columna
	table.SetColumnWidth(2, 150) // Tipo
	table.SetColumnWidth(3, 500) // Mensaje
	return table
}

func createSymbolTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return 0, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText("")
		},
	)
	table.SetColumnWidth(0, 150) // ID
	table.SetColumnWidth(1, 100) // Tipo
	table.SetColumnWidth(2, 100) // Tipo de dato
	table.SetColumnWidth(3, 100) // Ámbito
	table.SetColumnWidth(4, 50)  // Línea
	return table
}

func updateErrorReportTable(table *widget.Table, errors []ErrorReport) {
	table.Length = func() (int, int) { return len(errors), 4 }
	table.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		label := cell.(*widget.Label)
		err := errors[id.Row]
		switch id.Col {
		case 0:
			label.SetText(fmt.Sprintf("%d", err.Line))
		case 1:
			label.SetText(fmt.Sprintf("%d", err.Column))
		case 2:
			label.SetText(err.ErrorType)
		case 3:
			label.SetText(err.Message)
		}
	}
	table.Refresh()
}

func updateSymbolTable(table *widget.Table, symbols []SymbolEntry) {
	table.Length = func() (int, int) { return len(symbols), 5 }
	table.UpdateCell = func(id widget.TableCellID, cell fyne.CanvasObject) {
		label := cell.(*widget.Label)
		sym := symbols[id.Row]
		switch id.Col {
		case 0:
			label.SetText(sym.ID)
		case 1:
			label.SetText(sym.Type)
		case 2:
			label.SetText(sym.DataType)
		case 3:
			label.SetText(sym.Scope)
		case 4:
			label.SetText(fmt.Sprintf("%d", sym.Line))
		}
	}
	table.Refresh()
}

func sendToBackend(code string) (*ExecutionResponse, error) {
	request := ExecutionRequest{Code: code}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la solicitud: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/execute", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error al conectar con el backend: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en el backend: %s", resp.Status)
	}

	var response ExecutionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	return &response, nil
}
*/
