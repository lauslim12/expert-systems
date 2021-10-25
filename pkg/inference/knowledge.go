package inference

// All diseases that is in this expert system.
func getDiseases() []Disease {
	diseases := []Disease{
		{
			ID:          "D01",
			Name:        "Tuberculosis",
			Description: "Tuberculosis (TB) is a bacterial infection spread through inhaling tiny droplets from the coughs or sneezes of an infected person. It mainly affects the lungs, but it can affect any part of the body, including the tummy (abdomen), glands, bones and nervous system. TB is a potentially serious condition, but it can be cured if it's treated with the right antibiotics.",
			Treatment:   "With treatment, TB can almost always be cured. A course of antibiotics will usually need to be taken for 6 months.",
			Prevention:  "You can perform several precautions by providing good ventilation, natural light, and keeping everything clean.",
			Source: []string{
				"https://www.nhs.uk/conditions/tuberculosis-tb/",
				"https://www.tbalert.org/about-tb/what-is-tb/prevention/",
			},
			Symptoms: []string{
				"S1",
				"S2",
				"S3",
				"S4",
				"S5",
				"S6",
				"S7",
				"S8",
				"S9",
				"S10",
				"S11",
				"S12",
				"S13",
			},
		},
	}

	return diseases
}

// All symptoms that is here in the expert system.
// Dataset is processed further from: https://www.kaggle.com/victorcaelina/tuberculosis-symptoms.
func getSymptoms() []Symptom {
	symptoms := []Symptom{
		{
			ID:     "S1",
			Name:   "Fever for two weeks or more",
			Weight: 0.513,
		},
		{
			ID:     "S2",
			Name:   "Coughing blood",
			Weight: 0.475,
		},
		{
			ID:     "S3",
			Name:   "Sputum mixed with blood",
			Weight: 0.519,
		},
		{
			ID:     "S4",
			Name:   "Night sweats",
			Weight: 0.514,
		},
		{
			ID:     "S5",
			Name:   "Chest pain",
			Weight: 0.494,
		},
		{
			ID:     "S6",
			Name:   "Back pain in certain parts",
			Weight: 0.511,
		},
		{
			ID:     "S7",
			Name:   "Shortness of breath",
			Weight: 0.487,
		},
		{
			ID:     "S8",
			Name:   "Weight loss",
			Weight: 0.521,
		},
		{
			ID:     "S9",
			Name:   "Body feels tired",
			Weight: 0.496,
		},
		{
			ID:     "S10",
			Name:   "Lumps that appear around the armpits and neck",
			Weight: 0.484,
		},
		{
			ID:     "S11",
			Name:   "Cough and phlegm continuously for two weeks to four weeks",
			Weight: 0.493,
		},
		{
			ID:     "S12",
			Name:   "Swollen lymph nodes",
			Weight: 0.478,
		},
		{
			ID:     "S13",
			Name:   "Loss of apetite",
			Weight: 0.488,
		},
	}

	return symptoms
}
