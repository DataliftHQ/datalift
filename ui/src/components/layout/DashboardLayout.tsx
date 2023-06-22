import { Container, styled, useTheme } from '@mui/material';
import React from 'react';
import { Outlet } from 'react-router-dom';

import Header from './Header';

const StyledRoot = styled('div')({
  display: 'flex',
  minHeight: '100%',
  overflow: 'hidden',
});

const Main = styled('main')(({ theme }) => ({
  //   ...theme.typography.mainContent,
}));

const DashboardLayout = () => {
  const theme = useTheme();

  return (
    <StyledRoot>
      <Header />

      <Main theme={theme}>
        <Container maxWidth='xl'>
          <Outlet />
        </Container>
      </Main>
    </StyledRoot>
  );
};

export default DashboardLayout;
