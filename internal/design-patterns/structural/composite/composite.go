package composite

import "fmt"

type component interface {
	search(string)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("search in file %s", f.name)
}

type folder struct {
	name       string
	components []component
}

func (f *folder) search(keyword string) {
	fmt.Printf("search in folder %s", f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

func Run() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
