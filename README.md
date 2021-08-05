<p align="center">
	<a href="https://gokit.dianjiu.org.cn/"><img src="https://dianjiu.co/gotool.jpg" width="45%"></a>
</p>
<p align="center">
	<strong>A set of Go tool kit.</strong>
</p>
<p align="center">
	<a href="https://gokit.dianjiu.org.cn">https://gokit.dianjiu.org.cn/</a>
</p>

[**🌎 Chinese Documentation**](README.zh-CN.md)

-------------------------------------------------------------------------------

## 📚 Introduction
GoKit is a small and comprehensive Go tool class library. Through static method encapsulation, it reduces the learning cost of related APIs and improves work efficiency, so that Go has a Java language Hutool experience.

The tools and methods in GoKit come from the meticulous craftsmanship of each user. It covers all aspects of the underlying code of Go development. It is not only a weapon for solving small problems in large-scale project development, but also an efficiency responsibility in small-scale projects;

GoKit is a friendly alternative to the "util" package in the project. It saves developers the encapsulation time of public classes and public tool methods in the project, enables development to focus on business, and at the same time avoids bugs caused by imperfect encapsulation.

### 🎁 The origin of the name GoKit

GoKit = Go + Kit, is a Go language tool set created by the [Dianjiu Open Source](https://dianjiu.org.cn) team, "Go" stands for golang, and Kit stands for toolkit.

### 🍺 GoKit changes the coding method

GoKit's goal is to use a tool method to replace a piece of complex code, so as to avoid the problem of "copy and paste" code to the greatest extent, and completely change the way we write code.

Take the calculation of `MD5` as an example:

- [Before] Open the search engine -> search "Go MD5 encryption" -> open a blog -> copy and paste -> change to use
- [Now] Introduce GoKit -> md5:=security.MD5("Hello World")

>   GoKit exists to reduce the cost of code search and avoid bugs caused by uneven code on the Internet.

-------------------------------------------------------------------------------

## 📦 Use

Add the following to the dependencies of the Go project:

```bash
go get github.com/dianjiu/gokit
```

Take the introduction of components under the ip package as an example, call the GetLocalIPv4() method to obtain the local IP as a demonstration
```go
package main

import (
	"github.com/dianjiu/gokit/ip"
	"log"
)

func main() {
	ip := ip.GetLocalIPv4()
	log.Printf("Local IP: %v\n", ip)
}
```

result

>   2021/07/30 15:49:03 Local IP: 192.168.56.1

## 🛠️ Module

A Go basic tool class, which encapsulates methods such as files, streams, encryption and decryption, transcoding, regular, JSON, UUID, etc., to form various Go tool classes, and provides the following components:
| Module                 |     Introduce                                                                           |
| -------------------|---------------------------------------------------------------------------------- |
| assertions | Assertion processing tools. |
| bytes | Byte byte conversion tools. |
| caller | Get function name path tools. |
| compress | Compression and decompression tools, providing support for tar and zip formats . |
| file | Common file-related tools such as file path, whether the file exists, whether it is a folder. |
| hello       | The Hello World package of GoKit provides SayHello() method support. |
| http | http tools, commonly used Get and Post requests |
| ip | IP tools |
| json        |     JSON tools                                                            |
| logger | Provide log output tools with multiple risk levels |
| mail | Mail sending tools |
| map | map collection tools |
| poi         |     Encapsulation for Excel and Word in POI               |
| random | Random number tools |
| security | Provide encryption and decryption methods such as MD5, AES, RSA, etc. |
| set | Set collection tools |
| slice | Slicing tools |
| socket      |     Socket encapsulation of NIO and AIO based on Go   |
| string | string tool |
| system      |     System parameter call tool                          |
| time | Provide commonly used date tools and date formatting packages |
| uuid | uuid tools |

Each module can be introduced and used separately according to requirements, as follows: 

```
import (
	"github.com/dianjiu/GoKit/hello"
)
```

It can also be used by introducing all modules of `GoKit`, as follows: 

```
import (
	"github.com/dianjiu/GoKit"
)
```

-------------------------------------------------------------------------------

## 📝Documentation 

[📘 Get started quickly ](https://GoKit.dianjiu.org.cn/docs/doc01.html)

[📙 Update log ](https://GoKit.dianjiu.org.cn/version/ver01.html)

-------------------------------------------------------------------------------

## 🏗️ Contribute 

### 🎋 Branch description 

The source code of GoKit is divided into two branches, with the following functions: 

| Branch | Effect                                                  |
|-----------|---------------------------------------------------------------|
| master | The main branch, the branch used by the release version, is consistent with the submission of the central library, and does not receive any pr or modification |
| dev    | Development branch, default is the SNAPSHOT version of the next version, accept modification or pr |

### 🐞 Provide bug feedback or suggestions 

Please indicate the Go version, GoKit version, and related dependency library version you are using when submitting problem feedback. 

- [Gitee issue](https://gitee.com/dianjiu/gokit/issues)
- [Github issue](https://github.com/dianjiu/gokit/issues)
- [Gitlab issue](https://gitlab.com/dianjiu/gokit/issues)


### 🧬 Steps to contribute code 

Email to dianjiu@dianjiu.cc to get a developer account 

### 📐PR principles 

GoKit welcomes anyone to contribute code to GoKit. The pr (pull request) that needs to be submitted conforms to some specifications, and the specifications are as follows: 

1. The comments are complete, especially each new method should be marked with the method description, parameter description, return value description and other information according to the Go document specification. If you want, you can also add your name. 
2. The GoKit project recommends that Go programmers use vscode for development. Please pull request to the `dev` branch. 
3. Do not use third-party library methods for newly added methods. GoKit follows the principle of no dependency (unless you add method tools to extra modules). 
4. GoKit will use a new branch after version 1.x: `master` is the main branch, which means that the version of the central library has been released. This branch does not allow pr or modification. 

## ⭐ Star GoKit
[![Giteye chart](https://chart.giteye.net/gitee/dianjiu/gokit/5BTC4MSR.png)](https://giteye.net/chart/5BTC4MSR)

[![Stargazers over time](https://starchart.cc/dianjiu/gokit.svg)](https://starchart.cc/dianjiu/gokit)

## 📌 Comminicate 

### 🧡  EMail dianjiu@dianjiu.cc
