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
import { useTranslation } from 'react-i18next';
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
 * This modal will render the results of the inference.
 *
 * @param param - ChakraUI's modal props, and the API response
 * @returns React Functional Component
 */
const ResultModal = ({ isOpen, onClose, results }: Props) => {
  const [showRawData, setShowRawData] = useState(false);
  const { t } = useTranslation();

  /**
   * Renders the suggestion of what the user should do based on their probability rate.
   *
   * @param probability - The probability of having TB
   * @returns JSX Element
   */
  const renderSuggestion = (probability: number) => {
    if (probability <= 30) {
      return <Text>{t('inference.suggestionOkay')}</Text>;
    }

    if (probability > 30 && probability < 70) {
      return <Text>{t('inference.suggestionVisit')}</Text>;
    }

    return <Text>{t('inference.suggestionDangerous')}</Text>;
  };

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
        <ModalHeader>{t('inference.title')}</ModalHeader>
        <ModalCloseButton />

        <ModalBody>
          <VStack as="article" align="stretch">
            <Alert as="section" status="success" variant="solid">
              <AlertIcon />
              {t('inference.alert')}
            </Alert>

            <VStack as="section" align="start" spacing={4}>
              {results.data && (
                <>
                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.verdict')}
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
                        ? t('inference.verdictTB')
                        : t('inference.verdictNoTB')}
                    </Text>

                    <Text>{t('inference.verdictCalculation')}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.probability')}
                    </Text>

                    <Text fontWeight="bold" color="pink.400">
                      {t('inference.probabilityResult', {
                        probability: (results.data.probability * 100).toFixed(
                          2
                        ),
                      })}
                    </Text>

                    <Text>{t('inference.probabilityCalculation')}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.suggestion')}
                    </Text>

                    {renderSuggestion(
                      parseFloat((results.data.probability * 100).toFixed(2))
                    )}
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.information')}
                    </Text>

                    <Text>{results.data.disease.description}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.prevention')}
                    </Text>

                    <Text>{results.data.disease.prevention}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.treatment')}
                    </Text>

                    <Text>{results.data.disease.treatment}</Text>
                  </VStack>

                  <VStack align="start">
                    <Text fontSize="lg" fontWeight="bold">
                      {t('inference.sources')}
                    </Text>

                    <Text>{t('inference.sourcesBeginning')}</Text>

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
                        {t('inference.json')}
                      </Text>

                      <Text>{t('inference.jsonBeginning')}</Text>

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
              {showRawData === true
                ? t('inference.hideCode')
                : t('inference.showCode')}
            </Button>

            <Button
              colorScheme="red"
              leftIcon={<AiOutlineClose />}
              onClick={onClose}
            >
              {t('inference.close')}
            </Button>
          </ButtonGroup>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(ResultModal);
