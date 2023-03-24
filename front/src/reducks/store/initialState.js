// Store の初期状態を書いてる
const initialState = {
  utils: {
    windowSize: {
      width: window.innerWidth,
      height: window.innerHeight,
    },
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
    cartList: [],
    isLogedIn: false,
    orders: [],
    id: '',
    name: '',
  },
};

export default initialState;
