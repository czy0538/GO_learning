package main

import (
	"encoding/csv"
	"os"

	"github.com/axgle/mahonia"
)

// GetCSV 读取csv文件.
func GetCSV(csvPath string) [][]string {
	file, err := os.Open(csvPath)
	if err != nil {
		glog.Error("Error when open file in GetCSV():", err)
		return nil
	}
	defer file.Close()

	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	// r := csv.NewReader(file)
	r := csv.NewReader(decoder.NewReader(file)) // 这样，最终返回的字符串就是utf-8了。（go只认utf8）
	res, err := r.ReadAll()
	if err != nil {
		glog.Error("Error when read csv in GetCSV():", err)
		return nil
	}
	return res
}
