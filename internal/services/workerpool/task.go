package workerpool

type Task interface {
	Execute() error
}
