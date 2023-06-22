import { Provider } from 'react-redux';
import { BrowserRouter as Router } from 'react-router-dom';

import NavigationScroll from 'components/ui/NavigationScroll';
import Routes from 'routes';
import Snackbar from 'components/extended/Snackbar';
import Store from 'store';
import Theme from 'theme';

function App() {
  return (
    <Provider store={Store}>
      <Router>
        <Theme>
          <NavigationScroll>
            <>
              <Routes />
              <Snackbar />
            </>
          </NavigationScroll>
        </Theme>
      </Router>
    </Provider>
  );
}

export default App;
