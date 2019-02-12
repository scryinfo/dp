var Wallet = require('ethereumjs-wallet');
var key = Buffer.from('2d2328d9c7d762d52188ee972ae04a6524ec029be005195197059a8c4a04cb13', 'hex');
var wallet = Wallet.fromPrivateKey(key);
var keystore = wallet.toV3String('12345');
console.log(keystore);
