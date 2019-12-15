package unicornhat

type unicornError string

const (
	errBufferOverflow     unicornError = "buffer overflow"
	errIncompatibleDevice unicornError = "incompatible device size"
)

func (e unicornError) Error() string {
	return string(e)
}
