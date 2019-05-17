let Wallet = require('ethereumjs-wallet');

let privKey = Buffer.from('42c1b9105fe47c6862559041750a7616cdaad7d8fdef5db265370da5884498f7', 'hex');
let password = "111111";

console.log(Wallet.fromPrivateKey(privKey).toV3String(password));
