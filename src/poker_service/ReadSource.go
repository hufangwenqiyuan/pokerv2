package poker_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}
func (*File) ReadFile(filePath string) (matches map[string][]map[string]string, err error) {
	if len(filePath) <= 0 {
		return nil, nil
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if err := json.Unmarshal(file, &matches); err != nil {
		return nil, err
	}
	return matches, nil
}

func (*File) WriteJsonFile(FilePath string, matches *map[string][]map[string]string) {
	result, err := json.MarshalIndent(matches, "", "    ")
	if err != nil {
		fmt.Printf("WriteFile # Json.MarshalIndent =====" + err.Error())
		return
	}
	err = ioutil.WriteFile(FilePath, result, 0777)
	if err != nil {
		fmt.Printf("WhriteFile#WriteFile 出错啦 快去看看 ====" + err.Error())
		return
	}
	return
}
