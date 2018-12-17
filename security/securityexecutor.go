package security


type SecurityExecutor struct {
}

func (securityExecutor *SecurityExecutor) Encrypt(src *[]byte) (dst *[]byte, err error) {
	return src, nil
}

func (securityExecutor *SecurityExecutor) Decrypt(src *[]byte) (dst *[]byte, err error) {
	return src, nil
}