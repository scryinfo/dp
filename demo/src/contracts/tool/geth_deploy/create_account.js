//set 0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8 as coinbase
personal.importRawKey("1a5037946a7d4717a6dcaa638995064495cae1912a32af8e0af9490232542647", "111111")
coinbase = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
console.log("coinbase:", coinbase)
personal.unlockAccount(coinbase, "111111", 1000000)

//start mining
miner.setEtherbase(coinbase)
miner.start(2)
