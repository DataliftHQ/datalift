import { Action, combineReducers, configureStore, ThunkAction } from '@reduxjs/toolkit';
import logger from 'redux-logger';

import { listenerMiddleware } from './middleware';
import authReducer from './slices/auth';
import menuReducer from './slices/menu';
import snackbarReducer from './slices/snackbar';
import userReducer from './slices/user';

const reducer = combineReducers({
  snackbar: snackbarReducer,
  user: userReducer,
  auth: authReducer,
  menu: menuReducer,
});

const Store = configureStore({
  reducer: reducer,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger, listenerMiddleware.middleware),
  devTools: process.env.NODE_ENV !== 'production',
});

export type AppDispatch = typeof Store.dispatch;
export type RootState = ReturnType<typeof Store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType, RootState, unknown, Action<string>>;

export default Store;
