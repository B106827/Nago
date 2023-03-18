export const LOGIN = 'LOGIN';
export const loginAction = (userState) => {
  return {
    type: LOGIN,
    payload: {
      isLogedIn: true,
      id: userState.id,
      name: userState.name,
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
      role: '',
      uid: '',
      name: '',
    },
  };
};

export const FETCH_MY_CART = 'FETCH_MY_CART';
export const fetchMyCartAction = (cartList) => {
  return {
    type: FETCH_MY_CART,
    payload: cartList,
  };
};

export const FETCH_ORDERS_HISTORY = 'FETCH_ORDERS_HISTORY';
export const fetchOrdersHistoryAction = (history) => {
  return {
    type: FETCH_ORDERS_HISTORY,
    payload: history,
  };
};
