package rpc

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

const HeaderContentLength = "Content-Length: "

func errMalformedHeader(header []byte, msg string) error {
	return fmt.Errorf("Malformed header, %s: %s", msg, string(header))
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
