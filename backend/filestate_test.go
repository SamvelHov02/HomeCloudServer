package backend

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTreeBuilder(t *testing.T) {
	tr := Tree{}
	tr.Init()
	tr.Build()

	expected := Tree{
		Root: "/Users/samvelhovhannisyan/Documents/dev/Personal/HomeCloud/server/Vault",
		Children: []Leaf{
			{Category: "dir", Name: VaultPath + "/TestFolder", Children: []Leaf{
				{Category: "file", Name: VaultPath + "/TestFolder/test3.md"},
			}},
			{Category: "file", Name: VaultPath + "/test1.md"},
			{Category: "file", Name: VaultPath + "/test2.md"},
		},
	}

	if !reflect.DeepEqual(tr, expected) {
		fmt.Printf("Expected %+v\n", expected)
		fmt.Println("--------------------------")
		fmt.Printf("Expected %+v\n", tr)
	}
}
