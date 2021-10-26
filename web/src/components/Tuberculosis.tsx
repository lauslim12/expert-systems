import { Button, Text, useToast, VStack } from '@chakra-ui/react';
import { lazy, memo, Suspense, useState } from 'react';

import type Request from '../types/Request';
import type Response from '../types/Response';
import type { StateCertaintyWeight } from '../types/UserCertaintyWeight';
import request from '../utils/request';
import FailedToast from './FailedToast';
import AnswerInput from './Input/AnswerInput';

/**
 * Lazy-load modal, as it is not displayed right away.
 */
const ResultModal = lazy(() => import('../components/Modal/ResultModal'));

/**
 * Tuberculosis component to infer, render, and handle user inputs about this disease.
 *
 * @returns React Functional Component
 */
const Tuberculosis = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [openResult, setOpenResult] = useState(false);
  const [result, setResult] = useState({} as Response);
  const [fever, setFever] = useState(null as StateCertaintyWeight);
  const [coughBlood, setCoughBlood] = useState(null as StateCertaintyWeight);
  const [spBloody, setSpBloody] = useState(null as StateCertaintyWeight);
  const [nightSweat, setNightSweat] = useState(null as StateCertaintyWeight);
  const [chestPain, setChestPain] = useState(null as StateCertaintyWeight);
  const [backPain, setBackPain] = useState(null as StateCertaintyWeight);
  const [shortBreath, setShortBreath] = useState(null as StateCertaintyWeight);
  const [weightLoss, setWeightLoss] = useState(null as StateCertaintyWeight);
  const [bodyTired, setBodyTired] = useState(null as StateCertaintyWeight);
  const [lumps, setLumps] = useState(null as StateCertaintyWeight);
  const [coughing, setCoughing] = useState(null as StateCertaintyWeight);
  const [swollen, setSwollen] = useState(null as StateCertaintyWeight);
  const [lossApetite, setLossApetite] = useState(null as StateCertaintyWeight);
  const toast = useToast();

  const submitResult = () => {
    // Not '!' as 0 equals false as well. We need that literal 0 value.
    if (
      fever === null ||
      coughBlood === null ||
      spBloody === null ||
      nightSweat === null ||
      chestPain === null ||
      backPain === null ||
      shortBreath === null ||
      weightLoss === null ||
      bodyTired === null ||
      lumps === null ||
      coughing === null ||
      swollen === null ||
      lossApetite === null
    ) {
      FailedToast(
        toast,
        'Please input all of the information before continuing!'
      );
      return;
    }

    const requestBody: Request = {
      symptoms: [
        {
          symptomId: 'S1',
          weight: fever,
        },
        {
          symptomId: 'S2',
          weight: coughBlood,
        },
        {
          symptomId: 'S3',
          weight: spBloody,
        },
        {
          symptomId: 'S4',
          weight: nightSweat,
        },
        {
          symptomId: 'S5',
          weight: chestPain,
        },
        {
          symptomId: 'S6',
          weight: backPain,
        },
        {
          symptomId: 'S7',
          weight: shortBreath,
        },
        {
          symptomId: 'S8',
          weight: weightLoss,
        },
        {
          symptomId: 'S9',
          weight: bodyTired,
        },
        {
          symptomId: 'S10',
          weight: lumps,
        },
        {
          symptomId: 'S11',
          weight: coughBlood,
        },
        {
          symptomId: 'S12',
          weight: swollen,
        },
        {
          symptomId: 'S13',
          weight: lossApetite,
        },
      ],
    };

    setIsLoading(true);
    request('/api/v1', requestBody, 'POST')
      .then((data: Response) => setResult(data))
      .then(() => setOpenResult(true))
      .catch((err) => FailedToast(toast, err.message))
      .finally(() => setIsLoading(false));
  };

  return (
    <>
      <Suspense fallback={null}>
        <ResultModal
          isOpen={openResult}
          onClose={() => setOpenResult(false)}
          results={result}
        />
      </Suspense>

      <VStack as="form" w={['full', '70vw']} spacing={5}>
        <AnswerInput
          state={fever}
          setState={setFever}
          title="I have a fever for two weeks or more"
        />

        <AnswerInput
          state={coughBlood}
          setState={setCoughBlood}
          title="I cough blood"
        />

        <AnswerInput
          state={spBloody}
          setState={setSpBloody}
          title="My sputum is mixed with blood"
        />

        <AnswerInput
          state={nightSweat}
          setState={setNightSweat}
          title="I sweat at night"
        />

        <AnswerInput
          state={chestPain}
          setState={setChestPain}
          title="I suffer from chest pain constantly"
        />

        <AnswerInput
          state={backPain}
          setState={setBackPain}
          title="I suffer from back pain constantly"
        />

        <AnswerInput
          state={shortBreath}
          setState={setShortBreath}
          title="I feel a shortness of breath"
        />

        <AnswerInput
          state={weightLoss}
          setState={setWeightLoss}
          title="I have experienced weight loss"
        />

        <AnswerInput
          state={bodyTired}
          setState={setBodyTired}
          title="My body feels like always tired"
        />

        <AnswerInput
          state={lumps}
          setState={setLumps}
          title="Around my armpits and neck, lumps appeared"
        />

        <AnswerInput
          state={coughing}
          setState={setCoughing}
          title="I cough continously from two weeks to four weeks"
        />

        <AnswerInput
          state={swollen}
          setState={setSwollen}
          title="My lymph nodes are swollen"
        />

        <AnswerInput
          state={lossApetite}
          setState={setLossApetite}
          title="I experienced a loss of apetite"
        />

        <VStack w="full" align="start">
          <Text fontWeight="bold">Analyze answers</Text>

          <Button
            colorScheme="pink"
            w="full"
            variant="solid"
            onClick={submitResult}
            isLoading={isLoading}
          >
            Results
          </Button>
        </VStack>
      </VStack>
    </>
  );
};

export default memo(Tuberculosis);
