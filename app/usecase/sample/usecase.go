package sample

type UseCase interface {
	MultiplicationTable(in int) (out map[string]interface{}, err error)
}
