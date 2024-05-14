package lsp

import (
	"encoding/json"
	"fmt"
)

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

func DecodeCompletionParams(data []byte) (int, *CompletionParams, error) {
	var msg CompletionRequest
	if err := json.Unmarshal(data, &msg); err != nil {
		return 0, nil, fmt.Errorf("decode %s message: %w", CompletionMethod, err)
	}
	return msg.ID, &msg.Params, nil
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string        `json:"label"`
	Kind          int           `json:"kind"` // 1-25
	Detail        string        `json:"detail"`
	Documentation MarkupContent `json:"documentation"`
	// InsertText       string        `json:"insertText"`
	// InsertTextFormat int           `json:"insertTextFormat"`
	// InsertTextMode   int           `json:"insertTextMode"`
}
