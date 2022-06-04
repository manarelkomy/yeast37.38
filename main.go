package main

import (
	//"dataccess/stc/controllers/instructorcontroller"
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	header string
	p      string
}

var tp *template.Template

func main() {
	//http.HandleFunc("/", instructorcontroller.Index)
	//http.HandleFunc("/instructor", instructorcontroller.Index)
	//http.HandleFunc("/instructor/index", instructorcontroller.Index)
	//http.HandleFunc("/instructor/add", instructorcontroller.Add)
	//http.HandleFunc("/instructor/processAdd", instructorcontroller.ProcessAdd)
	//http.HandleFunc("/instructor/delet", instructorcontroller.Delete)
	//http.HandleFunc("/instructor/edit", instructorcontroller.Edit)
	//http.HandleFunc("/instructor/update", instructorcontroller.Update)
	//http.ListenAndServe(":3000", nil)
	http.HandleFunc("/", login)
	http.HandleFunc("/main", login)
	http.HandleFunc("/main/page", login)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":3000", nil)

}
func login(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("views/main/page.html")
	// if err!=nil{
	// 	fmt.Println("error",err)
	// 	return

	// }
	err2 := tp.Execute(w, "views/main/page.html")
	if err2 != nil {
		fmt.Println("error2", err2)
		return
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "post" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	h := r.FormValue("header")
	p := r.FormValue("p")
	d := struct {
		header string
		parag  string
	}{
		header: h,
		parag:  p,
	}

	tp.ExecuteTemplate(w, "p2.html", d)

}
