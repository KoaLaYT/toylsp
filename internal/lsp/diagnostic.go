package lsp

type PublishDiagnosticNotification struct {
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct {
	URI         string        `json:"uri"`
	Diagnostics []Diagnostics `json:"diagnostics"`
}

type Diagnostics struct {
	Range    Range              `json:"range"`
	Severity DiagnosticSeverity `json:"severity"`
	Source   string             `json:"source"`
	Message  string             `json:"message"`
}

type DiagnosticSeverity uint8

const (
	DiagnosticSeverityError DiagnosticSeverity = iota + 1
	DiagnosticSeverityWarning
	DiagnosticSeverityInformation
	DiagnosticSeverityHint
)
