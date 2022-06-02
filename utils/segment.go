package utils

import (
	"fmt"
	"github.com/huichen/sego"
	"strings"
)

//使用sego进行分词
//分词

var seg sego.Segmenter

func init() {
	seg.LoadDictionary("dictionary.txt")
}

var useless = "dpcueyahokw"

func Segment(text string) []string {
	segments := seg.Segment([]byte(text))
	str := sego.SegmentsToString(segments, true)
	arr := strings.Split(str, " ")
	var result []string
	for i := range arr {
		if len(arr[i]) > 1 {
			token := strings.Split(arr[i], "/")
			cx := token[1][0]
			flag := false
			for x := range useless {
				if cx == useless[x] {
					flag = true
				}
			}
			if flag {
				fmt.Println(arr[i])
				continue
			}
			if arr[i][0] == '/' {
				continue
			}
			result = append(result, arr[i])
		}
	}
	return result
}
