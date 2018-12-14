package transaction


type TransactionManager interface {
	CreateTransaction() (bool)
	CloseTransaction() (bool)
}