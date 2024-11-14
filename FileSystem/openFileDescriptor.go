package filesystem

type OpenFileDescriptor struct {
	Desc *fileDescriptor
	Offset int
	Flags string
}