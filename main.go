package main

import (
	c "OS_lab4/Core"
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
	core.Write(fd, 20)
	core.Write(fd, 10)
	core.Read(fd, 10)
	core.Read(fd, 20)

	core.Truncate("file2.txt",40)
	core.Write(fd, 20)
	core.Read(fd, 35)
	core.Read(fd, 30)
	core.Close(fd)
} 