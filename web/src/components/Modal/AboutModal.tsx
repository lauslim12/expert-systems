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
import { useTranslation } from 'react-i18next';

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
const AboutModal = ({ isOpen, onClose }: Props) => {
  const { t } = useTranslation();

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
        <ModalHeader>{t('about.title')}</ModalHeader>
        <ModalCloseButton />

        <ModalBody>
          <VStack as="article" align="stretch">
            <Alert as="section" status="success" variant="left-accent">
              <AlertIcon />
              {t('about.alert')}
            </Alert>

            <VStack as="section" align="start" spacing={4}>
              <Text>{t('about.content')}</Text>
            </VStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <Button colorScheme="green" onClick={onClose}>
            {t('about.close')}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(AboutModal);
