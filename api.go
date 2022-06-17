package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Users struct {
	Id int
	Username string
	Email string
	Password string
	Age int
	CreateAt time.Time
	UpdateAt time.Time
}

type UserResponse struct {
	Id    int    `json:"id"`
	Username  string `json:"name"`
	Email string `json:"email"`
}

var users = map[int]Users{
	1: {
		Id: 1, 
		Username: "Andi",
		Email: "andi@gmail.com", 
		Password: "1231231", 
		Age: 25,
	},
	2: {
		Id: 2, 
		Username: "Anda",
		Email: "andi@gmail.com", 
		Password: "1231231", 
		Age: 25,
	},
}

var PORT = ":8080"

func main()  {
	http.HandleFunc("/users/", getUsers)

	http.HandleFunc("/user/", createUser)

	http.HandleFunc("/user/edit/", editUser)

	http.HandleFunc("/user/delete/", deleteUser)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getUsers(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		fmt.Println(r.URL.Path[1:])
		paths := strings.Split(r.URL.Path[1:], "/")
		fmt.Printf("%+v", paths)
		fmt.Println(len(paths))

		if len(paths) == 2 && paths[1] != "" {
			if idInt, err := strconv.Atoi(paths[1]); err == nil {
			jsonData, _ := json.Marshal(users[idInt])
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			}
		} else {
			var userSlice []Users
			for _, v := range users {
				userSlice = append(userSlice, v)
			}

			fmt.Println("No Params")
			jsonData, _ := json.Marshal(userSlice)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
		fmt.Println(r.Method)
		return
	} 
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		var req Users
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		users[req.Id] = req

		resp := UserResponse{
			Id:    req.Id,
			Username:  req.Username,
			Email: req.Email,
		}
		
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	fmt.Println(r.Method)
	return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func editUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "PUT" {
		var req Users
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		users[req.Id] = req

		resp := UserResponse{
			Id:    req.Id,
			Username:  req.Username,
			Email: req.Email,
		}
		
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(r.Method)
		return
	} 
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func deleteUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "DELETE" {
		fmt.Println(r.URL.Path[1:])
		paths := strings.Split(r.URL.Path[1:], "delete/")
		fmt.Printf("%+v", paths)
		fmt.Println(len(paths))

		if len(paths) == 2 && paths[1] != "" {
			if idInt, err := strconv.Atoi(paths[1]); err == nil {
				delete(users, idInt)
			}
		}

		resp := "Data Deleted"
		
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(r.Method)
		return
	} 
	http.Error(w, "Invalid method", http.StatusBadRequest)
}
