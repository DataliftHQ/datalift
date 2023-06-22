import { Box, Button, Container, styled, Typography } from '@mui/material';
import React from 'react';
import { Link as RouterLink } from 'react-router-dom';

const StyledContent = styled('div')(({ theme }) => ({
  maxWidth: 480,
  margin: 'auto',
  minHeight: '100vh',
  display: 'flex',
  justifyContent: 'center',
  flexDirection: 'column',
  padding: theme.spacing(12, 0),
}));

const NotFound: React.FC = () => {
  return (
    <Container>
      <StyledContent sx={{ textAlign: 'center', alignItems: 'center' }}>
        <Typography variant='h1' paragraph>
          Sorry, page not found!
        </Typography>

        <Typography sx={{ color: 'text.secondary' }}>
          Sorry, we couldn’t find the page you’re looking for. Perhaps you’ve mistyped the URL? Be sure to check your
          spelling.
        </Typography>

        <Box component='img' src='/images/404.svg' sx={{ height: 260, mx: 'auto', my: { xs: 5, sm: 10 } }} />

        <Button to='/' size='large' variant='contained' component={RouterLink}>
          Go to Home
        </Button>
      </StyledContent>
    </Container>
  );
};

export default NotFound;
