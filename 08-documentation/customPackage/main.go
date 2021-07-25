// single line comment which describes our custom package
package customPackage

import(
  "fmt"
)

// single line commmen which describes out custom function
func CustomFunction() {
  fmt.Println("Namaste !")
}

/*

Usage:
  1. go doc <packageName> : go doc documentation
  2. go doc <packageName>.<medhodName> : go doc documentation.DoSomeStuff
  3. Ex: go doc fmt.Println
  4. Note the diff bw 'go doc', 'godoc' and 'godoc.org'
  5. godoc usage: locally hosts the website
  6. run in terminal: 'godoc -http :8080', navigate to 'http://localhost:8080/pkg/' in browser
  7. godoc -src fmt printf


cd /Users/username/workspace/backend/go/src/go_workspace/src/8-documentation/customPackage
Observe that the file nams should be main.go where which is actually a custom package
go doc customPackage
go doc customPackage.CustomFunction

go to godoc.org, paste this url and hit enter, you'll be able to see documentation of entered code repo
Example url: https://github.com/GoesToEleven/go-programming/tree/master/code_samples/007-documentation/01/mymath

How to create a document:
When there is a whole bunch of information (Ex: package fmt), document it in doc.go file.
Ex:  refer https://godoc.org/fmt
/*
  100 of lines of code
*\/ (remove escape character)
package <pkg_name>
*/