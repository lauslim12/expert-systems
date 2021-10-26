/**
 * Represents the symptom ID and its user confidence.
 * Check out the Go library in order to know about the 'symptomId' attribute.
 */
type SymptomAndWeight = {
  symptomId: string;
  weight: number;
};

/**
 * Type to represent the request body that will be sent to the API.
 */
type Request = {
  symptoms: SymptomAndWeight[];
};

export default Request;
