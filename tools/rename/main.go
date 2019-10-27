package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	dir  *string = flag.String(`d`, `E:\Pictures\mov`, `-d c:\path\to`)
	mode *string = flag.String(`m`, `bl`, `-m mov/add-dir/mo/pic`)
)

func main() {
	flag.Parse()
	if *dir == `` {
		println(`need dir: -d d:\path\to`)
		println(`-m mov/add-dir/bl`)
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
		// println(folder)
		switch mode {
		case `mov`:
			renameMov(path, folder, fileName)
		case `add-dir`:
			renameAddDir(path, folder, fileName)
		case `mo`:
			moveOut(path, folder, fileName)
		default:
			renamePic(path, folder, fileName)
		}
		return nil
	})
}

//子文件往外移一层
func moveOut(path, folder, fileName string) error {
	newFolder := filepath.Dir(folder)
	newPath := newFolder + `\` + filepath.Base(folder) + "-" + fileName
	println(newPath)
	return os.Rename(path, newPath)
}

func renameAddDir(path, folder, fileName string) error {
	newPath := folder + `\` + filepath.Base(folder) + "-" + fileName
	//println(newPath)
	return os.Rename(path, newPath)
}

func renameMov(path, folder, fileName string) error {
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

func renamePic(path, folder, fileName string) error {
	regs := []*regexp.Regexp{
		regexp.MustCompile(`No\.(\d+)\[`),
		regexp.MustCompile(`N[oO]\.([^\s]+)`),
		regexp.MustCompile(`N0([^\s]+)`),
		regexp.MustCompile(`Vol\.(\d*)\s`),
		regexp.MustCompile(`VOL([^\s]+)`),
		regexp.MustCompile(`(\d{4}[\.\-]\d{2}[\.\-]\d{2})`),
	}

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
