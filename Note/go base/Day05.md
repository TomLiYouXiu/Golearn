# 面相对象

* go语言仅支持封装，不支持继承和多态
* go语言没有class，只有struct

<img src="https://pic.imgdb.cn/item/656801c9c458853aeff6dab4.jpg" style="zoom:50%;" />

~~~go
package main

import "fmt"

// 自定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//new 内建函数
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
~~~

## 工厂函数

![](https://pic.imgdb.cn/item/656802a4c458853aeff951e3.jpg)

在go中局部变量的地址也是可以被返回的

**在其他语言中局部变量一般是分配在栈上，函数一旦退出，局部变量就是被销毁，如果要是想要长期使用的话就需要分配在堆上，但是需要手动释放**

## 为结构定义方法

<img src="https://pic.imgdb.cn/item/656806fbc458853aef045453.jpg" style="zoom:50%;" />

<img src="C:\Users\57343\AppData\Roaming\Typora\typora-user-images\image-20231130115247139.png" alt="image-20231130115247139" style="zoom:50%;" />

## 值接收者vs指针接受者

<img src="https://pic.imgdb.cn/item/656822d6c458853aef5bf6e6.jpg" style="zoom:50%;" />

# 封装（包）

<img src="https://pic.imgdb.cn/item/65682361c458853aef5dd257.jpg" style="zoom:50%;" />

<img src="https://pic.imgdb.cn/item/6568239cc458853aef5ea66c.jpg" style="zoom:50%;" />

<img src="https://pic.imgdb.cn/item/656823b5c458853aef5f05e0.jpg" style="zoom:50%;" />

## 如何扩充系统类型或者别人的类型

<img src="https://pic.imgdb.cn/item/656826efc458853aef6a140c.jpg" style="zoom:50%;" />

<img src="https://pic.imgdb.cn/item/65685c48c458853aef166b92.jpg" style="zoom:50%;" />