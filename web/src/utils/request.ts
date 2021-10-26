/**
 * Handles errors from Fetch API.
 *
 * @param response - Fetch response object.
 * @returns The result of the request, or an error if applicable
 */
const handleErrors = async (response: Response) => {
  if (!response.ok) {
    const error = await response.json();

    throw Error(error.message);
  }

  return await response.json();
};

/**
 * Customized Fetch API to have better error handling.
 *
 * @param url - URL of the request
 * @param requestBody - Body of the request
 * @param requestMethod - Method of the request
 * @returns Results of the request, or an error if applicable
 */
const request = async (
  url: string,
  requestBody: any,
  requestMethod: 'GET' | 'POST'
) => {
  const options: RequestInit = {
    method: requestMethod,
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestBody),
  };

  return await fetch(url, options).then(handleErrors);
};

export default request;
