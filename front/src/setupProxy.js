const { createProxyMiddleware } = require('http-proxy-middleware');

if (process.env.REACT_APP_ENV === 'development') {
  module.exports = function (app) {
    app.use(
      '/api',
      createProxyMiddleware({
        target: 'http://nago-backend:8081',
        changeOrigin: true,
      })
    );
  };
}
