package lsp

import (
	"encoding/json"
	"fmt"
)

type InitializeMessage struct {
	Request
	Params InitializeParam `json:"params"`
}

type InitializeParam struct {
	ProcessID  *int        `json:"processId"`
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version"`
}

func (ci ClientInfo) String() string {
	return fmt.Sprintf("Client [%s %s]", ci.Name, *ci.Version)
}

func DecodeInitializeParams(msg []byte) (int, *InitializeParam, error) {
	var im InitializeMessage
	if err := json.Unmarshal(msg, &im); err != nil {
		return 0, nil, fmt.Errorf("decode initialize message: %w", err)
	}
	return im.ID, &im.Params, nil
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   TextDocumentSyncOptions `json:"textDocumentSync"`
	CompletionProvider CompletionOptions       `json:"completionProvider"`
	HoverProvider      bool                    `json:"hoverProvider"`
}

type TextDocumentSyncOptions struct {
	OpenClose bool                 `json:"openClose"`
	Change    TextDocumentSyncKind `json:"change"`
}

type TextDocumentSyncKind uint8

const (
	TextDocumentSyncKindNone = iota
	TextDocumentSyncKindFull
	TextDocumentSyncKindIncremental
)

type CompletionOptions struct{}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func EncodeInitializeResponse(id int) ([]byte, error) {
	resp := InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   TextDocumentSyncOptions{OpenClose: true, Change: TextDocumentSyncKindFull},
				CompletionProvider: CompletionOptions{},
				HoverProvider:      true,
			},
			ServerInfo: ServerInfo{
				Name:    "Toy LSP",
				Version: "The greatest version ever",
			},
		},
	}
	return json.Marshal(resp)
}
