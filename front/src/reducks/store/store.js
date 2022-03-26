// store を定義するファイル
import {
  createStore as reduxCreateStore,
  combineReducers,
  applyMiddleware,
} from 'redux';
// ルーティングライブラリ
import { connectRouter, routerMiddleware } from 'connected-react-router';
// 非同期ライブラリ
import thunk from 'redux-thunk';

// import reducers
import { UtilsReducer } from '../utils/reducers';
import { MessageReducer } from '../messages/reducers';
import { ProductsReducer } from '../products/reducers';
import { UsersReducer } from '../users/reducers';

// logger
import logger from 'redux-logger';

export default function createStore(history) {
  // 下記は state の構造と一致するように
  const reducers = {
    utils: UtilsReducer,
    users: UsersReducer,
    message: MessageReducer,
    products: ProductsReducer,
    router: connectRouter(history), // React が持つhistoryをstateで管理できるようにする
  };
  if (process.env.REACT_APP_ENV === 'development') {
    return reduxCreateStore(
      combineReducers(reducers),
      applyMiddleware(
        // アプリケーションのミドルウェアとしてrouterMiddlewareを使うということを設定している
        routerMiddleware(history),
        thunk,
        logger
      )
    );
  } else {
    return reduxCreateStore(
      combineReducers(reducers),
      applyMiddleware(routerMiddleware(history), thunk)
    );
  }
}
