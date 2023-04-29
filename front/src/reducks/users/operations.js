import { push } from 'connected-react-router';
import { fetchWrapper } from '../../utils/http';
import { showMessageAction } from '../messages/actions';
import { customValidErrAction } from '../utils/actions';
import {
  loginAction,
  logoutAction,
  fetchUserTmpEmailAction,
  updateCartAction,
  resetCartAction,
} from './actions';

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
        if (!json) return;
        if (json.status === 200) {
          const user = json.result.user;
          const cartList = json.result.cartList;
          if (user) {
            dispatch(
              loginAction({
                id: user.id,
                name: user.name,
                cartList: cartList,
              })
            );
            dispatch(showMessageAction('success', json.result.message));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// メールアドレス登録
export const registerEmail = (email) => {
  return (dispatch) => {
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
        if (!json) return;
        if (json.status === 200) {
          dispatch(
            showMessageAction(
              'success',
              json.result.message
            )
          );
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// 新規登録URL有効性確認
export const checkRegisterUrl = (tmpId) => {
  return async (dispatch) => {
    if (!tmpId) {
      dispatch(showMessageAction('error', '無効なURLです'));
      return false;
    }
    const params = {
      tmpId,
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
        if (!json) return;
        if (json.status === 200) {
          dispatch(showMessageAction('info', '新規登録を進めてください'));
          if (json.result.email) {
            dispatch(fetchUserTmpEmailAction(json.result.email));
          } else {
            dispatch(showMessageAction('error', '登録用URLを確認してください'));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
          dispatch(push('/'));
        }
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
        if (!json) return;
        if (json.status === 200) {
          const user = json.result.user;
          if (user) {
            // 新規登録およびログインが完了
            dispatch(
              loginAction({
                id: user.id,
                name: user.name,
              })
            );
            dispatch(showMessageAction('success', json.result.message));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
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
        if (!json) return;
        if (json.status === 200) {
          dispatch(logoutAction());
          dispatch(showMessageAction('success', json.result.message));
          dispatch(push('/'));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// ユーザー情報取得
export const fetchMyInfo = () => {
  return (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/user/myinfo',
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          dispatch(showMessageAction('success', json.result.message));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};


// ユーザーカート情報更新
export const updateCart = (productId, cartNum) => {
  return (dispatch) => {
    if (!productId || !cartNum) {
      dispatch(showMessageAction('error', 'カートを更新できません'));
      return false;
    }
    const params = {
      productId: Number(productId),
      cartNum: cartNum,
    };
    fetchWrapper(
      {
        type: 'PUT',
        url: '/cart',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          dispatch(updateCartAction(json.result.updatedCartList));
          dispatch(showMessageAction('success', json.result.message));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// ユーザーカート情報削除
export const deleteCart = (cartId) => {
  return (dispatch) => {
    if (!cartId) {
      dispatch(showMessageAction('error', '商品が指定されていません'));
      return false;
    }
    const params = {
      cartId
    };
    fetchWrapper(
      {
        type: 'DELETE',
        url: '/cart',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          dispatch(updateCartAction(json.result.updatedCartList));
          dispatch(showMessageAction('success', json.result.message));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// 決済情報入力へ
export const createOrder = (total, address) => {
  const params = {
    total,
    ...address,
  };
  return (dispatch) => {
    fetchWrapper(
      {
        type: 'POST',
        url: '/order/create',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          location.href = json.result.orderSession.url;
        } else {
          dispatch(showMessageAction('error', json.messages));
          if (json.isCustomValidErr) {
            dispatch(customValidErrAction(json.result));
          }
        }
      });
  };
};

// 決済完了後の決済結果確認
export const checkCheckoutResult = (sessionId, orderId) => {
  return async (dispatch) => {
    const params = {
      sessionId,
      orderId: Number(orderId),
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/order/check',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          dispatch(resetCartAction());
          dispatch(showMessageAction('success', json.result.message));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

// 決済のキャンセル処理
export const cancelCheckout = (orderId) => {
  return async (dispatch) => {
    const params = {
      orderId: Number(orderId),
    };
    fetchWrapper(
      {
        type: 'POST',
        url: '/order/cancel',
        params: params,
      },
      dispatch
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          dispatch(showMessageAction('success', json.result.message));
          dispatch(push('/'));
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};
