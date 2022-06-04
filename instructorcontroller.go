package instructorcontroller

import (
	"dataccess/stc/enteties"
	"dataccess/stc/models"
	"html/template"
	"net/http"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var InstructorModel models.InstructorModel
	instructors, _ := InstructorModel.FindAll()
	data := map[string]interface{}{
		"instructor": instructors,
	}
	tmp, _ := template.ParseFiles("views/instructor/index.html")
	tmp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/instructor/add.html")
	tmp.Execute(response, nil)
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var instructor enteties.Instructor
	instructor.Inst_name = request.Form.Get("inst_name")
	instructor.Inst_dept = request.Form.Get("inst_dept")
	instructor.Salary, _ = strconv.ParseInt(request.Form.Get("salary"), 10, 64)
	instructor.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	var InstructorModel models.InstructorModel
	InstructorModel.Create(&instructor)
	http.Redirect(response, request, "/instructor", http.StatusSeeOther)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var inststructorModel models.InstructorModel
	inststructorModel.Delete(id)
	http.Redirect(response, request, "/instructor", http.StatusSeeOther)

}

func Edit(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var inststructorModel models.InstructorModel
	instuctor, _ := inststructorModel.Find(id)
	data := map[string]interface{}{
		"instuctor": instuctor,
	}
	tmp, _ := template.ParseFiles("views/instructor/edit.html")
	tmp.Execute(response, data)

}

func Update(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	var instructor enteties.Instructor

	instructor.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	instructor.Salary, _ = strconv.ParseInt(request.Form.Get("salary"), 10, 64)
	instructor.Inst_name = request.Form.Get("inst_name")
	instructor.Inst_dept = request.Form.Get("inst_dept")

	var inststructorModel models.InstructorModel
	inststructorModel.Update(instructor)
	http.Redirect(response, request, "/instructor", http.StatusSeeOther)

}
