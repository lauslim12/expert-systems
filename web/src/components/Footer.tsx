import {
  HStack,
  Link,
  StackDivider,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import { memo } from 'react';

type Props = {
  setOpenDisclaimer: () => void;
};

const Footer = ({ setOpenDisclaimer }: Props) => {
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

      <Text as="button" onClick={setOpenDisclaimer}>
        Terms
      </Text>
    </HStack>
  );
};

export default memo(Footer);
