import {
  HStack,
  Link,
  StackDivider,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import { memo } from 'react';

/**
 * Will accept a function to open up modals again.
 */
type Props = {
  setOpenDisclaimer: () => void;
  setOpenAbout: () => void;
};

/**
 * Footer component.
 *
 * @param param - Setter functions to open up modals
 * @returns React Functional Component
 */
const Footer = ({ setOpenDisclaimer, setOpenAbout }: Props) => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <HStack
      as="footer"
      fontSize="xs"
      divider={<StackDivider orientation="vertical" />}
      justify="center"
      p={2}
    >
      <Text as="button" onClick={toggleColorMode}>
        {colorMode === 'light' ? 'Darken' : 'Lighten'}
      </Text>

      <Link href="https://github.com/lauslim12/expert-systems" isExternal>
        GitHub
      </Link>

      <Text as="button" onClick={setOpenAbout}>
        About
      </Text>

      <Text as="button" onClick={setOpenDisclaimer}>
        Terms
      </Text>
    </HStack>
  );
};

export default memo(Footer);
