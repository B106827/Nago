import { push } from 'connected-react-router';
import { fetchWrapper } from '../../utils/http';
import { showMessageAction } from '../messages/actions';
import {
  fetchProductsInCartAction,
  loginAction,
  logedInAction,
  logoutAction,
  fetchUserTmpEmail,
} from './actions';

export const addProductToCart = () => {
  return async () => {
    // const uid = getState().users.uid;
    //const cartRef = db.collection('users').doc(uid).collection('cart').doc();
    //addedProduct['cartId'] = cartRef.id;
    //await cartRef.set(addedProduct);
    //dispatch(push('/'));
  };
};

export const fetchOrdersHistory = () => {
  return async () => {
    //const uid = getState().users.uid;
    //const list = [];
    // db.collection('users').doc(uid)
    //   .collection('orders')
    //   .orderBy('updated_at', 'desc')
    //   .get()
    //   .then((snapshots) => {
    //     snapshots.forEach(snapshot => {
    //       const data = snapshot.data();
    //       list.push(data);
    //     });
    //     dispatch(fetchOrdersHistoryAction(list));
    //   })
  };
};

export const fetchProductsInCart = (products) => {
  return async (dispatch) => {
    dispatch(fetchProductsInCartAction(products));
  };
};

export const resetPassword = (email) => {
  return async () => {
    if (email === '') {
      alert('必須項目が未入力です。');
      return false;
    } else {
      return;
      //auth.sendPasswordResetEmail(email)
      //  .then(() => {
      //    alert('入力されたメールアドレスにパスワードリセット用URLを送信しました');
      //    dispatch(push('/login'));
      //  }).catch(() => {
      //    alert('パスワードリセットに失敗しました。');
      //  });
    }
  };
};

// ログイン
export const login = (email, password) => {
  return (dispatch) => {
    if (email === '' || password === '') {
      dispatch(showMessageAction('error', '必須項目が未入力です'));
      return false;
    }
    const params = {
      email,
      password,
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/login',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        console.log('test', json);
        if (json.status === 200) {
          const user = json.data.user;
          if (user) {
            dispatch(
              loginAction({
                id: user.id,
                name: user.name,
              })
            );
            dispatch(showMessageAction('success', 'ログインしました'));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      })
      .catch((error) => {
        console.log('error :', error);
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
      });
  };
};

// 認証状態確認
export const listenAuth = (isRedirect = true) => {
  return async (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/listen_auth',
      },
      dispatch
    )
      .then((json) => {
        if (json.status === 200) {
          // 認証済み
          dispatch(logedInAction());
        } else if (isRedirect) {
          // 未認証
          dispatch(showMessageAction('info', 'ログインし直してください'));
          dispatch(push('/login'));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
      });
  };
};

// メールアドレス登録
export const registerEmail = (email) => {
  return async (dispatch) => {
    if (email === '') {
      dispatch(showMessageAction('error', 'メールアドレスが未入力です'));
      return false;
    }

    const params = {
      email,
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/register_email',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (json.status === 200) {
          dispatch(
            showMessageAction(
              'success',
              '登録用URLを送信しました。メールを確認してください'
            )
          );
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
      });
  };
};

// 新規登録URL有効性確認
export const checkRegisterUrl = (id) => {
  return async (dispatch) => {
    if (!id) {
      dispatch(showMessageAction('error', '無効なURLです'));
      return false;
    }

    const params = {
      id,
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/register_url_check',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (json.status === 200) {
          dispatch(showMessageAction('info', '新規登録を進めてください'));
          if (json.data.email) {
            dispatch(fetchUserTmpEmail(json.data.email));
          } else {
            dispatch(showMessageAction('error', '登録用URLを確認してください'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
      });
  };
};

// 新規登録
export const register = (tmpId, name, email, password, confirmPassword) => {
  return async (dispatch) => {
    if (!name || !email || !password || !confirmPassword) {
      dispatch(showMessageAction('error', '必須項目が未入力です'));
      return false;
    }
    if (password !== confirmPassword) {
      dispatch(showMessageAction('error', 'パスワードが一致しません'));
      return false;
    }

    const params = {
      tmpId,
      name,
      email,
      password,
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/register',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (json.status === 200) {
          const user = json.data.user;
          if (user) {
            // 新規登録およびログインが完了
            dispatch(
              loginAction({
                id: user.id,
                name: user.name,
              })
            );
            dispatch(showMessageAction('success', json.data.message));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
      });
  };
};

// ログアウト
export const logout = () => {
  return (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/logout',
      },
      dispatch
    )
      .then((json) => {
        if (json.status === 200) {
          dispatch(logoutAction());
          dispatch(showMessageAction('success', 'ログアウトしました'));
          dispatch(push('/'));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
      });
  };
};
