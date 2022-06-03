package main

import (
	"dataccess/stc/controllers/instructorcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", instructorcontroller.Index)
	http.HandleFunc("/instructor", instructorcontroller.Index)
	http.HandleFunc("/instructor/index", instructorcontroller.Index)
	http.HandleFunc("/instructor/add", instructorcontroller.Add)
	http.HandleFunc("/instructor/processAdd", instructorcontroller.ProcessAdd)
	http.HandleFunc("/instructor/delet", instructorcontroller.Delete)
	http.HandleFunc("/instructor/edit", instructorcontroller.Edit)
	http.HandleFunc("/instructor/update", instructorcontroller.Update)
	http.ListenAndServe(":3000", nil)
}
