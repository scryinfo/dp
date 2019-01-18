package security


type CryptoBehavior interface {
    SignTransaction(message []byte, pubkey string, password string) ([]byte, error)
    Encrypt(plainText []byte, pubkey string, password string) ([]byte, error)
    ReEncryptWithAnotherPubkey(ciphertext []byte, pubkey1 string, pubkey2 string, password string) ([]byte, error)
}

