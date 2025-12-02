package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

func (S *Storage[T]) Save(data T) error {
	filedata, err := json.MarshalIndent(data,"","    ")
	if err!=nil{
		return err
	}
	return os.WriteFile(S.Filename,filedata,0644)
}

func (S *Storage[T]) Load(data *T) error{
	filedata,err := os.ReadFile(S.Filename)
	if err!=nil{
		return err
	}
	return json.Unmarshal(filedata,data)
}