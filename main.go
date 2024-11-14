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
	/*
	filesystem.Stat("file.txt")
	filesystem.Link("file.txt","file2.txt")
	filesystem.Stat("file.txt")
	filesystem.Stat("file2.txt")
	filesystem.Unlink("file.txt")
	filesystem.Stat("file2.txt")
	*/
} 