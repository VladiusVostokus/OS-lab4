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

func (c *Core) Open(fileName, flags string) *fs.OpenFileDescriptor{
	if (!c.fs.Find(fileName)) {
		fmt.Println("Error: File",fileName,"to open does not exist")
		return nil
	}
	index := c.findFreeIndex()
	if (index == -1) {
		fmt.Println("No free descriptor indexes")
		return nil
	}
	fmt.Println("Open file", fileName)
	descriptor := c.fs.GetDescriptor(fileName)
	openFileDescriptor := &fs.OpenFileDescriptor{Desc: descriptor, Offset: 0, Flags: flags, Id: index}
	c.openFileDescriptors[index] = openFileDescriptor
	openFileDescriptor.Desc.Data = make(map[int]*fs.Block)
	return openFileDescriptor
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


func (c *Core) Close(fd *fs.OpenFileDescriptor) {
	if (fd == nil) {
		fmt.Println("Error: closing of non-existing file")
		return
	}
	fmt.Println("Closing file")
	c.openFileDescriptors[fd.Id] = nil
}

func (c *Core) Truncate(fileName string, size int) {
	if (size <= 0) {
		fmt.Println("Error: Incorrect size to truncate, must be bigger than 0")
		return
	}
	if (!c.fs.Find(fileName)) {
		fmt.Println("Error: File",fileName,"to truncate does not exist")
		return
	}
	descriptor := c.fs.GetDescriptor(fileName)
	descriptor.Size = size
}

func (c *Core) Read(fd *fs.OpenFileDescriptor, size int) {
	if (size <= 0) {
		fmt.Println("Error: Incorrect size to read, must be bigger than 0")
		return
	}
	if (size > fd.Desc.Size) {
		fmt.Println("Error: Incorrect size to read, must not be bigger than file size")
		return
	}
	curOffset := fd.Offset
	totalSize := size
	bytesToRead := 0
	res := ""
	for totalSize > 0 {
		curBlock := curOffset / 32
		offsetInsideBlock := curOffset % 32
		if (fd.Desc.Data[curBlock] == nil) {
			nullBlock := "00000000000000000000000000000000"
			res += nullBlock
			curOffset += 32
			totalSize -= 32
			continue
		}
		if (totalSize > (32 - offsetInsideBlock)) {
			bytesToRead = 32 - offsetInsideBlock
		} else {
			bytesToRead = totalSize
		}
		block := fd.Desc.Data[curBlock]
		for i := offsetInsideBlock; i < offsetInsideBlock + bytesToRead; i++ {
			res += string(block[i])
		}
		curOffset += bytesToRead
		totalSize -= bytesToRead
	}
	fmt.Println(res)
}

func (c *Core) Write(fd *fs.OpenFileDescriptor, size int) {
	if (size > fd.Desc.Size) {
		fmt.Println("Error: Incorrect size to write, must be less than file size")
		return
	}
	if (fd.Desc.Nblock == 0) {
		fd.Desc.Data = make(map[int]*fs.Block)
	}
	curOffset := fd.Offset
	totalSize := size
	bytesToWrite := 0
	for totalSize > 0 {
		curBlock := curOffset / 32
		offsetInsideBlock := curOffset % 32
		if (fd.Desc.Data[curBlock] == nil) {
			block := new(fs.Block)
			fd.Desc.Data[curBlock] = block
			fd.Desc.Nblock++
		}
		if (totalSize > (32 - offsetInsideBlock)) {
			bytesToWrite = 32 - offsetInsideBlock
		} else {
			bytesToWrite = totalSize
		}
		block := fd.Desc.Data[curBlock]
		for i := offsetInsideBlock; i < offsetInsideBlock + bytesToWrite; i++ {
			block[i] = 'a'
		}
		curOffset += bytesToWrite
		totalSize -= bytesToWrite
	}
}

func (c *Core) Seek(fd *fs.OpenFileDescriptor, offset int) {
	if (offset < 0) {
		fmt.Println("Error: Offset can not be less than 0")
		return
	}
	fd.Offset = offset
}
