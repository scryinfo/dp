// Scry Info.  All rights reserved.
// license that can be found in the license file.

let chalk = require('chalk');
let semver = require('semver');
let packageConfig = require('../package.json');
let shell = require('shelljs');

function exec (cmd) {
  return require('child_process').execSync(cmd).toString().trim();
}

let versionRequirements = [
  {
    name: 'node',
    currentVersion: semver.clean(process.version),
    versionRequirement: packageConfig.engines.node
  }
];

if (shell.which('npm')) {
  versionRequirements.push({
    name: 'npm',
    currentVersion: exec('npm --version'),
    versionRequirement: packageConfig.engines.npm
  });
}

module.exports = function () {
  let warnings = [];
  for (let i = 0; i < versionRequirements.length; i++) {
    let mod = versionRequirements[i];
    if (!semver.satisfies(mod.currentVersion, mod.versionRequirement)) {
      warnings.push(mod.name + ': ' +
        chalk.red(mod.currentVersion) + ' should be ' +
        chalk.green(mod.versionRequirement)
      )
    }
  }

  if (warnings.length) {
    console.log('');
    console.log(chalk.yellow('To use this template, you must update following to modules:'));
    console.log();
    for (let i = 0; i < warnings.length; i++) {
      let warning = warnings[i];
      console.log('  ' + warning);
    }
    console.log();
    process.exit(1);
  }
};
