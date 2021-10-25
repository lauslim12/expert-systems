package inference

// Disease is the representation of the diseases data in this Expert System.
type Disease struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Treatment   string   `json:"treatment"`
	Prevention  string   `json:"prevention"`
	Source      []string `json:"source"`
	Symptoms    []string `json:"symptoms"`
}

// Inferred is the object that will be returned after all of the calculations.
type Inferred struct {
	Verdict     bool      `json:"verdict"`
	Probability float64   `json:"probability"`
	Message     string    `json:"message"`
	Disease     []Disease `json:"disease"`
	Symptoms    []Symptom `json:"symptoms"`
}

// Input is used as a representative of a user's input.
type Input struct {
	Symptoms []struct {
		SymptomID string  `json:"symptomId"`
		Weight    float64 `json:"weight"`
	} `json:"symptoms"`
}

// Symptom is an object that represents the symptoms data in this library.
type Symptom struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}
