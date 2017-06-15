package main

import "os"
import "fmt"


func printUsage() {
	fmt.Println(`
  Usage:
    htmls <html_template_path>
`)
}

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	template := os.Args[1]
	fmt.Println(template)
}

