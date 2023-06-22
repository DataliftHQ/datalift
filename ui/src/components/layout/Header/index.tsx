import { Menu as MenuIcon } from '@mui/icons-material';
import { AppBar, Avatar, Box, Toolbar } from '@mui/material';

import LogoSection from './LogoSection';
import NotificationSection from './NotificationSection';
import ProfileSection from './ProfileSection';

const Header = () => {
  return (
    <AppBar enableColorOnDark position='fixed' color='inherit' elevation={0} sx={{ backgroundColor: 'steelblue' }}>
      <Toolbar>
        <Box sx={{ width: 228, display: 'flex' }}>
          <Box
            component='span'
            sx={{
              display: 'flex',
              alignItems: 'center',
              flexGrow: 1,
            }}
          >
            <LogoSection />
          </Box>
          <Avatar variant='rounded' onClick={() => console.log('menu clicked')}>
            <MenuIcon />
          </Avatar>
        </Box>

        <Box sx={{ flexGrow: 1 }} />

        <NotificationSection />
        <ProfileSection />
      </Toolbar>
    </AppBar>
  );
};

export default Header;
