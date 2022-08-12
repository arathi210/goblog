package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Employee struct {
	Id       int
	Name     string
	City     string
	Province string
	Image    string
}
type Province struct {
	Id       int
	Province string
}
type Province_City struct {
	Id       int
	Province string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

//var Company = template.Must(template.ParseGlob("Company/*"))

func Home(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id1, province int
		var name, city, image string
		err = selDB.Scan(&id1, &name, &city, &province, &image)
		if err != nil {
			panic(err.Error())
		}

		//}

		nId1 := province
		selDB, err := db.Query("SELECT id, province_city FROM provinces WHERE id=?", nId1)
		if err != nil {
			panic(err.Error())
		}
		empProvince := Province{}
		for selDB.Next() {
			var id int
			var province_city string
			err = selDB.Scan(&id, &province_city)
			if err != nil {
				panic(err.Error())
			}

			emp.Id = id1
			emp.Name = name
			emp.City = city
			emp.Province = province_city
			emp.Image = image
			res = append(res, emp)

			empProvince.Province = province_city

			//res = append(res, empProvince)

		}
	}

	tmpl.ExecuteTemplate(w, "Home", res)
	defer db.Close()

}

func Blog(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id1, province int
		var name, city, image string
		err = selDB.Scan(&id1, &name, &city, &province, &image)
		if err != nil {
			panic(err.Error())
		}

		//}

		nId1 := province
		selDB, err := db.Query("SELECT id, province_city FROM provinces WHERE id=?", nId1)
		if err != nil {
			panic(err.Error())
		}
		empProvince := Province{}
		for selDB.Next() {
			var id int
			var province_city string
			err = selDB.Scan(&id, &province_city)
			if err != nil {
				panic(err.Error())
			}

			emp.Id = id1
			emp.Name = name
			emp.City = city
			emp.Province = province_city
			emp.Image = image
			res = append(res, emp)

			empProvince.Province = province_city

			//res = append(res, empProvince)

		}
	}

	tmpl.ExecuteTemplate(w, "Blog", res)
	defer db.Close()

}

func BlogSingle(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,description,image FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city, image string
		err = selDB.Scan(&id, &name, &city, &image)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		emp.Image = image

	}
	tmpl.ExecuteTemplate(w, "BlogSingle", emp)
	defer db.Close()

}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id1, province int
		var name, city, image string
		err = selDB.Scan(&id1, &name, &city, &province, &image)
		if err != nil {
			panic(err.Error())
		}

		//}

		nId1 := province
		selDB, err := db.Query("SELECT id, province_city FROM provinces WHERE id=?", nId1)
		if err != nil {
			panic(err.Error())
		}
		empProvince := Province{}
		for selDB.Next() {
			var id int
			var province_city string
			err = selDB.Scan(&id, &province_city)
			if err != nil {
				panic(err.Error())
			}

			emp.Id = id1
			emp.Name = name
			emp.City = city
			emp.Province = province_city
			emp.Image = image
			res = append(res, emp)

			empProvince.Province = province_city

			//res = append(res, empProvince)

		}
	}

	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,description FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city

	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT id , province_city FROM provinces WHERE parent_id=0")
	if err != nil {
		panic(err.Error())
	}
	emp := Province{}
	res := []Province{}
	for selDB.Next() {
		var id int
		var province_city string
		err = selDB.Scan(&id, &province_city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Province = province_city

		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "New", res)
	defer db.Close()
	//tmpl.ExecuteTemplate(w, "New", res)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,description,province,image FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, province int
		var name, city, image string
		err = selDB.Scan(&id, &name, &city, &province, &image)
		if err != nil {
			panic(err.Error())
		}

		nId1 := province
		selDB, err := db.Query("SELECT id, province_city FROM provinces WHERE id=?", nId1)
		if err != nil {
			panic(err.Error())
		}
		//empProvince := Province{}
		for selDB.Next() {
			var id1 int
			var province_city string
			err = selDB.Scan(&id1, &province_city)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.City = city
			emp.Image = image
			emp.Province = province_city
		}
	}

	// empProvince := Province{}
	// resProvince := []Province{}
	// for selDB.Next() {
	// 	var id int
	// 	var province_city string
	// 	err = selDB.Scan(&id, &province_city)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	empProvince.Id = id
	// 	empProvince.Province = province_city

	// 	emp = append(resProvince, empProvince)
	// }

	//empProvince.Province = province_city
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		r.ParseMultipartForm(10 << 20)

		file, handler, err1 := r.FormFile("myFile")
		if err1 != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err1)
			return
		}
		defer file.Close()
		fmt.Printf("File Size: %+v\n", handler.Size)

		tempFile, err2 := ioutil.TempFile("public/images", "upload-*.png")
		if err2 != nil {
			fmt.Println(err2)
		}
		//fmt.Printf("File Name: %+v\n",tempFile.Name())
		defer tempFile.Close()
		fmt.Printf("File Name: %+v\n", tempFile.Name())
		fileBytes, err3 := ioutil.ReadAll(file)
		if err3 != nil {
			fmt.Println(err3)
		}
		tempFile.Write(fileBytes)

		name := r.FormValue("name")
		city := r.FormValue("city")

		province := r.FormValue("province")
		//image := "test1"
		image := tempFile.Name()
		insForm, err := db.Prepare("INSERT INTO Employee(name, description,province,image) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, province, image)
		log.Println("INSERT: Name: " + name + " | City: " + city + " | Province: " + province)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		r.ParseMultipartForm(10 << 20)

		file, handler, err1 := r.FormFile("myFile")
		if err1 != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err1)
			return
		}
		defer file.Close()
		fmt.Printf("File Size: %+v\n", handler.Size)

		tempFile, err2 := ioutil.TempFile("public/images", "upload-*.png")
		if err2 != nil {
			fmt.Println(err2)
		}
		//fmt.Printf("File Name: %+v\n",tempFile.Name())
		defer tempFile.Close()
		fmt.Printf("File Name: %+v\n", tempFile.Name())
		fileBytes, err3 := ioutil.ReadAll(file)
		if err3 != nil {
			fmt.Println(err3)
		}
		tempFile.Write(fileBytes)
		image := tempFile.Name()

		name := r.FormValue("name")
		city := r.FormValue("city")

		id := r.FormValue("uid")

		insForm, err := db.Prepare("UPDATE Employee SET name=?, description=?, image=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, image, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city + "| Image: " + image)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func City(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	cat := r.URL.Query().Get("cat")
	selDB, err := db.Query("SELECT id , province_city FROM provinces WHERE parent_id=?", cat)
	if err != nil {
		panic(err.Error())
	}
	emp := Province_City{}
	res := []Province_City{}
	for selDB.Next() {
		var id int
		var province_city string
		err = selDB.Scan(&id, &province_city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Province = province_city

		res = append(res, emp)
	}
	//return json_encode(res)
	defer db.Close()
	//tmpl.ExecuteTemplate(w, "New", res)
}

func main() {
	//fs := http.FileServer(http.Dir("/css/"))
	//http.Handle("/css/", fs)

	log.Println("Server started on: http://localhost:8380")
	//	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("http/css"))))

	http.Handle("/css/",
		http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.Handle("/Company/",
		http.StripPrefix("/Company/", http.FileServer(http.Dir("Company"))))

	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/index", Index)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/blog", Blog)
	http.HandleFunc("/blogSingle", BlogSingle)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/city", City)
	http.ListenAndServe(":8380", nil)
}
