package main

import "fmt"

//Prototype interface
type iNode interface {
	print(string)
	clone() iNode
}

//Concrete Prototype File
type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}
func (f *file) clone() iNode {
	return &file{name: f.name + "_clone"}
}

//Concrete Prototype Folder
type folder struct {
	children []iNode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() iNode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []iNode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []iNode{file1},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []iNode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("--")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("--")
}
