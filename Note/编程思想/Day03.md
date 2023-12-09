# 资源管理与出错处理

<img src="https://pic.imgdb.cn/item/6572d513c458853aef222799.jpg" style="zoom: 50%;" />

# defer

* 确保在函数结束时发生
* 参数在defer语句时计算
* defer列表为先进后出（类似栈）

## 何时调用

* open、close
* lock，unlock
* printHeader，printFooter