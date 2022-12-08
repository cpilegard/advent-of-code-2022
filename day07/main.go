package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1Result := P1(utils.ParseInput(os.Args[2]))
	p2Result := P2(utils.ParseInput(os.Args[2]))

	fmt.Println(p1Result)
	fmt.Println(p2Result)
}

type directory struct {
	name     string
	parent   *directory
	children []directory
}

func P1(input []string) int {
	dirSizes := getDirSizes(input)

	sumOfSmallDirSizes := 0
	for _, v := range dirSizes {
		if v < 100000 {
			sumOfSmallDirSizes += v
		}
	}
	return sumOfSmallDirSizes
}

func P2(input []string) int {
	dirSizes := getDirSizes(input)
	totalSpace := 70000000
	unusedSpace := totalSpace - dirSizes["/"]
	spaceNeededToFreeUp := 30000000 - unusedSpace

	spaceOfSmallestDirNeeded := dirSizes["/"]
	for _, v := range dirSizes {
		if v > spaceNeededToFreeUp && v < spaceOfSmallestDirNeeded {
			spaceOfSmallestDirNeeded = v
		}
	}

	return spaceOfSmallestDirNeeded
}

func getDirSizes(input []string) map[string]int {
	currentDir := &directory{name: "/"}
	dirSizes := map[string]int{}

	for _, cmd := range input {
		cmds := strings.Split(cmd, " ")
		switch cmds[0] {
		case "$":
			switch cmds[1] {
			case "cd":
				// "cd .." moves up one directory
				if cmds[2] == ".." {
					currentDir = currentDir.parent
					continue
				}

				// look for directory in currentDir's children
				for _, dir := range currentDir.children {
					s := strings.Split(dir.name, "/")
					childDirName := s[len(s)-2]

					// set new current directory
					if childDirName == cmds[2] {
						currentDir = &dir
						break
					}
				}
			case "ls":
				// no-op
			}
		case "dir":
			// distinguish between same names in different paths
			dirName := currentDir.name + cmds[1] + "/"
			newDir := directory{name: dirName, parent: currentDir}
			currentDir.children = append(currentDir.children, newDir)
		default: // is a filename
			filesize, _ := strconv.Atoi(cmds[0])

			if _, ok := dirSizes[currentDir.name]; !ok {
				dirSizes[currentDir.name] = filesize
			} else {
				dirSizes[currentDir.name] += filesize
			}

			// add to parent sizes all the way up
			parentCursor := currentDir.parent
			for parentCursor != nil {
				if _, ok := dirSizes[parentCursor.name]; !ok {
					dirSizes[parentCursor.name] = filesize
				} else {
					dirSizes[parentCursor.name] += filesize
				}

				parentCursor = parentCursor.parent
			}
		}
	}

	return dirSizes
}
