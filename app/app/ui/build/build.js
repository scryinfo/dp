// Scry Info.  All rights reserved.
// license that can be found in the license file.

require('./check-versions')();

process.env.NODE_ENV = 'production';

let ora = require('ora');
let rm = require('rimraf');
let path = require('path');
let config = require('../config');
let webpack = require('webpack');
let webpackConfig = require('./webpack.prod.conf.js');
let chalk = require('chalk');

let spinner = ora('building for production...');
spinner.start();

rm(path.join(config.build.assetsRoot, config.build.assetsSubDirectory), err => {
  if (err) throw err;
  webpack(webpackConfig, function (err, stats) {
    spinner.stop();
    if (err) throw err;
    process.stdout.write(stats.toString({
      colors: true,
      modules: false,
      children: false,
      chunks: false,
      chunkModules: false
    }) + '\n\n');

    console.log(chalk.cyan('  Build complete.\n'))
  });
});
