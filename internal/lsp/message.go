package lsp

const (
	InitializeMethod            = "initialize"
	DidChangeTextDocumentMethod = "textDocument/didChange"
	DidOpenTextDocumentMethod   = "textDocument/didOpen"
	HoverMethod                 = "textDocument/hover"
	CompletionMethod            = "textDocument/completion"
	PublishDiagnosticMethod     = "textDocument/publishDiagnostics"
)

type BaseMessage struct {
	Method string `json:"method"`
}

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
