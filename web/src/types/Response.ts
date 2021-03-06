/**
 * Core data from the API to be read in the front-end.
 */
type Inferred = {
  verdict: boolean;
  probability: number;
  disease: {
    id: string;
    name: string;
    description: string;
    treatment: string;
    prevention: string;
    source: {
      name: string;
      link: string;
    }[];
    symptoms: {
      symptomId: string;
      weight: number;
    }[];
  };
};

/**
 * Data type to represent the response from the API.
 */
type Response = {
  code: number;
  message: string;
  status: 'success' | 'fail';
  data?: Inferred;
};

export default Response;
