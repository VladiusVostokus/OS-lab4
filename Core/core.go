package core

import (
	fs "OS_lab4/FileSystem"
	"fmt"
)

type Core struct {
	fs *fs.FileSystem
	openFileDescriptors []*fs.OpenFileDescriptor
}

func (c *Core) Mkfs (descriptorsCount int) {
	fmt.Println("System initialization...")
	c.openFileDescriptors = make([]*fs.OpenFileDescriptor, descriptorsCount)
	fmt.Println("Create core with", descriptorsCount, "possible open file descpriptors")
	c.fs = &fs.FileSystem{}
	c.fs.Mkfs()
	fmt.Println("System is ready to work!\n")
}

func (c *Core) Create(fileName string) {
	if (c.fs.Find(fileName)) {
		fmt.Println("Error: File",fileName,"exist already")
		return
	}
	c.fs.Create(fileName)
}

func (c *Core) Ls() {
	c.fs.Ls()
}
