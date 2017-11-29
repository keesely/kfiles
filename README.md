# kfiles
Go文件操作 - Go学习实例

```
go get github.com/keesely/kfiles
```

## Usage:

```
 import "github.com/keesely/kfiles"
 _, err := Put("./test.txt", "这是一段测试文本")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("文本写入")
  }

  _, err = Append("./test.txt", "\n追加一段文本")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("文本追加")
  }

  exists := Exists("./test.txt")
```
