# Light-view

先试试后端

## 8.11 hello.go

通过拖拽上传
choose your files 并不能找到src也不能找到go文件，bin找到gocode
但在文件夹里没有看到

## 8.13 tshhja.go

还是没有找到传文件时写的说明，只好复制在下面了：

调了很久才搞定。幸好tour的连贯性很强。以及想多了不是一个套一个，应该是二维就行了。

这种网络包似乎需要下载才能用？还是说能帮忙联网的插件我少下了？不对我运行是在终端里啊

PS：还是说那里不是写说明的？

我应该找找这里怎么换行:(

## 8.14
终于知道怎么换行了

还是没找到跟文件一起写的说明，不过标题倒是看到了。复制在下面：

刚看完教程就犯了一个超蠢的错误：`else`另起了一行，应该跟在`}`之后才对

这个练习再进一步把大小写不区分应该更实用，有空再说

## 8.15
根本是按递归的思路写的？闭包第一层还用不了参数？也不能用`main`里的`f`，只是把能调用的部分缩小了所以叫闭包？
> 紗霧：看过我的闭包教程 https://github.com/jihezhi/BackEnd/blob/master/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3closure.md 了吗？看完有什么地方不明白吗？另外Go不鼓励你使用递归。

是理解`f(i)`出错了，以为是`f(i)`需要调用一次`fibonacci`，实际上只是按`i`值填上而已？以及真没想到可以`mid, ans = ans, ans+mid`这么用，还是习惯一个个赋，不过这样真的很好用。

## 8.18
明天再细想强制类型转换的细节

## 8.20
如果只是输出就已经搞定了，然而不是。`print`和`println`并不是只有加换行符的区别，还有空格

`s := string(v[0]) + "." + string(v[1]) + "." + string(v[2]) + "." + string(v[3])`的输出只有符号没有数字`...`，但是
`s := string(v[0], ".", v[1], ".", v[2], ".", v[3])`却是语法错误？

## 8.21
`print`和`Sprint`两个字符间没有空格，非常理想。 // 紗霧：Go语言没有立场、没有预设，给你最大的自由度。因此它不会干擅自加空格这种事。

# 二周目
感谢时间机器

## 12.11
给自己建个教程还是很必要，二周目几乎等于重来

翻一周目找到了将黑框光标返回正常情况，按i即可（纱雾：你说的是Vim的使用吗？使用Vim应当按照Vim的思维方式，那种“不正常”情况才是正常情况，i是输入模式。）

## 12.14
感觉还是不太懂interface，但是照着练习前的例子仿写还是能做到的

## 12.17
不懂就是不懂，虽然能做出来~~按前面例子仿写的~~，但是根本没用上io【。（纱雾：你可能需要问我，interface不好理解。不过Gitter那边之前我讲过一次，你或许可以翻翻）
