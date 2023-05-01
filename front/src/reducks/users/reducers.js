import * as Actions from './actions';
import initialState from '../store/initialState';

export const UsersReducer = (state = initialState.users, action) => {
  switch (action.type) {
    case Actions.LOGIN:
      // スプレッド構文は後から渡したもので上書きされる
      // 下記であれば state と action.payload で同じキーがあれば action.payload 側で上書きされる
      return {
        ...state,
        ...action.payload,
      };
    case Actions.LOGEDIN:
      return {
        ...action.payload,
      };
    case Actions.FETCH_USER_TMP_EMAIL:
      return {
        ...action.payload,
      };
    case Actions.LOGOUT:
      return {
        ...action.payload,
      };
    case Actions.UPDATE_CART:
      return {
        ...state,
        cartList: Array.isArray(action.payload) ? [...action.payload] : [],
      };
    case Actions.RESET_CART:
      return {
        ...state,
        cartList: [],
      };
    case Actions.FETCH_ORDER_HISTORY:
      return {
        ...state,
        orderHistoryList: [...action.payload],
      };
    default:
      return state;
  }
};
