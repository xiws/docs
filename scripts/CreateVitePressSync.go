package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Text  string `json:"text"`
	Link  string `json:"link,omitempty"`
	Items []Item `json:"items,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: CreateVitePressSync <path>")
		return
	}
	root := os.Args[1]
	sidebar := createSidebar(root)
	output, err := json.MarshalIndent(sidebar, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(root, ".vitepress/SideBar.mts"), []byte("export default "+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func createSidebar(root string) []Item {
	var items []Item
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return items
	}

	for _, file := range files {
		if file.IsDir() {
			subItems := createSidebar(filepath.Join(root, file.Name()))
			items = append(items, Item{
				Text:  file.Name(),
				Items: subItems,
			})
		} else if strings.HasSuffix(file.Name(), ".md") {
			items = append(items, Item{
				Text: strings.TrimSuffix(file.Name(), ".md"),
				Link: filepath.Join(root, file.Name()),
			})
		}
	}
	return items
}
