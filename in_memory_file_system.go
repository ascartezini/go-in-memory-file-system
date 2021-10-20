package main

import (
	"strings"
)

type File struct {
	IsDir    bool
	Children map[string]*File
	Name     string
	Content  string
}

type InMemoryFileSystem struct {
	Root *File
}

func (fs InMemoryFileSystem) Ls(path string) []string {
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
	for k := range fileNode.Children {
		keys = append(keys, k)
	}

	return keys
}

func (fs InMemoryFileSystem) MkDir(path string) {
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
