import { Button, VStack } from '@chakra-ui/react';
import { lazy, memo, Suspense, useState } from 'react';

import type Request from '../types/Request';
import type Response from '../types/Response';
import type UserCertaintyWeight from '../types/UserCertaintyWeight';
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
  const [openResult, setOpenResult] = useState(false);
  const [result, setResult] = useState({} as Response);
  const [isFever, setIsFever] = useState(0 as UserCertaintyWeight);
  const [isCoughBlood, setIsCoughBlood] = useState(0 as UserCertaintyWeight);
  const [isSpBloody, setIsSpBloody] = useState(0 as UserCertaintyWeight);
  const [isNightSweat, setIsNightSweat] = useState(0 as UserCertaintyWeight);
  const [isChestPain, setIsChestPain] = useState(0 as UserCertaintyWeight);
  const [isBackPain, setIsBackPain] = useState(0 as UserCertaintyWeight);
  const [isShortBreath, setIsShortBreath] = useState(0 as UserCertaintyWeight);
  const [isWeightLoss, setIsWeightLoss] = useState(0 as UserCertaintyWeight);
  const [isBodyTired, setIsBodyTired] = useState(0 as UserCertaintyWeight);
  const [isLumps, setIsLumps] = useState(0 as UserCertaintyWeight);
  const [isCoughing, setIsCoughing] = useState(0 as UserCertaintyWeight);
  const [isSwollen, setIsSwollen] = useState(0 as UserCertaintyWeight);
  const [isLossApetite, setIsLossApetite] = useState(0 as UserCertaintyWeight);

  const submitResult = () => {
    const requestBody: Request = {
      symptoms: [
        {
          symptomId: 'S1',
          weight: isFever,
        },
        {
          symptomId: 'S2',
          weight: isCoughBlood,
        },
        {
          symptomId: 'S3',
          weight: isSpBloody,
        },
        {
          symptomId: 'S4',
          weight: isNightSweat,
        },
        {
          symptomId: 'S5',
          weight: isChestPain,
        },
        {
          symptomId: 'S6',
          weight: isBackPain,
        },
        {
          symptomId: 'S7',
          weight: isShortBreath,
        },
        {
          symptomId: 'S8',
          weight: isWeightLoss,
        },
        {
          symptomId: 'S9',
          weight: isBodyTired,
        },
        {
          symptomId: 'S10',
          weight: isLumps,
        },
        {
          symptomId: 'S11',
          weight: isCoughBlood,
        },
        {
          symptomId: 'S12',
          weight: isSwollen,
        },
        {
          symptomId: 'S13',
          weight: isLossApetite,
        },
      ],
    };

    fetch('/api/v1', {
      method: 'POST',
      body: JSON.stringify(requestBody),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then((data: Response) => setResult(data))
      .then(() => setOpenResult(true));
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

      <VStack as="form" spacing={5}>
        <AnswerInput
          state={isFever}
          setState={setIsFever}
          title="I have a fever for two weeks or more"
        />

        <AnswerInput
          state={isCoughBlood}
          setState={setIsCoughBlood}
          title="I cough blood"
        />

        <AnswerInput
          state={isSpBloody}
          setState={setIsSpBloody}
          title="My sputum is mixed with blood"
        />

        <AnswerInput
          state={isNightSweat}
          setState={setIsNightSweat}
          title="I sweat at night"
        />

        <AnswerInput
          state={isChestPain}
          setState={setIsChestPain}
          title="I suffer from chest pain constantly"
        />

        <AnswerInput
          state={isBackPain}
          setState={setIsBackPain}
          title="I suffer from back pain constantly"
        />

        <AnswerInput
          state={isShortBreath}
          setState={setIsShortBreath}
          title="I feel a shortness of breath"
        />

        <AnswerInput
          state={isWeightLoss}
          setState={setIsWeightLoss}
          title="I have experienced weight loss"
        />

        <AnswerInput
          state={isBodyTired}
          setState={setIsBodyTired}
          title="My body feels like always tired"
        />

        <AnswerInput
          state={isLumps}
          setState={setIsLumps}
          title="Around my armpits and neck, lumps appeared"
        />

        <AnswerInput
          state={isCoughing}
          setState={setIsCoughing}
          title="I cough continously from two weeks to four weeks"
        />

        <AnswerInput
          state={isSwollen}
          setState={setIsSwollen}
          title="My lymph nodes are swollen"
        />

        <AnswerInput
          state={isLossApetite}
          setState={setIsLossApetite}
          title="I experienced a loss of apetite"
        />

        <Button
          colorScheme="green"
          w="full"
          variant="outline"
          onClick={submitResult}
        >
          Results
        </Button>
      </VStack>
    </>
  );
};

export default memo(Tuberculosis);
