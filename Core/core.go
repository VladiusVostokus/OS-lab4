package core

import (
	fs "OS_lab4/FileSystem"
	"fmt"
)

type Core struct {
	fs *fs.FileSystem
	openFileDescriptors []*fs.OpenFileDescriptor
	blockSize int
}

func (c *Core) Mkfs (descriptorsCount int) {
	fmt.Println("System initialization...")
	c.openFileDescriptors = make([]*fs.OpenFileDescriptor, descriptorsCount)
	c.blockSize = fs.BlockSize
	fmt.Println("Create core with", descriptorsCount, "possible open file descpriptors")
	c.fs = &fs.FileSystem{}
	c.fs.Mkfs()
	fmt.Println("System is ready to work!")
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
	descriptor := c.fs.GetDescriptor(fileName)
	c.fs.Unlink(fileName)
	if (descriptor.Nlink == 0 && !descriptor.IsOpen) {
		descriptor = nil
	}
}

func (c *Core) Open(fileName string) *fs.OpenFileDescriptor{
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
	descriptor.IsOpen = true
	openFileDescriptor := &fs.OpenFileDescriptor{Desc: descriptor, Offset: 0, Id: index}
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


func (c *Core) Close(fd *fs.OpenFileDescriptor) *fs.OpenFileDescriptor {
	if (fd == nil) {
		fmt.Println("Error: closing of non-existing file")
		return nil
	}
	fmt.Println("Closing file")
	c.openFileDescriptors[fd.Id] = nil
	fd.Desc.IsOpen = false
	if(fd.Desc.Nlink == 0 && !fd.Desc.IsOpen) {
		fd.Desc = nil
	}
	return nil
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
	if (descriptor.Size > size) {
		newBlockCount := size / c.blockSize
		remainingBytes := size % c.blockSize
		if (remainingBytes > 0) {
			newBlockCount++
		}
		for i := newBlockCount; descriptor.Nblock > newBlockCount; i++ {
			if (descriptor.Data[i] == nil) {
				continue
			}
			delete(descriptor.Data, i)
			descriptor.Nblock--
		}
	}
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
		curBlock := curOffset / c.blockSize
		offsetInsideBlock := curOffset % c.blockSize
		if (totalSize > (c.blockSize - offsetInsideBlock)) {
			bytesToRead = c.blockSize - offsetInsideBlock
		} else {
			bytesToRead = totalSize
		}
		if (fd.Desc.Data[curBlock] == nil) {
			for i := 0; i < bytesToRead; i++ {
				res += "0"
			}
			curOffset += bytesToRead
			totalSize -= bytesToRead
			continue
		}
		block := fd.Desc.Data[curBlock]
		readTo := offsetInsideBlock + bytesToRead
		res += string(block[offsetInsideBlock:readTo])
		curOffset += bytesToRead
		totalSize -= bytesToRead
	}
	fmt.Println(res)
}

func (c *Core) Write(fd *fs.OpenFileDescriptor, data []byte) {
	totalSize := len(data)
	if (totalSize > fd.Desc.Size) {
		fmt.Println("Error: Incorrect size to write, must be less than file size")
		return
	}
	curOffset := fd.Offset
	bytesToWrite := 0
	for totalSize > 0 {
		curBlock := curOffset / c.blockSize
		offsetInsideBlock := curOffset % c.blockSize
		if (fd.Desc.Data[curBlock] == nil) {
			block := new(fs.Block)
			fd.Desc.Data[curBlock] = block
			fd.Desc.Nblock = len(fd.Desc.Data)
		}
		if (totalSize > (c.blockSize - offsetInsideBlock)) {
			bytesToWrite = c.blockSize - offsetInsideBlock
		} else {
			bytesToWrite = totalSize
		}
		block := fd.Desc.Data[curBlock]
		writeTo := offsetInsideBlock + bytesToWrite
		getDataFrom := curOffset - fd.Offset
		getDataTo := getDataFrom + bytesToWrite
		copy(block[offsetInsideBlock:writeTo], data[getDataFrom:getDataTo])
		curOffset += bytesToWrite
		totalSize -= bytesToWrite
	}
}

func (c *Core) Seek(fd *fs.OpenFileDescriptor, offset int) {
	if (offset < 0) {
		fmt.Println("Error: Offset can not be less than 0")
		return
	}
	if (offset > fd.Desc.Size) {
		fmt.Println("Error: Offset can not be bigger tnah file size")
		return
	}
	fd.Offset = offset
}
