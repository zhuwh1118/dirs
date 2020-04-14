package process

import (
	"container/list"
	"io/ioutil"
)

type (
	result struct {
		Path string
		Dirs []dirs
	}

	dirs struct {
		Name  string
		IsDir bool
		Size  int
	}
)

func GetAllFile(path string) (result, error) {
	s := list.New()
	var r result
	r, err := getAllFile(path, s, r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func getAllFile(path string, list *list.List, r result) (result, error) {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}
	var dir dirs
	for _, fi := range rd {
		if fi.IsDir() {
			fullName := path + "/" + fi.Name()
			dir = dirs{Name: fullName, IsDir: true, Size: 0}
			r.Dirs = append(r.Dirs, dir)
			r, err = getAllFile(fullName, list, r)
		} else {
			fullName := path + "/" + fi.Name()
			blob, err := ioutil.ReadFile(fullName)
			if err != nil {
				return r, err
			}
			dir = dirs{Name: fullName, IsDir: false, Size: len(blob)}
			r.Dirs = append(r.Dirs, dir)
			list.PushBack(fullName)
		}
	}
	return r, nil
}
