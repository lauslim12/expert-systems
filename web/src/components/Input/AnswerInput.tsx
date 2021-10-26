import { Button, FormControl, FormLabel, SimpleGrid } from '@chakra-ui/react';
import { Dispatch, memo, SetStateAction } from 'react';
import { AiOutlineCheck } from 'react-icons/ai';

import { StateCertaintyWeight } from '../../types/UserCertaintyWeight';

/**
 * Will accept a state, the setter, and the title.
 */
type Props = {
  state: StateCertaintyWeight;
  setState: Dispatch<SetStateAction<StateCertaintyWeight>>;
  title: string;
};

/**
 * Reusable component to handle user inputs according to the Expert System.
 *
 * @param param - the state, state's setter, and the title
 * @returns React Functional Component
 */
const AnswerInput = ({ state, setState, title }: Props) => (
  <FormControl as="fieldset">
    <FormLabel as="legend">{title}</FormLabel>

    <SimpleGrid columns={[1, 2, 4]} spacing="10px">
      <Button
        colorScheme="green"
        leftIcon={state === 0 ? <AiOutlineCheck /> : undefined}
        variant={state === 0 ? 'solid' : 'outline'}
        onClick={() => setState(0)}
      >
        I do not/never feel so
      </Button>

      <Button
        colorScheme="blue"
        leftIcon={state === 0.25 ? <AiOutlineCheck /> : undefined}
        variant={state === 0.25 ? 'solid' : 'outline'}
        onClick={() => setState(0.25)}
      >
        I sometimes feel so
      </Button>

      <Button
        colorScheme="orange"
        leftIcon={state === 0.75 ? <AiOutlineCheck /> : undefined}
        variant={state === 0.75 ? 'solid' : 'outline'}
        onClick={() => setState(0.75)}
      >
        I often feel so
      </Button>

      <Button
        colorScheme="red"
        leftIcon={state === 1 ? <AiOutlineCheck /> : undefined}
        variant={state === 1 ? 'solid' : 'outline'}
        onClick={() => setState(1)}
      >
        I strongly feel so
      </Button>
    </SimpleGrid>
  </FormControl>
);

export default memo(AnswerInput);
