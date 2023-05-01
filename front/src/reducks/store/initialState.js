// Store の初期状態を書いてる
const initialState = {
  utils: {
    windowSize: {
      width: window.innerWidth,
      height: window.innerHeight,
    },
    prefMaster: [],
    validation: {
      error: false,
      errorResult: null,
    }
  },
  message: {
    state: false,
    severity: '',
    text: [],
  },
  products: {
    list: [],
    product: null,
  },
  users: {
    isLogedIn: false,
    id: '',
    name: '',
    cartList: [],
    orderHistoryList: [],
  },
};

export default initialState;
