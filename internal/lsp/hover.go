package lsp

import (
	"encoding/json"
	"fmt"
)

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

func DecodeHoverParams(data []byte) (int, *HoverParams, error) {
	var msg HoverRequest
	if err := json.Unmarshal(data, &msg); err != nil {
		return 0, nil, fmt.Errorf("decode %s message: %w", HoverMethod, err)
	}
	return msg.ID, &msg.Params, nil
}

type HoverResponse struct {
	Response
	Result *Hover `json:"result"`
}

type Hover struct {
	Contents MarkupContent `json:"contents"`
	Range    *Range        `json:"range"`
}

type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

type MarkupKind string

const (
	MarkupKindPlainText MarkupKind = "plaintext"
	MarkupKindMarkdown  MarkupKind = "markdown"
)

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
