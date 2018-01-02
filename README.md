# Hello-World
This is a go lang project, that is used to practice.

# 2017年12月29日
## 下载并安装go tools 配置环境变量如下：
```
#go lang
#export GOROOT=/usr/local/go
export GOPATH=$HOME/workspace-go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
$GOROOT是默认路径，故注释掉，写不写都一样。[学习链接](https://golang.org/doc/install)

## 完成自己第一个程序和完成自己的第一个lib

### workspaces介绍
workspaces默认的工作空间，$HOME/go on Unix, $home/go on Plan 9, and %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows。
都在家目录下的go文件夹。

通过$GOPATH环境变量修改默认的workspaces的路径，如上图环境变量中指定了新的目录

workspaces目录下必须包含如下的目录结构（必须要有这3个文件夹）：
```
\src contains Go source files,
\pkg contains package objects, and
\bin contains executable commands.
```

### import paths
导入路径是唯一确定一个包的导入路径。


(类似java的import package)。\src下面可以有多个包，每个包是一个go工程，每个包有独立的版本控制管理。所有工程编译的命令都放在\bin目录下，所有工程产生的lib都放在\pkg中。

---
[go和其他语言的不同(Java)](/doc/diff.md)

---
[gotour的练习题](/doc/goTour.md)




