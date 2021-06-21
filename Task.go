package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id int
	Name string
	DateOfStart time.Time
	DateOfEnd time.Time
	Information string
	Status Status
}

type Status string

const (
	Undefined Status = "Undefined"
	Expect Status = "Expect"
	Finished Status = "Finished"
)

func CreateTask(name string, dateOfStart time.Time, dateOfEnd time.Time, information string, status Status) (Task, error){
	if len(name) == 0{
		return Task{}, fmt.Errorf("Invalid name\n")
	}

	if status == ""{
		return Task{}, fmt.Errorf("Invalid status\n")
	}

	return Task{
		Id: 0,
		Name: name,
		DateOfStart: dateOfStart,
		DateOfEnd: dateOfEnd,
		Information: information,
		Status: status,
	}, nil
}

func (t *Task) Update(newName string, newDateOfStart time.Time, newDateOfEnd time.Time, newInformation string, newStatus Status) error {
	if len(newName) == 0{
		return fmt.Errorf("Invalid name\n")
	}

	if newStatus == ""{
		return fmt.Errorf("Invalid status\n")
	}

	t.Name = newName
	t.DateOfStart = newDateOfStart
	t.DateOfEnd = newDateOfEnd
	t.Information = newInformation
	t.Status = newStatus
	return nil
}

func (t Task) print(){
	fmt.Println("Id:",t.Id)
	fmt.Println("Name:",t.Name)
	fmt.Println("Date of start:",t.DateOfStart)
	fmt.Println("Date of end:",t.DateOfEnd)
	fmt.Println("Information:",t.Information)
	fmt.Println("Status:",t.Status)
}