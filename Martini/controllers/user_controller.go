package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

// GetAllUser...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			users = append(users, user)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if len(users) < 5 {
		var response UsersResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	}
}

// UpdateUser
func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	userId := params["idUser"]

	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	sqlStatement := `
		UPDATE users 
		SET name = ?, age = ?, address =  ?
		WHERE id = ?`

	_, errQuery := db.Exec(sqlStatement,
		name,
		age,
		address,
		userId,
	)

	var response ErrorResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Update Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// InsertUser...
func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	// Read from Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	userType, _ := strconv.Atoi(r.Form.Get("type"))

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, usertype) values (?,?,?,?)",
		name,
		age,
		address,
		userType,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		log.Println(errQuery)
		response.Status = 400
		response.Message = "Insert Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	userId := params["idUser"]

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	var response ErrorResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
