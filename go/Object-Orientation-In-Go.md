# Go语言中的面向对象
> 什么是面向对象呢？

这里我试着分享我对 go 语言如何实现面向对象的理解，以及让它看起来比其他传统的面向对象语言更加面向对象。

在之前的一篇文章中，我探讨了用函数来表达一切的想法，而对象实际上只是一种安全的方式来表示一组函数，它总是与相同的闭包一起运行，并且可以在相同的状态下运行。

安全地表示在同一个状态下运行的一组函数是非常有用的，但是如何真正的创建具有多种不同实现的抽象呢？

对象本身不支持这一点，每种对象类型都有它自己的函数集（我们称这组函数为方法）。这被称为多态，但是当我从传统的面向对象的概念（关于对昂和类型）离开，我发现，在设计软件时，考虑协议而不是多态性更符合应该关注的内容（稍后会详细介绍）。


## 如何定义一个协议？
首先让我们定义一下什么是协议，对于我来说协议就是实现预期结果所需的一系列操作。

这个概念似乎是晦涩难懂的，让我举个栗子，一个更容易理解的关于一个抽象概念的栗子，它就是 I/O。

假如你想读取一个文件，协议将是：

* 打开文件

* 读取文件的内容

* 关闭文件

为了实现读取一个文件所有内容的简单需求，你需要这三个操作，因此这些操作就构成了你“阅读文件”的协议。

现在让我们来看看这个例子，并且一起完成剩余的实现。

如果函数在一门编程语言中是一等公民，结构化函数和结构化数据并无区别。

我们有一种合成数据的方法，那就是结构体(structs)，我们也可以使用这种方法来合成函数，例如：

```go
type Reader func(data []byte) (int, error)

type Closer func() error

type ReadCloser struct {
    Read Reader
    Close Closer
}

type Opener func() (*ReadCloser, error)
func useFileProtocol(open Opener) {
    f, _ := open()
    data := make([]byte, 50)
    f.Read(data)
    f.Close()
}

func main() {
    useFileProtocol(func() (*ReadCloser, error) {
        return &ReadCloser{}, nil
    })
}
```

使用编译时安全的方法来表达一个协议是非常困难的（假如这并不是不可能的）。为了说明这个问题，这个例子造成了一个分割错误。

另一个问题是实现这个协议的代码需要知道协议被显示地实现，以便正确地初始化结构体（就像继承的过程一样），或者把结构体的初始化委托给系统的其他部分，该部分将围绕如何正确的初始化结构体展开。当你考虑实现多种协议的相同函数时，将会变得更加糟糕。

需要一些第三个协议的对象需要一种方法：

* 清楚的表达出它所需要的协议。

* 确保当它开始与一个实现交互时，没有任何功能丢失。

实现服务的对象需要：

* 能够安全地表示它具有满足协议的必须功能。

* 能够满足一个协议，甚至对它不是很清晰的了解。

不需要两个对象交互来实现一个特定类型的共同目标，最重要的是它们之间的协议是否匹配。

这就是 Go interfaces（接口）的由来。它提供了一个编译时安全的方式来表现协议，通过适当的函数来消除初始化结构的所有模板。它会为我们初始化结构，甚至对初始化结构优化一次，这在 Go 中被称为 iface 。类似于 C++ 的 vtable 。

它还允许代码更加的解耦，因为你不需要了解定义这个接口 （interface）并且实现这个接口的包。与相同的编译安全型语言Java、C++比较，在 Go 允许的基础上，Go更加的灵活。


让我们重新审视之前的带有接口（interfaces）的文件协议：

```go
package main

type Reader interface {
    Read(data []byte) (int, error)
}

type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}

type Opener func() (ReadCloser, error)

type File struct {}

func (f *File) Read(data []byte)(int, error){
    return 0, nil
}

func (f *File) Close() error {
    return nil
}

func useFileProtocol(open Opener) {
    f, _ := open()
    data := make([]byte, 50)
    f.Read(data)
    f.Close()
}

func main(){
    useFileProtocol(func() (ReadCloser, error) {
                return &File{}, nil
    })
}
```

一个关键的不同是, 这段代码使用了接口（interface），现在是安全的。那个`useFileProtocol`不必担心调用函数是否为nil，go编译器将会创建一个结构体，通过一个 `iface`描述符来保持一个指针，该指针具有满足协议的所有功能。它会按类型<->接口的每一个匹配项执行此操作，就像它被使用的那样（它第一次初始化使用的那样）。


如果你这样做，仍然会造成一个分割错误，如下：

```go
useFileProtocol(func() (ReadCloser, error) {
        var a ReadCloser
        return a, nil
})
```






----------------

via: https://katcipis.github.io/blog/object-orientation-go/

作者：[TIAGO KATCIPIS](https://katcipis.github.io/)
译者：[译者ID](https://github.com/MengYP)
校对：[校对者ID](https://github.com/校对者ID)

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创编译，[Go 中文网](https://studygolang.com/) 荣誉推出
