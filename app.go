package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var curDir = []string{"/home/saman", "/Documents", "/samples" }
var dir_items = []File{}

type File struct {
	Name string `json:"name"`
	Is_directory bool `json:"is_directory"`
}

func CreateFileItem(i os.DirEntry) File {
	f := File{
		Name: i.Name(),
		Is_directory: i.IsDir(),
	}

	return f
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CurrentDir() []string {
	return curDir
}

func RenderDir(i []string) string {
	x := ""
	for _, e := range i {
		x = x + e
	}
	return x
}

func (a *App) ReturnDirItems(dir []string) string {
	itemNames := []string{}
	entries, err := os.ReadDir(RenderDir(dir))
	check(err)

	dir_items = []File{}

	for _, e := range entries {
		dir_items = append(dir_items, CreateFileItem(e))
		itemNames = append(itemNames, e.Name())
	}

	data, err := json.Marshal(dir_items)
	check(err)
	return string(data)
}

func (a* App) MoveDir(dir string) {
	if (slices.Contains(curDir, dir)) {
		index := Index(curDir, dir)
		curDir = curDir[:index + 1]
	} else {
		if (slices.Contains(dir_items, File{Name: dir, Is_directory: true})) {
			curDir = append(curDir, "/" + dir)
		} else {
			fmt.Println(dir + " is not a directory")
		}
	}
}

/* func CheckDir(i string) bool {
	file, err := os.Open(i)
	check(err);

	fileInfo, err := file.Stat()
	check(err)

	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
} */

func Index(s []string, value string) int {
	for i := range s {
		if value == s[i] {
			return i
		}
	}
	return -1
}
