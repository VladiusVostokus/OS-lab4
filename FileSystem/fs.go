package filesystem

import "fmt"

type FileSystem struct {
	descriptorsCount int
	descriptors []int
}

func (fs * FileSystem) Mkfs (descriptorsCount int) {
	fs.descriptorsCount = descriptorsCount
	fs.descriptors = make([]int, fs.descriptorsCount)
	for i := 0; i < fs.descriptorsCount; i++ {
		fs.descriptors[i] = i
	}
	fmt.Println("Creating file system with", fs.descriptorsCount, "descriptors")
}