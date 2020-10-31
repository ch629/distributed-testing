package ast

import (
	"distributed-testing/scenario/scanner"
)

type (
	Feature struct {
		Name      scanner.StepText
		Scenarios []*Scenario
	}

	Scenario struct {
		Name       scanner.StepText
		References scanner.StepText
		Steps      []*scanner.Step
	}
)

func BuildAst(steps []*scanner.Step) []*Feature {
	features := make([]*Feature, 1)
	var currentFeatureName scanner.StepText = ""
	//var currentScenarioName parser.StepText = ""
	scenarios := make([]*Scenario, 1)
	//steps := make([]*parser.Step, 1)
	//var references parser.StepText = ""

	for _, step := range steps {
		if step.Token == scanner.FEATURE {
			if len(currentFeatureName) > 0 {
				features = append(features, &Feature{
					Name:      currentFeatureName,
					Scenarios: scenarios,
				})
			}

			currentFeatureName = step.Text
			continue
		}

		if step.Token == scanner.SCENARIO {
			continue
		}

		if step.Token == scanner.REFERENCES {
			//references = step.Text
			continue
		}

		steps = append(steps, step)
	}
	return features
}
