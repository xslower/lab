package main

import (
	`flag`
	`io/ioutil`
	`os`
	"os/exec"
	`path/filepath`
	`strings`
)

var (
	usage_i = `-i /path/to/input/files/dir/`
	usage_o = `-o /path/to/output/dir/or_filename.js`
	// usage_f = `-f output_filename.js`
	usage_c = `-c //whether compress`
	i       = flag.String(`i`, ``, usage_i)
	o       = flag.String(`o`, ``, usage_o)
	// f       = flag.String(`f`, `output.js`, usage_f)
	c = flag.Bool(`c`, false, usage_c)
)

func main() {
	flag.Parse()
	if *i == `` || *o == `` {
		println(usage_i)
		println(usage_o)
		// println(usage_f)
		println(usage_c)
	}
	var dir, file = filepath.Split(*o)
	if file == `` {
		file = dir + `output.js`
	} else {
		file = *o
	}
	var ext = filepath.Ext(*i)
	if ext == `.json` {
		compressJson(*i, file)
	} else {
		uglifyJs(*i, file, *c)
	}
}

//undone
func compressJson(in, out string) {
	var bytes, err = ioutil.ReadFile(in)
	if err != nil {
		println(err.Error())
	}
	var compressed = strings.Replace(string(bytes), `  `, ``, -1)
	compressed = strings.Replace(compressed, "\r\n", ``, -1)
	compressed = strings.Replace(compressed, `: `, `:`, -1)
	err = ioutil.WriteFile(out, []byte(compressed), 0666)
	if err != nil {
		println(err.Error())
	}
}

func uglifyJs(dir, filename string, c bool) {
	input := []string{}
	filepath.Walk(dir, func(fullpath string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		var ext = filepath.Ext(fi.Name())
		if ext != `.js` {
			return nil
		}
		input = append(input, fullpath)
		return nil
	})
	input = append(input, `-o`, filename)
	if c {
		input = append(input, `--source-map`, filename+`.map`)
		input = append(input, `--source-map-root`, `./`)
		input = append(input, `--source-map-url`, filepath.Base(filename)+`.map`)
		input = append(input, `-p`, `5`, `-c`, `-m`)
	}
	cmd := exec.Command(`uglifyjs`, input...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		println(err.Error())
	}
}
