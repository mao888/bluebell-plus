module.exports = {
    assetsDir: "static",
    devServer: {
        proxy: {
            '/api/v1': {
              target: 'http://127.0.0.1:8081',
              changeOrigin: true,
            }
        }
    }
  }
