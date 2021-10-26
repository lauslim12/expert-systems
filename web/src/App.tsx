import { chakra, Flex } from '@chakra-ui/react';
import { memo, useEffect, useState } from 'react';

import Footer from './components/Footer';
import DisclaimerModal from './components/Modal/DisclaimerModal';
import Response from './types/Response';

const App = () => {
  const [openDisclaimer, setOpenDisclaimer] = useState(true);
  const [response, setResponse] = useState({} as Response);

  useEffect(() => {
    fetch('/api/v1')
      .then((res) => res.json())
      .then((data: Response) => setResponse(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <DisclaimerModal
        isOpen={openDisclaimer}
        onClose={() => setOpenDisclaimer(false)}
      />

      <Flex h="100vh" direction="column" maxW="1200px" mx="auto">
        <chakra.div as="main" flex={1} mt={5} mb={5}>
          <p>Hello, Expert Systems!</p>
          {response && <p>JSON API response: {JSON.stringify(response)}</p>}
        </chakra.div>

        <Footer setOpenDisclaimer={() => setOpenDisclaimer(true)} />
      </Flex>
    </>
  );
};

export default memo(App);
