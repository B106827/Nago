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

export const FETCH_ORDERS_HISTORY = 'FETCH_ORDERS_HISTORY';
export const fetchOrdersHistoryAction = (history) => {
  return {
    type: FETCH_ORDERS_HISTORY,
    payload: history,
  };
};
