package day7

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Folder struct {
	folders map[string]*Folder
	parent  *Folder
	size    int
}

// 1611443
func PartA(input []byte) any {
	tree := buildTree(input)
	folderSizes := []int{}
	tree.folderSizes(&folderSizes)

	totalSize := 0
	for _, size := range folderSizes {
		if size <= 100000 {
			totalSize += size
		}
	}
	return totalSize
}

// 2086088
func PartB(input []byte) any {
	tree := buildTree(input)
	folderSizes := []int{}
	tree.folderSizes(&folderSizes)
	slices.Sort(folderSizes)

	totalSize := folderSizes[len(folderSizes)-1]
	spaceFree := 70000000 - totalSize
	spaceNeeded := 30000000 - spaceFree

	for _, size := range folderSizes {
		if size >= spaceNeeded {
			return size
		}
	}
	return 0
}

func (folder *Folder) folderSizes(folderSizes *[]int) {
	for _, subFolder := range folder.folders {
		subFolder.folderSizes(folderSizes)
		folder.size += subFolder.size
	}
	*folderSizes = append(*folderSizes, folder.size)
}

func buildTree(input []byte) *Folder {
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	root := &Folder{
		folders: make(map[string]*Folder)}
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
		} else if parts[0] == "dir" {
			// add folder
			newFolder := &Folder{
				parent:  cwd,
				folders: make(map[string]*Folder)}
			cwd.folders[parts[1]] = newFolder
		} else {
			// add file
			size, _ := strconv.Atoi(parts[0])
			cwd.size += size
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}
