# 同步 vitePress 的文件夹结构

```go

package main

import (
	"DirStructEcho/Serialization"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func main() {
	rootPath := os.Args[1]
	//rootPath = " ~/workspace/document/docs/docs/"
	var result = GetDirModel(rootPath, "./")
	text := result.Print("/Skill/")
	fmt.Println(text)
}

func Debug(s interface{}) {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func GetDirModel(rootPath string, relative string) DirModel {

	var dirName = path.Base(rootPath)
	var result DirModel = DirModel{
		Name:     dirName,
		Items:    []FileModel{},
		Children: []DirModel{},
	}
	result.Relative = path.Join(relative, dirName)

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return result
	}

	for _, file := range files {
		if file.IsDir() {
			var dirPath = path.Join(rootPath, file.Name())
			var currentDir = GetDirModel(dirPath, path.Join(relative, file.Name()))
			result.Children = append(result.Children, currentDir)
		} else {
			var filePath = path.Join(rootPath, file.Name())
			var fileItem = FileModel{
				Name:     file.Name(),
				Path:     filePath,
				Relative: path.Join(relative, file.Name()),
			}
			result.Items = append(result.Items, fileItem)
		}
	}

	return result
}

type DirModel struct {
	Name     string      `json:"name"`
	Items    []FileModel `json:"items"`
	Children []DirModel  `json:"children"`
	Relative string      `json:"relative"`
}

type FileModel struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Relative string `json:"relative"`
}

type VitePress struct {
	Title string          `json:"text"`
	Items []VitePressItem `json:"items"`
}

type VitePressItem struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func (dir DirModel) Print(path string) string {

	var result = []VitePress{}

	var rootVitePress = VitePress{
		Title: dir.Name,
		Items: []VitePressItem{},
	}
	var rootItems = DirModelToVitePressItems(dir, path)
	rootVitePress.Items = append(rootVitePress.Items, rootItems...)

	for _, child := range dir.Children {
		var vitePress = VitePress{
			Title: child.Name,
			Items: []VitePressItem{},
		}
		var childItems = DirModelToVitePressItems(child, path)
		vitePress.Items = append(vitePress.Items, childItems...)
		result = append(result, vitePress)
	}

	var jsonData, err = Serialization.SerializeJson(result)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return jsonData
}

func DirModelToVitePressItems(model DirModel, relative string) []VitePressItem {
	var result []VitePressItem
	for _, item := range model.Items {
		var title = ReadMarkdownFirstTitle(item.Path)
		var vitePressItem = VitePressItem{
			Text: title,
			Link: relative + item.Relative,
		}
		result = append(result, vitePressItem)
	}

	for _, child := range model.Children {
		var childItems = DirModelToVitePressItems(child, relative)
		result = append(result, childItems...)
	}

	return result
}

func ReadMarkdownFirstTitle(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return path.Base(filePath)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		if line[0] == '#' {
			return line[1:]
		}
	}

	return path.Base(filePath)
}

func isMarkdownFile(dirPath string) bool {
	fileExtension := path.Ext(dirPath)
	if fileExtension == ".md" {
		return true
	}
	return false
}

```