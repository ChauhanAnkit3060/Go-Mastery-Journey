package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type TODO struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []TODO

func (todo *Todos) add(title string){
	t := TODO{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todo = append(*todo,t)
}

func (todo *Todos) validateIndex(ind int) error{
	if ind < 0 || ind >= len(*todo){
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	} 
	return nil
}

func (todo *Todos) delete(ind int) error{
	if err := todo.validateIndex(ind);err!=nil{
		fmt.Println(err)
		return err
	}
	
	t := *todo
	*todo = append(t[:ind],t[ind+1:]...)
	return nil
}

func (todo *Todos) toggle(ind int) error{
	if err := todo.validateIndex(ind);err!=nil{
		fmt.Println(err)
		return  err
	}
	t := *todo
	
	isCompleted := t[ind].Completed
	if !isCompleted{
		completedAt := time.Now()
		t[ind].CompletedAt = &completedAt
	}

	t[ind].Completed = !t[ind].Completed

	return nil
}

func (todo *Todos) edit(ind int, title string) error{
	if err := todo.validateIndex(ind);err!=nil{
		return  err
	}
	t := *todo

	t[ind].Title = title

	return nil
}

func (todo *Todos) print(){
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Title","Status","Created at","Completed On")

	
	for ind,t := range *todo{
		completed := "❌"
		completedAT := ""
		if t.Completed{
			completed = "✅"
			if t.CompletedAt != nil{
				completedAT = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(ind),t.Title,completed,t.CreatedAt.Format(time.RFC1123),completedAT)
	}

	table.Render()
}