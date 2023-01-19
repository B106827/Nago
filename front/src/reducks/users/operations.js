import { push } from 'connected-react-router';
import { fetchWrapper } from '../../utils/http';
import { showMessageAction } from '../messages/actions';
import {
  loginAction,
  logoutAction,
  fetchUserTmpEmail,
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
        if (json.status === 200) {
          const user = json.result.user;
          if (user) {
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
      })
      .catch((error) => {
        console.log('error :', error);
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
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
      })
      .catch((error) => {
        console.log('error: ', error);
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
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
        if (json.status === 200) {
          dispatch(showMessageAction('info', '新規登録を進めてください'));
          if (json.result.email) {
            dispatch(fetchUserTmpEmail(json.result.email));
          } else {
            dispatch(showMessageAction('error', '登録用URLを確認してください'));
            dispatch(push('/'));
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
          dispatch(push('/'));
        }
      })
      .catch((error) => {
        dispatch(showMessageAction('error', '予期せぬエラーが発生しました'));
        console.log(error);
        dispatch(push('/'));
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
          dispatch(showMessageAction('success', json.result.message));
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
        if (json.status === 200) {
          console.log(json);
          dispatch(showMessageAction('success', json.result.message));
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
