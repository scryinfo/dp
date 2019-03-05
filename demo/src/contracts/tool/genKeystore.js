var Wallet = require('ethereumjs-wallet');
var key = Buffer.from('1a5037946a7d4717a6dcaa638995064495cae1912a32af8e0af9490232542647', 'hex');
var wallet = Wallet.fromPrivateKey(key);
var keystore = wallet.toV3String('12345');
console.log(keystore);
