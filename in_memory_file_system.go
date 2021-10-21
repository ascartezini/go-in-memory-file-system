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
	lastNode := fs.getLastNode(path)

	// reads all children filenodes of the last sub path
	keys := []string{}
	for k := range lastNode.Children {
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

			fileNode.Children[v] = &File{IsDir: true, Name: v}
		}

		// sets the current filenode as root
		fileNode = fileNode.Children[v]
	}
}

func (fs InMemoryFileSystem) WriteFile(path string, content string) {
	paths := strings.Split(path, "/")
	fileName := paths[len(paths)-1]
	newPath := strings.Replace(path, "/"+fileName, "", 1)
	lastNode := fs.getLastNode(newPath)

	if lastNode.Children == nil {
		lastNode.Children = map[string]*File{}
	}

	lastNode.Children[fileName] = &File{Name: fileName, Content: content}
}

func (fs InMemoryFileSystem) ReadFile(path string) *File {
	paths := strings.Split(path, "/")
	fileName := paths[len(paths)-1]
	newPath := strings.Replace(path, "/"+fileName, "", 1)
	lastNode := fs.getLastNode(newPath)

	return lastNode.Children[fileName]
}

func (fs InMemoryFileSystem) getLastNode(path string) *File {
	paths := strings.Split(path, "/")
	fileNode := fs.Root

	if path != "/" {
		// loops over the sub paths till it reaches the last one
		for i := 1; i < len(paths); i++ {
			// sets the current filenode as root
			fileNode = fileNode.Children[paths[i]]
		}
	}

	return fileNode
}
