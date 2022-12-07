package day7

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type File struct {
	name string
	size int
}

type Folder struct {
	name    string
	files   map[string]*File
	folders map[string]*Folder
	parent  *Folder
	size    int
}

// 1611443
func PartA(input []byte) {
	tree := buildTree(input)
	folderSizes := []int{}
	tree.folderSizes(&folderSizes)

	totalSize := 0
	for _, size := range folderSizes {
		if size <= 100000 {
			totalSize += size
		}
	}
	fmt.Println(totalSize)
}

// 2086088
func PartB(input []byte) {
	tree := buildTree(input)
	folderSizes := []int{}
	tree.folderSizes(&folderSizes)
	slices.Sort(folderSizes)

	totalSize := folderSizes[len(folderSizes)-1]
	spaceFree := 70000000 - totalSize
	spaceNeeded := 30000000 - spaceFree

	for _, size := range folderSizes {
		if size >= spaceNeeded {
			fmt.Println(size)
			break
		}
	}
}

func (folder *Folder) folderSizes(folderSizes *[]int) {
	size := 0
	for _, file := range folder.files {
		size += file.size
	}
	for _, folder := range folder.folders {
		folder.folderSizes(folderSizes)
		size += folder.size
	}
	folder.size = size
	*folderSizes = append(*folderSizes, size)
}

func buildTree(input []byte) *Folder {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	root := &Folder{
		name:    "/",
		folders: make(map[string]*Folder),
		files:   make(map[string]*File)}
	cwd := root

	for scanner.Scan() {
		var line = scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "$" && parts[1] == "cd" {
			if parts[2] == "/" {
				// root folder, ignore
			} else if parts[2] == ".." {
				// level up
				cwd = cwd.parent
			} else {
				// change cwd
				cwd = cwd.folders[parts[2]]
			}
		} else if parts[0] == "$" && parts[1] == "ls" {
			// ignore
		} else if parts[0] == "dir" {
			// add folder
			newFolder := &Folder{
				name:    parts[1],
				parent:  cwd,
				folders: make(map[string]*Folder),
				files:   make(map[string]*File)}
			cwd.folders[parts[1]] = newFolder
		} else {
			// add file
			size, _ := strconv.Atoi(parts[0])
			newFile := &File{
				name: parts[1],
				size: size}
			cwd.files[parts[1]] = newFile
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}
