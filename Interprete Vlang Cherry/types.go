package main

type ExecutionRequest struct {
	Code string `json:"code"`
}

type ExecutionResponse struct {
	Output       string          `json:"output"`
	Errors       []ErrorReport   `json:"errors"`
	SymbolTable  []SymbolEntry   `json:"symbolTable"`
	AST          string          `json:"ast"`
}

type ErrorReport struct {
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	Message     string `json:"message"`
	ErrorType   string `json:"errorType"`
}

type SymbolEntry struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	DataType string `json:"dataType"`
	Scope    string `json:"scope"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
}