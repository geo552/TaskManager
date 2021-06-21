package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func homePage(res http.ResponseWriter, req *http.Request) {

	tmpl, _ := template.ParseFiles("HomePage.html")
	tmpl.Execute(res, Dictionary)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("data_input")

	task, err := CreateTask(name, time.Now(), time.Now(), "", Undefined)
	if err != nil{
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	Dictionary.AddTask(&task)

	fmt.Printf("Create task (Id = %d)\n", task.Id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.FormValue("data_key")
	Name := r.FormValue("data_name")
	DateStart := r.FormValue("data_dateStart")
	DateEnd := r.FormValue("data_dateEnd")
	Information := r.FormValue("data_information")
	Stat := r.FormValue("data_status")

	key, _ := strconv.ParseInt(Id, 0, 64)

	layOut := "2006-01-02T15:04"
	start, _ := time.Parse(layOut, DateStart)
	end, _ := time.Parse(layOut, DateEnd)

	result, err := Dictionary.FindTask(int(key))
	if err != nil{
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = result.Update(Name, start, end, Information, Status(Stat))
	if err != nil{
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	result.print()

	fmt.Printf("Update task (Id = %s)\n", Id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.FormValue("listTask")

	key, err := strconv.ParseInt(Id, 0, 64)

	err = Dictionary.DeleteTask(int(key))
	if err != nil{
		fmt.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Printf("Delete task (Id = %s)\n", Id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/edit", editHandler)
	http.ListenAndServe(":8000", nil)
}