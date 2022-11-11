package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"

	"tawesoft.co.uk/go/dialog"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type Employee struct {
	Id       int
	Name     string
	City     string
	Province string
	Image    string
}
type User struct {
	Id       int
	Name     string
	Username string
	Password string
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
	dbUser := "go-test"
	dbPass := "8NxLtOLiag8d7bY"
	dbName := "go_test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

//var Company = template.Must(template.ParseGlob("Company/*"))

func Home(w http.ResponseWriter, r *http.Request) {

	res := 0

	tmpl.ExecuteTemplate(w, "Home", res)

}

func Blog(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
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
	selDB, err := db.Query("SELECT id,name,description,image FROM employee WHERE id=?", nId)
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
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
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
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,description FROM employee WHERE id=?", nId)
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
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
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
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,description,province,image FROM employee WHERE id=?", nId)
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
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
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
		insForm, err := db.Prepare("INSERT INTO employee(name, description,province,image) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, province, image)
		log.Println("INSERT: Name: " + name + " | City: " + city + " | Province: " + province)
	}

	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
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

		insForm, err := db.Prepare("UPDATE employee SET name=?, description=?, image=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, image, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city + "| Image: " + image)
	}
	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}
func LoginPage(w http.ResponseWriter, r *http.Request) {

	res := 0

	tmpl.ExecuteTemplate(w, "Login", res)

}
func Logincheck(res http.ResponseWriter, req *http.Request) {
	db := dbConn()
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {

		dialog.Alert("Username Incorrect")
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		dialog.Alert("Password incorrect")
		http.Redirect(res, req, "/login", 301)
		return
	}

	selDB, err := db.Query("SELECT name  FROM users  WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for selDB.Next() {

		var name string
		err = selDB.Scan(&name)
		if err != nil {
			panic(err.Error())
		}

		user.Name = name

	}
	session, _ := store.Get(req, "cookie-name")
	session1, _ := store.Get(req, "username")
	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session1.Values["username"] = user
	session.Save(req, res)
	session1.Save(req, res)
	defer db.Close()
	//res.Write([]byte("Hello" + databaseUsername))
	http.Redirect(res, req, "/dashboard", 301)

}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		http.Error(w, "Forbidden", http.StatusForbidden)
		http.Redirect(w, r, "/login", 301)
		return
	}
	//userID, _ := store.Get(r, "username")
	userID := "Admin"

	// store data (here in a map[string]interface{})
	data := make(map[string]interface{})
	data["userID"] = userID
	//res := 0

	tmpl.ExecuteTemplate(w, "Dashboard", data)

}
func Register(w http.ResponseWriter, r *http.Request) {

	res := 0

	tmpl.ExecuteTemplate(w, "Register", res)

}
func RegisterUser(res http.ResponseWriter, req *http.Request) {
	db := dbConn()
	if req.Method != "POST" {
		//http.ServeFile(res, req, "signup.html")
		tmpl.ExecuteTemplate(res, "Register", res)
		return
	}
	name := req.FormValue("name")
	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(name,username, password) VALUES(?, ?, ?)", name, username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}
		defer db.Close()
		//res.Write([]byte("User created!"))
		dialog.Alert("User created!")
		http.Redirect(res, req, "/register", 301)
		return

	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/register", 301)

	}

}
func Logout(res http.ResponseWriter, req *http.Request) {

	//res := 0
	session, _ := store.Get(req, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(req, res)

	http.Redirect(res, req, "/login", 301)

}

func nav(res http.ResponseWriter, req *http.Request) {
	// use a lib like gorilla sessions
	//session, _ := store.Get(req, "cookie-name")
	//userID, _ := store.Get(req, "cookie-name")

	// store data (here in a map[string]interface{})
	data := make(map[string]interface{})
	data["userID"] = "admin"

	// Send this data to the view
	//t.Execute(w, data) // Sets the . variable in templates
	tmpl.ExecuteTemplate(res, "Menu", data)
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
	http.HandleFunc("/", Home)
	http.HandleFunc("/blog", Blog)
	http.HandleFunc("/blogSingle", BlogSingle)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/login-check", Logincheck)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/register-user", RegisterUser)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/nav", nav)

	//http.ListenAndServe(":8380", nil)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8380" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}
