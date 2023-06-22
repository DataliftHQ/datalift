import { Link } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';

import Logo from 'components/ui/Logo';

const LogoSection = () => (
  <Link component={RouterLink} to={'/'}>
    <Logo />
  </Link>
);

export default LogoSection;
