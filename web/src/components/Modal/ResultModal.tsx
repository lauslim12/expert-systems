import {
  Alert,
  AlertIcon,
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  VStack,
} from '@chakra-ui/react';
import { memo } from 'react';
import { AiFillCheckCircle } from 'react-icons/ai';

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
 * This modal will render the results of the inference.
 *
 * @param param - ChakraUI's modal props, and the API response
 * @returns React Functional Component
 */
const ResultModal = ({ isOpen, onClose, results }: Props) => (
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
      <ModalHeader>Results</ModalHeader>
      <ModalCloseButton />

      <ModalBody>
        <VStack as="article" align="stretch">
          <Alert as="section" status="info" variant="left-accent">
            <AlertIcon />
            Your results are as follows.
          </Alert>

          <VStack as="section" align="start" spacing={4}>
            <Text>{JSON.stringify(results)}</Text>
          </VStack>
        </VStack>
      </ModalBody>

      <ModalFooter>
        <Button
          colorScheme="green"
          leftIcon={<AiFillCheckCircle />}
          onClick={onClose}
        >
          I got it!
        </Button>
      </ModalFooter>
    </ModalContent>
  </Modal>
);

export default memo(ResultModal);
