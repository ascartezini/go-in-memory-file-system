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

	if path != "/" {
		// loops over the sub paths till it reaches the last one
		for i := 1; i < len(paths); i++ {
			// sets the current filenode as root
			fileNode = fileNode.Children[paths[i]]
		}
	}

	// reads all children filenodes of the last sub path
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
	fileNode := fs.Root // file node starts with root

	for i := 1; i < len(paths); i++ {
		v := paths[i]

		// create a new entry in the map if the path does not exist
		if _, exists := fileNode.Children[v]; !exists {
			if fileNode.Children == nil {
				fileNode.Children = map[string]*File{}
			}

			fileNode.Children[v] = &File{IsDir: true}
		}

		// sets the current filenode as root
		fileNode = fileNode.Children[v]
	}
}

func main() {
	fs := FileSystem{&File{IsDir: true, Name: "/"}}
	fs.MkDir("/apps/golang/pointers")
	fs.MkDir("/apps/golang/concurrency")
	fs.MkDir("/apps/nodejs/streams")
	fs.MkDir("/apps/nodejs/event-loop")
	fs.MkDir("/docs")
	fs.MkDir("/images")
	fmt.Println(fs.Ls("/"))
}
