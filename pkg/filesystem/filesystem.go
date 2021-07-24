package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/a8m/tree"
	"github.com/a8m/tree/ostree"
)

type resource struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}

type resourceList struct {
	List []resource
}

func Test(filename string, resourceType string, content []byte) resource {
	r := createResource(filename, resourceType, content)

	return r
}
func createResource(filename string, resourceType string, content []byte) resource {

	r := &resource{
		Name:    filename,
		Type:    resourceType,
		Content: content}
	return *r
	// rjson, _ := json.Marshal(r)
	// fmt.Println(string(rjson))
	// jsondat := &resourceList{List: []resource{*r, *r}}
	// encjson, _ := json.Marshal(jsondat)
	// fmt.Println(string(encjson))
}

func Test2(r resource) {
	// var resources *resourceList
	// resources = new(resourceList)

	rl := []resource{}
	resources := resourceList{rl}

	resources.Addresource(r)

	WriteFSFile(resources)
}
func (resources *resourceList) Addresource(r resource) []resource {
	resources.List = append(resources.List, r)
	return resources.List
}

func WriteFSFile(resources resourceList) {
	file, _ := json.MarshalIndent(resources, "", " ")
	// file, _ := json.Marshal(resources)

	_ = ioutil.WriteFile("test.arafs", file, 0644)
}

func MyVisitTree() {
	tr := tree.New("./test") //node
	opts := &tree.Options{
		Fs: new(ostree.FS),
	}
	dir, files := tr.Visit(opts)
	fmt.Println(dir)
	fmt.Println(files)
}

func MyVisitWalk() {
	nodes := make(map[string]bool)
	// err := filepath.Walk("./test",
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(path, info.Size(), info.IsDir())
	// 		return nil
	// 	})
	// if err != nil {
	// 	log.Println(err)
	// }
	err := filepath.Walk("./test",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			nodes[path] = info.IsDir()
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	for node, isDir := range nodes {
		fmt.Println("node:", node, "=>", "isDir:", isDir)
	}
}
