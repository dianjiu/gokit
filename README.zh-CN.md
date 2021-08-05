<p align="center">
	<a href="https://GoKit.dianjiu.org.cn/"><img src="https://dianjiu.co/gotool.jpg" width="45%"></a>
</p>
<p align="center">
	<strong>一套 Go 工具包。</strong>
</p>
<p align="center">
	<a href="https://gkit.dianjiu.org.cn">https://gkit.dianjiu.org.cn/</a>
</p>

-------------------------------------------------------------------------------

[**🌎 English Documentation**](README.zh-CN.md)

-------------------------------------------------------------------------------

## 📚 简介
GoKit是一个小而全的Go工具类库，通过静态方法封装，降低相关API的学习成本，提高工作效率，使Go拥有Java语言Hutool般地体验。

GoKit中的工具方法来自每个用户的精雕细琢，它涵盖了Go开发底层代码中的方方面面，它既是大型项目开发中解决小问题的利器，也是小型项目中的效率担当；

GoKit是项目中“util”包友好的替代，它节省了开发人员对项目中公用类和公用工具方法的封装时间，使开发专注于业务，同时可以最大限度的避免封装不完善带来的bug。

### 🎁 GoKit名称的由来

GoKit = Go + Kit，是由[点九开源](https://dianjiu.org.cn)团队创建的Go语言工具集，“Go”是golang的表示，Kit表示工具包。

### 🍺 GoKit改变coding方式

GoKit的目标是使用一个工具方法代替一段复杂代码，从而最大限度的避免“复制粘贴”代码的问题，彻底改变我们写代码的方式。

以计算`MD5`为例：

- 【以前】打开搜索引擎 -> 搜“Go MD5加密” -> 打开某篇博客-> 复制粘贴 -> 改改使用
- 【现在】引入GoKit -> md5:=security.MD5("Hello World")

>   GoKit的存在就是为了减少代码搜索成本，避免网络上参差不齐的代码出现导致的bug。

-------------------------------------------------------------------------------

## 📦 使用

在Go项目的dependencies中加入以下内容:

```bash
go get github.com/dianjiu/gokit
```

以引入ip包下组件为例，调用GetLocalIPv4()方方法获取本地IP作为演示

```go
package main

import (
	"github.com/dianjiu/gokit/ip"
	"log"
)

func main() {
	ip := ip.GetLocalIPv4()
	log.Printf("本机IP: %v\n", ip)
}
```

运行结果

>   2021/07/30 15:49:03 本机IP: 192.168.56.1

## 🛠️ 组件

一个Go基础工具类，对文件、流、加密解密、转码、正则、JSON、UUID等方法进行封装，组成各种Go工具类，同时提供以下组件：

| 模块                |     介绍                                                                          |
| -------------------|---------------------------------------------------------------------------------- |
| assertions | 断言处理工具类 |
| bytes | byte字节转换工具类 |
| caller | 获取函数名路径工具类 |
| compress | 压缩、解压缩封装，提供tar和zip两种格式支持 |
| file | 文件路径、文件是否存在、是否为文件夹等常用文件相关工具类 |
| hello       |     GoKit的Hello World封装，提供SayHello()方法支持                |
| http | http工具类，常用的Get、Post请求 |
| ip | IP工具类 |
| json        |     JSON操作实现                                                                       |
| logger | 提供多种风险等级的日志输出封装 |
| mail | 邮件发送工具类 |
| map | map集合工具类 |
| poi         |     针对POI中Excel和Word的封装                                                       |
| random | 随机数工具类 |
| security | 提供MD5、AES、RSA等方式的加解密方法 |
| set | set集合工具类 |
| slice | 切片工具类 |
| socket      |     基于Go的NIO和AIO的Socket封装                                                   |
| string | string字符串工具类 |
| system      |     系统参数调用封装                                                     |
| time | 提供常用的日期工具类及日期格式化封装 |
| uuid | uuid工具类 |

可以根据需求对每个模块单独引入使用，如下：

```
import (
	"github.com/dianjiu/GoKit/hello"
)
```

也可以通过引入`GoKit`所有模块使用，如下：

```
import (
	"github.com/dianjiu/GoKit"
)
```

-------------------------------------------------------------------------------

## 📝文档

[📘 快速上手](https://GoKit.dianjiu.org.cn/docs/doc01.html)

[📙 更新日志](https://GoKit.dianjiu.org.cn/version/ver01.html)

-------------------------------------------------------------------------------

## 🏗️ 贡献

### 🎋 分支说明

GoKit的源码分为两个分支，功能如下：

| 分支       | 作用                                                          |
|-----------|---------------------------------------------------------------|
| master | 主分支，release版本使用的分支，与中央库提交一致，不接收任何pr或修改 |
| dev    | 开发分支，默认为下个版本的SNAPSHOT版本，接受修改或pr                 |

### 🐞 提供bug反馈或建议

提交问题反馈请说明正在使用的Go版本、GoKit版本和相关依赖库版本。

- [Gitee issue](https://gitee.com/dianjiu/gokit/issues)
- [Github issue](https://github.com/dianjiu/gokit/issues)
- [Gitlab issue](https://gitlab.com/dianjiu/gokit/issues)


### 🧬 贡献代码的步骤

邮件联系 dianjiu@dianjiu.cc 获取开发者账号

### 📐 PR遵照的原则

GoKit欢迎任何人为GoKit贡献代码，需要提交的pr（pull request）符合一些规范，规范如下：

1. 注释完备，尤其每个新增的方法应按照Go文档规范标明方法说明、参数说明、返回值说明等信息，如果愿意，也可以加上你的大名。
2. GoKit项目建议Go程序员使用vscode开发，请pull request到`dev`分支。
3. 新加的方法不要使用第三方库的方法，GoKit遵循无依赖原则（除非在extra模块中加方法工具）。
4. GoKit在1.x版本后将使用新的分支：`master`是主分支，表示已经发布中央库的版本，这个分支不允许pr，也不允许修改。

## ⭐ Star GoKit
## 📌 交流

### 🧡 公众号

![点九先生](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210729173820.jpg)

### 🧡 交流群

![点九开源](https://gitee.com/dianjiu/typora-imgs/raw/master/imgs/20210729173825.jpg)