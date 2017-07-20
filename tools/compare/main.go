package main

import (
	`biz/tools/tools`
	`flag`
	"fmt"
	`github.com/astaxie/beego/config`
	`greentea/utils`
	`os`
	`path/filepath`
	`strings`
)

var (
	php      = flag.Bool(`php`, false, `-php`)
	gol      = flag.Bool(`gol`, false, `-gol`)
	sync     = flag.Bool(`sync`, false, `-sync`)
	override = flag.Bool(`override`, false, `-override`)
)

func main() {
	flag.Parse()
	conf, err := config.NewConfig(`ini`, `config.ini`)
	throw(err)
	phpcnf, _ := conf.GetSection(`php`)
	if *php {
		compare(phpcnf)
	}
	gocnf, _ := conf.GetSection(`go`)
	if *gol {
		compare(gocnf)
	}
}

func compare(conf map[string]string) {
	prefixA := conf[`prefix_a`]
	prefixB := conf[`prefix_b`]
	folder := strings.Split(conf[`folders`], `;`)
	for _, fld := range folder {
		f := strings.Trim(fld, ` `)
		CompareDir(prefixA+f, prefixB+f, *sync, *override)
	}
}

func CompareDir(dirA, dirB string, sync bool, override bool) {
	//用来存放dirA中所有的文件以供对比
	fileA := map[string]os.FileInfo{}
	filepath.Walk(dirA, func(fullPath string, fiA os.FileInfo, err error) error {
		if fiA.IsDir() {
			return nil
		}
		path := fullPath[len(dirA)+1:]
		if neglect(path, fiA) {
			return nil
		}
		fileA[path] = fiA
		return nil
	})
	//用来存放dirA中没有的文件
	fileB := map[string]os.FileInfo{}
	filepath.Walk(dirB, func(fullPath string, fiB os.FileInfo, err error) error {
		if fiB.IsDir() {
			return nil
		}
		path := fullPath[len(dirB)+1:]
		if neglect(path, fiB) {
			return nil
		}
		fiA, ok := fileA[path]
		if !ok {
			fileB[path] = fiA
			return nil
		}
		delete(fileA, path)
		if fiB.Size() != fiA.Size() || !fiB.ModTime().Equal(fiA.ModTime()) {
			showFileInfo(path, dirA, dirB, fiA, fiB)
			if override {
				src := dirB + `\` + path
				dst := dirA + `\` + path
				if fiB.ModTime().Before(fiA.ModTime()) {
					src, dst = dst, src
				}
				//以防错误覆盖
				moveFile(dst, dst+`.bak`)
				copyFile(src, dst)
			}
		}
		return nil
	})

	fmt.Println(`[lonely files]:`)
	listFiles(fileA, dirA, dirB, sync)
	listFiles(fileB, dirB, dirA, sync)
	fmt.Println(`[End]`)
}

func listFiles(files map[string]os.FileInfo, dirSrc, dirDst string, sync bool) {
	if len(files) == 0 {
		return
	}
	fmt.Println(dirSrc)
	for path, fi := range files {
		showFileInfo(path, ``, ``, nil, fi)
		if sync {
			src := dirSrc + `\` + path
			dst := dirDst + `\` + path
			copyFile(src, dst)
		}
	}
}

func neglect(path string, fi os.FileInfo) bool {
	if string(path[0]) == `.` {
		return true
	}
	ext := filepath.Ext(path)
	if ext == `.bak` || ext == `.ini` {
		return true
	}
	if fi.Size() > 2*1024*1024 { //>2M的文件忽略
		return true
	}
	return false
}

func showFileInfo(path, dirA, dirB string, fiA, fiB os.FileInfo) {
	layout := `2006-01-02 15:04:05`
	fmt.Println(path)
	if fiA != nil {
		fmt.Println(fiA.ModTime().Format(layout), fiA.Size(), dirA)
	}
	if fiB != nil {
		fmt.Println(fiB.ModTime().Format(layout), fiB.Size(), dirB)
	}

}

func copyFile(src, dst string) {
	os.MkdirAll(filepath.Dir(dst), 0755)
	// os.Rename(src, dst)
	tools.CopyFile(src, dst)
	fi, err := os.Stat(src)
	throw(err)
	err = os.Chtimes(dst, fi.ModTime(), fi.ModTime())
	check(err)

}

func moveFile(src, dst string) {
	_, err := os.Stat(dst)
	if err == nil || os.IsExist(err) {
		err = os.Remove(dst)
		check(err)
	}
	err = os.Rename(src, dst)
	check(err)

}

func throw(err error, msg ...string) {
	if err != nil {
		panic(err.Error())
	}

}

func echo(i ...interface{}) {
	utils.Echo(i...)
}

func check(err error, msg ...interface{}) {
	utils.Check(err, msg...)
}
