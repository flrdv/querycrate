package qc

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
	Pretty naive function
*/
func makePathUnixLike(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

func getFilteredFiles(dir Dir, filters ...Filter) (files []File) {
	for _, file := range dir.Files {
		allowed := true

		for _, filter := range filters {
			if !filter.IsAllowed(file) {
				allowed = false
				break
			}
		}

		if allowed {
			files = append(files, file)
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
func buildFilesTree(rootPath string, maxRecursionDepth int) (dir Dir, err error) {
	dir.Path = rootPath
	files, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return dir, err
	}

	for _, f := range files {
		fname := f.Name()

		if f.IsDir() {
			if maxRecursionDepth <= 0 {
				continue
			}

			var newDir Dir
			newDir, err = buildFilesTree(path.Join(rootPath, fname), maxRecursionDepth-1)

			if err != nil {
				return dir, err
			}

			dir.Dirs = append(dir.Dirs, newDir)
		} else {
			dir.Files = append(dir.Files, newFile(rootPath, fname))
		}
	}

	return dir, nil
}

func newFile(path, name string) File {
	ext := filepath.Ext(name)

	return File{
		RelPath:   path,
		Name:      removeExt(name, ext),
		Extension: ext,
	}
}

func getFile(path string) (string, string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return "", "", err
	}

	return removeExt(makePathUnixLike(path), filepath.Ext(path)), string(content), nil
}

/*
	Just makes sure path is unix-like and doesn't have an extension
*/
func getQueryPath(file File) string {
	return removeHeadingFolder(makePathUnixLike(path.Join(file.RelPath, file.Name)))
}

func removeExt(path, ext string) string {
	return path[:len(path)-len(ext)]
}

func removeHeadingFolder(path string) string {
	return path[strings.Index(path, "/")+1:]
}
