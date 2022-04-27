package notifier

type Notifier interface {
	Observe() error
}
