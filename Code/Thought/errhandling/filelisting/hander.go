package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

// HandleFileList 是一个处理函数，它读取文件的内容并将其写入 HTTP 响应。
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must satrt with" + prefix)
	}
	// 从 URL 中提取路径，去除 "/list/" 前缀
	path := request.URL.Path[len(prefix):]

	// 打开与提取的路径对应的文件
	file, err := os.Open(path)
	if err != nil {
		// 如果打开文件时发生错误，返回错误
		return err
	}
	defer file.Close() // 在函数返回时确保文件关闭

	// 读取文件的全部内容
	all, err := ioutil.ReadAll(file)
	if err != nil {
		// 如果读取文件内容时发生错误，返回错误
		return err
	}

	// 将文件内容写入 HTTP 响应
	writer.Write(all)

	// 返回 nil 表示成功（没有错误）
	return nil
}
