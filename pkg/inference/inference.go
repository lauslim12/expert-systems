package inference

// Disease is the representation of the diseases data in this Expert System.
type Disease struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Treatment   string          `json:"treatment"`
	Prevention  string          `json:"prevention"`
	Source      []SourceAndLink `json:"source"`
	Symptoms    []string        `json:"symptoms"`
}

// Inferred is the object that will be returned after all of the calculations.
type Inferred struct {
	Verdict     bool    `json:"verdict"`
	Probability float64 `json:"probability"`
	Disease     Disease `json:"disease"`
}

// SymptomAndWeight is a struct representative of the members of 'symptoms' array in 'Input' struct.
type SymptomAndWeight struct {
	SymptomID string  `json:"symptomId"`
	Weight    float64 `json:"weight"`
}

// Input is used as a representative of a user's input.
type Input struct {
	Symptoms []SymptomAndWeight `json:"symptoms"`
}

// SourceAndLink represents the source name and its link for the information regarding a disease.
type SourceAndLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// Symptom is an object that represents the symptoms data in this library.
type Symptom struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

// ForwardChaining is used to perform inference by using the Forward Chaining Algorithm.
// A weight of zero means that the user is NOT sick.
// This forward chaining will be true only and only if the user has experienced 7 symptoms.
// This is because of our knowledge base - the average of symptoms had by each patient.
func ForwardChaining(input *Input) bool {
	numberOfPositives := 0

	for _, symptom := range input.Symptoms {
		if symptom.Weight > 0.0 {
			numberOfPositives += 1
		}
	}

	return numberOfPositives > 7
}

// CertaintyFactor is used to perform analysis and to find the certainty probability.
// First, match the user symptoms' and the available expert symptom' from the knowledge base.
// Second, calculate the real probability.
func CertaintyFactor(input *Input, symptoms []Symptom) float64 {
	certainties := make([]float64, 0)
	probability := 0.0

	// Match and calculate certainty between the expert and the user.
	for _, userSymptom := range input.Symptoms {
		for _, expertSymptom := range symptoms {
			if userSymptom.SymptomID == expertSymptom.ID {
				certainties = append(certainties, userSymptom.Weight*expertSymptom.Weight)
			}
		}
	}

	// Calculate probability from the certainty array.
	probability = certainties[0]
	for i := 1; i < len(certainties); i += 1 {
		probability = probability + certainties[i]*(1-probability)
	}

	return probability
}

// Infer is used to calculate based on an input to decide whether the user is infected or not.
// We will use Forward Chaining and Certainty Factor algorithms in order to decide that.
// Algorithm: Get knowledge base -> Forward Chaining -> Certainty Factor -> Result.
func Infer(input *Input) *Inferred {
	// 0. Fetch all data from the knowledge base.
	diseases := getDiseases()
	symptoms := getSymptoms()

	// 1. Infer if the user is diagnosed with TB or not with Forward Chaining.
	isSick := ForwardChaining(input)

	// 2. Calculate certainty factor.
	certaintyProbability := CertaintyFactor(input, symptoms)

	// 3. Create result structure.
	inferred := &Inferred{
		Verdict:     isSick,
		Probability: certaintyProbability,
		Disease:     diseases[0],
	}

	// 4. Return result.
	return inferred
}
