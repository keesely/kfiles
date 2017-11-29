/*************************************************************************
    > File Name: test.go
    > Author: Kee
    > Mail: chinboy2012@gmail.com 
    > Created Time: 2017.11.29
 ************************************************************************/
package kfiles

import (
  "fmt"
  //"keesely/kfiles"
  "testing"
)

func Test (t *testing.T) {
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

  if true == exists {
    fmt.Println("文件已存在")
    content, _ := Get("./test.txt")
    
    fmt.Println("文本内容: \n```\n", content, "\n```");
  } else {
    fmt.Println("文件不存在")
  }

}
