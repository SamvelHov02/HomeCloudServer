package backend

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const VaultPath = "/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/server/"

type Tree struct {
	Root     string `json:"root"`
	RootHash string `json:"hash"`
	Children []Leaf `json:"children"`
}

type Leaf struct {
	Category string `json:"type"`
	Children []Leaf `json:"children"`
	Name     string `json:"name"`
	Hash     string `json:"hash"`
}

func (t *Tree) Init() {
	t.Root = "Vault"
}

func (t *Tree) Build() {
	// Want to recurse the Vault for each level create a Leaf
	t.Children = []Leaf{}

	contents, err := os.ReadDir(t.Root)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range contents {
		path := filepath.Join(t.Root, entry.Name())

		if entry.IsDir() {
			l := &Leaf{
				Category: "dir",
				Name:     entry.Name(),
				Children: []Leaf{},
			}
			t.Children = append(t.Children, *GetAllNames(path, l))
		} else {
			t.Children = append(t.Children, Leaf{
				Category: "file",
				Name:     entry.Name(),
			})
		}
	}

}

func Test() {
	fmt.Println("Test")
}

func GetAllNames(root string, l *Leaf) *Leaf {
	l.Children = []Leaf{}
	contents, err := os.ReadDir(root)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range contents {
		path := filepath.Join(root, entry.Name())

		if entry.IsDir() {
			l2 := &Leaf{
				Category: "dir",
				Name:     entry.Name(),
				Children: []Leaf{},
			}
			l.Children = append(l.Children, *GetAllNames(path, l2))
		} else {
			l.Children = append(l.Children, Leaf{
				Category: "file",
				Name:     entry.Name(),
			})
		}
	}
	return l
}

func TurnFilesToLeafs(files []string) []Leaf {
	var Leaves []Leaf

	for _, file := range files {
		Leaves = append(Leaves, Leaf{Category: "file", Name: file})
	}
	return Leaves
}
