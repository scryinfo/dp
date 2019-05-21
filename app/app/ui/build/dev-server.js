// Scry Info.  All rights reserved.
// license that can be found in the license file.

require('./check-versions')();

let config = require('../config');
if (!process.env.NODE_ENV) {
  process.env.NODE_ENV = JSON.parse(config.dev.env.NODE_ENV);
}

let path = require('path');
let webpack = require('webpack');
let webpackConfig = require('./webpack.dev.conf.js');

let compiler = webpack(webpackConfig);

let devMiddleware = require('webpack-dev-middleware')(compiler, {
  publicPath: webpackConfig.output.publicPath,
  quiet: true
});

let hotMiddleware = require('webpack-hot-middleware')(compiler, {
  log: () => {}
});
compiler.plugin('compilation', function (compilation) {
  compilation.plugin('html-webpack-plugin-after-emit', function (data, cb) {
    hotMiddleware.publish({ action: 'reload' });
    cb();
  });
});


let express = require('express');
let app = express();

app.use(devMiddleware);
app.use(hotMiddleware);

let staticPath = path.posix.join(config.dev.assetsPublicPath, config.dev.assetsSubDirectory);
app.use(staticPath, express.static('./static'));

let _resolve;
let readyPromise = new Promise(resolve => {
  _resolve = resolve;
});

let port = process.env.PORT || config.dev.port;

let opn = require('opn');
console.log('> Starting dev server...');
devMiddleware.waitUntilValid(() => {
  let url = 'http://localhost:' + port;
  console.log('> Listening at ' + url + '\n');
  if (!!config.dev.autoOpenBrowser && process.env.NODE_ENV !== 'testing') {
    opn(url);
  }
  _resolve();
});

let server = app.listen(port);

module.exports = {
  ready: readyPromise,
  close: () => {
    server.close();
  }
};
