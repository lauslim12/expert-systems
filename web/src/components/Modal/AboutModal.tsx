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

/**
 * Accepts ChakraUI's modal props.
 */
type Props = {
  isOpen: boolean;
  onClose: () => void;
};

/**
 * Describes the application.
 *
 * @param param - ChakraUI's modal props
 * @returns React Functional Component
 */
const AboutModal = ({ isOpen, onClose }: Props) => (
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
      <ModalHeader>About</ModalHeader>
      <ModalCloseButton />

      <ModalBody>
        <VStack as="article" align="stretch">
          <Alert as="section" status="success" variant="left-accent">
            <AlertIcon />
            About the app
          </Alert>

          <VStack as="section" align="start" spacing={4}>
            <Text>Pre-memory.</Text>
          </VStack>
        </VStack>
      </ModalBody>

      <ModalFooter>
        <Button colorScheme="green" onClick={onClose}>
          Okay!
        </Button>
      </ModalFooter>
    </ModalContent>
  </Modal>
);

export default memo(AboutModal);
