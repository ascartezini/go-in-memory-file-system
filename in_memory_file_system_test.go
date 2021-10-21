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
		{
			name: "golang_folder_should_have_two_sub_folders",
			path: "/apps/golang",
			want: map[string]bool{
				"pointers":    true,
				"concurrency": true},
		},
		{
			name: "pointers_folder_should_have_no_sub_folders",
			path: "/apps/golang/pointers",
			want: map[string]bool{},
		},
		{
			name: "concurrency_folder_should_have_no_sub_folders",
			path: "/apps/golang/concurrency",
			want: map[string]bool{},
		},
		{
			name: "nodejs_folder_should_have_two_sub_folders",
			path: "/apps/nodejs",
			want: map[string]bool{
				"streams":    true,
				"event-loop": true},
		},
		{
			name: "streams_folder_should_have_no_sub_folders",
			path: "/apps/nodejs/streams",
			want: map[string]bool{},
		},
		{
			name: "event-loop_folder_should_have_no_sub_folders",
			path: "/apps/nodejs/event-loop",
			want: map[string]bool{},
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

func TestInMemoryFileSystem_WriteFile(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		wantFileName    string
		wantFileContent string
	}{
		{
			name:            "golang_folder_should_have_file_named_main.go",
			path:            "/golang/main.go",
			wantFileName:    "main.go",
			wantFileContent: "package main",
		},
		{
			name:            "docs_folder_should_have_file_named_help.txt",
			path:            "/docs/help.txt",
			wantFileName:    "help.txt",
			wantFileContent: "Golang help",
		},
	}

	fs := InMemoryFileSystem{&File{IsDir: true, Name: "/"}}
	fs.MkDir("/golang")
	fs.WriteFile("/golang/main.go", "package main")

	fs.MkDir("/docs")
	fs.WriteFile("/docs/help.txt", "Golang help")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fs.ReadFile(tt.path)

			if got.Name != tt.wantFileName {
				t.Errorf(fmt.Sprintf(`got "%s", want "%s"`, got.Name, tt.wantFileName))
			}

			if got.Content != tt.wantFileContent {
				t.Errorf(fmt.Sprintf(`got "%s", want "%s"`, got.Content, tt.wantFileContent))
			}

		})
	}
}
