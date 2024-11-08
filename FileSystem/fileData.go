package filesystem

const blockSize int = 32

type block [blockSize]byte

type fileData struct {
	data []block
	offset int
}