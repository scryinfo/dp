// Scry Info.  All rights reserved.
// license that can be found in the license file.

let utils = require('./utils.js');
let config = require('../config');
let isProduction = process.env.NODE_ENV === 'production';

module.exports = {
  loaders: utils.cssLoaders({
    sourceMap: isProduction
      ? config.build.productionSourceMap
      : config.dev.cssSourceMap,
    extract: isProduction
  })
};
