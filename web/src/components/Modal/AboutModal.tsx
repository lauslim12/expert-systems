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
import { AiFillLike } from 'react-icons/ai';

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
              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.purpose')}
                </Text>

                <Text>{t('about.content')}</Text>
              </VStack>

              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.dataProcessing')}
                </Text>

                <Text>{t('about.dataProcessingExplanation')}</Text>
              </VStack>

              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.algorithmOne')}
                </Text>

                <Text>{t('about.algorithmOneExplanation')}</Text>
              </VStack>

              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.algorithmTwo')}
                </Text>

                <Text>{t('about.algorithmTwoExplanation')}</Text>
              </VStack>

              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.process')}
                </Text>

                <Text>{t('about.inferenceProcess')}</Text>
              </VStack>

              <VStack align="start">
                <Text fontSize="lg" fontWeight="bold">
                  {t('about.references')}
                </Text>

                <Text>{t('about.referencesAndSource')}</Text>
              </VStack>
            </VStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <Button
            leftIcon={<AiFillLike />}
            colorScheme="green"
            onClick={onClose}
          >
            {t('about.close')}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(AboutModal);
