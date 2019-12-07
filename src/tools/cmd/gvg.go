package main

import (
	"fmt"
	"flag"
)

/**
命令行tool

NAME:
   gvg - go version management by go

USAGE:
   gvg [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   list       list go versions
   install    install a go version
   info       show go version info
   use        select a version
   uninstall  uninstall a go version
   get        get the latest code
   uninstall  uninstall a go version
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
 */

 var h bool
 var v bool
func init() {
	flag.BoolVar(&v, "v", false, "print the version")

	flag.BoolVar(&h, "h", false, "show help")

	var configPath string
	flag.StringVar(&configPath, "config-path", "", "config file path")

	// 设置长短选项 -h / --help
	flag.BoolVar(&v, "version", false, "print the version")
	flag.BoolVar(&h, "help", false, "show help")
}

func main() {
	flag.Parse()
	fmt.Printf("help:%v\n",h)
	fmt.Printf("version:%v\n",v)
	fmt.Printf("gvg version control tool")

}
