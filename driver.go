package unicornhat

type Driver interface {
	Render(buffer []byte) error
	Close() error
}
