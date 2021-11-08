import 'focus-visible/dist/focus-visible';
import './utils/i18n';

import { ChakraProvider, ColorModeScript, extendTheme } from '@chakra-ui/react';
import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import reportWebVitals from './reportWebVitals';

/**
 * Fallback fonts.
 */
const fallbackFonts =
  '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"';

/**
 * Renders whole application.
 */
ReactDOM.render(
  <React.StrictMode>
    <ChakraProvider
      resetCSS
      theme={extendTheme({
        styles: {
          global: {
            '::selection': {
              backgroundColor: '#FBB6CE',
              color: '#000',
            },
          },
        },
        fonts: {
          body: `Karla, ${fallbackFonts}`,
          heading: `Karla, ${fallbackFonts}`,
        },
        config: {
          useSystemColorMode: 'true',
        },
      })}
    >
      <ColorModeScript initialColorMode="system" />

      <App />
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
