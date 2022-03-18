package main

import "fmt"

/**
Composite is a structural design pattern
that lets you compose objects into tree
structures and then work these structures
as if they were individual objects
Ex: files and folders
*/

// Component Interface
type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Search for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

// Composite
type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Search recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

// Leaf
type component interface {
	search(string)
}

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{
		name: "folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "folder 2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
