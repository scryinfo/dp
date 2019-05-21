// Scry Info.  All rights reserved.
// license that can be found in the license file.

personal.newAccount("123456");
personal.unlockAccount(eth.accounts[0], "123456", 1000000);
miner.start(1);