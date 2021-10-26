import { UseToastOptions } from '@chakra-ui/react';

const FailedToast = (
  toast: (options?: UseToastOptions | undefined) => string | number | undefined,
  message: string
) =>
  toast({
    title: 'Failed!',
    description: message,
    status: 'error',
    duration: 1500,
    isClosable: true,
  });

export default FailedToast;
