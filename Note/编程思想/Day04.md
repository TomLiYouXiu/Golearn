# 测试

<img src="https://pic.imgdb.cn/item/65756471c458853aef356bdc.jpg" style="zoom:50%;" />

## 传统测试缺点

* 测试数据和测试逻辑混在一起
* 出错信息不明确
* 一旦一个数据出错测试全部结束

## 表格驱动测试

<img src="https://pic.imgdb.cn/item/657565f7c458853aef3e855d.jpg" style="zoom:50%;" />

* 分离的测试数据和测试逻辑
* 明确的出错信息
* 可以部分失败
* go语言的语法使得我们更易实践表格测试驱动

# 测试过程

## o 单测入门篇：Golang 单元测试基本使用

## Golang 单元测试规范

### Go 单元测试概要

Go 语言的单元测试默认采用官方自带的测试框架，通过引入 testing 包以及 执行 `go test` 命令来实现单元测试功能。

在源代码包目录内，所有以 `_test.go` 为后缀名的源文件会被 `go test` 认定为单元测试的文件，这些单元测试的文件不会包含在 `go build` 的源代码构建中，而是单独通过 go test 来编译并执行。

### Go 单元测试的基本规范

Go 单元测试的基本规范如下：

- 每个测试函数都必须导入 testing 包。测试函数的命名类似`func TestName(t *testing.T)`，入参必须是 `*testing.T`
- 测试函数的函数名必须以大写的 Test 开头，后面紧跟的函数名，要么是大写开关，要么就是下划线，比如 `func TestName(t *testing.T)` 或者  `func Test_name(t *testing.T)`  都是 ok 的， 但是 `func Testname(t *testing.T)`不会被检测到
- 通常情况下，需要将测试文件和源代码放在同一个包内。一般测试文件的命名，都是 `{source_filename}_test.go`，比如我们的源代码文件是allen.go ，那么就会在 allen.go 的相同目录下，再建立一个 allen_test.go 的单元测试文件去测试 allen.go 文件里的相关方法。

当运行 go test 命令时，go test 会遍历所有的 `*_test.go` 中符合上述命名规则的函数，然后生成一个临时的 main 包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

## 从一个简单测试用例来确认 go test 的各种使用方法

一个简单的 xxx_test.go 的单元测试文件如下，里面有两个测试方法：

```scss
scss复制代码package util

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2, 3) != 6 {
		t.Fatal("sum error")
	}
}

func TestAbs(t *testing.T) {
	if Abs(5) != 5 {
		t.Fatal("abs error, except:5, result:", Abs(5))
	}
}
```

### go test -v 执行单测并打印详情

运行方法：进入到包内，运行命令 go test -v ，参数 -v 可以打印详情。 也可以只运行某个方法的单元测试： go test -v -run="xxx" ，支持正则表达式。

```ruby
ruby复制代码allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go test -v
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok  	baseCodeExample/gotest	0.005s

allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go test -v -run="Abs"
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok  	baseCodeExample/gotest	0.006s
```

### go test -v -cover 执行单测并计算覆盖率

go test 工具还有个功能是测试单元测试的覆盖率，用法为 go test -v -cover， 示例如下：

```ruby
ruby复制代码allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go test -v -cover
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
coverage: 85.7% of statements
ok  	baseCodeExample/gotest	0.005s
```

从覆盖率来看（coverage: 85.7% of statements），单元测试没有覆盖全部的代码，只有 85.7% ，我们可以通过如下命令将 cover 的详细信息保存到cover.out 中。

```diff
diff复制代码go test -cover -coverprofile=cover.out -covermode=count
注：
-cover 允许代码分析
-covermode 代码分析模式（set：是否执行；count：执行次数；atomic：次数，并发执行）
-coverprofile 输出结果文件
```

然后再通过

```go
go
复制代码go tool cover -func=cover.out
```

查看每个方法的覆盖率。

```ruby
ruby复制代码allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go tool cover -func=cover.out
baseCodeExample/gotest/compute.go:5:	Sum		100.0%
baseCodeExample/gotest/compute.go:13:	Abs		66.7%
total:					(statements)	85.7%
```

这里发现是 Abs 方法没有覆盖完全，因为我们的用例只用到了正数的那个分支。 还可以使用 html 的方式查看具体的覆盖情况。

```ini
ini
复制代码go tool cover -html=cover.out
```

会默认打开浏览器，将覆盖情况显示到页面中:

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6276701cd6e9483a9dc847d70594550d~tplv-k3u1fbpfcp-zoom-in-crop-mark:1512:0:0:0.awebp?)

可以看出 Abs 方法的负数分支没有覆盖到。将 TestAbs 方法修改如下即可：

```scss
scss复制代码func TestAbs(t *testing.T) {
    if Abs(5) != 5 {
        t.Fatal("abs error, except:5, result:", Abs(5))
    }
    if Abs(-4) != 4 {
        t.Fatal("abs error, except:4, result:", Abs(-4))
    }
}
```

再次运行：

```ini
ini复制代码go test -cover -coverprofile=cover2.out -covermode=count
go tool cover -func=cover2.out
```

运行结果如下：

```ruby
ruby复制代码allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go test -cover -coverprofile=cover2.out -covermode=count
PASS
coverage: 100.0% of statements
ok  	baseCodeExample/gotest	0.006s
allen@MackBook:~/work/goDev/Applications/src/baseCodeExample/gotest$go tool cover -func=cover2.out
baseCodeExample/gotest/compute.go:5:	Sum		100.0%
baseCodeExample/gotest/compute.go:13:	Abs		100.0%
total:					(statements)	100.0%
```

这个说明已经达到了 100% 的覆盖率了。

Go 单测覆盖度的相关命令汇总如下：

```go
go复制代码go test -v -cover

go test -cover -coverprofile=cover.out -covermode=count

go tool cover -func=cover.out
```

## Go 单测常见使用方法

### 测试单个文件

通常，一个包里面会有多个方法，多个文件，因此也有多个 test 用例，假如我们只想测试某一个方法的时候，那么我们需要指定某个文件的某个方法

如下：

```ruby
ruby复制代码allen@MackBook:~/work/goDev/Applications/src/gitlab.allen.com/avatar/app_server/service/centralhub$tree .
.
├── msghub.go
├── msghub_test.go
├── pushhub.go
├── rtvhub.go
├── rtvhub_test.go
├── userhub.go
└── userhub_test.go

0 directories, 7 files
```

总共有7个文件，其中有三个test文件，如果直接运行 go test，就会测试所有test.go文件了。

但是，假如我们只更新了 rtvhub.go 里面的代码，所以我只想要测试 rtvhub.go 里面的某个方法，那么就需要指定文件，具体的方法就是同时指定我们需要测试的test.go 文件和 它的源文件，如下：

```go
go
复制代码go test -v msghub.go  msghub_test.go 
```

### 测试单个文件下的单个方法

在测试单个文件之下，假如我们单个文件下，有多个方法，我们还想只是测试单个文件下的单个方法，要如何实现？我们需要再在此基础上，用 -run 参数指定具体方法或者使用正则表达式。

假如 test 文件如下：

```go
go复制代码package centralhub

import (
	"context"
	"testing"
)

func TestSendTimerInviteToServer(t *testing.T) {
	ctx := context.Background()

	err := sendTimerInviteToServer(ctx, 1461410596, 1561445452, 2)
	if err != nil {
		t.Errorf("send to server friendship build failed. %v", err)
	}
}

func TestSendTimerInvite(t *testing.T) {
	ctx := context.Background()
	err := sendTimerInvite(ctx, "test", 1461410596, 1561445452)
	if err != nil {
		t.Errorf("send timeinvite to client failed:%v", err)
	}
}
go复制代码只测试 TestSendTimerInvite 方法
go test -v msghub.go  msghub_test.go -run TestSendTimerInvite

测试所有正则匹配 SendTimerInvite 的方法 
go test -v msghub.go  msghub_test.go -run "SendTimerInvite"
```

### 测试所有方法

直接 go test 就行

## 竞争检测(race detection)

### go run -race 执行竞争检测

当两个goroutine并发访问同一个变量，且至少一个goroutine对变量进行写操作时，就会发生数据竞争（data race）。 为了协助诊断这种bug，Go提供了一个内置的数据竞争检测工具。 通过传入-race选项，go tool就可以启动竞争检测。

```go
go复制代码$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```

### 示例代码

```go
go复制代码package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 0
	go func() {
		for {
			i++
			fmt.Println("subroutine: i = ", i)
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		i++
		fmt.Println("mainroutine: i = ", i)
		time.Sleep(1 * time.Second)
	}
}
```

### 演示结果

```ini
ini复制代码$ go run -race testrace.go
mainroutine: i =  1
==================
WARNING: DATA RACE
Read at 0x00c0000c2000 by goroutine 6:
  main.main.func1()
      /Users/wudebao/Documents/workspace/goDev/Applications/src/base-code-example/system/testrace/testrace.go:12 +0x3c

Previous write at 0x00c0000c2000 by main goroutine:
  main.main()
      /Users/wudebao/Documents/workspace/goDev/Applications/src/base-code-example/system/testrace/testrace.go:18 +0x9e

Goroutine 6 (running) created at:
  main.main()
      /Users/wudebao/Documents/workspace/goDev/Applications/src/base-code-example/system/testrace/testrace.go:10 +0x7a
==================
subroutine: i =  2
mainroutine: i =  3
subroutine: i =  4
mainroutine: i =  5
subroutine: i =  6
mainroutine: i =  7
subroutine: i =  8
subroutine: i =  9
mainroutine: i =  10
```

# pprof性能调优

<img src="https://pic.imgdb.cn/item/6575845ec458853aefd2f3b8.jpg" style="zoom:50%;" />

## 一、golang 程序性能调优[#](https://www.cnblogs.com/jiujuan/p/14588185.html#2349594506)

### 在 golang 程序中，有哪些内容需要调试优化？[#](https://www.cnblogs.com/jiujuan/p/14588185.html#2593256378)

一般常规内容：

1. **cpu**：程序对cpu的使用情况 - 使用时长，占比等
2. **内存**：程序对cpu的使用情况 - 使用时长，占比，内存泄露等。如果在往里分，程序堆、栈使用情况
3. **I/O**：IO的使用情况 - 哪个程序IO占用时间比较长

golang 程序中：

1. **goroutine**：go的协程使用情况，调用链的情况
2. **goroutine leak**：goroutine泄露检查
3. **go dead lock**：死锁的检测分析
4. **data race detector**：数据竞争分析，其实也与死锁分析有关

上面是在 golang 程序中，性能调优的一些内容。

### 有什么方法工具调试优化 golang 程序？[#](https://www.cnblogs.com/jiujuan/p/14588185.html#857540299)

比如 linux 中 cpu 性能调试，工具有 top，dstat，perf 等。

那么在 golang 中，有哪些分析方法？

**golang 性能调试优化方法：**

- **Benchmark**：**基准测试**，对特定代码的运行时间和内存信息等进行测试
- **Profiling**：**程序分析**，程序的运行画像，在程序执行期间，通过采样收集的数据对程序进行分析
- **Trace**：**跟踪**，在程序执行期间，通过采集发生的事件数据对程序进行分析

> profiling 和 trace 有啥区别？
> profiling 分析没有时间线，trace 分析有时间线。

在 golang 中，应用方法的工具呢？

这里介绍 pprof 这个 golang 工具，它可以帮助我们调试优化程序。

> 它的最原始程序是 [gperftools](https://github.com/gperftools/gperftools) - https://github.com/gperftools/gperftools，golang 的 pprof 是从它而来的。

## 二、pprof 介绍[#](https://www.cnblogs.com/jiujuan/p/14588185.html#1257255281)

### 简介[#](https://www.cnblogs.com/jiujuan/p/14588185.html#1599952447)

pprof 是 golang 官方提供的性能调优分析工具，可以对程序进行性能分析，并可视化数据，看起来相当的直观。
当你的 go 程序遇到性能瓶颈时，可以使用这个工具来进行调试并优化程序。

本文将对下面 golang 中 2 个监控性能的包 pprof 进行运用：

- [runtime/pprof](https://golang.org/pkg/runtime/pprof/)：采集程序运行数据进行性能分析，一般用于后台工具型应用，这种应用运行一段时间就结束。
- [net/http/pprof](https://golang.org/pkg/net/http/pprof/)：对 runtime/pprof 的二次封装，一般是服务型应用。比如 web server ，它一直运行。这个包对提供的 http 服务进行数据采集分析。

上面的 pprof 开启后，每隔一段时间就会采集当前程序的堆栈信息，获取函数的 cpu、内存等使用情况。通过对采样的数据进行分析，形成一个数据分析报告。

pprof 以 [profile.proto](https://github.com/google/pprof/blob/master/proto/profile.proto) 的格式保存数据，然后根据这个数据可以生成可视化的分析报告，支持文本形式和图形形式报告。
profile.proto 里具体的数据格式是 [protocol buffers](https://developers.google.com/protocol-buffers)。

那用什么方法来对数据进行分析，从而形成文本或图形报告？

用一个命令行工具 `go tool pprof` 。

### pprof 使用模式[#](https://www.cnblogs.com/jiujuan/p/14588185.html#4161658926)

- Report generation：报告生成
- Interactive terminal use：交互式终端
- Web interface：Web 界面

## 三、runtime/pprof[#](https://www.cnblogs.com/jiujuan/p/14588185.html#3358980221)

### 使用前的准备工作[#](https://www.cnblogs.com/jiujuan/p/14588185.html#2266577717)

调试分析 golang 程序，要开启 profile 然后开始采样数据。
然后安装：`go get github.com/google/pprof` ，后面分析会用到。

采样数据的方式：

- **第 1 种**，在 go 程序中添加如下代码：
  [StartCPUProfile](https://golang.org/pkg/runtime/pprof/#StartCPUProfile) 为当前 process 开启 CPU profiling 。
  [StopCPUProfile](https://golang.org/pkg/runtime/pprof/#StopCPUProfile) 停止当前的 CPU profile。当所有的 profile 写完了后它才返回。

```scss
Copy// 开启 cpu 采集分析：
pprof.StartCPUProfile(w io.Writer)

// 停止 cpu 采集分析：
pprof.StopCPUProfile()
```

[WriteHeapProfile](https://golang.org/pkg/runtime/pprof/#WriteHeapProfile) 把内存 heap 相关的内容写入到文件中

```lua
Copypprof.WriteHeapProfile(w io.Writer)
```

- **第 2 种**，在 benchmark 测试的时候

```sh
Copygo test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

还有就是对 web 服务（http server） 数据的采集

```bash
Copygo tool pprof $host/debug/pprof/profile
```

### 程序示例[#](https://www.cnblogs.com/jiujuan/p/14588185.html#1213617579)

> go version go1.13.9

#### 例子 1[#](https://www.cnblogs.com/jiujuan/p/14588185.html#1516818927)

我们用第 1 种方法，在程序中添加分析代码，demo.go :

```go
Copypackage main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write mem profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 200; i++ {
		go cyclenum(30000, &wg)
	}

	writeBytes()

	wg.Wait()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("cound not write memory profile: ", err)
		}
	}
}

func cyclenum(num int, wg *sync.WaitGroup) {
	slice := make([]int, 0)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 0; i < 30000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
```

编译程序、采集数据、分析程序：

1. 编译 demo.go

```go
Copygo build demo.go
```

1. 用 pprof 采集数据，命令如下：

```bash
Copy./demo.exe --cpuprofile=democpu.pprof  --memprofile=demomem.pprof
```

> 说明：我是 win 系统，这个 demo 就是 demo.exe ，linux 下是 demo

1. 分析数据，命令如下：

```go
Copygo tool pprof democpu.pprof
```

go tool pprof 简单的使用格式为：`go tool pprof [binary] [source]`

- binary： 是应用的二进制文件，用来解析各种符号
- source： 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址

> 要了解 go tool pprof 更多命令使用方法，请查看文档：`go tool pprof --help`

> **注意**：
>
> 获取的 Profiling 数据是动态获取的，如果想要获取有效的数据，需要保证应用或服务处于较大的负载中，比如正在运行工作中的服务，或者通过其他工具模拟访问压力。
> 否则如果应用处于空闲状态，比如 http 服务处于空闲状态，得到的结果可能没有任何意义。
> （后面会遇到这种问题，http 的 web 服务处于空闲状态，采集显示的数据为空）

分析数据，基本的模式有 2 种：

- 一个是**命令行交互分析**模式
- 一个是**图形可视化分析**模式

#### 命令行交互分析[#](https://www.cnblogs.com/jiujuan/p/14588185.html#3651001728)

##### A：命令行交互分析

1. 分析上面采集的数据，命令： `go tool pprof democpu.pprof`

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122110682-306537236.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122110682-306537236.png)

| 字段           | 说明                 |
| -------------- | -------------------- |
| **Type**：     | 分析类型，这里是 cpu |
| **Duration**： | 程序执行的时长       |

Duration 下面还有一行提示，这是交互模式（通过输入 help 获取帮助信息，输入 o 获取选项信息）。

可以看出，go 的 pprof 操作还有很多其他命令。

1. 输入 help 命令，出来很多帮助信息：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122128925-308223199.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122128925-308223199.png)

Commands 下有很多命令信息，text ，top 2个命令解释相同，输入这个 2 个看看：

1. 输入 top，text 命令

> top 命令：对函数的 cpu 耗时和百分比排序后输出

top后面还可以带参数，比如： top15

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122145607-2007829916.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122145607-2007829916.png)

输出了相同的信息。

| 字段   | 说明                                                |
| ------ | --------------------------------------------------- |
| flat   | 当前函数占用 cpu 耗时                               |
| flat % | 当前函数占用 cpu 耗时百分比                         |
| sum%   | 函数占用 cpu 时间累积占比，从小到大一直累积到 100%  |
| cum    | 当前函数加上调用当前函数的函数占用 cpu 的总耗时     |
| %cum   | 当前函数加上调用当前函数的函数占用 cpu 的总耗时占比 |

从字段数据我们可以看出哪一个函数比较耗费时间，就可以对这个函数进一步分析。
分析用到的命令是 `list` 。

> list 命令：可以列出函数最耗时的代码部分，格式：list 函数名

从上面采样数据可以分析出总耗时最长的函数是 `main.cycylenum`，用 `list cyclenum` 命令进行分析，如下图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122206105-453790225.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122206105-453790225.png)

发现最耗时的代码是 62 行：`slice = append(slice, j)` ，这里耗时有 1.47s ，可以对这个地方进行优化。

这里耗时的原因，应该是 slice 的实时扩容引起的。那我们空间换时间，固定 slice 的容量，make([]int, num * num)

##### B：命令行下直接输出分析数据

在命令行直接输出数据，基本命令格式为：

```sh
Copygo tool pprof <format> [options] [binary] <source>
```

输入命令：`go tool pprof -text democpu.pprof` ，输出：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330153439747-339748791.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330153439747-339748791.png)

#### 可视化分析[#](https://www.cnblogs.com/jiujuan/p/14588185.html#2208266135)

##### A. pprof 图形可视化

除了上面的命令行交互分析，还可以用图形化来分析程序性能。

图形化分析前，先要安装 graphviz 软件，

- 下载地址：[graphviz地址](https://graphviz.org/download/)，

下载对应的平台安装包，安装完成后，把执行文件 bin 放入 Path 环境变量中，然后在终端输入 `dot -version` 命令查看是否安装成功。

**生成可视化文件：**

有 2 个步骤，根据上面采集的数据文件 democpu.pprof 来进行可视化：

1. 命令行输入：go tool pprof democpu.pprof
2. 输入 web 命令

在命令行里输入 web 命令，就可以生成一个 svg 格式的文件，用浏览器打开即可查看 svg 文件。

执行上面 2 个命令如下图：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122233492-941911310.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122233492-941911310.png)

用浏览器查看生成的 svg 图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122243497-1225791807.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122243497-1225791807.png)

(文件太大，只截取了一小部分图，完整的图请自行生成查看)

**关于图形的一点说明：**

1. 每个框代表一个函数，理论上框越大表示占用的 cpu 资源越多
2. 每个框之间的线条代表函数之间的调用关系，线条上的数字表示函数调用的次数
3. 每个框中第一行数字表示当前函数占用 cpu 的百分比，第二行数字表示当前函数累计占用 cpu 的百分比

##### B. web可视化-浏览器上查看数据

运行命令：`go tool pprof -http=:8080 democpu.pprof`

```sh
Copy$ go tool pprof -http=:8080 democpu.pprof
Serving web UI on http://localhost:8080
```

命令运行完成后，会自动在浏览器上打开地址： `http://localhost:8080/ui/`，我们可以在浏览器上查看分析数据：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330041127682-216491175.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330041127682-216491175.png)
这张图就是上面用 web 命令生成的图。

> 如果你在 web 浏览时没有这么多菜单可供选择，那么请安装原生的 pprof 工具：
> `go get -u github.com/google/pprof` ，然后在启动 `go tool pprof -http=:8080 democpu.pprof` ，就会出来菜单。

还可以查看火焰图， http 地址：http://localhost:8080/ui/flamegraph，可直接点击 VIEW 菜单下的 Flame Graph 选项查看火焰图。当然还有其他选项可供选择，比如 Top，Graph 等等选项。你可以根据需要选择。
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330041153228-1227677383.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330041153228-1227677383.png)

##### C. 火焰图 Flame Graph

其实上面的 web 可视化已经包含了火焰图，把火焰图集成到了 pprof 里。但为了向性能优化专家 Bredan Gregg 致敬，还是来体会一下火焰图生成过程。

火焰图 ([Flame Graph](https://github.com/brendangregg/FlameGraph)) 是性能优化专家 Bredan Gregg 创建的一种性能分析图。Flame Graphs visualize profiled code。

火焰图形状如下：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122301550-1658411625.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122301550-1658411625.png)

（来自：https://github.com/brendangregg/FlameGraph）

上面用 pprof 生成的采样数据，要把它转换成火焰图，就要使用一个转换工具 [go-torch](https://www.cnblogs.com/jiujuan/p/github.com/uber/go-torch)，这个工具是 uber 开源，它是用 go 语言编写的，可以直接读取 pprof 采集的数据，并生成一张火焰图， svg 格式的文件。

1. 安装 go-torch：

> go get -v github.com/uber/go-torch

1. 安装 flame graph：

> git clone https://github.com/brendangregg/FlameGraph.git

并把 FlameGraph 安装目录位置添加进 Path 中。

1. 安装 perl 环境：

生成火焰图的程序 FlameGraph 是用 perl 写的，所以先要安装执行 perl 语言的环境。

- 安装 perl 环境：https://www.perl.org/get.html
- 把执行文件 bin 加入 Path 中
- 在终端下执行命令：`perl -h` ，输出了帮助信息，则说明安装成功

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122317270-658836005.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122317270-658836005.png)

1. 验证 FlameGraph 是否安装成功：

进入到 FlameGraph 安装目录，执行命令，`./flamegraph.pl --help`

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122328508-547114025.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122328508-547114025.png)

输出信息说明安装成功

1. 生成火焰图：

重新进入到文件 democpu.pprof 的目录，然后执行命令：

> go-torch -b democpu.pprof

上面命令默认生成名为 torch.svg 的文件，用浏览器打开查看：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122340804-1272269398.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122340804-1272269398.png)

自定义输出文件名，后面加 `-f` 参数：

> go-torch -b democpu.pprof -f cpu_flamegraph.svg

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122352322-482553191.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210329122352322-482553191.png)

火焰图说明：

> 火焰图 svg 文件，你可以点击上面的每个方块来查看分析它上面的内容。
>
> 火焰图的调用顺序从下到上，每个方块代表一个函数，它上面一层表示这个函数会调用哪些函数，方块的大小代表了占用 CPU 使用时长长短。

go-torch 的命令格式：

```sh
Copygo-torch [options] [binary] <profile source>
```

go-torch 帮助文档：

> 想了解更多 go-torch 用法，请用 help 命令查看帮助文档，`go-torch --help`。
>
> 或查看 [go-torch README](https://github.com/uber-archive/go-torch/blob/master/README.md) 文档 。

## 四、web 服务(http server)的分析 net/http/pprof[#](https://www.cnblogs.com/jiujuan/p/14598141.html#239151147)

### 4.1 代码例子 1[#](https://www.cnblogs.com/jiujuan/p/14598141.html#1538133357)

> go version go1.13.9

把上面的程序例子稍微改动下，命名为 demohttp.go:

```go
Copypackage main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/pprof-test", handler)

	fmt.Println("http server start")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 200; i++ {
		go cyclenum(30000, &wg)
	}

	wg.Wait()

	wb := writeBytes()
	b, err := ioutil.ReadAll(wb)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	resp.Write(b)
}

func cyclenum(num int, wg *sync.WaitGroup) {
	slice := make([]int, 0)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 0; i < 30000; i++ {
		buff.Write([]byte{'a' + byte(rand.Intn(10))})
	}
	return &buff
}
```

### 4.2 开始分析[#](https://www.cnblogs.com/jiujuan/p/14598141.html#3237244745)

#### 4.2.1 在 web 界面上分析[#](https://www.cnblogs.com/jiujuan/p/14598141.html#1596166763)

先运行上面的 demohttp.go 程序，执行命令：

> go run demohttp.go

然后在浏览器输入：http://localhost:8090/debug/pprof/，查看服务运行情况，如下图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181541666-1171115210.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181541666-1171115210.png)

| 名称         | url                                    | 说明                                                         |
| ------------ | -------------------------------------- | ------------------------------------------------------------ |
| allocs       | $host/debug/pprof/allocs?debug=1       | 过去所有内存抽样情况                                         |
| block        | $host/debug/pprof/block?debug=1        | 同步阻塞时程序栈跟踪的一些情况                               |
| heap         | $host/debug/pprof/heap?debug=1         | 活动对象的内存分配情况                                       |
| mutex        | $host/debug/pprof/mutex?debug=1        | 互斥锁持有者的栈帧情况                                       |
| profile      | $host/debug/pprof/profile              | cpu profile，点击时会得到一个文件，然后可以用 go tool pprof 命令进行分析 |
| threadcreate | $host/debug/pprof/threadcreate?debug=1 | 创建新 OS 线程的堆栈跟踪情况                                 |
| trace        | $host/debug/pprof/trace                | 当前程序执行的追踪情况，点击时会得到一个文件，可以用 go tool trace 命令来分析这个文件 |

点击上面的链接，就可以查看具体的分析情况。
不断刷新网页，可以看到数据在不断变化。

#### 4.2.2 命令行交互分析[#](https://www.cnblogs.com/jiujuan/p/14598141.html#883942169)

在命令行上运行 demohttp.go 程序，执行命令:

> go run demohttp.go

**A. 分析 cpu profile**

在开启另外一个命令行终端，执行如下命令：

> go tool pprof http://localhost:8090/debug/pprof/profile?seconds=70

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181600702-127232672.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181600702-127232672.png)

参数 seconds = 70：进行 70s 的数据样本采集，这个参数可以根据实际情况调整。

上面的命令执行后，会等待 70s ， 然后才会进入命令交互界面，如上图

输入 `top` 命令：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181615969-2144123551.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181615969-2144123551.png)

大家发现没，其实与上面 runtime/pprof 在命令行交互时是一样的操作，可以参考上面的字段参数说明。

找出耗时代码部分，也可以用命令：`list`。

在 `top` 命令执行后，发现什么问题没？这个 top 命令显示的信息都是系统调用信息耗时，没有用户定义的函数。为什么？下面进行分析。

**B. 分析 memory profile**

执行命令：

> go tool pprof http://localhost:8090/debug/pprof/heap

然后同样输入 `top` 命令查看函数使用情况，如下图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181636344-736191248.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181636344-736191248.png)

其余的跟踪分析命令类似，就不一一分析了。

把上面在终端命令行下交互分析的数据进行可视化分析。

#### 4.2.3 图形可视化分析[#](https://www.cnblogs.com/jiujuan/p/14598141.html#1582412224)

##### A. pprof 图形可视化

在前面可视化分析中，我们了解到可视化最重要有 2 步：1.采集数据 2.图形化采集的数据。

在上面第三节 runtime/pprof 中，进入终端命令行交互操作，然后输入 web 命令，就可以生成一张 svg 格式的图片，用浏览器可以直接查看该图片。我们用同样的方法来试一试。

1. 输入命令：

> go tool pprof http://localhost:8090/debug/pprof/profile?seconds=30

1. 等待 30s 后输入 `web` 命令

如下图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181654769-298765602.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181654769-298765602.png)

果然生成了一个 svg 文件，在浏览器查看该图片文件，啥有用信息也没有，如下图：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181713771-1008781385.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181713771-1008781385.png)

为什么没有有用信息？前面有讲到过，没有用户访问 http server ，需要的程序没有运行，一直阻塞在那里等待客户端的访问连接，所以 go tool pprof 只能采集部分代码运行的信息，而这部分代码又没有消耗多少 cpu。

那怎么办？

一个方法就是用 http 测试工具模拟用户访问。这里用 https://github.com/rakyll/hey 这个工具。
安装 hey：

> go get -u github.com/rakyll/hey

安装完成后，进行 http 测试：

> hey -n 1000 http://localhost:8090/pprof-test

同时开启另一终端执行命令：

> go tool pprof http://localhost:8090/debug/pprof/profile?seconds=120

等待 120s 后，采集信息完成，如下图：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181737404-825994524.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181737404-825994524.png)

输入 `top` 命令查看统计信息：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181750407-678718185.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181750407-678718185.png)

可以看到用户定义的一个最耗时函数是：`main.cyclenum`。如果要查看这个函数最耗时部分代码，可以用 `list cyclenum` 命令查看。

我们这里是要生成一张图片，所以输入 `web` 命令生成图片：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181803508-766169522.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181803508-766169522.png)

在浏览器上查看 svg 图片：
[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182343933-1263805848.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182343933-1263805848.png)

(图片较大，只截取了部分)

这张图完整的展示了 `top` 命令的信息。

##### B. web 可视化

执行命令：

> go tool pprof -http=":8080" http://localhost:8090/debug/pprof/profile

同时开启另一终端执行测试命令：

> hey -n 200 -q 5 http://localhost:8090/pprof-test

上面 `go tool pprof` 执行完成后，会自动在浏览器打开一个 http 地址，http://localhost:8080/ui/，如下图：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181955303-144225281.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330181955303-144225281.png)
(截取部分图片)

这样就可以在web浏览器上查看分析数据了。

##### C. 火焰图

用 http 测试框架 [hey](http://github.com/rakyll/hey) 访问，命令为：

> hey -n 200 -q 5 http://localhost:8090/pprof-test

在压测的同时开启另一终端执行命令：

> go-torch -u [http://localhost:8090](http://localhost:8090/)

来生成火焰图。

运行命令时在终端输出了信息 ：

> Run pprof command: go tool pprof -raw -seconds 30 http://localhost:8090/debug/pprof/profile

可以看到 `go-torch` 的原始命令也是用到了 `go tool pprof`

上面这个命令默认生成了 torch.svg 的火焰图文件，如下：

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182015328-1804123499.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182015328-1804123499.png)
(截取一部分图展示)

点击方块可以查看更详细信息:

[![img](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182028695-762559956.png)](https://img2020.cnblogs.com/blog/650581/202103/650581-20210330182028695-762559956.png)

## 参考[#](https://www.cnblogs.com/jiujuan/p/14598141.html#2429723964)

- pprof
  - [README](https://github.com/google/pprof/blob/master/doc/README.md)
- [Profiling Go Programs](https://blog.golang.org/pprof)
- [runtime/pprof](https://golang.org/pkg/runtime/pprof/)
- [net/http/pprof](https://golang.org/pkg/net/http/pprof/)
- [go-torch](https://github.com/uber-archive/go-torch)
- [Flame Graph](https://github.com/brendangregg/FlameGraph)
- [http 压测工具 hey](http://github.com/rakyll/hey)

# Http测试

<img src="https://pic.imgdb.cn/item/6575ab2bc458853aef63fbaa.jpg" style="zoom:50%;" />

# 注释

<img src="https://pic.imgdb.cn/item/6575ac08c458853aef689d18.jpg" style="zoom:50%;" />