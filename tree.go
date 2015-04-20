package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Goのflagパッケージで同名の引数を複数受け取る - Qiita
// http://qiita.com/hironobu_s/items/96e8397ec453dfb976d4
type strslice []string

func (s *strslice) String() string {
    return fmt.Sprintf("%v", *s)
}

func (s *strslice) Set(v string) error {
    *s = append(*s, v)
    return nil
}

func main() {
	var ignore_dirs strslice
	var empty_dirs strslice
	input_dir := flag.String( "top", ".", "tree TOP directory" )
	flag.Var( &ignore_dirs,  "xd", "eXclude Directory" )
	file_display := flag.Bool( "f", false, "Dispaly files" )
	flag.Var( &empty_dirs, "emd", "Behavior as empty directory" )
	flag.Parse()

	fmt.Println(*input_dir)
	tree(*input_dir, *input_dir, 0, "", *file_display, ignore_dirs, empty_dirs)
}

func tree(rootPath, searchPath string, depth int, parent string, file_display bool, ignore_dirs []string, empty_dirs []string) {

	searchName := filepath.Base(searchPath)
	for _, emd := range empty_dirs {
		if searchName == emd {
			// 空ディレクトリ扱いの場合、以降の処理は不要
			return
		}
	}

	fis, err := ioutil.ReadDir(searchPath)

	if err != nil {
		//fmt.Println( searchPath, " is error" )
		//panic(err)
		return
	}

	dirlist  := make([]string, 0)
	filelist := make([]string, 0)
	for _, fi := range fis {
		fullPath := filepath.Join(searchPath, fi.Name())

		if fi.IsDir() {
			ignore := false
			for _, ignore_dir := range ignore_dirs {
				if ignore_dir == fi.Name() {
					ignore = true
					break
				}
			}
			if !ignore {
				dirlist = append(dirlist, fullPath)
			}
		} else if file_display {
			filelist = append(filelist, fullPath)
		}
	}

	has_dir := ( 0 < len(dirlist) )
	for _, file := range filelist {
		rel, err := filepath.Rel(rootPath, file)

		if err != nil {
			//panic(err)
			return
		}

		base := filepath.Base(rel)

		if has_dir {
			fmt.Println(parent + "│  ", base )
		} else {
			fmt.Println(parent + "    ", base )
		}
	}
	if 0 < len(filelist) {
		if has_dir {
			fmt.Println(parent + "│" )
		} else {
			fmt.Println(parent)
		}
	}

	for idx, dir := range dirlist {
		rel, err := filepath.Rel(rootPath, dir)
		has_young_brother := ( 0 < len(dirlist[idx+1:]) )
		next := ""

		if err != nil {
			//panic(err)
			return
		}

		base := filepath.Base(rel)

		if has_young_brother {
			fmt.Println(parent + "├─", base )
			next += "│  "
		} else {
			fmt.Println(parent + "└─", base )
			next += "    "
		}
		tree(rootPath, dir, depth + 1, parent + next, file_display, ignore_dirs, empty_dirs)
	}

}
