import CloseIcon from '@mui/icons-material/Close';
import { Alert, Button, Fade, Grow, IconButton, Slide, SlideProps } from '@mui/material';
import MuiSnackbar from '@mui/material/Snackbar';
import React from 'react';

import { useAppDispatch, useAppSelector } from '../../hooks/state';
import * as snackbar from '../../store/slices/snackbar';

const TransitionSlideLeft = (props: SlideProps) => {
  return <Slide {...props} direction='left' />;
};

const TransitionSlideUp = (props: SlideProps) => {
  return <Slide {...props} direction='up' />;
};

const TransitionSlideRight = (props: SlideProps) => {
  return <Slide {...props} direction='right' />;
};

const TransitionSlideDown = (props: SlideProps) => {
  return <Slide {...props} direction='down' />;
};

const GrowTransition = (props: SlideProps) => {
  return <Grow {...props} />;
};

export type KeyedObject = {
  [key: string]: string | number | KeyedObject | any;
};

const animation: KeyedObject = {
  SlideLeft: TransitionSlideLeft,
  SlideUp: TransitionSlideUp,
  SlideRight: TransitionSlideRight,
  SlideDown: TransitionSlideDown,
  Grow: GrowTransition,
  Fade,
};

const Snackbar = () => {
  const dispatch = useAppDispatch();
  const { actionButton, anchorOrigin, alert, close, message, open, transition, variant } = useAppSelector(
    snackbar.selectSnackbar,
  );

  const handleClose = (event: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    dispatch(snackbar.closeSnackbar());
  };

  return (
    <>
      {variant === 'default' && (
        <MuiSnackbar
          anchorOrigin={anchorOrigin}
          open={open}
          autoHideDuration={5000}
          onClose={handleClose}
          message={message}
          TransitionComponent={animation[transition]}
          action={
            <>
              <Button color='secondary' size='small' onClick={handleClose}>
                UNDO
              </Button>
              <IconButton size='small' aria-label='close' color='inherit' onClick={handleClose} sx={{ mt: 0.25 }}>
                <CloseIcon fontSize='small' />
              </IconButton>
            </>
          }
        />
      )}

      {variant === 'alert' && (
        <MuiSnackbar
          TransitionComponent={animation[transition]}
          anchorOrigin={anchorOrigin}
          open={open}
          autoHideDuration={5000}
          onClose={handleClose}
        >
          <Alert
            variant={alert.variant}
            color={alert.color}
            action={
              <>
                {actionButton !== false && (
                  <Button size='small' onClick={handleClose} sx={{ color: 'background.paper' }}>
                    UNDO
                  </Button>
                )}
                {close !== false && (
                  <IconButton sx={{ color: 'background.paper' }} size='small' aria-label='close' onClick={handleClose}>
                    <CloseIcon fontSize='small' />
                  </IconButton>
                )}
              </>
            }
            sx={{
              ...(alert.variant === 'outlined' && {
                bgcolor: 'background.paper',
              }),
            }}
          >
            {message}
          </Alert>
        </MuiSnackbar>
      )}
    </>
  );
};

export default Snackbar;
