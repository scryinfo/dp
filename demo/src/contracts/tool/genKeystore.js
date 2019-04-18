let Wallet = require('ethereumjs-wallet');
let key = Buffer.from('42c1b9105fe47c6862559041750a7616cdaad7d8fdef5db265370da5884498f7', 'hex');
let wallet = Wallet.fromPrivateKey(key);
let keystore = wallet.toV3String('111111');
console.log(keystore);
