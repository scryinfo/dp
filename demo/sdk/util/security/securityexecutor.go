package security

import (
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/crypto"
    "strings"
)

type CryptExecutor struct {
}

func (ce CryptExecutor) SignTransaction(message []byte, pubkey string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    keyJson := ""
    pubkey = strings.ToLower(pubkey)

    if pubkey == "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8" {
        keyJson = `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
    }

    if pubkey == "0x2d13d4faba031e66a36a6b307fce2087db55c43d" {
        keyJson = `{"version":3,"id":"7f01defb-3543-4459-bcc3-3b86197a4e17","address":"2d13d4faba031e66a36a6b307fce2087db55c43d","crypto":{"ciphertext":"ce55ddca8d430b2aef68a5fafda00f37a6df0aad45a6fa50c67920722c5a06b7","cipherparams":{"iv":"68dbb3ec9294ed20129807e46da74d72"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"b88dfead45a58f8d73645a7c458f4130cd51d4511f8b5861a756c840dc50bd66","n":262144,"r":8,"p":1},"mac":"6d90b054f4a1c8e3278fc47119c684fbb63257d3726df26b6f69fcd3f2f087dc"}}`
    }

    if pubkey == "0x3d6f42489e5283c95af70169373d85ba5799bb6f" {
        keyJson = `{"version":3,"id":"aaf00f30-3689-499f-9bd3-3ce9fc02c731","address":"3d6f42489e5283c95af70169373d85ba5799bb6f","crypto":{"ciphertext":"d02ee56a66f7688e761bb70b48bee7fdf4c6b1e78e34bc412756a0d94247e8b6","cipherparams":{"iv":"86b9f6ec3f8ae68c4801925b706814ca"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"902a7aec29df2ffe54fc455dfc6c4addc42e02aa495c891fed196c0021d88b0e","n":262144,"r":8,"p":1},"mac":"d2dc1bdd173c8d63403e60838116a985b44e1f3912784316de352bdb99cba471"}}`
    }

    if pubkey == "0xaef420b44068f363fa9905f3fa2d3eb047d8570c" {
        keyJson = `{"version":3,"id":"8d13ca05-0c86-40f4-b79e-d13735b8424e","address":"aef420b44068f363fa9905f3fa2d3eb047d8570c","crypto":{"ciphertext":"43237f0c2eebc3a78cacb99134b2a5b2ea436a98415e5d15781c3b6ceb0d1b29","cipherparams":{"iv":"0bd59781851b2dca98f79f1db06519b0"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"2bad47416714e60e41d81437ab5d1c1007842405722eb01826fc7a41ac790b1d","n":262144,"r":8,"p":1},"mac":"e1afb059ca485604c3bed6e768386b15c2722921429d06cd4930069c8c0ca129"}}`
    }

    if pubkey == "0x8c091c18bf57db0896d077b1b778301cab48bc37" {
        keyJson = `{"version":3,"id":"857cb57c-cdcb-4a19-a3c2-fe5a2ea3b3c2","address":"8c091c18bf57db0896d077b1b778301cab48bc37","crypto":{"ciphertext":"90141512a4b938921f67f943426939b20a21bef08cb9767f09a0ee3bbf33f8f7","cipherparams":{"iv":"568781bb771bb2aedef69a21eedb3e4d"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"c8740dd4d4ba72b5a513bb8986342a47ac93c9c89ea2bb9420d924d78243b343","n":262144,"r":8,"p":1},"mac":"1760cc87aa8eddb269e582cea5213db55c5f615f3527b2941d1d6b6b2ada200d"}}`
    }

    if pubkey == "0x30c04b0ded7c042c09ad884bdcb8ddb38e536f0e" {
        keyJson = `{"version":3,"id":"2ffd7c1d-e948-4f44-ade4-10b8a6619e0a","address":"30c04b0ded7c042c09ad884bdcb8ddb38e536f0e","crypto":{"ciphertext":"8cd9db8c430dd55fbeb56bfd93c91033972229b783c678f6864bcaf4d1291723","cipherparams":{"iv":"471f25a18e9cb8eddf1965d31f2989b8"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"ba77cda872eb0ef19882c099ec47f33dbfc8a27335cb99d1ede4c2ea1963a5a0","n":262144,"r":8,"p":1},"mac":"b2d34b7e879effbe30185e9f6222607733741e94f04444caa608496f4dcbfcad"}}`
    }

    key, err := keystore.DecryptKey([]byte(keyJson), "12345")
    if err != nil {
        return nil, err
    }

    sig, err := crypto.Sign(message, key.PrivateKey)
    return sig, err
}

func (ce CryptExecutor) Encrypt(plainText []byte, pubKey string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    return plainText, nil
}

func (ce CryptExecutor) ReEncryptWithAnotherPubKey(cipherText []byte, pubKey1 string,
                                        pubKey2 string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    return cipherText, nil
}

