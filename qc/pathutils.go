package qc

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

func joinUnix(paths ...string) string {
	return strings.Join(paths, "/")
}

/*
	Pretty naive function
*/
func makePathUnixLike(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

func getFilteredFiles(dir Dir, filters ...Filter) (files []File) {
	for _, filter := range filters {
		for _, file := range dir.Files {
			if filter.IsAcceptable(file) {
				files = append(files, file)
			}
		}
	}

	for _, nestedDir := range dir.Dirs {
		files = append(files, getFilteredFiles(nestedDir, filters...)...)
	}

	return files
}

/*
	Builds a tree of files, where returned Dir is a headPath
*/
func buildFilesTree(rootPath string) (dir Dir, err error) {
	dir.Path = rootPath
	files, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return dir, err
	}

	for _, f := range files {
		fname := f.Name()

		if f.IsDir() {
			var newDir Dir
			newDir, err = buildFilesTree(path.Join(rootPath, fname))

			if err != nil {
				return dir, err
			}

			dir.Dirs = append(dir.Dirs, newDir)
		} else {
			dir.Files = append(dir.Files, File{
				RelPath: rootPath,
				Name:    fname,
				// Be careful: extension always starts with a dot
				Extension: filepath.Ext(fname),
			})
		}
	}

	return dir, nil
}

/*
	Just makes sure path is unix-like and doesn't have an extension
*/
func getQueryPath(file File) string {
	return removeExt(makePathUnixLike(path.Join(file.RelPath, file.Name)), file.Extension)
}

func removeExt(path, ext string) string {
	return path[:len(path)-len(ext)]
}
