import {
  Alert,
  AlertIcon,
  Button,
  ButtonGroup,
  chakra,
  Link,
  List,
  ListIcon,
  ListItem,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  Textarea,
  VStack,
} from '@chakra-ui/react';
import { memo, useState } from 'react';
import {
  AiFillInfoCircle,
  AiOutlineClose,
  AiOutlineCode,
} from 'react-icons/ai';

import type Response from '../../types/Response';

/**
 * Accepts ChakraUI's modal props, plus the response from the API.
 */
type Props = {
  isOpen: boolean;
  onClose: () => void;
  results: Response;
};

/**
 * Renders the suggestion of what the user should do based on their probability rate.
 *
 * @param probability - The probability of having TB
 * @returns JSX Element
 */
const renderSuggestion = (probability: number) => {
  if (probability <= 30) {
    return (
      <Text>
        Judging from your probability rate, you are not at risk of getting TB.
        However, it's always nice to go to have a medical check up with a doctor
        or any authorized medical personnel. Stay safe and healthy!
      </Text>
    );
  }

  if (probability > 30 && probability < 70) {
    return (
      <Text>
        Judging from your probability rate, you are at a mild risk of getting
        TB. It is recommended for you to go and have a medical check up with a
        doctor as soon as possible. Hopefully, you are not at risk. Stay safe
        and healthy!
      </Text>
    );
  }

  return (
    <Text>
      From your probability rate, you have a high chance of having TB right now.
      It is absolutely recommended for you to go to a hospital and have a
      medical check up with a doctor right now. Please stay safe and healthy!
    </Text>
  );
};

/**
 * This modal will render the results of the inference.
 *
 * @param param - ChakraUI's modal props, and the API response
 * @returns React Functional Component
 */
const ResultModal = ({ isOpen, onClose, results }: Props) => {
  const [showRawData, setShowRawData] = useState(false);

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      size="5xl"
      motionPreset="slideInBottom"
      closeOnEsc={false}
      closeOnOverlayClick={false}
    >
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>Inference Result</ModalHeader>
        <ModalCloseButton />

        <ModalBody>
          <VStack as="article" align="stretch">
            <Alert as="section" status="success" variant="solid">
              <AlertIcon />
              Your results are as follows.
            </Alert>

            <VStack as="section" align="start" spacing={4}>
              {results.data && (
                <>
                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Verdict
                    </Text>

                    <Text
                      color={
                        results.data.verdict === true
                          ? 'orange.400'
                          : 'twitter.400'
                      }
                      fontWeight="bold"
                    >
                      {results.data.verdict === true
                        ? 'From our analysis and our system, you may have been infected by TB.'
                        : 'From our analysis and our system, you have not been infected by TB.'}
                    </Text>

                    <Text>
                      The calculation was done by using the Forward Chaining
                      Algorithm and provides above result.
                    </Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Probability
                    </Text>

                    <Text fontWeight="bold" color="pink.400">
                      The human-readable probability of you being infected with
                      TB stands at {(results.data.probability * 100).toFixed(2)}
                      %.
                    </Text>

                    <Text>
                      The calculation was done by using the Certainty Factor
                      Algorithm and provides above result.
                    </Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Suggestion
                    </Text>

                    {renderSuggestion(
                      parseFloat((results.data.probability * 100).toFixed(2))
                    )}
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Information
                    </Text>

                    <Text>{results.data.disease.description}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Prevention
                    </Text>

                    <Text>{results.data.disease.prevention}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Treatment
                    </Text>

                    <Text>{results.data.disease.treatment}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      Sources
                    </Text>

                    <Text>
                      The sources of the data regarding of the disease are taken
                      from the following sources:
                    </Text>

                    <List>
                      {results.data.disease.source.map((source) => (
                        <ListItem key={source.name}>
                          <ListIcon as={AiFillInfoCircle} color="green.500" />

                          <Link
                            color="twitter.500"
                            href={source.link}
                            isExternal
                          >
                            {source.name}
                          </Link>
                        </ListItem>
                      ))}
                    </List>
                  </VStack>

                  {showRawData && (
                    <VStack align="start" w="full">
                      <Text fontSize="lg" fontWeight="bold">
                        JSON
                      </Text>

                      <Text>Raw data for those interested.</Text>

                      <chakra.code w="full">
                        <Textarea
                          value={JSON.stringify(results, null, 2)}
                          w="full"
                          h="50vh"
                          readOnly
                        />
                      </chakra.code>
                    </VStack>
                  )}
                </>
              )}
            </VStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <ButtonGroup>
            <Button
              colorScheme="twitter"
              leftIcon={<AiOutlineCode />}
              onClick={() => setShowRawData(!showRawData)}
            >
              {showRawData === true ? 'Hide raw data' : 'Show raw data'}
            </Button>

            <Button
              colorScheme="red"
              leftIcon={<AiOutlineClose />}
              onClick={onClose}
            >
              Close
            </Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(ResultModal);
