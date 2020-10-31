package scanner

import "fmt"

type (
	Token    int
	StepText string

	Scenario struct {
		BackgroundSteps []Step
		Steps           []Step
	}

	Step struct {
		Token Token
		Text  StepText
	}
)

func (t Token) String() string {
	switch t {
	case FEATURE:
		return "FEATURE"
	case SCENARIO:
		return "SCENARIO"
	case BACKGROUND:
		return "BACKGROUND"
	case GIVEN:
		return "GIVEN"
	case WHEN:
		return "WHEN"
	case THEN:
		return "THEN"
	case AND:
		return "AND"
	case WITH:
		return "WITH"
	case REFERENCES:
		return "REFERENCES"
	}
	return ""
}

func (s Step) String() string {
	return fmt.Sprintf("%v: %v", s.Token, s.Text)
}

const (
	FEATURE Token = iota
	SCENARIO
	BACKGROUND
	GIVEN
	WHEN
	THEN
	AND
	WITH
	REFERENCES
)
