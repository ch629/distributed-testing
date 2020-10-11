package parser

import "strings"

var steps []Step
var token *Token

var sb strings.Builder

func HandleByte(b byte) {
	if b == '\n' {
		steps = append(steps, makeStep())
		token = nil
		sb.Reset()
		return
	}

	if b == ':' {
		token = mapToken()
		sb.Reset()
		return
	}

	sb.WriteByte(b)
}

var tokenMap = map[string]Token{
	"feature":    FEATURE,
	"scenario":   SCENARIO,
	"background": BACKGROUND,
	"given":      GIVEN,
	"when":       WHEN,
	"then":       THEN,
	"and":        AND,
}

func mapToken() *Token {
	str := sb.String()
	sb.Reset()

	if value, ok := tokenMap[strings.ToLower(str)]; ok {
		return &value
	}
	return nil
}

func makeStep() Step {
	step := Step{Token: *token, Text: StepText(strings.TrimSpace(sb.String()))}
	token = nil
	sb.Reset()
	return step
}

func GetSteps() []Step {
	returnSteps := steps
	steps = []Step{}
	if sb.Len() > 0 {
		returnSteps = append(returnSteps, makeStep())
	}
	return returnSteps
}
