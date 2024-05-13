package lsp

import (
	"encoding/json"
	"fmt"
)

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type VersionedTextDocumentIdentifier struct {
	URI     string `json:"uri"`
	Version int    `json:"version"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

func DecodeDidChangeTextDocumentParams(data []byte) (*DidChangeTextDocumentParams, error) {
	var msg DidChangeTextDocumentNotification
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("decode %s message: %w", DidChangeTextDocumentMethod, err)
	}
	return &msg.Params, nil
}

type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

func DecodeDidOpenTextDocumentParams(data []byte) (*DidOpenTextDocumentParams, error) {
	var msg DidOpenTextDocumentNotification
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("decode %s message: %w", DidOpenTextDocumentMethod, err)
	}
	return &msg.Params, nil
}
