package main

import (
	"fmt"
	"testing"
)

func TestInMemoryFileSystem_MkDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want map[string]bool
	}{
		{
			name: "root_folder_should_have_three_sub_folders",
			path: "/",
			want: map[string]bool{
				"apps":   true,
				"docs":   true,
				"images": true},
		},
		{
			name: "apps_folder_should_have_two_sub_folders",
			path: "/apps",
			want: map[string]bool{
				"golang": true,
				"nodejs": true},
		},
	}

	fs := InMemoryFileSystem{&File{IsDir: true, Name: "/"}}
	fs.MkDir("/apps/golang/pointers")
	fs.MkDir("/apps/golang/concurrency")
	fs.MkDir("/apps/nodejs/streams")
	fs.MkDir("/apps/nodejs/event-loop")
	fs.MkDir("/docs")
	fs.MkDir("/images")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fs.Ls(tt.path)
			want := tt.want
			errorMessage := fmt.Sprintf("got %v, want %v", got, want)

			if len(got) != len(want) {
				t.Errorf(errorMessage)
			}

			for _, v := range got {
				if _, exists := tt.want[v]; !exists {
					t.Errorf(errorMessage)
				}
			}
		})
	}
}
