package session

import (
	"Go-Session-Clustering/src/user"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	key   []byte
	store *sessions.CookieStore
)

func init() {
	key = []byte(os.Getenv("SESSION_KEY"))
	store = sessions.NewCookieStore(key)
	store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: true,
	}
}

func Find(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "johnny-cookie")
	if err != nil {
		fmt.Println(err)
	}

	if !verification(session) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	userID := getUser(session)
	if userID == "" {
		http.Error(w, "There is no session", http.StatusUnauthorized)
		return
	}

	if userID == user.UserJohnny.ID {
		if _, err = fmt.Fprintln(w, user.UserJohnny.Name); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if _, err = fmt.Fprintln(w, userID); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func verification(session *sessions.Session) bool {
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		return true
	}
	return false
}

func getUser(session *sessions.Session) string {
	userID, ok := session.Values["user"].(string)
	if ok {
		return userID
	}
	return ""
}

func Login(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "mr-cookie")

	if req.Method != "POST" {
		http.Error(w, "Not Allow Method", http.StatusMethodNotAllowed)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Not Allow Content Type", http.StatusBadRequest)
		return
	}

	var u user.User
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		http.Error(w, "There is no user", http.StatusNotFound)
		return
	}

	session.Values["authenticated"] = true
	session.Values["user"] = u.ID
	if err := session.Save(req, w); err != nil {
		fmt.Println(err)
		return
	}
}

func Logout(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "mr-cookie")
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values["authenticated"] = false
	if err = session.Save(req, w); err != nil {
		fmt.Println(err)
		return
	}
}
