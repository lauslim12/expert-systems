package inference

// All diseases that is in this expert system.
func getDiseases() []Disease {
	diseases := []Disease{
		{
			ID:          "D01",
			Name:        "Tuberculosis",
			Description: "Tuberculosis (TB) is a bacterial infection spread through inhaling tiny droplets from the coughs or sneezes of an infected person. It mainly affects the lungs, but it can affect any part of the body, including the tummy (abdomen), glands, bones and nervous system. TB is a potentially serious condition, but it can be cured if it is treated with the right antibiotics. TB is spread from person to person through the air. When people with lung TB cough, sneeze or spit, they propel the TB germs into the air. A person needs to inhale only a few of these germs to become infected. About one-quarter of the world's population has a TB infection, which means people have been infected by TB bacteria but are not (yet) ill with the disease and cannot transmit it.",
			Treatment:   "TB is a treatable and curable disease. Active, drug-susceptible TB disease is treated with a standard 6-month course of 4 antimicrobial drugs that are provided with information and support to the patient by a health worker or trained volunteer. Without such support, treatment adherence is more difficult. Since 2000, an estimated 66 million lives were saved through TB diagnosis and treatment.",
			Prevention:  "You can perform several precautions by providing good ventilation, natural light, and keeping everything clean. It is recommended to take vaccinations of TB as well. Vaccinations help you to keep your immune system in prime condition, thus allowing you to resist the virus for a longer time (some up to 15 years). Depending on some external factors, some people are more at risk for being exposed into developing the TB disease.",
			Source: []SourceAndLink{
				{
					Name: "NHS",
					Link: "https://www.nhs.uk/conditions/tuberculosis-tb/",
				},
				{
					Name: "WHO",
					Link: "https://www.who.int/news-room/fact-sheets/detail/tuberculosis",
				},
				{
					Name: "TBAlert",
					Link: "https://www.tbalert.org/about-tb/what-is-tb/prevention/",
				},
				{
					Name: "CDC Government",
					Link: "https://www.cdc.gov/tb/topic/basics/tbprevention.htm",
				},
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
