package parser

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

const (
	FEATURE Token = iota
	SCENARIO
	BACKGROUND
	GIVEN
	WHEN
	THEN
	AND
)
