export const LOGIN = 'LOGIN';
export const loginAction = (userState) => {
  return {
    type: LOGIN,
    payload: {
      isLogedIn: true,
      id: userState.id,
      name: userState.name,
      cartList: userState.cartList,
    },
  };
};

export const LOGEDIN = 'LOGEDIN';
export const logedInAction = () => {
  return {
    type: LOGEDIN,
    payload: {
      isLogedIn: true,
    },
  };
};

export const FETCH_USER_TMP_EMAIL = 'FETCH_USER_TMP_EMAIL';
export const fetchUserTmpEmailAction = (email) => {
  return {
    type: FETCH_USER_TMP_EMAIL,
    payload: {
      tmpEmail: email,
    },
  };
};

export const LOGOUT = 'LOGOUT';
export const logoutAction = () => {
  return {
    type: LOGOUT,
    payload: {
      isLogedIn: false,
      id: '',
      name: '',
    },
  };
};

export const UPDATE_CART = 'UPDATE_CART';
export const updateCartAction = (updatedCartList) => {
  return {
    type: UPDATE_CART,
    payload: updatedCartList,
  };
};

export const RESET_CART = 'RESET_CART';
export const resetCartAction = () => {
  return {
    type: RESET_CART,
  };
};

export const FETCH_ORDER_HISTORY = 'FETCH_ORDER_HISTORY';
export const fetchOrderHistoryAction = (orderHistoryList) => {
  return {
    type: FETCH_ORDER_HISTORY,
    payload: orderHistoryList,
  };
};
