package main

import (
	c "OS_lab4/Core"
	"fmt"
)

func main() {
	core := c.Core{}
	core.Mkfs(100)

	fmt.Println("\n=====================Test creation of FS and files=========================")
	core.Create("file.txt")
	core.Create("file.txt")
	core.Create("a.txt")
	core.Ls()
	core.Stat("file.txt")
	core.Stat("bbbbb.txt")

	fmt.Println("\n===========================Test link/unlink================================")
	core.Link("file.txt","file.txt")
	core.Link("file3123.txt","file1.txt")
	core.Link("file.txt","file2.txt")
	core.Stat("file.txt")
	core.Stat("file2.txt")
	core.Unlink("fileaaaa.txt")
	core.Unlink("file.txt")
	core.Stat("file2.txt")

	fmt.Println("\n============================Test open/close================================")
	fd := core.Open("file2.txt")
	core.Close(fd)

	fmt.Println("\n============================Test truncate==================================")
	core.Truncate("file2.txt",-10)
	core.Truncate("file2.txt",10)
	core.Stat("file2.txt")

	fmt.Println("\n==============================Test write/read==============================")
	fd = core.Open("file2.txt")
	str1 := []byte("20 len str is here !")
	str2 := []byte("10 len str")
	core.Write(fd, str1)
	core.Write(fd, str2)
	core.Read(fd, 10)
	core.Read(fd, 20)

	fmt.Println("\n=======================Test write/read with offset=========================")
	core.Truncate("file2.txt",40)
	str := []byte("This string contains 32 symbols 35!")
	core.Write(fd, str)
	core.Read(fd, 32)
	core.Read(fd, 35)
	core.Seek(fd, 40)
	core.Read(fd, 32)
	core.Seek(fd, 5)
	core.Read(fd, 5)
	core.Seek(fd, 30)
	core.Read(fd, 5)

	fmt.Println("====================Test offset < 0 and offset < size======================")
	core.Seek(fd, -10)
	core.Seek(fd, 100)
	fd = core.Close(fd)

	fmt.Println("\n=======================Test truncate with less size========================")
	fd = core.Open("file2.txt")
	core.Truncate("file2.txt", 65)
	str = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	core.Write(fd, str)
	core.Stat("file2.txt")
	core.Read(fd, 55)
	core.Truncate("file2.txt", 10)
	core.Read(fd, 10)
	core.Stat("file2.txt")
	core.Truncate("file2.txt", 200)
	core.Seek(fd, 100)
	core.Read(fd, 4)
	fd = core.Close(fd)

	fmt.Println("\n===================Test write/read after open and unlink===================")
	core.Create("unlink.txt")
	core.Truncate("unlink.txt", 23)
	fdd := core.Open("unlink.txt")
	fdd2 := core.Open("unlink.txt")
	core.Unlink("unlink.txt")
	core.Stat("unlink.txt")
	str = []byte("Content of deleted file")
	core.Write(fdd, str)
	core.Read(fdd, 23)
	fdd = core.Close(fdd)

	
	core.Seek(fdd2, 3)
	aaa := []byte("aaa")
	core.Write(fdd2, aaa)
	core.Read(fdd2, 20)
	fdd2 = core.Close(fdd2)
	//core.Read(fdd2, 23) Should give and error
}