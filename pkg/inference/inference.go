package inference

// Disease is the representation of the diseases data in this Expert System.
type Disease struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Treatment   string          `json:"treatment"`
	Prevention  string          `json:"prevention"`
	Source      []SourceAndLink `json:"source"`
	Symptoms    []Symptom       `json:"symptoms"`
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
	DiseaseID string             `json:"diseaseId"`
	Locale    string             `json:"locale"`
	Symptoms  []SymptomAndWeight `json:"symptoms"`
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

// NewInput creates a new input instance that's already validated.
// If default ID is not inside, then we are going to assume Tuberculosis.
// If locale is not inside, then we are going to assume English.
func NewInput(input *Input) *Input {
	if input.DiseaseID == "" {
		input.DiseaseID = "D01"
	}

	if input.Locale == "" {
		input.Locale = "en"
	}

	if input.Symptoms == nil {
		input.Symptoms = []SymptomAndWeight{}
	}

	newInput := &Input{
		DiseaseID: input.DiseaseID,
		Locale:    input.Locale,
		Symptoms:  input.Symptoms,
	}

	return newInput
}

// GetDiseaseByID is used to fetch a disease data by its ID.
func GetDiseaseByID(ID string, diseases []Disease) *Disease {
	for _, disease := range diseases {
		if disease.ID == ID {
			return &disease
		}
	}

	return nil
}

// ForwardChaining is used to perform inference by using the Forward Chaining Algorithm.
// A weight of zero means that the user is NOT sick.
// This forward chaining will be true only and only if the user has experienced 7 symptoms.
// This is because of our knowledge base - the average of symptoms had by each patient.
func ForwardChaining(input *Input, disease *Disease) bool {
	numberOfPositives := 0

	for _, userSymptom := range input.Symptoms {
		for _, diseaseSymptom := range disease.Symptoms {
			if userSymptom.Weight > 0.0 && userSymptom.SymptomID == diseaseSymptom.ID {
				numberOfPositives += 1
			}
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

	// If invalid input, return zero probability.
	if len(certainties) == 0 {
		return probability
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
	// Initial preparation: if no locale, set it to be English as default.
	processedInput := NewInput(input)

	// 0. Fetch all data from the knowledge base.
	diseases := getDiseases(processedInput.Locale)

	// 1. Get disease from the identifier in the input request body.
	disease := GetDiseaseByID(processedInput.DiseaseID, diseases)

	// 2. Infer if the user is diagnosed with TB or not with Forward Chaining.
	isSick := ForwardChaining(processedInput, disease)

	// 3. Calculate certainty factor based on the symptoms.
	certaintyProbability := CertaintyFactor(processedInput, disease.Symptoms)

	// 4. Create result structure.
	inferred := &Inferred{
		Verdict:     isSick,
		Probability: certaintyProbability,
		Disease:     *disease,
	}

	// 5. Return result.
	return inferred
}
