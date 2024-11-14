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

func (c *Core) Stat(fileName string) {
	if (c.fs.Find(fileName)) {
		c.fs.Stat(fileName)
		return
	}
	fmt.Println("Error: File",fileName,"does not exist")
}

func (c *Core) Link(linkWith, toLink string) {
	if (c.fs.Find(toLink)) {
		fmt.Println("Error: File",toLink,"to create link exist already")
		return
	}
	if (!c.fs.Find(linkWith)) {
		fmt.Println("Error: File ",toLink,"to create link with does not exist")
		return
	}
	c.fs.Link(linkWith, toLink)
}

func (c *Core) Unlink(fileName string) {
	if (!c.fs.Find(fileName)) {
		fmt.Println("Error: File",fileName,"to delete does not exist")
		return
	}
	c.fs.Unlink(fileName)
}

func (c *Core) Open(fileName, flags string)  {
	if (!c.fs.Find(fileName)) {
		fmt.Println("Error: File",fileName,"to open does not exist")
		return
	}
	fmt.Println("Open file", fileName)
	descriptor := c.fs.GetDescriptor(fileName)
	openFileDescriptor := &fs.OpenFileDescriptor{Desc: descriptor, Offset: 0, Flags: flags}
	index := c.findFreeIndex()
	if (index == -1) {
		fmt.Println("No free descriptor indexes")
		return
	}
	c.openFileDescriptors[index] = openFileDescriptor
}

func (c *Core) findFreeIndex() int {
	freeIndex := -1
	for i, v := range c.openFileDescriptors {
		if (v == nil) {
			freeIndex = i
			break
		}
	}
	return freeIndex
}
