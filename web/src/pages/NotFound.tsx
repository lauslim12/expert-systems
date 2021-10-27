import { Button, Heading, Text, VStack } from '@chakra-ui/react';
import { memo } from 'react';
import { useTranslation } from 'react-i18next';
import { AiFillHome } from 'react-icons/ai';
import { Link as ReactLink } from 'react-router-dom';

/**
 * Renders if a user tries a non-existent page.
 *
 * @returns React Functional Component
 */
const NotFound = () => {
  const { t } = useTranslation();

  return (
    <VStack p={2} h="100vh" w="100vw" align="center" justify="center">
      <Heading textAlign="center">{t('notFound.title')}</Heading>

      <Text textAlign="center">{t('notFound.message')}</Text>

      <Button
        as={ReactLink}
        to="/"
        colorScheme="green"
        leftIcon={<AiFillHome />}
        variant="outline"
      >
        {t('notFound.back')}
      </Button>
    </VStack>
  );
};

export default memo(NotFound);
