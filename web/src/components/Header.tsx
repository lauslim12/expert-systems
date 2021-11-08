import {
  HStack,
  IconButton,
  Link,
  Spacer,
  Text,
  Tooltip,
  useColorMode,
} from '@chakra-ui/react';
import type { ReactElement } from 'react';
import { memo, MouseEventHandler } from 'react';
import { useTranslation } from 'react-i18next';
import {
  AiFillBook,
  AiFillBulb,
  AiFillGithub,
  AiFillInfoCircle,
  AiFillZhihuCircle,
} from 'react-icons/ai';

/**
 * Will accept a function to open up modals again.
 */
type Props = {
  setOpenDisclaimer: () => void;
  setOpenAbout: () => void;
};

/**
 * Stylized icon button to create custom and constant icon button based on props.
 *
 * @param params - Set of parameters native for this component
 * @returns Customized icon button component
 */
const StylizedIconButton = ({
  icon,
  label,
  onClick,
}: {
  icon: ReactElement<any, string | React.JSXElementConstructor<any>>;
  label: string;
  onClick: MouseEventHandler<HTMLButtonElement>;
}) => (
  <Tooltip label={label} hasArrow>
    <IconButton
      colorScheme="pink"
      icon={icon}
      size="sm"
      aria-label={label}
      variant="outline"
      onClick={onClick}
    />
  </Tooltip>
);

/**
 * Header component for the whole application.
 *
 * @param param - Setter function to open modals
 * @returns React Functional Component
 */
const Header = ({ setOpenDisclaimer, setOpenAbout }: Props) => {
  const { toggleColorMode } = useColorMode();
  const { t, i18n } = useTranslation();

  return (
    <HStack as="nav" px={6} py={4} spacing={2}>
      <Text fontWeight="bold" color="pink.400" fontSize="xl">
        ES
      </Text>

      <Spacer />

      <StylizedIconButton
        icon={<AiFillBulb />}
        label={t('tooltips.color')}
        onClick={toggleColorMode}
      />

      <Link href="https://github.com/lauslim12/expert-systems" isExternal>
        <StylizedIconButton
          icon={<AiFillGithub />}
          label={t('tooltips.github')}
          onClick={() => null}
        />
      </Link>

      <StylizedIconButton
        icon={<AiFillInfoCircle />}
        label={t('tooltips.about')}
        onClick={setOpenAbout}
      />

      <StylizedIconButton
        icon={<AiFillBook />}
        label={t('tooltips.terms')}
        onClick={setOpenDisclaimer}
      />

      <StylizedIconButton
        icon={<AiFillZhihuCircle />}
        label={t('tooltips.lang')}
        onClick={() => {
          if (i18n.language === 'id') {
            i18n.changeLanguage('en');
            return;
          }

          i18n.changeLanguage('id');
        }}
      />
    </HStack>
  );
};

export default memo(Header);
