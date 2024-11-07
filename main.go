package main

import (
	fs "OS_lab4/FileSystem"
)

func main() {
	filesystem := fs.FileSystem{}
	filesystem.Mkfs(100)
	filesystem.Create("file.txt")
	filesystem.Create("a.txt")
	filesystem.Ls()
	filesystem.Stat("file.txt")
	filesystem.Link("file.txt","file2.txt")
	filesystem.Stat("file.txt")
	filesystem.Stat("file2.txt")
	filesystem.Unlink("file.txt")
	filesystem.Stat("file2.txt")
} 