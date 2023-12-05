package register

import (
	"auth/internals/models"
	"auth/internals/models/user"
	"auth/internals/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


// The HandleRegister function handles the registration of a user by reading the request body,
// unmarshaling it into a User struct, saving the user, and returning an appropriate response.
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	resp, err := io.ReadAll(r.Body)
	var user user.User
	var message models.Message

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(resp, &user); err != nil {
		message := models.Message{
			Error: "Bad Request" + err.Error(),
		}
		utils.RespondWithJSON(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	if err := user.Save(); err != nil {

		if shouldReturn := CheckExist(err, message, w); shouldReturn {
			return
		}

		message := models.Message{
			Error: fmt.Sprintf("Status Internal Server Error: %v", err.Error()),
		}
		utils.RespondWithJSON(w, message, http.StatusInternalServerError)
		fmt.Println(message)
		return
	}

	message.Message = fmt.Sprintf("User %s registered successfully", user.Username)

	utils.RespondWithJSON(w, message, http.StatusOK)

	fmt.Println(message)

}

// The function checks if a specific error message exists and returns a boolean value indicating
// whether the error exists or not.
func CheckExist(err error, message models.Message, w http.ResponseWriter) bool {
	if err.Error() == "UNIQUE constraint failed: users.email" {
		message.Error = "Email already in use"
		utils.RespondWithJSON(w, message, http.StatusBadRequest)
		return true
	} else if err.Error() == "UNIQUE constraint failed: users.username" {
		message.Error = "Username already in use"
		utils.RespondWithJSON(w, message, http.StatusBadRequest)
		return true
	}
	return false
}
