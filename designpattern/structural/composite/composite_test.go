package composite

import "testing"

func TestFactory(t *testing.T) {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{
		name: "/",
	}
	folder1.add(file1)

	folder2 := &folder{
		name: "/sub",
	}
	folder2.add(file2)
	folder2.add(file3)

	folder1.add(folder2)

	folder1.display()

}

