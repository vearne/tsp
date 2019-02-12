package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	// 从文件中装载城市信息

}

var m2 map[string]Point

func init() {
	m2 := map[string]Point{}

	fi, err := os.Open("./data.csv")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		buf, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(buf)
		line = strings.TrimSpace(line)
		itemList := strings.Split(line, ";")
		// 北京、上海
		pos := itemList[0]
		m2[pos] = Point{}
	}
}
