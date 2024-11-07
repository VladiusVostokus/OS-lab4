package filesystem

import "fmt"

type FileSystem struct {
	descriptorsCount int
	descriptors []int
	directory map[string]*fileDescriptor
}

func (fs * FileSystem) Mkfs (descriptorsCount int) {
	fs.directory = make(map[string]*fileDescriptor)
	fs.descriptorsCount = descriptorsCount
	fs.descriptors = make([]int, fs.descriptorsCount)
	for i := 0; i < fs.descriptorsCount; i++ {
		fs.descriptors[i] = i
	}
	fmt.Println("Creating file system with", fs.descriptorsCount, "descriptors")
}

func (fs* FileSystem) Create (fileName string) {
	descriptor := &fileDescriptor{ FileType:"reg", Nlink: 1, Size: 0, Id: 0}
	fs.directory[fileName] = descriptor
	fmt.Println("Create file:", fileName,"| Descriptor id:", descriptor.Id)
}