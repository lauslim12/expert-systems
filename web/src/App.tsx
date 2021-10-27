import { Suspense } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import AppLoader from './pages/AppLoader';
import Home from './pages/Home';
import NotFound from './pages/NotFound';

/**
 * Application starting point.
 *
 * @returns React Functional Component
 */
const App = () => (
  <Suspense fallback={<AppLoader />}>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  </Suspense>
);

export default App;
