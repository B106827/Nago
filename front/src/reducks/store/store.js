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
// store永続化ライブラリ
import { persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
// import reducers
import { UtilsReducer } from '../utils/reducers';
import { MessageReducer } from '../messages/reducers';
import { ProductsReducer } from '../products/reducers';
import { UsersReducer } from '../users/reducers';
// logger
import logger from 'redux-logger';

export default function createStore(history) {
  // reducers（ store の構造と一致させる）
  const reducers = combineReducers({
    utils:    UtilsReducer,
    users:    UsersReducer,
    message:  MessageReducer,
    products: ProductsReducer,
    router:   connectRouter(history), // historyをstateで管理できるようにする
  });
  // redux-persist設定
  const persistConfig = {
    key: process.env.REACT_APP_ENV + ':root', // storageに保存されるキー名
    storage,                                  // 保存先はlocalstorage
    whitelist: ['users'],                     // storeの`users`のみ永続化する
  }
  // 永続化されたReducerとして定義
  const persistedReducers = persistReducer(persistConfig, reducers)
  if (process.env.REACT_APP_ENV === 'development') {
    return reduxCreateStore(
      persistedReducers,
      applyMiddleware(
        // アプリケーションのミドルウェアとしてrouterMiddlewareを使うということを設定している
        routerMiddleware(history),
        thunk,
        logger
      )
    );
  } else {
    return reduxCreateStore(
      persistedReducers,
      applyMiddleware(routerMiddleware(history), thunk)
    );
  }
}
