package generadorcodigo

import (
	"fmt"
	"strings"
)

type CodeBuilder struct {
	textLines  []string
	dataLines  []string
	labelCount int

	globalVars map[string]bool

	UsesPrintInt bool // ← nuevo campo

	breakStack []string
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		textLines:  []string{},
		dataLines:  []string{},
		labelCount: 0,

		globalVars: make(map[string]bool),
	}
}

func (cb *CodeBuilder) AddText(line string) {
	cb.textLines = append(cb.textLines, line)
}

func (cb *CodeBuilder) AddData(line string) {
	cb.dataLines = append(cb.dataLines, line)
}

func (cb *CodeBuilder) AddLine(line string) {
	cb.AddText(line) // alias por compatibilidad
}

func (cb *CodeBuilder) AddComment(comment string) {
	cb.AddText(fmt.Sprintf("# %s", comment))
}

func (cb *CodeBuilder) NewLabel(prefix string) string {
	cb.labelCount++
	return fmt.Sprintf("%s%d", prefix, cb.labelCount)
}

func (cb *CodeBuilder) String() string {
	var out strings.Builder
	out.WriteString(".section .text\n")
	for _, l := range cb.textLines {
		out.WriteString(l + "\n")
	}
	out.WriteString("\n.section .data\n")
	for _, l := range cb.dataLines {
		out.WriteString(l + "\n")
	}
	return out.String()
}

func (cb *CodeBuilder) DeclareGlobalVar(name string) {
	if !cb.globalVars[name] {
		cb.globalVars[name] = true
		cb.dataLines = append(cb.dataLines, fmt.Sprintf("_%s: .quad 0", name))
	}
}

func (cb *CodeBuilder) EmitAll() string {
	var full []string
	full = append(full, ".section .text")
	full = append(full, cb.textLines...)
	full = append(full, "")
	full = append(full, ".section .data")
	full = append(full, cb.dataLines...)
	return strings.Join(full, "\n")
}

func (cb *CodeBuilder) PushBreakLabel(label string) {
	cb.breakStack = append(cb.breakStack, label)
}

func (cb *CodeBuilder) PopBreakLabel() {
	if len(cb.breakStack) > 0 {
		cb.breakStack = cb.breakStack[:len(cb.breakStack)-1]
	}
}

func (cb *CodeBuilder) CurrentBreakLabel() string {
	if len(cb.breakStack) == 0 {
		return "⚠️ no-break-label"
	}
	return cb.breakStack[len(cb.breakStack)-1]
}
