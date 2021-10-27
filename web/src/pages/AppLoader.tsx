import { Spinner, Text, VStack } from '@chakra-ui/react';
import { memo } from 'react';

const AppLoader = () => (
  <VStack w="100vw" h="100vh" align="center" justify="center">
    <Spinner
      size="xl"
      thickness="5px"
      emptyColor="gray.200"
      color="twitter.300"
    />

    <Text textAlign="center">Loading localizations...</Text>
  </VStack>
);

export default memo(AppLoader);