<p align="center">
	<a href="https://gotool.dianjiu.org.cn/"><img src="https://dianjiu.co/gotool.jpg" width="45%"></a>
</p>
<p align="center">
	<strong>A set of tools that keep Go sweet.</strong>
</p>
<p align="center">
	<a href="https://Gotool.cn">https://gotool.dianjiu.org.cn/</a>
</p>


-------------------------------------------------------------------------------

[**English Documentation**](README-EN.md)

-------------------------------------------------------------------------------

## 简介
Gotool是一个小而全的Go工具类库，通过静态方法封装，降低相关API的学习成本，提高工作效率，使Go拥有函数式语言般的优雅，让Go语言也可以“甜甜的”。

Gotool中的工具方法来自每个用户的精雕细琢，它涵盖了Go开发底层代码中的方方面面，它既是大型项目开发中解决小问题的利器，也是小型项目中的效率担当；

Gotool是项目中“util”包友好的替代，它节省了开发人员对项目中公用类和公用工具方法的封装时间，使开发专注于业务，同时可以最大限度的避免封装不完善带来的bug。

### Gotool名称的由来

Gotool = Go + tool，是由点九开源团队创建的Go语言工具集，“Go”是golang的表示，tool表示工具。

### Gotool如何改变我们的coding方式

Gotool的目标是使用一个工具方法代替一段复杂代码，从而最大限度的避免“复制粘贴”代码的问题，彻底改变我们写代码的方式。

以计算MD5为例：

- 【以前】打开搜索引擎 -> 搜“Go MD5加密” -> 打开某篇博客-> 复制粘贴 -> 改改好用
- 【现在】引入Gotool -> SecureUtil.md5()

Gotool的存在就是为了减少代码搜索成本，避免网络上参差不齐的代码出现导致的bug。

-------------------------------------------------------------------------------

## 包含组件
一个Go基础工具类，对文件、流、加密解密、转码、正则、线程、XML等JDK方法进行封装，组成各种Util工具类，同时提供以下组件：

| 模块                |     介绍                                                                          |
| -------------------|---------------------------------------------------------------------------------- |
| gotool-aop         |     JDK动态代理封装，提供非IOC下的切面支持                                              |
| gotool-bloomFilter |     布隆过滤，提供一些Hash算法的布隆过滤                                                |
| gotool-cache       |     简单缓存实现                                                                     |
| gotool-core        |     核心，包括Bean操作、日期、各种Util等                                               |
| gotool-cron        |     定时任务模块，提供类Crontab表达式的定时任务                                          |
| gotool-crypto      |     加密解密模块，提供对称、非对称和摘要算法封装                                          |
| gotool-db          |     JDBC封装后的数据操作，基于ActiveRecord思想                                         |
| gotool-dfa         |     基于DFA模型的多关键字查找                                                         |
| gotool-extra       |     扩展模块，对第三方封装（模板引擎、邮件、Servlet、二维码、Emoji、FTP、分词等）            |
| gotool-http        |     基于HttpUrlConnection的Http客户端封装                                            |
| gotool-log         |     自动识别日志实现的日志门面                                                         |
| gotool-script      |     脚本执行封装，例如Goscript                                                     |
| gotool-setting     |     功能更强大的Setting配置文件和Properties封装                                        |
| gotool-system      |     系统参数调用封装（JVM信息等）                                                      |
| gotool-json        |     JSON实现                                                                       |
| gotool-captcha     |     图片验证码实现                                                                   |
| gotool-poi         |     针对POI中Excel和Word的封装                                                       |
| gotool-socket      |     基于Go的NIO和AIO的Socket封装                                                   |

可以根据需求对每个模块单独引入，也可以通过引入`go get gotool`方式引入所有模块。

-------------------------------------------------------------------------------

## 文档 

[中文文档](https://gotool.dianjiu.org.cn/docs/)

-------------------------------------------------------------------------------

## 安装

### Github
在Go项目的dependencies中加入以下内容:

```
go get https://github.com/dianjiu/gotool
```

```
import (
	"github.com/dianjiu/gotool"
)
```


### 编译安装

访问Gotool的Gitee主页：[https://gitee.com/dianjiu/gotool](https://gitee.com/dianjiu/gotool)   
下载整个项目源码 然后进入Gotool项目目录执行：

```sh
 go run xxx
```

然后就可以使用

-------------------------------------------------------------------------------

## 添砖加瓦

### 分支说明

Gotool的源码分为两个分支，功能如下：

| 分支       | 作用                                                          |
|-----------|---------------------------------------------------------------|
| master | 主分支，release版本使用的分支，与中央库提交的jar一致，不接收任何pr或修改 |
| dev    | 开发分支，默认为下个版本的SNAPSHOT版本，接受修改或pr                 |

### 提供bug反馈或建议

提交问题反馈请说明正在使用的JDK版本呢、Gotool版本和相关依赖库版本。

- [Gitee issue](https://gitee.com/dianjiu/gotool/issues)
- [Github issue](https://github.com/dianjiu/gotool/issues)


### 贡献代码的步骤

邮件联系 dianjiu@dianjiu.cc 获取开发者账号

### PR遵照的原则

Gotool欢迎任何人为Gotool添砖加瓦，贡献代码，不过维护者是一个强迫症患者，为了照顾病人，需要提交的pr（pull request）符合一些规范，规范如下：

1. 注释完备，尤其每个新增的方法应按照Go文档规范标明方法说明、参数说明、返回值说明等信息，如果愿意，也可以加上你的大名。
2. Gotool项目建议Java程序员使用vscode开发，请pull request到`dev`分支。
3. 新加的方法不要使用第三方库的方法，Gotool遵循无依赖原则（除非在extra模块中加方法工具）。
4. Gotool在1.x版本后使用了新的分支：`master`是主分支，表示已经发布中央库的版本，这个分支不允许pr，也不允许修改。

-------------------------------------------------------------------------------



## 公众号

#### 欢迎关注点九先生的公众号

![点九先生](./imgs/wechat.jpg)

#### 点九开源项目QQ交流群

![点九开源](./imgs/qqchat.jpg)