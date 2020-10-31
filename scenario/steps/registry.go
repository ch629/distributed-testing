package steps

import "strconv"

type (
	StepArguments struct {
		Args []string
	}

	StepDefinition struct {
		Name string
		Run  func(args *StepArguments) error
	}
)

var definedSteps = map[string]*StepDefinition{}

// TODO: Add errors for index out of bounds, invalid type etc
func (args StepArguments) GetString(index int) (string, error) {
	return args.Args[index], nil
}

func (args StepArguments) GetInt(index int) (int, error) {
	return strconv.Atoi(args.Args[index])
}

func (args StepArguments) GetBool(index int) (bool, error) {
	return strconv.ParseBool(args.Args[index])
}

// TODO: Figure a good way to hash these for a key with arguments
func Register(definitions ...*StepDefinition) {
	for _, definition := range definitions {
		definedSteps[definition.Name] = definition
	}
}

func GetDefinition(name string) (*StepDefinition, bool) {
	def, ok := definedSteps[name]
	return def, ok
}
