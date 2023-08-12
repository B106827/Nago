import config from '../config/base';
import { showMessageAction } from '../reducks/messages/actions';
import { logoutAction } from '../reducks/users/actions';

export const fetchWrapper = (args, dispatch) => {
  let { type, url, params = {} } = args; // eslint-disable-line prefer-const
  let options             = null;
  const networkTimeout    = 10000;
  let networkErrorFlg     = false;

  url = config.apiUrl + url;

  switch (type) {
    case 'GET':
      options = {
        method: 'GET',
      };
      break;
    case 'POST':
      options = {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(params),
      };
      break;
    case 'PUT':
      options = {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(params),
      };
      break;
    case 'DELETE':
      options = {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(params),
      };
      break;
    default:
      return;
  }

  const result = new Promise((resolve, reject) => {
    setTimeout(() => {
      reject(new Error('timeout'));
    }, networkTimeout);
    fetch(url, options).then(resolve, reject);
  }).then((res) => handleStatusErrors(res))
    .then((json) => handleResult(json))
    .catch((e) => handleServerError(e));

  const handleStatusErrors = (res) => {
    if (!res) return;
    switch (Number(res.status)) {
      case 200:
        return res.json();                  // API側バリデーションエラー等も含む
      case 400:
        return handleStatus400(res);        // API側不正なリクエストエラー
      case 401:
        return handleStatus401(res);        // API側認証エラー
      case 404:
        return handleStatus404(res);        // API側Not Foundエラー
      case 500:
        return handeleStatus500(res);       // API側サーバーエラー
      default:
        return handleUnexpectedErrors(res); // API側想定外のエラー
    }
  };

  const handleStatus400 = (res) => {
    console.log('400:bad request', res);
    // TODO: API側でmissing or malformed jwtの場合のレスポンスをリファクタリングし、resをthenで繋げなくても判別するようにする
    res.json().then(data => {
      if (data.message && data.message === 'missing or malformed jwt') {
        dispatch(logoutAction());
        throw new Error('ログインしてださい');
      } else {
        throw new Error('不正なリクエストです');
      }
    }).catch((e) => {
      dispatch(showMessageAction('error', e.message));
    });
  };

  const handleStatus401 = (res) => {
    console.log('401:unauthorized', res);
    dispatch(logoutAction());
    throw new Error('ログインしてください');
  };

  const handleStatus404 = (res) => {
    console.log('404:not found', res);
    throw new Error('ページが見つかりません');
  };

  const handeleStatus500 = (res) => {
    console.log('500:internal server error', res);
    throw new Error('エラーが発生しました');
  };

  const handleUnexpectedErrors = (res) => {
    console.log('unexpected error', res);
    throw new Error('予期せぬエラーが発生しました');
  };

  // HTTPリクエストは成功し（ステータス200）、ビジネスロジックの処理結果を扱う
  const handleResult = (json) => {
    if (!json) return;
    switch (Number(json.status)) {
      case 200: // API側処理成功
        return json;
      case 400:
      case 401:
      case 404:
      case 500: // API側の処理に問題あり
        return fixResponse(json);
      default:
        return json;
    }
  };

  const fixResponse = (json) => {
    if (!json.result) return json;
    if (json.isCustomValidErr) {
      // カスタムバリデーションエラー
      json.messages = ['入力内容に誤りがあります'];
    } else {
      json.messages = Object.values(json.result);
    }
    return json;
  };

  const handleServerError = (error) => {
    if (networkErrorFlg) return;
    // 連続でエラーメッセージが出るのを制限
    networkErrorFlg = true;
    dispatch(
      showMessageAction(
        'error',
        error.message
      )
    );
    setTimeout(() => {
      networkErrorFlg = false;
    }, 500);
  };

  return result;
};
