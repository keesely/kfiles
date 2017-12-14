/*************************************************************************
> File Name: file_search.go
> Author: Kee
> Mail: chinboy2012@gmail.com
> Created Time: 2017.12.13
************************************************************************/
package kfiles

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindFile(fileName string) ([]string, error) {
	pattern := fileName

	if strings.Index(fileName, "*") != -1 {
		fileName, pattern = getPattern(fileName)
	}

	var list []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if !fi.IsDir() {
		list = append(list, fileName)
		return list, nil
	}

	reg := regexp.MustCompile(pattern)

	// 遍历目录
	filepath.Walk(fileName,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if f.IsDir() {
				path = filepath.ToSlash(path)
				if !strings.HasSuffix(path, "/") {
					path += "/"
				}
				return nil
			}
			// 匹配目录
			pAName := strings.Replace(path, fileName, "", -1)
			matched := reg.MatchString(pAName)
			if matched {
				pflag := true
				match_maps := reg.FindAllStringSubmatch(pAName, -1)
				for k, mms := range match_maps[0] {
					if k <= 0 {
						continue
					}
					smm := fmt.Sprintf("%v", mms)

					if strings.Index(string(smm), "/") != -1 {
						pflag = false
					}
				}
				if pflag == true {
					list = append(list, path)
				}
			}
			return nil
		})

	return list, nil
}

func getPattern(fileName string) (string, string) {
	part_1 := strings.Index(fileName, "*")
	fn_1 := substr(fileName, 0, part_1)

	part_2 := strings.LastIndex(fn_1, "/") + 1
	fn_2 := substr(fn_1, 0, part_2)

	fn_2_last := substr(fn_1, part_2, len(fn_2))

	fn_last := substr(fileName, part_1, len(fileName))

	pattern := fn_2_last + fn_last
	fileName = fn_2

	if strings.Index(pattern, ".") != -1 {
		pattern = strings.Replace(pattern, ".", `\.`, -1)
	}

	pattern = strings.Replace(pattern, "*", `(.*)`, -1)
	pattern = "^" + pattern + "$"
	pattern = "(?U)" + pattern

	return fileName, pattern
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
