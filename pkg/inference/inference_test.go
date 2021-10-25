package inference

import "testing"

func TestInfer(t *testing.T) {
	tests := []struct {
		name                    string
		input                   Input
		expectedCertaintyFactor float64
		expectedVerdict         bool
	}{
		{
			name: "test_basic_input",
			input: Input{
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.5,
					},
				},
			},
			expectedCertaintyFactor: 0.2565,
			expectedVerdict:         false,
		},
		{
			name: "test_advanced_input",
			input: Input{
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.5,
					},
					{
						SymptomID: "S2",
						Weight:    0.4,
					},
					{
						SymptomID: "S3",
						Weight:    0.2,
					},
					{
						SymptomID: "S4",
						Weight:    0.6,
					},
					{
						SymptomID: "S5",
						Weight:    0.2,
					},
					{
						SymptomID: "S6",
						Weight:    0.4,
					},
					{
						SymptomID: "S7",
						Weight:    0.8,
					},
					{
						SymptomID: "S8",
						Weight:    0.2,
					},
					{
						SymptomID: "S9",
						Weight:    0.2,
					},
					{
						SymptomID: "S10",
						Weight:    0.4,
					},
					{
						SymptomID: "S11",
						Weight:    0.2,
					},
					{
						SymptomID: "S12",
						Weight:    0.2,
					},
					{
						SymptomID: "S13",
						Weight:    0.8,
					},
				},
			},
			expectedCertaintyFactor: 0.9471713614230385,
			expectedVerdict:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Infer(&tt.input)

			if tt.expectedCertaintyFactor != output.Probability {
				t.Errorf("Expected and actual certainty factor values are different! Expected: %v. Got: %v", tt.expectedCertaintyFactor, output.Probability)
			}

			if tt.expectedVerdict != output.Verdict {
				t.Errorf("Expected and actual verdict values are different! Expected: %v. Got: %v", tt.expectedVerdict, output.Verdict)
			}

		})
	}
}
