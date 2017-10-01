package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	dir  *string = flag.String(`d`, ``, `-d c:\path\to`)
	mode *string = flag.String(`m`, `bl`, `-m cherry_mov/cherry_pic/bl`)
)

func main() {
	flag.Parse()
	if *dir == `` {
		println(`need dir: -d d:\path\to`)
		println(`-m cherry_mov/cherry_pic/bl`)
	}
	PatchRename(*dir, *mode)
}

func PatchRename(dir string, mode string) {

	filepath.Walk(dir, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo == nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		fileName := fileInfo.Name()
		folder := filepath.Dir(path)
		switch mode {
		case `cherry_mov`:
			renameCherryMov(path, folder, fileName)
		case `cherry_pic`:
			renameCherryPic(path, folder, fileName)
		case `bl`:
			renameBL(path, folder, fileName)
		}
		return nil
	})
}

func renameCherryPic(path, folder, fileName string) error {
	newPath := folder + `\` + filepath.Base(folder) + "-" + fileName
	//println(newPath)
	return os.Rename(path, newPath)
}

func renameCherryMov(path, folder, fileName string) error {
	ext := filepath.Ext(fileName)
	if ext == ".jpg" {
		os.Remove(path)
		return nil
	}
	newPath := filepath.Dir(folder) + `\` + filepath.Base(folder) + ext
	fmt.Println(newPath)
	return os.Rename(path, newPath)
	//os.Remove(folder)
}

func renameBL(path, folder, fileName string) error {
	regs := []*regexp.Regexp{}
	regs = append(regs, regexp.MustCompile(`No\.(.*)\[`))
	regs = append(regs, regexp.MustCompile(`Vol\.(.*)\s`))
	regs = append(regs, regexp.MustCompile(`N[oO]\.([^\s]+)`))
	vol := ``
	for _, reg := range regs {
		prefix := reg.FindStringSubmatch(filepath.Base(folder))
		if len(prefix) < 2 {
			continue
		} else {
			vol = prefix[1]
			break
		}
	}
	if vol == `` {
		return nil
	}
	newPath := folder + `\` + vol + "-" + fileName
	println(newPath)
	return os.Rename(path, newPath)
}
