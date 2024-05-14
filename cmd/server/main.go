package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/KoaLaYT/toylsp/internal/lsp"
	"github.com/KoaLaYT/toylsp/internal/rpc"
)

func main() {
	s := newServer()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		raw := scanner.Bytes()
		method, msg, err := rpc.Decode(raw)
		if err != nil {
			s.log.Fatal(err)
		}
		s.handleMsg(method, msg)
	}
}

type server struct {
	log   *log.Logger
	state *lsp.State
}

func newServer() *server {
	s := new(server)

	s.log = initLogger("/Users/koalayt/Projects/toylsp/log.txt")
	s.state = lsp.NewState()

	s.log.Println("started")
	return s
}

func (s *server) fatalOn(err error) {
	if err != nil {
		s.log.Fatal(err)
	}
}

func (s *server) reply(resp []byte) {
	raw := rpc.Encode(resp)
	if _, err := os.Stdout.Write(raw); err != nil {
		s.log.Fatal("write to stdout: ", err)
	}
}

func (s *server) handleMsg(method string, msg []byte) {
	switch method {
	case lsp.InitializeMethod:
		s.handleInitializeRequest(msg)
	case lsp.DidOpenTextDocumentMethod:
		s.handleDidOpenTextDocumentNotification(msg)
	case lsp.DidChangeTextDocumentMethod:
		s.handleDidChangeTextDocumentNotification(msg)
	case lsp.CompletionMethod:
		s.handleCompletionRequest(msg)
	case lsp.HoverMethod:
		s.handleHoverRequest(msg)
	default:
		s.log.Printf("unknown: %s, %v", method, string(msg))
	}
}

func (s *server) handleCompletionRequest(msg []byte) {
	id, param, err := lsp.DecodeCompletionParams(msg)
	s.fatalOn(err)

	s.log.Printf("%s (line: %d, char: %d): %s\n", lsp.CompletionMethod,
		param.Position.Line, param.Position.Character, param.TextDocument.URI)

	resp, err := s.state.ResolveCompletion(id)
	s.fatalOn(err)
	s.reply(resp)
}

func (s *server) handleHoverRequest(msg []byte) {
	id, param, err := lsp.DecodeHoverParams(msg)
	s.fatalOn(err)

	s.log.Printf("%s (line: %d, char: %d): %s\n",
		lsp.HoverMethod, param.Position.Line, param.Position.Character, param.TextDocument.URI)

	resp, err := s.state.ResolveHover(id, param.TextDocument.URI, param.Position)
	s.fatalOn(err)

	s.reply(resp)
}

func (s *server) handleDidOpenTextDocumentNotification(msg []byte) {
	param, err := lsp.DecodeDidOpenTextDocumentParams(msg)
	s.fatalOn(err)

	s.log.Printf("%s (version: %d): %s\n", lsp.DidOpenTextDocumentMethod,
		param.TextDocument.Version, param.TextDocument.URI)
	s.state.AddFile(param.TextDocument.URI, param.TextDocument.Text)
}

func (s *server) handleDidChangeTextDocumentNotification(msg []byte) {
	param, err := lsp.DecodeDidChangeTextDocumentParams(msg)
	s.fatalOn(err)

	s.log.Printf("%s (version: %d): %s\n", lsp.DidChangeTextDocumentMethod,
		param.TextDocument.Version, param.TextDocument.URI)
	s.state.UpdateFile(param.TextDocument.URI, param.ContentChanges[0].Text)
}

func (s *server) handleInitializeRequest(msg []byte) {
	// s.log.Println(string(msg))
	id, param, err := lsp.DecodeInitializeParams(msg)
	s.fatalOn(err)

	s.log.Printf("Connected to %s (PID: %d)\n", param.ClientInfo, *param.ProcessID)

	resp, err := lsp.EncodeInitializeResponse(id)
	s.fatalOn(err)

	s.reply(resp)
}

func initLogger(filepath string) *log.Logger {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("openfile: %v", err))
	}
	return log.New(f, "[toylsp]", log.LstdFlags)
}
