package scanner

import (
	"bufio"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

var tokenMap = map[string]Token{
	"feature":    FEATURE,
	"scenario":   SCENARIO,
	"background": BACKGROUND,
	"given":      GIVEN,
	"when":       WHEN,
	"then":       THEN,
	"and":        AND,
	"references": REFERENCES,
	"with":       WITH,
}

var eof = rune(0)

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) Scan() *Step {
	token := s.scanToken()

	if token == nil {
		return nil
	}

	text := s.scanText()

	return &Step{
		Token: *token,
		Text:  text,
	}
}

// TODO: Error handling
func (s *Scanner) scanToken() *Token {
	var sb strings.Builder
	for {
		if c := s.read(); c != ':' && c != eof {
			sb.WriteRune(c)
			continue
		}
		break
	}

	if token, ok := tokenMap[strings.ToLower(strings.TrimSpace(sb.String()))]; ok {
		return &token
	}
	return nil
}

func (s *Scanner) scanText() StepText {
	var sb strings.Builder
	for {
		if c := s.read(); c != '\n' && c != eof {
			sb.WriteRune(c)
			continue
		}
		break
	}
	return StepText(strings.TrimSpace(sb.String()))
}
