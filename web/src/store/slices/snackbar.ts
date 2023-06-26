import { AlertProps, SnackbarOrigin } from '@mui/material';
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

import { RootState } from '../../store';

export interface SnackbarState {
  action: boolean;
  open: boolean;
  message: string;
  anchorOrigin: SnackbarOrigin;
  variant: string;
  alert: AlertProps;
  transition: string;
  close: boolean;
  actionButton: boolean;
}

const initialState: SnackbarState = {
  action: false,
  open: false,
  message: 'Note archived',
  anchorOrigin: {
    vertical: 'bottom',
    horizontal: 'center',
  },
  variant: 'default',
  alert: {
    // color: 'primary',
    variant: 'filled',
  },
  transition: 'Fade',
  close: true,
  actionButton: false,
};

const snackbarSlice = createSlice({
  name: 'snackbar',
  initialState,
  reducers: {
    openSnackbar(state, action: PayloadAction<any>) {
      const { open, message, anchorOrigin, variant, alert, transition, close, actionButton } = action.payload;

      state.action = !state.action;
      state.open = open || initialState.open;
      state.message = message || initialState.message;
      state.anchorOrigin = anchorOrigin || initialState.anchorOrigin;
      state.variant = variant || initialState.variant;
      state.alert = {
        color: alert?.color || initialState.alert.color,
        variant: alert?.variant || initialState.alert.variant,
      };
      state.transition = transition || initialState.transition;
      state.close = close === false ? close : initialState.close;
      state.actionButton = actionButton || initialState.actionButton;
    },
    closeSnackbar(state) {
      state.open = false;
    },
  },
});

export const { closeSnackbar, openSnackbar } = snackbarSlice.actions;

export const selectSnackbar = (state: RootState) => state.snackbar;

export default snackbarSlice.reducer;
