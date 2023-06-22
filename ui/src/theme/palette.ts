import { createTheme, PaletteMode } from '@mui/material';
import colors from 'assets/scss/_theme.module.scss';

const Palette = (paletteMode: PaletteMode) => {
  return createTheme({
    palette: {
      mode: paletteMode,
      common: {
        black: colors.darkPaper,
      },
      primary: {
        light: paletteMode === 'dark' ? colors.darkPrimaryLight : colors.primaryLight,
        main: paletteMode === 'dark' ? colors.darkPrimaryMain : colors.primaryMain,
        dark: paletteMode === 'dark' ? colors.darkPrimaryDark : colors.primaryDark,
        200: paletteMode === 'dark' ? colors.darkPrimary200 : colors.primary200,
        800: paletteMode === 'dark' ? colors.darkPrimary800 : colors.primary800,
      },
      secondary: {
        light: paletteMode === 'dark' ? colors.darkSecondaryLight : colors.secondaryLight,
        main: paletteMode === 'dark' ? colors.darkSecondaryMain : colors.secondaryMain,
        dark: paletteMode === 'dark' ? colors.darkSecondaryDark : colors.secondaryDark,
        200: paletteMode === 'dark' ? colors.darkSecondary200 : colors.secondary200,
        800: paletteMode === 'dark' ? colors.darkSecondary800 : colors.secondary800,
      },
      error: {
        light: colors.errorLight,
        main: colors.errorMain,
        dark: colors.errorDark,
      },
      warning: {
        light: colors.warningLight,
        main: colors.warningMain,
        dark: colors.warningDark,
      },
      success: {
        light: colors.successLight,
        200: colors.success200,
        main: colors.successMain,
        dark: colors.successDark,
      },
      grey: {
        50: colors.grey50,
        100: colors.grey100,
        500: paletteMode === 'dark' ? colors.darkTextSecondary : colors.grey500,
        600: paletteMode === 'dark' ? colors.darkTextTitle : colors.grey900,
        700: paletteMode === 'dark' ? colors.darkTextPrimary : colors.grey700,
        900: paletteMode === 'dark' ? colors.darkTextPrimary : colors.grey900,
      },
      text: {
        primary: paletteMode === 'dark' ? colors.darkTextPrimary : colors.grey700,
        secondary: paletteMode === 'dark' ? colors.darkTextSecondary : colors.grey500,
      },
      divider: paletteMode === 'dark' ? colors.darkTextPrimary : colors.grey200,
      background: {
        paper: paletteMode === 'dark' ? colors.darkLevel2 : colors.paper,
        default: paletteMode === 'dark' ? colors.darkPaper : colors.paper,
      },
    },
  });
};

export default Palette;
