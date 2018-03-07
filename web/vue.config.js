module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8081',
        pathRewrite: { '^/api': '' },
        changeOrigin: true
      }
    }
  }
};
