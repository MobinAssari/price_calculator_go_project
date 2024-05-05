package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	readPath  string
	writePath string
}

func New(inPath, outPath string) FileManager {
	return FileManager{
		readPath:  inPath,
		writePath: outPath,
	}
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.readPath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed reading file")
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		println(err)
		file.Close()
		return nil, errors.New("failed reading")
	}
	file.Close()
	return lines, nil
}
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.writePath)
	if err != nil {
		file.Close()
		return errors.New("write failed")
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("write failed")
	}
	file.Close()
	return nil
}
