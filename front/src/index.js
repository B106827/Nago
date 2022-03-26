// Store と React の接続
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import createStore from './reducks/store/store';
import { ConnectedRouter } from 'connected-react-router';
import * as History from 'history';
import { MuiThemeProvider } from '@material-ui/core';
import { theme } from './assets/theme';
import App from './App';
import { windowResizeAction } from './reducks/utils/actions';

// 下記で store が作られる
const history = History.createBrowserHistory();
export const store = createStore(history);

// 画面サイズを監視
window.addEventListener('resize', () => {
  store.dispatch(
    windowResizeAction({ width: window.innerWidth, height: window.innerHeight })
  );
});

ReactDOM.render(
  // React に store が使えるようにしている
  // AppコンポーネントをProviderやConnectedRouterでラッピングしている
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <MuiThemeProvider theme={theme}>
        <App />
      </MuiThemeProvider>
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root')
);
