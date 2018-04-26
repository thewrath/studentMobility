package dbManager

import(
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//add support for google drive and switcher bw local and online save or use something like R.sync 

var(
	db *sql.DB
	dbInit  bool
)

//define structure of a student 
type Student struct{
	Id int `json:"id"`
	FirstnameRU  string `json:"firstnameRU"` 
	FirstnameFR  string `json:"firstnameFR"`
	SecondNameRU string `json:"secondNameRU"`
	SecondNameFR string `json:"secondNameFR"`
	Email string `json:"email"`
	Department int `json:"department"`
	FromYear string `json:"fromYear"`
	ToYear string `json:"toYear"`

}

//define structure of a year
type Year struct{
	Id int `json:"id"`
	StudentId int `json:"studenId"`  
	NSUDepartment int `json:"NSUDepartment"`  
	NSUProgram int `json:"NSUProgram"`  
	NSUGrade int `json:"NSUGrade"`  
	Diploma int `json:"diploma"`  
	RecvOrgaRU string `json:"recvOrgaRU"`  
	RecvOrgaFR string `json:"recvOrgaFR"` 
	FromYear string `json:"fromYear"`
	ToYear string `json:"toYear"`
	CertificateTypeEdu int `json:"certificateTypeEdu"` 
	CertificateTypePro int `json:"certificateTypePro"` 
}

//used to initialize database access and components
func OpenDBAccess(path string) error{
	var err error 
	db, err = sql.Open("sqlite3",path)
	checkErr(err, "DB connection error")
	err = db.Ping() 
	checkErr(err, "DB ping error")
	if err == nil{
		dbInit = true 
	}
	fmt.Println(err)
	return err 
	
}

func CloseDBAccess() error{
	err := db.Close()
	return err
}

//Select student 
func SelectStudent(id int) error{
	if dbInit {
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		//add a selection of what is take by the select command 
		rq, err := db.Prepare("select * from students where id = ?");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		//exec the request you can add .scan to have what you want 
		row := rq.QueryRow(string(id))
		fmt.Println(row)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//select all student for the table
func SelectAllStudents() error{
	if dbInit {
		rows, err := db.Query("select * from students")
		checkErr(err)
		defer rows.Close()

		//5.1 Iterate through result set
		for rows.Next() {
			var name string
			var id int
			err := rows.Scan(&id, &name)
			checkErr(err)
			fmt.Printf("id=%d, name=%s\n", id, name)
		}

		//5.2 check error, if any, that were encountered during iteration
		err = rows.Err()
		checkErr(err)
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to remove student from DBs 
func RemoveStudent(student Student) error{
	if dbInit {
		fmt.Println("Start remove student")
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		rq, err := db.Prepare("delete from students where id = ?");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		_, err = rq.Exec(student.Id)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		fmt.Println("Student removed")
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to udpate student data in DBs  
func UpdateStudent(student Student) error{
	if dbInit {
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		rq, err := db.Prepare("update ");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		_, err = rq.Exec(student.Id)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to create a new student and insert it into DBs
func CreateStudent(student Student) error{
	fmt.Println(db)
	if dbInit {
		fmt.Println("Create a new student")
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		rq, err := db.Prepare("insert into students(firstnameRU, firstnameFR, secondNameRU, secondNameFR, email, department, fromYear, toYear) values(?,?,?,?,?,?,?,?)");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		_, err = rq.Exec(student.FirstnameRU, student.FirstnameFR, student.SecondNameRU, student.SecondNameFR, student.Email, student.Department, student.FromYear, student.ToYear)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//select one year data 
func SelectYear(id int) error{
	if dbInit {
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		//add a selection of what is take by the select command 
		rq, err := db.Prepare("select * from years where id = ?");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		//exec the request you can add .scan to have what you want 
		row := rq.QueryRow(string(id))
		fmt.Println(row)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//select all year of student 
func SelectAllYears(studentId int) error{
	if dbInit {
		rows, err := db.Query("select * from years")
		checkErr(err)
		defer rows.Close()

		//5.1 Iterate through result set
		for rows.Next() {

		}

		//5.2 check error, if any, that were encountered during iteration
		err = rows.Err()
		checkErr(err)
		return err
	}
	err := errors.New("Database not initialized.")
	return err
} 

//used to remove year from DBs 
func RemoveYear(year Year) error{
	if dbInit {
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		rq, err := db.Prepare("remove * from years where id = ?");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		_, err = rq.Exec(year.Id)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to created year 
func CreateYear(year Year) error{
	if dbInit {
		//return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to update year 
func UpdateYear(year Year) error{
	if dbInit {
		//return err
	}
	err := errors.New("Database not initialized.")
	return err
}

//used to find ID in DB 
func FindID(id int) (bool, error) {
	if dbInit {
		//tx means transaction 
		//begin transaction 
		tx, err := db.Begin()
		checkErr(err, "Begin transaction error")	
		//rq means sql request
		//prepare request 
		//add a selection of what is take by the select command 
		rq, err := db.Prepare("select id from students where id = ?");  
		//close request 
		defer rq.Close()
		//execute request 
		checkErr(err, "preparation error")
		//exec the request you can add .scan to have what you want 
		row := rq.QueryRow(string(id))
		fmt.Println(row)
		checkErr(err, "Request  execution error")
		//Commit transaction 
		tx.Commit()
		if(row != nil){
			return true, nil
		}
		return false, err
	}
	err := errors.New("Database not initialized.")
	return false, err
}

//used to check and handle error (replace print by logPrint)
func checkErr(err error, args ...string){
	if err != nil {
		fmt.Println("Error")
		fmt.Print(err)
		fmt.Println(args)

	}
}