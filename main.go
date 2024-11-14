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
	/*
	filesystem.Unlink("file.txt")
	filesystem.Stat("file2.txt")
	*/
} 