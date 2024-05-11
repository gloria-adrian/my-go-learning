package main

import (
"bufio"
"fmt"
"io/ioutil"
"os"
"path/filepath"
"strings"
)

func main(){
//Prompt the user to enter their name
fmt.Print("Enter your name: ")
reader :=
bufio.NewReader(os.Stdin)
name, _ :=
reader.ReadString('\n')
name = strings.TrimSpace(name)
// wite a name to a file
desktop,err:=os.UserHomeDir()
if err != nil {
fmt.Println("Error:", err)
return
}
filePath :=
filepath.Join(desktop, "Desktop","people.txt")
err = writeToFile(filePath, name)
if err != nil {
fmt.Println("Error:", err)
return
}
fmt.Println("Name saved successfully!")
//Ask the user if they want to search for a name in the file
fmt.Print("Do you want to search for a name in the file?(yes/no):")
searchOption, _ :=
reader.ReadString('\n')
searchOption =
strings.TrimSpace(strings.ToLower(searchOption))
if searchOption == "yes" {
  fmt.Print("Enter the name you want to search for: ")
  searchName, _ :=
  reader.ReadString('\n')
  searchName =
  strings.TrimSpace(searchName)

  //search for the name in the file
  found, err :=
  searchInFile(filePath, searchName)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  if found {
    fmt.Println("Name not found in the file!")
  }else {
    fmt.Println("Name not found in the file.")
  }
}
}
func writeToFile(filePath,name string) error {
  f,err :=os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
  if err != nil {
    return err
  }
  defer f.Close()
  _, err = f.WriteString(name + "\n")
  if err != nil {
    return err
  }
  return nil
}
func searchInFile(filePath,searchName string)(bool,error) {
  content,err :=
  ioutil.ReadFile(filePath)
  if err !=nil {
    return false,err
  }
  names :=
  strings.Split(string(content),"\n")
  for _,n := range names {
    if n ==searchName{
      return true,nil
    }
  }
  return false, nil
}
