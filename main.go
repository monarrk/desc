package main

import (
  "fmt"
  "io/ioutil"
  "bufio"
  "os"
  "strings"
  "flag"
)

func check_for_file(file string) bool {
  if _, err := os.Stat(file); err == nil {
    return true
  } else {
    return false
  }
}

func read() string {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func main() {
  flag.Parse()
  var dir string
  if flag.Arg(0) == "" {
    fmt.Print("Enter the dir: ")
    dir = read()
  } else {
    dir = flag.Arg(0)
  }

  //check for a DESCRIPTION file
  if check_for_file(dir + "/DESCRIPTION") {
    content, _ := ioutil.ReadFile(dir + "/DESCRIPTION")
    fmt.Printf("DESCRIPTION\n%s\n", content)
    return
  }

  // check for readmes
  files, err := ioutil.ReadDir(dir)
  if err != nil {
      fmt.Printf("%v", err)
      return
  }

  for _, f := range files {
    if strings.HasPrefix(f.Name(), "README") {
      content, _ := ioutil.ReadFile(dir + "/" + f.Name())
      fmt.Printf("Found a README:\n%s\n", content)
      return
    }
  }
  fmt.Println("No DESCRIPTION or README found")
}
