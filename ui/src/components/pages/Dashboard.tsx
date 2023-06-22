import { Grid } from '@mui/material';
import React, { useState } from 'react';

const Dashboard: React.FC = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  return (
    <>
      <div>
        {error ? (
          <span>
            We&apos;re having issues loading your home page. You may see a different home page in the meantime as we
            resolve this issue.
          </span>
        ) : (
          <Grid item xs={12} zeroMinWidth>
            <h1>Dashboard</h1>
          </Grid>
        )}
      </div>
    </>
  );
};

export default Dashboard;
