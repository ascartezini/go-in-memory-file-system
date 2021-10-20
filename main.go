package main

import (
	"fmt"
	"strings"
)

type File struct {
	IsDir    bool
	Children map[string]*File
	Name     string
	Content  string
}

type FileSystem struct {
	Root *File
}

func (fs FileSystem) Ls(path string) []string {
	paths := strings.Split(path, "/")
	fileNode := fs.Root

	for i := 1; i < len(paths); i++ {
		fileNode = fileNode.Children[paths[i]]
	}

	keys := []string{}
	for k, v := range fileNode.Children {
		keys = append(keys, k)
		doNothing(v)
	}

	return keys
}

func doNothing(v interface{}) {}

func (fs FileSystem) MkDir(path string) {
	paths := strings.Split(path, "/")
	fileNode := fs.Root

	for i := 1; i < len(paths); i++ {
		v := paths[i]

		if _, exists := fileNode.Children[v]; !exists {
			if fileNode.Children == nil {
				fileNode.Children = map[string]*File{}
			}

			fileNode.Children[v] = &File{IsDir: true}
		}

		fileNode = fileNode.Children[v]
	}
}

func main() {
	fs := FileSystem{&File{IsDir: true, Name: "/"}}
	fs.MkDir("/apps/golang/gorocks")
	fs.MkDir("/apps/golang/go-concurrency")
	fmt.Println(fs.Ls("/apps/golang"))
}
