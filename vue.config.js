const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
  // devServer: {

  //   proxy: {
  //     // 和上面自定义的 baseURL 保持一致
  //     '/api': {
  //       target: 'http://localhost:8088',
  //       changeOrigin: true,
  //       pathRewrite: { '^/api': '' }
  //     }
  //   }
  // }
})
