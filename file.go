/*************************************************************************
    > File Name: file.go
    > Author: Kee
    > Mail: chinboy2012@gmail.com 
    > Created Time: 2017.11.29
 ************************************************************************/
package kfiles;

import (
  "os"
  "io"
  "io/ioutil"
)

const(
  APPEND = 1
)

func GetFile (file string) (string, error) {
  fl, err := os.Open(file)

  if err != nil {
    return "", err
  }

  defer fl.Close()
  
  content, err := ioutil.ReadAll(fl)

  return string(content), err
}

func putFile (file string, content string, opt int) (bool, error) {
  if true != Exists(file) {
     _, err := os.Create(file)

     if err != nil {
       return false, err
     }
  }

  opts := 0

  if opt == APPEND {
    opts = os.O_APPEND | os.O_WRONLY
  } else {
    opts = os.O_WRONLY
  }

  fl, err := os.OpenFile(file, opts, 0755)

  n, err := io.WriteString(fl, content)
  defer fl.Close()

  if err != nil {
    return false, err
  }
  return bool(n > len(content)), err
}

func Put (file string, content string) (bool, error) {
  return putFile(file, content, 0)
}

func Append (file string, content string) (bool, error) {
  return putFile(file, content, APPEND)
} 

func Exists (file string) bool {
  if _, err := os.Stat(file); os.IsNotExist(err) {
    return false;
  }
  return true;
}
