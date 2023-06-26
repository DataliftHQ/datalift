import React from 'react';
import { useRoutes } from 'react-router-dom';

import DashboardLayout from '../components/layout/DashboardLayout';
import SimpleLayout from '../components/layout/SimpleLayout';
import Dashboard from '../components/pages/Dashboard';
import NotFound from '../components/pages/NotFound';

const routes = {
  path: '/',
  element: <DashboardLayout />,
  children: [{ path: '/', element: <Dashboard /> }],
};

const Routes = () => {
  return useRoutes([
    routes,
    {
      element: <SimpleLayout />,
      children: [{ path: '*', element: <NotFound /> }],
    },
  ]);
};

export default Routes;
