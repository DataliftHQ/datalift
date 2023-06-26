import { Theme } from '@mui/material';

const ComponentStyleOverrides = (theme: Theme) => {
  return {
    MuiButton: {
      styleOverrides: {
        root: {
          fontWeight: 500,
          borderRadius: '4px',
        },
      },
    },
  };
};

export default ComponentStyleOverrides;
