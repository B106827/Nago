module.exports = {
  root: true,
  // 環境特有の変数を理解できるようにする
  env: {
    // DOM API を理解できるように( document や onload 等)
    browser: true,
    // Node.js 固有の変数を理解できるように( process や require 等)
    node: true,
    // ES6 など新しい JavaScript の書き方を理解できるように( Promise 等)
    es6: true,
  },
  parser: 'babel-eslint',
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    sourceType: 'module',
  },
  extends: ['eslint:recommended', 'prettier'],
  plugins: ['react', 'react-hooks'],
  rules: {
    // 0:チェックしない、1:警告、2:エラー
    // 宣言しているにも関わらず使用していない変数
    'no-unused-vars': process.env.REACT_APP_ENV === 'production' ? 2 : 1,
    // React 独自表記はサポート外なので下記を追記
    'react/jsx-uses-vars': 1,
    // debuggerステートメントの記述を許可しない。
    'no-debugger': process.env.REACT_APP_ENV === 'production' ? 2 : 0,
    // let と const の区別を厳格化
    'prefer-const': process.env.REACT_APP_ENV === 'production' ? 2 : 1,
    // === の使用を推奨
    'eqeqeq': [
      process.env.REACT_APP_ENV === 'production' ? 2 : 1,
      'always',
      { null: 'ignore' },
    ],
  },
};
