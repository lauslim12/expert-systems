import 'react-i18next';

import ns1 from '../public/locales/en/translation.json';

declare module 'react-i18next' {
  interface CustomTypeOptions {
    defaultNS: 'ns1';
    resources: {
      ns1: typeof ns1;
    };
  }
}
