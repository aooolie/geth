package module

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
	"util"
)

func ReadTable() map[string]string {
	f, err := os.Open("/home/aolie/project/test/src/input.txt")
	if err != nil {
		fmt.Println("os Open error: ", err)
		return nil
	}
	defer f.Close()

	var inputTable map[string]string /*创建集合 */
	inputTable = make(map[string]string)
	//按行读取文件全部内容,并且将文件的表名转换为hash值
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('*') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		tracer := line
		comma := strings.Index(tracer, "#")//每行的表名以#结尾
		key:=sha256.New()
		key.Write([]byte(tracer))
		r:=key.Sum(nil)
		fmt.Printf("sha256 key hash-A:%x\n",r)
		inputTable [util.Convert(r)] = tracer[comma:]
	}
	return inputTable
}
