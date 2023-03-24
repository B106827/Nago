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
        cartList: [...action.payload],
      };
    case Actions.FETCH_ORDERS_HISTORY:
      return {
        ...state,
        orders: [...action.payload],
      };
    default:
      return state;
  }
};
