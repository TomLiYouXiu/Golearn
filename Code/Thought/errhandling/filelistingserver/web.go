package main

import (
	"Thought/errhandling/filelisting"
	"log"
	"log/slog"
	"net/http"
	"os"
)

// 定义一个类型，表示处理请求的函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 错误包装函数，接受一个 appHandler，并返回一个带错误处理的函数
func errWrapper(handle appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// defer 关键字用于在函数执行过程中调用一个函数，该函数会在当前函数返回之前执行。
		// 这里使用 defer 关键字调用一个匿名函数，该函数在函数执行过程中发生异常时被调用。
		defer func() {
			// recover() 函数用于捕获当前函数上下文中的异常。如果当前函数发生异常，recover() 函数会返回异常值。
			if r := recover(); r != nil {

				// log.Panicf() 函数用于打印带有格式化字符串的错误信息，并将程序panic。
				// 这里使用 log.Panicf() 函数打印异常信息，格式为 "Panic: %v"
				log.Panicf("Panic: %v", r)

				// http.Error() 函数用于返回一个 HTTP 错误响应，状态码为 http.StatusInternalServerError（500 Internal Server Error）
				// 这里使用 http.Error() 函数将 HTTP 错误响应写入输出流（writer），并将状态码设置为 http.StatusInternalServerError
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 调用原始的处理函数
		err := handle(writer, request)
		if err != nil {
			// 记录错误日志
			slog.Warn("Error handling request:", err.Error())

			//if userErr, ok := err.(userError); ok {
			//	http.Error(writer,
			//		userError.Message(),
			//		http.StatusBadRequest)
			//	return
			//}
			// 根据错误类型确定 HTTP 状态码
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				// 如果文件不存在，设置 HTTP 404 Not Found 状态码
				code = http.StatusNotFound
			case os.IsPermission(err):
				// 如果权限不足，设置 HTTP 403 Forbidden 状态码
				code = http.StatusForbidden
			default:
				// 其他错误，设置 HTTP 500 Internal Server Error 状态码
				code = http.StatusInternalServerError
			}

			// 响应相应的 HTTP 状态码和错误消息
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	// 将 "/list/" 路径映射到带错误处理的 filelisting.HandleFileList 处理函数
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	// 启动 HTTP 服务器，监听端口 8888
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		// 启动失败时，抛出异常
		panic(err)
	}
}
