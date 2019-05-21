personal.newAccount("123456");
personal.unlockAccount(eth.accounts[0], "123456", 1000000);
miner.start(1);