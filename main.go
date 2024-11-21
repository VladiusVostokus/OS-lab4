package main

import (
	c "OS_lab4/Core"
	"fmt"
)

func main() {
	core := c.Core{}
	core.Mkfs(100)
	core.Create("file.txt")
	core.Create("file.txt")
	core.Create("a.txt")
	core.Ls()
	core.Stat("file.txt")
	core.Stat("bbbbb.txt")
	core.Link("file.txt","file.txt")
	core.Link("file3123.txt","file1.txt")
	core.Link("file.txt","file2.txt")
	core.Stat("file.txt")
	core.Stat("file2.txt")
	core.Unlink("fileaaaa.txt")
	core.Unlink("file.txt")
	core.Stat("file2.txt")
	fd := core.Open("file2.txt","rw")
	core.Close(fd)
	core.Truncate("file2.txt",-10)
	core.Truncate("file2.txt",10)
	core.Stat("file2.txt")

	fd = core.Open("file2.txt","rw")

	str1 := []byte("aaaaaaaaaaaaaaaaaaaa")
	str2 := []byte("10 len str")
	core.Write(fd, str1)
	core.Write(fd, str2)
	core.Read(fd, 10)
	core.Read(fd, 20)

	
	core.Truncate("file2.txt",40)
	str := []byte("This string contains 32 symbols 123")
	core.Write(fd, str)
	core.Read(fd, 35)
	core.Seek(fd, 40)
	core.Read(fd, 32)
	core.Seek(fd, 5)
	core.Read(fd, 5)
	core.Seek(fd, -10)
	core.Seek(fd, 100)
	core.Close(fd)

	fmt.Println("======================")
	fd = core.Open("file2.txt","rw")
	core.Seek(fd, 0)
	core.Truncate("file2.txt", 65)
	str = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	core.Write(fd, str)
	core.Stat("file2.txt")
	core.Read(fd, 55)
	core.Truncate("file2.txt", 10)
	core.Read(fd, 10)
	core.Stat("file2.txt")
	core.Close(fd)
} 