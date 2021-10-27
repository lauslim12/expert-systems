import {
  HStack,
  Link,
  StackDivider,
  Text,
  useColorMode,
} from '@chakra-ui/react';
import { memo } from 'react';
import { useTranslation } from 'react-i18next';

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
  const { t, i18n } = useTranslation();

  return (
    <HStack
      as="footer"
      fontSize="xs"
      divider={<StackDivider orientation="vertical" />}
      justify="center"
      p={2}
    >
      <Text as="button" onClick={toggleColorMode}>
        {colorMode === 'light' ? t('general.darken') : t('general.lighten')}
      </Text>

      <Link href="https://github.com/lauslim12/expert-systems" isExternal>
        {t('footer.github')}
      </Link>

      <Text as="button" onClick={setOpenAbout}>
        {t('footer.about')}
      </Text>

      <Text as="button" onClick={setOpenDisclaimer}>
        {t('footer.terms')}
      </Text>

      <Text
        as="button"
        onClick={() => {
          if (i18n.language === 'id') {
            i18n.changeLanguage('en');
            return;
          }

          i18n.changeLanguage('id');
        }}
      >
        Lang
      </Text>
    </HStack>
  );
};

export default memo(Footer);
