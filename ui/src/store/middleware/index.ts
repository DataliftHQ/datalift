import { createListenerMiddleware } from '@reduxjs/toolkit';

import { setThemeMode } from '../slices/user';
import { RootState } from 'store';

export const listenerMiddleware = createListenerMiddleware();

listenerMiddleware.startListening({
  actionCreator: setThemeMode, // isAnyOf for multiples
  effect: (action, listenerApi) => {
    localStorage.setItem('user.preferences', JSON.stringify((listenerApi.getState() as RootState).user.preferences));
  },
});
