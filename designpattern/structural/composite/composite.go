package composite

import "fmt"

type fileFolderComponent interface {
	display()
}

type file struct {
	name string
}

func (f *file) display()  {
	fmt.Println(f.name)
}

type folder struct {
	components []fileFolderComponent
	name string
}

func (f *folder) display()   {
	fmt.Println(f.name)
	for _, composite := range f.components {
		composite.display()
	}
}

func (f *folder) add(c fileFolderComponent)  {
	f.components = append(f.components, c)
}