package rest_user_v1

import (
	"actify_srv/internal/db"
	rest "actify_srv/internal/rest/result"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
    ID            int        `json:"id"`
    Username      string     `json:"username"`
    Email         string     `json:"email"`
    PhoneNumber   string     `json:"phone_number"`
    PhoneVerified bool       `json:"phone_verified"`
    Password      string     `json:"password"`
    Birthdate     string     `json:"birthdate"`
    Gender        int     	 `json:"gender"`
    ProfileURL    string     `json:"profile_url"`
    IsActive      bool       `json:"is_active"`
    IsAdmin       bool       `json:"is_admin"`
    LastLogin     *time.Time `json:"last_login"`
    CreatedAt     time.Time  `json:"created_at"`
    UpdatedAt     time.Time  `json:"updated_at"`
}


type SignUpRequest struct {
	Name			string 	`json:"name"`
	Email			string	`json:"email"`
	Phone_number 	string	`json:"phone_number"`
	Phone_verified  bool	`json:"phone_verified"`
	Password		string	`json:"password"`
	Birthdate		string	`json:"birthdate"`
	Gender			int		`json:"gender"`
	Profile_url		string	`json:"profile_url"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	code := rest.Success

	if r.Method != http.MethodGet {
		code = http.StatusMethodNotAllowed
	}

	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	user_id := query.Get("id")

	resp := rest.RestResultJson{
		Code: code,
		Message: rest.MessageFromCode(code),
	}

	selectQuery := `
		SELECT id, username, email, phone_number, phone_verified, password, birthdate, gender, profile_url, is_active, is_admin, last_login, created_at, updated_at
		FROM users
	`

	if user_id != "" {
		selectQuery += " WHERE id = $1"
		var user User
		row := pDb.DB.QueryRow(selectQuery, user_id)
		err := row.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PhoneNumber,
			&user.PhoneVerified,
			&user.Password,
			&user.Birthdate,
			&user.Gender,
			&user.ProfileURL,
			&user.IsActive,
			&user.IsAdmin,
			&user.LastLogin,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			fmt.Printf("Search error..(error: %v)\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			resp = *rest.NewRestJsonResp(rest.DBGetError)
			if err == sql.ErrNoRows {
				resp.Message += "\nCan not find user id.\n"
			}
			resp.Message += "\n" + err.Error()
			return
		}
		resp = *rest.NewRestJsonResp(rest.Success)
		resp.Data = user
		
	} else {
		var users []User

		rows, err := pDb.DB.Query(selectQuery)
		if err != nil {
			fmt.Printf("Search error..(error: %v)\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			resp = *rest.NewRestJsonResp(rest.DBGetError)
			resp.Message += "\n" + err.Error()
			return
		}

		for rows.Next() {
			var user User
			err := rows.Scan(
				&user.ID,
				&user.Username,
				&user.Email,
				&user.PhoneNumber,
				&user.PhoneVerified,
				&user.Password,
				&user.Birthdate,
				&user.Gender,
				&user.ProfileURL,
				&user.IsActive,
				&user.IsAdmin,
				&user.LastLogin,
				&user.CreatedAt,
				&user.UpdatedAt,
			)

			if err != nil {
				fmt.Printf("Search error..(error: %v)\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				resp = *rest.NewRestJsonResp(rest.DBGetError)
				resp.Message += "\n" + err.Error()
				return
			}

			users = append(users, user)
			resp = *rest.NewRestJsonResp(rest.Success)
			resp.Data = users
		}
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func SignUpHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	query := r.URL.Query()
	signUpType := query.Get("type")

	if signUpType == "" || signUpType == "actify" {
		SignUpNormalHandler(w, r, pDb)
	} else {
		http.Error(w, "Not support other sign up process", http.StatusNotImplemented)
	}
}


func SignUpNormalHandler(w http.ResponseWriter, r *http.Request, pDb *db.PostgresDB) {
	var req SignUpRequest

	w.Header().Set("Content-Type", "application/json")
	resp := rest.RestResultJson{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp = *rest.NewRestJsonResp(rest.InvalidJson)
		json.NewEncoder(w).Encode(resp)
		return
	}
	
	birthdate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp = *rest.NewRestJsonResp(rest.InvalidDateFormat)
		json.NewEncoder(w).Encode(resp)
		return
	}

	binPassword := []byte(req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(binPassword, bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp = *rest.NewRestJsonResp(rest.WrongPasswordEncrypt)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, binPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp = *rest.NewRestJsonResp(rest.WrongPasswordDecrypt)
		json.NewEncoder(w).Encode(resp)
		return
	}

	query := `
		INSERT INTO users
			(username, email, phone_number, phone_verified, password,
			 birthdate, gender, profile_url, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id
	`

	var newUserID int64
	err = pDb.DB.QueryRow(
		query,
		req.Name,
		req.Email,
		req.Phone_number,
		req.Phone_verified,
		string(hashedPassword),
		birthdate,
		req.Gender,
		req.Profile_url,
	).Scan(&newUserID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp = *rest.NewRestJsonResp(rest.DBInsertError)
		resp.Message = fmt.Sprintf("%v\nError: %v", resp.Message, err)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	data := map[string] interface{} {
		"user_id" : newUserID,
	}

	resp.Data = data
	json.NewEncoder(w).Encode(resp)
}