/**
 * Enum of weights, can be used at state to prevent rendering the default value.
 */
export type StateCertaintyWeight = 0 | 0.25 | 0.75 | 1 | null;

/**
 * Enum of weights used to define the 'certainty' of a symptom according to the user.
 */
type UserCertaintyWeight = 0 | 0.25 | 0.75 | 1;

export default UserCertaintyWeight;
