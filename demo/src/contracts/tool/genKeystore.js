var Wallet = require('ethereumjs-wallet');
var key = Buffer.from('42c1b9105fe47c6862559041750a7616cdaad7d8fdef5db265370da5884498f7', 'hex');
var wallet = Wallet.fromPrivateKey(key);
var keystore = wallet.toV3String('12345');
console.log(keystore);
