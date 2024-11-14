package filesystem

type fileDescriptor struct{
	FileType string
	Nlink, Size, Id int
	Data *fileData
}