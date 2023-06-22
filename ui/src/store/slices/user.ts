import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { RootState } from 'store';

const userPreferenceState = JSON.parse(localStorage.getItem('user.preferences')!);

export type ThemeMode = 'light' | 'dark' | 'system';

export interface UserState {
  isAdmin?: boolean;
  preferences: {
    themeMode?: ThemeMode;
  };
}

const initialState: UserState = {
  isAdmin: false,
  preferences:
    userPreferenceState === null
      ? {
          themeMode: 'system',
        }
      : userPreferenceState,
};

const userSlice = createSlice({
  name: 'theme',
  initialState,
  reducers: {
    setThemeMode(state, action: PayloadAction<ThemeMode>) {
      state.preferences.themeMode = action.payload || initialState.preferences.themeMode;
    },
  },
});

export const { setThemeMode } = userSlice.actions;

export const selectUser = (state: RootState) => state.user;
export const selectThemeMode = (state: RootState) => state.user.preferences.themeMode;

export default userSlice.reducer;
