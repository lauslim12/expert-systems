import { Button, Heading, Text, VStack } from '@chakra-ui/react';
import { memo } from 'react';
import { AiFillHome } from 'react-icons/ai';
import { Link as ReactLink } from 'react-router-dom';

/**
 * Renders if a user tries a non-existent page.
 *
 * @returns React Functional Component
 */
const NotFound = () => (
  <VStack p={2} h="100vh" w="100vw" align="center" justify="center">
    <Heading textAlign="center">Not Found!</Heading>

    <Text textAlign="center">
      The page you are looking for may or may not have existed in the past, but
      it certainly doesn't anymore.
    </Text>

    <Button
      as={ReactLink}
      to="/"
      colorScheme="green"
      leftIcon={<AiFillHome />}
      variant="outline"
    >
      Back to Home
    </Button>
  </VStack>
);

export default memo(NotFound);
