package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Information string
	Status      Status
}

type Status string

const (
	Undefined Status = "Undefined"
	Expect Status = "Expect"
	Finished Status = "Finished"
)

func CreateTask(name string, startDate time.Time, endDate time.Time, information string, status Status) (Task, error){
	if len(name) == 0{
		return Task{}, fmt.Errorf("Invalid name\n")
	}

	if status == ""{
		return Task{}, fmt.Errorf("Invalid status\n")
	}

	return Task{
			Id:          0,
			Name:        name,
			StartDate:   startDate,
			EndDate:     endDate,
			Information: information,
			Status:      status,
		}, nil
}

func (t *Task) Update(newName string, newStartDate time.Time, newEndDate time.Time, newInformation string, newStatus Status) error {
	if len(newName) == 0{
		return fmt.Errorf("Invalid name\n")
	}

	if newStatus == ""{
		return fmt.Errorf("Invalid status\n")
	}

	t.Name = newName
	t.StartDate = newStartDate
	t.EndDate = newEndDate
	t.Information = newInformation
	t.Status = newStatus
	return nil
}

func (t Task) print(){
	fmt.Println("Id:",t.Id)
	fmt.Println("Name:",t.Name)
	fmt.Println("StartDate:",t.StartDate)
	fmt.Println("EndDate:",t.EndDate)
	fmt.Println("Information:",t.Information)
	fmt.Println("Status:",t.Status)
}