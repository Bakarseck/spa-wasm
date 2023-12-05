package login

import (
	"auth/internals/models"
	"auth/internals/models/user"
	"auth/internals/utils"
	"encoding/json"
	"io"
	"net/http"
)

var GlobalSessionManager = models.App{
	SessionHandler: models.CreateSessionManager(),
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	rsp, err := io.ReadAll(r.Body)
	if err != nil {
		data := map[string]interface{}{"error": err.Error()}
		utils.RespondWithJSON(w, data, http.StatusInternalServerError)
		return
	}

	var u user.UserLogin
	if err_json := json.Unmarshal(rsp, &u); err_json != nil {
		data := map[string]interface{}{"error": err_json.Error()}
		utils.RespondWithJSON(w, data, http.StatusInternalServerError)
		return
	}

	userFromDb, err := user.GetUser(u.Login)
	if err != nil {
		data := map[string]interface{}{"error": err.Error()}
		utils.RespondWithJSON(w, data, http.StatusInternalServerError)
		return
	}

	if userFromDb.Password != u.Password {
		data := map[string]interface{}{"error": "Invalid username or password"}
		utils.RespondWithJSON(w, data, http.StatusBadRequest)
		return
	}

	GlobalSessionManager.OpenSession(w, userFromDb)

	data := map[string]interface{}{
		"message": "User logged in successfully",
	}
	utils.RespondWithJSON(w, data, http.StatusOK)
}
