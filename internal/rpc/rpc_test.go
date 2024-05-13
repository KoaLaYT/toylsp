package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentLength(t *testing.T) {
	for _, tt := range []struct {
		input       string
		expected    int
		expectedErr error
	}{
		{
			"Content-Length: 16",
			16,
			nil,
		},
		{
			"Content-Length: 16\r\nContent-Type: application/vscode-jsonrpc; charset=utf-8",
			16,
			nil,
		},
	} {
		got, gotErr := getContentLength([]byte(tt.input))
		assert.Equal(t, tt.expected, got)
		if tt.expectedErr != nil {
			assert.Equal(t, tt.expectedErr.Error(), gotErr.Error())
		} else {
			assert.Nil(t, gotErr)
		}
	}
}
