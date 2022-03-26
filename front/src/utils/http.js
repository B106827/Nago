import config from '../config/base';
import { showMessageAction } from '../reducks/messages/actions';

const networkTimeout = 10000;

let networkErrorFlg = false;

export const commonActions = {
  networkError(dispatch) {
    dispatch(
      showMessageAction(
        'error',
        '予期せぬエラーが発生しました。解消されない場合はお手数ですが運営元にお問い合わせください'
      )
    );
  },
};

export const fetchWrapper = (args, dispatch) => {
  let { type, url, params = {} } = args; // eslint-disable-line prefer-const
  let options = null;

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
    default:
      return;
  }

  const result = new Promise((resolve, reject) => {
    setTimeout(() => {
      reject(new Error('timeout'));
    }, networkTimeout);
    fetch(url, options).then(resolve, reject);
  })
    // リクエスト以前の段階でのエラー処理
    .catch((e) => handleNetworkError(e))
    // ヘッダーステータスごとに処理
    .then((res) => handleStatusErrors(res))
    // メイン処理
    .then((json) => handleResult(json));

  const handleNetworkError = (error) => {
    if (networkErrorFlg) return;
    // 連続でエラーメッセージが出るのを制限
    console.log('Promise error', error);
    commonActions.networkError(dispatch);
    networkErrorFlg = true;
    setTimeout(() => {
      networkErrorFlg = false;
    }, 500);
  };

  const handleStatusErrors = (res) => {
    if (!res) return;
    switch (Number(res.status)) {
      case 200:
        return res.json(); // API側バリデーションエラー等も含む
      case 401:
        return handle401(res); // API側認証エラー
      case 404:
        return handle404(res); // API側Not Found エラー
      case 500:
        return handle500(res); // API側サーバーエラー
      default:
        return handleErrors(res); // API側想定外のエラー
    }
  };

  const handle401 = (res) => {
    console.log('401:unauthorized', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handle404 = (res) => {
    console.log('404:not found', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handle500 = (res) => {
    console.log('500:internal server error', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handleErrors = (res) => {
    console.log('unknown error', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handleResult = (json) => {
    return json;
  };

  return result;
};
