package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func ReadFiles(RootDir string, MaxDepth int) (map[string]string, error) {

	retmap := make(map[string]string)

	err := filepath.WalkDir(RootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// make sure I don't dig deeper than maxdepth
		relpath, err := filepath.Rel(RootDir, path)
		if MaxDepth > 0 && strings.Count(relpath, string(os.PathSeparator)) >= MaxDepth {
			return nil
		}

		// resolve all symlinks
		link, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}

		ls, err := os.Lstat(link)
		if err != nil {
			return err
		}
		if !ls.IsDir() {
			contents, err := os.ReadFile(link)
			if err != nil {
				return err
			}
			retmap[path] = string(contents)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return retmap, nil
}

func GetBinding(RootDir string, Type string) (map[string]string, error) {
	files, err := os.ReadDir(RootDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		fc, err := ReadFiles(filepath.Join(RootDir, f.Name()), 1)
		if err != nil {
			return nil, err
		}

		// flatten key paths
		b := make(map[string]string)
		for k, v := range fc {
			b[filepath.Base(k)] = strings.TrimSpace(v)
		}

		if val, ok := b["type"]; ok {
			if val == Type {
				return b, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("No bindings of type '%s' have been found", Type))
}