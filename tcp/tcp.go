package tcp

type Tcp interface {
	Hit(data [][]byte) (name []byte, ok bool)
}
