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
        <ModalHeader>Disclaimer & Terms</ModalHeader>

        <ModalBody>
          <VStack as="article" align="stretch">
            <Alert as="section" status="info" variant="left-accent">
              <AlertIcon />
              Please read the terms and conditions before using this
              application!
            </Alert>

            <VStack as="section" align="start" spacing={4}>
              <Text>
                Even though this application is made based on available data and
                the opinion of the pertaining experts, this application is still
                not a fully correct way to diagnose yourself with. You can use
                this application as a first-aid kind of situation if you need to
                get to know about your condition as fast as possible. This
                application is made for research and educational purposes only.
              </Text>

              <Text>
                In this website, we use cookies in order to store your color
                mode preferences. We do not use your data for anything. As for
                licenses, all of the usages are cited properly, either in this
                app or in the repository.
              </Text>

              <Text as="strong">
                This application IS NOT to be construed as a complete medical
                advice. Consult a medical expert in your area if you need to
                diagnose yourself in a more detailed way that is more catered to
                your needs. Use this tool with responsibility. If you need a
                specialized medical advice for your condition, DO NOT USE THIS
                APPLICATION.
              </Text>

              <Text as="strong">
                This software is provided "as is", without warranty of any kind,
                express or implied, including but not limited to the warranties
                of merchantability, fitness for a particular purpose and
                noninfringement. In no event shall the authors or copyright
                holders be liable for any claim, damages or other liability,
                whether in an action of contract, tort or otherwise, arising
                from, out of or in connection with the software or the use or
                other dealings in the software.
              </Text>

              <Text>
                Please check below boxes to continue with this application.
              </Text>
            </VStack>

            <VStack as="section" align="start">
              <Checkbox
                colorScheme="green"
                isChecked={isAgreeWithTerms}
                onChange={({ target: { checked } }) =>
                  setIsAgreeWithTerms(checked)
                }
              >
                I understand the terms written above.
              </Checkbox>

              <Checkbox
                colorScheme="green"
                isChecked={isResponsible}
                onChange={({ target: { checked } }) =>
                  setIsResponsible(checked)
                }
              >
                All of the usage is my own responsibility.
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
            I understand.
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default memo(DisclaimerModal);
