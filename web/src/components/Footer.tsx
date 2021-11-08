import { HStack, Text } from '@chakra-ui/react';
import { memo } from 'react';
import { useTranslation } from 'react-i18next';

/**
 * Footer component.
 *
 * @param param - Setter functions to open up modals
 * @returns React Functional Component
 */
const Footer = () => {
  const { t } = useTranslation();

  return (
    <HStack as="footer" fontSize="xs" justify="center" p={2}>
      <Text>{t('footer')}</Text>
    </HStack>
  );
};

export default memo(Footer);
