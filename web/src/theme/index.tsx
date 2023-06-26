import {
  createTheme,
  CssBaseline,
  PaletteMode,
  StyledEngineProvider,
  Theme as MuiTheme,
  ThemeOptions,
  ThemeProvider,
  useMediaQuery,
} from '@mui/material';
import { TypographyOptions } from '@mui/material/styles/createTypography';
import React, { ReactElement, useMemo } from 'react';

import { useAppSelector } from '../hooks/state';
import ComponentStyleOverrides from './compStyleOverride';
import Palette from './palette';
import Typography from './typography';

const Theme = ({ children }: { children: ReactElement | null }) => {
  const { preferences } = useAppSelector((state) => state.user);
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');

  let paletteMode: PaletteMode = (prefersDarkMode ? 'dark' : 'light') as PaletteMode;
  if (preferences.themeMode !== 'system') {
    paletteMode = preferences.themeMode as PaletteMode;
  }

  const theme: MuiTheme = useMemo<MuiTheme>(() => Palette(paletteMode), [paletteMode]);
  const themeTypography: TypographyOptions = useMemo<TypographyOptions>(() => Typography(theme), [theme]);
  const themeOptions: ThemeOptions = useMemo(
    () => ({
      palette: theme.palette,
      typography: themeTypography,
    }),
    [theme, themeTypography],
  );

  const customTheme: MuiTheme = createTheme(themeOptions);
  customTheme.components = useMemo(() => ComponentStyleOverrides(customTheme), [customTheme]);

  return (
    <StyledEngineProvider injectFirst>
      <ThemeProvider theme={customTheme}>
        <CssBaseline />
        {children}
      </ThemeProvider>
    </StyledEngineProvider>
  );
};

export default Theme;
