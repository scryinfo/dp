package accountsmanager

type AccountBehavior interface {
    CreateAccount(password string) (string, error)
    DestroyAccount(pubkey string, password string) (error)
    SignTransaction(message []byte, pubkey string, password string) ([]byte, error)
    Encrypt(plainText []byte, pubkey string, password string) ([]byte, error)
    ReEncryptWithAnother(ciphertext []byte, pubkey1 string, pubkey2 string, password string) ([]byte, error)
}

type AccountsManager struct {
    accounts []string
}

func NewAccountManager() (*AccountsManager) {
    return new(AccountsManager)
}

func (am *AccountsManager) CreateAccount(password string) (string, error) {

}

func (am *AccountsManager) DestroyAccount(pubkey string, password string) (error) {

}

func (am AccountsManager) SignTransaction(message []byte, pubkey string, password string) ([]byte, error)

}

func (am AccountsManager) Encrypt(plainText []byte, pubkey string, password string) ([]byte, error)
func (am AccountsManager) ReEncryptWithAnother(ciphertext []byte, pubkey1 string, pubkey2 string, password string) ([]byte, error)