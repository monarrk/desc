package main

import (
  "fmt"
  "io/ioutil"
  "bufio"
  "os"
  "strings"
  "flag"
  "os/exec"
)

func open_vim() {
  // use vim to edit DESCRIPTION
  cmd := exec.Command("vim", flag.Arg(0) + "/DESCRIPTION")
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr
  if cmd.Run() != nil {
    fmt.Printf("ERR: %v", cmd.Run())
  }
}

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
    if flag.Arg(1) == "edit" ||  flag.Arg(1) == "-e" {
      open_vim()
      return
    }
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
      fmt.Printf("%v\n", err)
      return
  }

  for _, f := range files {
    if strings.HasPrefix(f.Name(), "README") {
      content, _ := ioutil.ReadFile(dir + "/" + f.Name())
      fmt.Printf("Found a README:\n%s\n", content)
      return
    }
  }
  fmt.Print("No DESCRIPTION or README found\n\nCreate a DESCRIPTION? [Y/n]: ")
  make_desc := read()
  if strings.ToLower(make_desc) != "n" {
    open_vim()
  } 
}
