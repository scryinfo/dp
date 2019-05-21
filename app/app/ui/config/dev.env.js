// Scry Info.  All rights reserved.
// license that can be found in the license file.

let merge = require('webpack-merge');
let prodEnv = require('./prod.env.js');

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"'
});
