package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

var imageSuffix = map[string]bool {
	".jpg": true,
	".jpeg": true,
	".png": true,
	".webp": true,
	".gif": true,
	".tiff": true,
	".eps": true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`Find images

Usage:
  %s [path to find]
`, os.Args[0])
		return
	}

	root := os.Args[1]
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == "_build" {
					return  filepath.SkipDir
				}
				return nil
			}
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if imageSuffix[ext] {
				// ルートパスからの相対パスに変換
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return  nil
				}
				fmt.Printf("%s\n", rel)
			}
			return nil
		})
	if err != nil {
		fmt.Println(1, err)
	}
}
