package main

import "fmt"

func main() {
	file := File{
		name: "file-1",
	}
	file.printName()

	clone := file.clone()
	clone.printName()
}

type Inode interface {
	printName()
	clone() Inode
}

type File struct {
	name string
}

func (f *File) printName() {
	fmt.Println(f.name)
}

func (f *File) clone() Inode {
	return &File{f.name + "_name"}
}
