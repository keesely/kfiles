# kfiles
Go文件操作 - Go学习实例

### Get:

```
go get github.com/keesely/kfiles
```

### Usage:

```go
 import (
     "fmt"
     "github.com/keesely/kfiles"
 )
 
 // 写入文本
 _, err := kfiles.Put("./test.txt", "这是一段测试文本")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("文本写入")
  }

  // 追加写入
  _, err = kfiles.Append("./test.txt", "\n追加一段文本")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("文本追加")
  }

  // 判断是否存在
  exists := kfiles.Exists("./test.txt")

  if true == exists {
    fmt.Println("文件已存在")

    // 获取文本内容
    content, _ := kfiles.Get("./test.txt")
    
    fmt.Println("文本内容: \n```\n", content, "\n```");

  } else {
    fmt.Println("文件不存在")
  }
```
