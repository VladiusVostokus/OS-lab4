package filesystem

type OpenFileDescriptor struct {
	desc *fileDescriptor
	offset int
	flags string
}