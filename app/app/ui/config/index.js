// Scry Info.  All rights reserved.
// license that can be found in the license file.

let path = require('path');

module.exports = {
  build: {
    env: require('./prod.env.js'),
    index: path.resolve(__dirname, '../resources/app/index.html'),
    assetsRoot: path.resolve(__dirname, '../resources/app'),
    assetsSubDirectory: 'static',
    assetsPublicPath: './',
    productionSourceMap: true,
    // Before definition to `true`, make sure to:
    // npm install --save-dev compression-webpack-plugin
    productionGzip: false,
    productionGzipExtensions: ['js', 'css'],
    bundleAnalyzerReport: false // control show analyzer report or not.
  },
  dev: {
    env: require('./dev.env.js'),
    port: 9001,
    autoOpenBrowser: true,
    assetsSubDirectory: 'static',
    assetsPublicPath: '/',
    cssSourceMap: false
  }
};
