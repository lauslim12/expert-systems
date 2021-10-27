import {
  Alert,
  AlertIcon,
  Button,
  Checkbox,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  VStack,
} from '@chakra-ui/react';
import { memo, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { AiFillCheckCircle } from 'react-icons/ai';

/**
 * Accepts ChakraUI's basic props: 'isOpen' and 'onClose'.
 */
type Props = {
  isOpen: boolean;
  onClose: () => void;
};

/**
 * This modal is used to render the disclaimer.
 *
 * @param param - ChakraUI's modal props
 * @returns React Functional Component
 */
const DisclaimerModal = ({ isOpen, onClose }: Props) => {
  const [isAgreeWithTerms, setIsAgreeWithTerms] = useState(false);
  const [isResponsible, setIsResponsible] = useState(false);
  const focusRef = useRef(null);
  const { t } = useTranslation();

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      initialFocusRef={focusRef}
      size="5xl"
      closeOnEsc={false}
      closeOnOverlayClick={false}
    >
      <ModalOverlay />

      <ModalContent ref={focusRef}>
        <ModalHeader>{t('disclaimer.title')}</ModalHeader>

        <ModalBody>
          <VStack as="article" align="stretch">
            <Alert as="section" status="info" variant="left-accent">
              <AlertIcon />
              {t('disclaimer.alert')}
            </Alert>

            <VStack as="section" align="start" spacing={4}>
              <Text>{t('disclaimer.initial')}</Text>
              <Text>{t('disclaimer.privacy')}</Text>
              <Text as="strong">{t('disclaimer.warning')}</Text>
              <Text as="strong">{t('disclaimer.license')}</Text>
              <Text>{t('disclaimer.agreement')}</Text>
            </VStack>

            <VStack as="section" align="start">
              <Checkbox
                colorScheme="green"
                isChecked={isAgreeWithTerms}
                onChange={({ target: { checked } }) =>
                  setIsAgreeWithTerms(checked)
                }
              >
                {t('disclaimer.understand')}
              </Checkbox>

              <Checkbox
                colorScheme="green"
                isChecked={isResponsible}
                onChange={({ target: { checked } }) =>
                  setIsResponsible(checked)
                }
              >
                {t('disclaimer.responsibility')}
              </Checkbox>
            </VStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <Button
            colorScheme="green"
            leftIcon={<AiFillCheckCircle />}
            isDisabled={!isResponsible || !isAgreeWithTerms}
            onClick={onClose}
          >
            {t('disclaimer.close')}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(DisclaimerModal);
