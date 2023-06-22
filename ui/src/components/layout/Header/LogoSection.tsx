import { Link } from '@mui/material';
import React from 'react';
import { Link as RouterLink } from 'react-router-dom';

import Logo from '../../ui/Logo';

const LogoSection = () => (
  <Link component={RouterLink} to={'/'}>
    <Logo />
  </Link>
);

export default LogoSection;
