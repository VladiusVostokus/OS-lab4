package main

import (
	c "OS_lab4/Core"
)

func main() {
	core := c.Core{}
	core.Mkfs(100)
	/*
	filesystem := fs.FileSystem{}
	filesystem.Mkfs()
	filesystem.Create("file.txt")
	filesystem.Create("a.txt")
	filesystem.Ls()
	filesystem.Stat("file.txt")
	filesystem.Link("file.txt","file2.txt")
	filesystem.Stat("file.txt")
	filesystem.Stat("file2.txt")
	filesystem.Unlink("file.txt")
	filesystem.Stat("file2.txt")
	*/
} 