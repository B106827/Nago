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
  })
    // リクエスト以前の段階でのエラー処理
    .catch((e) => handleNetworkError(e))
    // HTTPリクエストのステータスごとに処理
    .then((res) => handleStatusErrors(res))
    // ビジネスロジックの処理結果を扱う
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
        return handleStatus401(res); // API側認証エラー
      case 404:
        return handleStatus404(res); // API側Not Found エラー
      case 500:
        return handeleStatus500(res); // API側サーバーエラー
      default:
        return handleErrors(res); // API側想定外のエラー
    }
  };

  const handleStatus401 = (res) => {
    console.log('401:unauthorized', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handleStatus404 = (res) => {
    console.log('404:not found', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handeleStatus500 = (res) => {
    console.log('500:internal server error', res);
    commonActions.networkError(dispatch);
    return;
  };

  const handleErrors = (res) => {
    console.log('unknown error', res);
    commonActions.networkError(dispatch);
    return;
  };

  // HTTPリクエスト自体は成功し、ビジネスロジックの処理結果を扱う
  const handleResult = (json) => {
    if (!json) return;
    switch (Number(json.status)) {
      case 200:
        return json; // API側処理成功
      case 400:
        return fixResponse400(json); // リクエスト内容に問題あり
      case 401:
        return fixResponse401(json); // 認証内容に問題あり
      case 404:
        return fixResponse404(json); // リクエスト先が見つからない
      case 500:
        return fixResponse500(json); // API側の処理に問題あり
      default:
        return json;
    }
  };

  const fixResponse400 = (json) => {
    const result = json.result;
    if (!result) return json;
    // バリデーションエラーの該当箇所を調整する
    json.errorKeys = Object.keys(result);
    json.messages  = Object.values(result);
    return json;
  };

  const fixResponse401 = (json) => {
    const result = json.result;
    if (!result) return json;
    json.messages = Object.values(result);
    return json;
  };

  const fixResponse404 = (json) => {
    const result = json.result;
    if (!result) return json;
    json.messages = Object.values(result);
    return json;
  };

  const fixResponse500 = (json) => {
    const result = json.result;
    if (!result) return json;
    json.messages = Object.values(result);
    return json;
  };

  return result;
};
