package inference

import (
	"reflect"
	"testing"
)

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
				DiseaseID: "D01",
				Locale:    "en",
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
			name: "test_advanced_input_locale_en",
			input: Input{
				DiseaseID: "D01",
				Locale:    "en",
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
		{
			name: "test_advanced_input_locale_en",
			input: Input{
				DiseaseID: "D01",
				Locale:    "en",
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
		{
			name: "test_advanced_input_locale_en",
			input: Input{
				DiseaseID: "D01",
				Locale:    "id",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.25,
					},
					{
						SymptomID: "S2",
						Weight:    0,
					},
					{
						SymptomID: "S3",
						Weight:    0.25,
					},
					{
						SymptomID: "S4",
						Weight:    0,
					},
					{
						SymptomID: "S5",
						Weight:    0,
					},
					{
						SymptomID: "S6",
						Weight:    0,
					},
					{
						SymptomID: "S7",
						Weight:    0,
					},
					{
						SymptomID: "S8",
						Weight:    0,
					},
					{
						SymptomID: "S9",
						Weight:    0,
					},
					{
						SymptomID: "S10",
						Weight:    0,
					},
					{
						SymptomID: "S11",
						Weight:    0,
					},
					{
						SymptomID: "S12",
						Weight:    0.5,
					},
					{
						SymptomID: "S13",
						Weight:    0.2,
					},
				},
			},
			expectedCertaintyFactor: 0.47902158346120005,
			expectedVerdict:         false,
		},
		{
			name:                    "test_invalid_input",
			input:                   Input{},
			expectedCertaintyFactor: 0.0,
			expectedVerdict:         false,
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

func TestNewInput(t *testing.T) {
	tests := []struct {
		name           string
		input          *Input
		expectedOutput *Input
	}{
		{
			name: "test_valid_input",
			input: &Input{
				DiseaseID: "D01",
				Locale:    "id",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S01",
						Weight:    0.25,
					},
				},
			},
			expectedOutput: &Input{
				DiseaseID: "D01",
				Locale:    "id",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S01",
						Weight:    0.25,
					},
				},
			},
		},
		{
			name:  "test_invalid_input",
			input: &Input{},
			expectedOutput: &Input{
				DiseaseID: "D01",
				Locale:    "en",
				Symptoms:  []SymptomAndWeight{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := NewInput(tt.input)

			if !reflect.DeepEqual(&tt.expectedOutput, &output) {
				t.Errorf("Expected and actual structs are not equal! Expected: %v. Got: %v", tt.expectedOutput, output)
			}
		})
	}
}

func TestGetDiseaseByID(t *testing.T) {
	diseases := getDiseases("en")

	tests := []struct {
		name           string
		diseaseID      string
		expectedOutput *Disease
	}{
		{
			name:           "test_valid_disease_id",
			diseaseID:      "D01",
			expectedOutput: &diseases[0],
		},
		{
			name:           "test_invalid_disease_id",
			diseaseID:      "404",
			expectedOutput: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GetDiseaseByID(tt.diseaseID, diseases)

			if !reflect.DeepEqual(&tt.expectedOutput, &output) {
				t.Errorf("Expected and actual structs are not equal! Expected: %v. Got: %v", tt.expectedOutput, output)
			}
		})
	}
}

func TestForwardChaining(t *testing.T) {
	disease := getDiseases("en")[0]

	tests := []struct {
		name           string
		input          Input
		expectedOutput bool
	}{
		{
			name: "test_forward_chaining_false",
			input: *NewInput(&Input{
				DiseaseID: "D01",
				Locale:    "en",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.25,
					},
					{
						SymptomID: "S2",
						Weight:    0,
					},
					{
						SymptomID: "S3",
						Weight:    0.25,
					},
					{
						SymptomID: "S4",
						Weight:    0,
					},
					{
						SymptomID: "S5",
						Weight:    0,
					},
					{
						SymptomID: "S6",
						Weight:    0,
					},
					{
						SymptomID: "S7",
						Weight:    0,
					},
					{
						SymptomID: "S8",
						Weight:    0,
					},
					{
						SymptomID: "S9",
						Weight:    0,
					},
					{
						SymptomID: "S10",
						Weight:    0,
					},
					{
						SymptomID: "S11",
						Weight:    0,
					},
					{
						SymptomID: "S12",
						Weight:    0.5,
					},
					{
						SymptomID: "S13",
						Weight:    0.2,
					},
				},
			}),
			expectedOutput: false,
		},
		{
			name: "test_forward_chaining_true",
			input: *NewInput(&Input{
				DiseaseID: "D01",
				Locale:    "en",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.25,
					},
					{
						SymptomID: "S2",
						Weight:    0.25,
					},
					{
						SymptomID: "S3",
						Weight:    0.25,
					},
					{
						SymptomID: "S4",
						Weight:    0.25,
					},
					{
						SymptomID: "S5",
						Weight:    0.25,
					},
					{
						SymptomID: "S6",
						Weight:    0.25,
					},
					{
						SymptomID: "S7",
						Weight:    0.25,
					},
					{
						SymptomID: "S8",
						Weight:    0,
					},
					{
						SymptomID: "S9",
						Weight:    0,
					},
					{
						SymptomID: "S10",
						Weight:    0,
					},
					{
						SymptomID: "S11",
						Weight:    0,
					},
					{
						SymptomID: "S12",
						Weight:    0.5,
					},
					{
						SymptomID: "S13",
						Weight:    0.2,
					},
				},
			}),
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ForwardChaining(&tt.input, &disease)

			if tt.expectedOutput != output {
				t.Errorf("Expected and actual verdict values are different! Expected: %v. Got: %v", tt.expectedOutput, output)
			}
		})
	}
}

func TestCertaintyFactor(t *testing.T) {
	symptoms := getDiseases("en")[0].Symptoms

	tests := []struct {
		name           string
		input          Input
		expectedOutput float64
	}{
		{
			name: "test_valid_certainty_factor",
			input: Input{
				DiseaseID: "D01",
				Locale:    "en",
				Symptoms: []SymptomAndWeight{
					{
						SymptomID: "S1",
						Weight:    0.25,
					},
					{
						SymptomID: "S2",
						Weight:    0.25,
					},
					{
						SymptomID: "S3",
						Weight:    0.25,
					},
					{
						SymptomID: "S4",
						Weight:    0.25,
					},
					{
						SymptomID: "S5",
						Weight:    0.25,
					},
					{
						SymptomID: "S6",
						Weight:    0.25,
					},
					{
						SymptomID: "S7",
						Weight:    0.25,
					},
					{
						SymptomID: "S8",
						Weight:    0,
					},
					{
						SymptomID: "S9",
						Weight:    0,
					},
					{
						SymptomID: "S10",
						Weight:    0,
					},
					{
						SymptomID: "S11",
						Weight:    0,
					},
					{
						SymptomID: "S12",
						Weight:    0.5,
					},
					{
						SymptomID: "S13",
						Weight:    0.2,
					},
				},
			},
			expectedOutput: 0.7313435264022431,
		},
		{
			name:           "test_invalid_certainty_factor",
			input:          Input{},
			expectedOutput: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := CertaintyFactor(&tt.input, symptoms)

			if tt.expectedOutput != output {
				t.Errorf("Expected and actual certainty factor values are different! Expected: %v. Got: %v", tt.expectedOutput, output)
			}
		})
	}
}
