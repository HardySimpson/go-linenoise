# Go-Linenoise

2014-10-27可用

Go语言的[linenoise](https://github.com/antirez/linenoise)包装，附带了原来所有的C源代码

My Golang binding of [linenoise](https://github.com/antirez/linenoise), with all C code inside and modified for go-callback

# Linenoise

或许很多人知道readline，一个命令行编辑库，Bash,Mysql,Mutt都在使用。Antirez(Redis的作者)编写了一个更加轻量级的行编辑库，只有一千多行，就实现了在命令行上移动、增删、复制、粘贴、搜索等功能，用于Redis/MongoDB/Andriod。

通过linenoise，你可以完成这些功能：

* 上下方向键切换历史命令
* ctrl+r 可以搜索历史命令，很常用的一个
* ctrl+a 到行首
* ctrl+e 到行尾
* ctrl+u 删除到行首
* ctrl+k 删除到行尾
* ctrl+l 类似 clear 命令效果
* ctrl+y 粘贴
