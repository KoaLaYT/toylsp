package lsp

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TODO: currently, only one thread, no need to add lock
type State struct {
	Files map[string]string
}

func NewState() *State {
	return &State{Files: make(map[string]string)}
}

func (s *State) AddFile(file, content string) {
	s.Files[file] = content
}

// TODO: probably need to consider the version
func (s *State) UpdateFile(file, content string) {
	s.Files[file] = content
}

func (s *State) GetFile(file string) string {
	content, found := s.Files[file]
	if !found {
		return ""
	}
	return content
}

// TODO: we are now just generating garbages
func (s *State) ResolveCompletion(id int) ([]byte, error) {
	resp := CompletionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: []CompletionItem{
			{
				Label: "Dinosaur",
				// InsertText:       "Dinosaur",
				// InsertTextFormat: 1,
				// InsertTextMode:   1,
				Kind:   2,
				Detail: "This is detail",
				Documentation: MarkupContent{
					Kind:  MarkupKindMarkdown,
					Value: "# Serious Docs\nYeah, it is!",
				},
			},
			{
				Label: "Tomas Train",
				// InsertText:       "Tomas Train",
				// InsertTextFormat: 1,
				// InsertTextMode:   1,
				Kind:   6,
				Detail: "This is detail",
				Documentation: MarkupContent{
					Kind:  MarkupKindMarkdown,
					Value: "# Serious Docs\nYeah, it is!",
				},
			},
		},
	}
	return json.Marshal(resp)
}

func (s *State) ResolveHover(id int, file string, pos Position) ([]byte, error) {
	content := s.GetFile(file)
	if content == "" {
		return nil, fmt.Errorf("ResolveHover cannot found file %s", file)
	}

	// Just return some nonesense
	// and with whole line as range
	lines := strings.Split(content, "\n")
	if pos.Line >= uint(len(lines)) || pos.Character >= uint(len(lines[pos.Line])) {
		return nil, fmt.Errorf("ResolveHover position out of range")
	}

	resp := HoverResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: &Hover{
			Contents: MarkupContent{
				Kind: MarkupKindMarkdown,
				Value: "# Toy Description\n\n" +
					"This is a nice *toy* to play with your Mom!\n" +
					" - TODO 1\n" +
					" - TODO 2\n" +
					" > some quotes\n" +
					"```javascript\n" +
					"// look at it!\n" +
					"someYawascriptCode();\n" +
					"```",
			},
			Range: &Range{
				Start: Position{
					Line:      pos.Line,
					Character: 0,
				},
				End: Position{
					Line:      pos.Line,
					Character: uint(len(lines[pos.Line])),
				},
			},
		},
	}

	return json.Marshal(resp)
}
