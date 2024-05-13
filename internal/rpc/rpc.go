package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/KoaLaYT/toylsp/internal/lsp"
)

const HeaderContentLength = "Content-Length: "

func errMalformedHeader(header []byte, msg string) error {
	return fmt.Errorf("malformed header, %s: %s", msg, string(header))
}

func Decode(data []byte) (string, []byte, error) {
	var req lsp.BaseMessage
	if err := json.Unmarshal(data, &req); err != nil {
		return "", nil, fmt.Errorf("decode message: %w", err)
	}
	return req.Method, data, nil
}

func Encode(data []byte) []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(data)))
	buf.Write(data)
	return buf.Bytes()
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		if atEOF && len(data) > 0 {
			return 0, nil, io.ErrUnexpectedEOF
		}
		return 0, nil, nil
	}

	contentLength, err := getContentLength(header)
	if err != nil {
		return 0, nil, err
	}

	// wait for more!
	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + len(content)
	return totalLength, data[len(header)+4 : totalLength], nil
}

func getContentLength(header []byte) (int, error) {
	conetentLengthIndex := bytes.Index(header, []byte(HeaderContentLength))
	if conetentLengthIndex == -1 {
		return 0, errMalformedHeader(header, "No 'Content-Length' found")
	}

	contentLengthBytes := header[conetentLengthIndex+len(HeaderContentLength):]
	endIndex := bytes.Index(contentLengthBytes, []byte{'\r', '\n'})
	if endIndex != -1 {
		contentLengthBytes = contentLengthBytes[:endIndex]
	}
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, errMalformedHeader(header, "Invalid content length")
	}

	return contentLength, nil
}
