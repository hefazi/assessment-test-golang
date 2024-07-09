package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userService = NewUserService()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", addUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}/add/{amount}", addUserBalance).Methods("PUT")
	router.HandleFunc("/user/{id}/subtract/{amount}", subtractUserBalance).Methods("PUT")
	router.HandleFunc("/user/{id}/random", randomUserBalance).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		// Log the error (optional)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := userService.GetUsers()
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	newUser := User{ID: len(userService.users) + 1, Balance: 0}
	err := userService.AddUser(newUser.ID, newUser.Balance)
	if err != nil {
		handleError(w, http.StatusMisdirectedRequest, "Failed to create user")
		return
	}
	handleSuccess(w)
	json.NewEncoder(w).Encode(newUser)

}

func addUserBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	amountStr, ok := params["amount"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Missing required parameters"})
		if err != nil {
			handleError(w, http.StatusInternalServerError, "Error!")
			return
		}
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID format"})
		return
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid amount format")
		return
	}

	user := userService.GetUser(userID)
	user.deposit(amount)

	handleSuccess(w)
	json.NewEncoder(w).Encode(user)
}

func subtractUserBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	amountStr, ok := params["amount"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Missing required parameters"})
		if err != nil {
			handleError(w, http.StatusInternalServerError, "Error!")
			return
		}
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID format"})
		return
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid amount format")
		return
	}

	user := userService.GetUser(userID)
	user.withdraw(amount)

	json.NewEncoder(w).Encode(user)
}

func randomUserBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Missing required parameters"})
		if err != nil {
			handleError(w, http.StatusInternalServerError, "Error!")
			return
		}
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID format"})
		return
	}

	user := userService.GetUser(userID)
	user.bet()

	json.NewEncoder(w).Encode(user)
}
