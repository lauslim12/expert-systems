import { Flex, VStack } from '@chakra-ui/react';
import { lazy, memo, Suspense, useState } from 'react';

import Footer from '../components/Footer';
import DisclaimerModal from '../components/Modal/DisclaimerModal';
import Tuberculosis from '../components/Tuberculosis';

/**
 * Lazy-load 'About' modal, as we have no need for it at render time.
 */
const AboutModal = lazy(() => import('../components/Modal/AboutModal'));

/**
 * Homepage of the application.
 *
 * @returns React Functional Component
 */
const Home = () => {
  const [openAbout, setOpenAbout] = useState(false);
  const [openDisclaimer, setOpenDisclaimer] = useState(true);

  return (
    <>
      <DisclaimerModal
        isOpen={openDisclaimer}
        onClose={() => setOpenDisclaimer(false)}
      />

      <Suspense fallback={null}>
        <AboutModal isOpen={openAbout} onClose={() => setOpenAbout(false)} />
      </Suspense>

      <Flex h="100vh" direction="column" maxW="1200px" mx="auto">
        <VStack as="main" p={3} spacing={5} flex={1} mt={5} mb={5}>
          <VStack as="header" w="full">
            <p>Hello, Expert Systems!</p>
          </VStack>

          <Tuberculosis />
        </VStack>

        <Footer
          setOpenDisclaimer={() => setOpenDisclaimer(true)}
          setOpenAbout={() => setOpenAbout(true)}
        />
      </Flex>
    </>
  );
};

export default memo(Home);
