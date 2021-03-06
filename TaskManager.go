package main

import (
	"fmt"
)

var Dictionary = TaskList{}
var key = staticKey()
func staticKey() (f func()(int)){
	var i int
	f = func()(int){
		i++
		return i
	}
	return
}

type TaskList []Task

func (list *TaskList) FindTask(index int) (*Task, error){
	for key,n := range *list {
		if n.Id == index{
			var tmp = *list
			return &tmp[key], nil
		}
	}

	return nil,fmt.Errorf("No value found\n")
}

func (list *TaskList) AddTask(task *Task){
	task.Id = key()
	*list = append(*list, *task)
}

func (list *TaskList) DeleteTask(index int) error{
	_, err:= list.FindTask(index)
	if err != nil{
		return err
	}

	var dict = *list
	index--
	dict[len(dict)-1], dict[index] = dict[index], dict[len(dict)-1]
	*list = dict[:len(dict)-1]
	return nil
}

func (list TaskList) print(){
	for _, val := range list{
		val.print()
	}
}