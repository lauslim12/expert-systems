import { memo, useEffect, useState } from 'react';

import Response from '../types/Response';

const App = () => {
  const [response, setResponse] = useState({} as Response);

  useEffect(() => {
    fetch('/api/v1')
      .then((res) => res.json())
      .then((data: Response) => setResponse(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <main>
      <p>Hello, Expert Systems!</p>
      {response && <p>JSON API response: {JSON.stringify(response)}</p>}
    </main>
  );
};

export default memo(App);
