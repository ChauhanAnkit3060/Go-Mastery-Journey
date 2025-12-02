package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Del    int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add,"add","","Add task to TODO, specify title")
	flag.StringVar(&cf.Edit,"edit","","Edit a TODO task by index & specify title, id:new_title")
	flag.IntVar(&cf.Del,"del",-1,"Specify a TODO by index to delete")
	flag.IntVar(&cf.Toggle,"toggle",-1,"Specify a TODO by index to toggle")
	flag.BoolVar(&cf.List,"list",false,"List all TODO task")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute (todos *Todos){
	switch{
	case cf.List :
		todos.print()

	case cf.Add != "" :
		todos.add(cf.Add)

	case cf.Edit != "" :
		parts := strings.SplitN(cf.Edit,":",2)
		if len(parts) != 2{
			fmt.Println("Error : invalid format for edit, Please use id:new_title")
			os.Exit(1)
		}
		ind,err := strconv.Atoi(parts[0])
		if err!=nil{
			fmt.Println("Error : invalid index for edit")
			os.Exit(1)
		}
		todos.Edit(ind,parts[1])
	case cf.Toggle > -1 :
		if err := todos.Toggle(cf.Toggle); err != nil{
			fmt.Println("Error : invalid index to toggle : ",err)
			os.Exit(1)
		}
	case cf.Del > -1 :
		if err := todos.delete(cf.Del); err != nil{
			fmt.Println("Error : invalid index to toggle : ",err)
			os.Exit(1)
		}
	default : fmt.Println("Invalid command")
	}
}