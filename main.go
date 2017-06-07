package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//Reads a file named source.txt and prints output to standard out
func main() {
	currentDir, _ := os.Getwd()
	sourcePath := filepath.Join(currentDir, "/source.txt")
	source, err := os.Open(sourcePath)
	defer source.Close()
	if err != nil {
		panic(err)
	}

	bufferedSource := bufio.NewReader(source)

	for {
		line, err := bufferedSource.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Trim(line, "\n")
		battle := CreateBattle(line)

		if err != nil {
			panic(err)
		}

		fmt.Println(battle.RageOn(&City{}))
	}
}
