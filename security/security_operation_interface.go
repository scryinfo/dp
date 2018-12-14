package security


type SecurityOperationInterface interface {
	Encrypt(src *[]byte) (dst *[]byte, err error)
	Decrypt(src *[]byte) (dst *[]byte, err error)
}

