package protocol

type Protocol interface {
	Hit(data [][]byte) (name []byte, ok bool)
}
