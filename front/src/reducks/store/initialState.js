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
  },
  users: {
    cart: [],
    isLogedIn: false,
    orders: [],
    role: '',
    uid: '',
    id: '',
    name: '',
  },
};

export default initialState;
