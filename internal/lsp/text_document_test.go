package lsp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeDidChangeTextDocumentNotification(t *testing.T) {
	input := `{"params":{"textDocument":{"languageId":"toy","text":"wdefw\nkqwfqqw\nkqwfqqw\nwef b\n","version":0,"uri":"file:\/\/\/Users\/koalayt\/Projects\/toylsp\/test.toy"}},"method":"textDocument\/didOpen","jsonrpc":"2.0"}`

	params, err := DecodeDidChangeTextDocumentParams([]byte(input))
	assert.Nil(t, err)
	_ = params
}

func TestDecodeDidOpenTextDocumentNotification(t *testing.T) {
	input := `{"params":{"textDocument":{"uri":"file:\/\/\/Users\/koalayt\/Projects\/toylsp\/test.toy","version":0,"text":"wdefw\nqjwkldjqlwk\n","languageId":"toy"}},"jsonrpc":"2.0","method":"textDocument\/didOpen"}`

	params, err := DecodeDidOpenTextDocumentParams([]byte(input))
	assert.Nil(t, err)
	_ = params
}
