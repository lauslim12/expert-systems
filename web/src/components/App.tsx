import { memo, useEffect } from 'react';

const App = () => {
  useEffect(() => {
    fetch('/api/v1')
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div>
      <p>Hello, Expert Systems!</p>
    </div>
  );
};

export default memo(App);
