// webpack.config.js
var webpack = require('webpack');

module.exports = {
  entry: './client/src/main.js',
  output: {
    // path: './server/assets/',
    filename: 'build.js'
  },
  module: {
    loaders: [
      { test: /\.vue$/, loader: 'vue' },
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({
      $: "jquery",
    })
  ]
}
