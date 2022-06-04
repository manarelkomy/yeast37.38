package models

import (
	"dataccess/stc/config"
	"dataccess/stc/enteties"
)

type InstructorModel struct {
}

func (*InstructorModel) FindAll() ([]enteties.Instructor, error) {
	db, err := config.GtDB()
	if err != nil {
		return nil, err

	} else {
		row, err2 := db.Query("select * from instructor")
		if err2 != nil {
			return nil, err2

		} else {
			var instructors []enteties.Instructor
			for row.Next() {
				var instructor enteties.Instructor
				row.Scan(&instructor.Id, &instructor.Inst_name, &instructor.Salary, &instructor.Inst_dept)
				instructors = append(instructors, instructor)

			}
			return instructors, nil
		}
	}
}
func (*InstructorModel) Find(id int64) (enteties.Instructor, error) {

	db, err := config.GtDB()
	if err != nil {
		return enteties.Instructor{}, err
	} else {
		rows, err2 := db.Query("select * from student where id=?", id)
		if err2 != nil {
			return enteties.Instructor{}, err
		} else {
			var instructor enteties.Instructor
			for rows.Next() {
				rows.Scan(&instructor.Id, &instructor.Inst_name, &instructor.Inst_dept, &instructor.Salary)
			}
			return instructor, nil
		}
	}

}

func (*InstructorModel) Create(instructor *enteties.Instructor) bool {
	db, err := config.GtDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into instructor(name,id,salary,department)vaues(?,?,?,?)",
			instructor.Inst_name, instructor.Id, instructor.Inst_dept, instructor.Salary)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}

}
func (*InstructorModel) Delete(id int64) bool {
	db, err := config.GtDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from instructor where id=?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}

	}
}

func (*InstructorModel) Update(instructor enteties.Instructor) bool {
	db, err := config.GtDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update instructor set inst_name=?, inst_dept=?,inst_dept=?, where id=?", instructor.Inst_name, instructor.Inst_dept, instructor.Salary, instructor.Id)
		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}

	}
}
