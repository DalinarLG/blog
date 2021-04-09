package routers

import (
	"io"
	"net/http"
	"os"

	"strconv"
	"strings"

	"github.com/DalinarLG/blog/db"
	"github.com/DalinarLG/blog/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}
	ext := strings.Split(handler.Filename, ".")[1]
	id := strconv.Itoa(UserID)
	path := "uploads/avatars/" + id + "." + ext

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	var user models.User
	user.Avatar = path

	status, err := db.UpdateUser(user, UserID)
	if !status {
		http.Error(w, "Could not update user", 400)
		return
	}

	if err != nil {
		http.Error(w, "Could not update user "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
