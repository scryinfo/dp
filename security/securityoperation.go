package security


type SecurityOperation interface {
	Encrypt(src *[]byte) (dst *[]byte, err error)
	Decrypt(src *[]byte) (dst *[]byte, err error)
}

