package main

import (
	fs "OS_lab4/FileSystem"
)

func main() {
	filesystem := fs.FileSystem{}
	filesystem.Mkfs(100)
} 