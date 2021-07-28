package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

// Serves for all incoming requests and matches incoming
// url with regex of url handled by different handlerFuncs
func Serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var h http.Handler
	var bookId int64
	var userId int64
	p := r.URL.Path
	switch {
	case match(p, "/library/users") && r.Method == http.MethodGet:
		h = get(getUsers)
	case match(p, "/library/users") && r.Method == http.MethodPost:
		h = post(addUser)
	case match(p, "/library/users/([0-9]+)", &userId) && r.Method == http.MethodGet:
		h = get(UserId(userId).getUser)
	case match(p, "/library/users/([0-9]+)", &userId) && r.Method == http.MethodPut:
		h = put(UserId(userId).updateUser)
	case match(p, "/library/users/([0-9]+)", &userId) && r.Method == http.MethodDelete:
		h = delete(UserId(userId).deleteUser)
	case match(p, "/library/books") && r.Method == http.MethodGet:
		h = get(getBooks)
	case match(p, "/library/books") && r.Method == http.MethodPost:
		h = post(addBook)
	case match(p, "/library/books/([0-9]+)", &bookId) && r.Method == http.MethodGet:
		h = get(BookId(bookId).getBook)
	case match(p, "/library/books/([0-9]+)", &bookId) && r.Method == http.MethodPut:
		h = put(BookId(bookId).updateBook)
	case match(p, "/library/books/([0-9]+)", &bookId) && r.Method == http.MethodDelete:
		h = delete(BookId(bookId).deleteBook)

	case match(p, "/library/users/([0-9]+)/books/([0-9]+)", &userId, &bookId) && r.Method == http.MethodPost:
		h = post(userBookAssociation{userId, bookId}.IssueBookToUser)
	case match(p, "/library/users/([0-9]+)/books/([0-9]+)", &userId, &bookId) && r.Method == http.MethodDelete:
		h = delete(userBookAssociation{userId, bookId}.ReleaseBookFromUser)
	default:
		http.NotFound(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

// match reports whether path matches ^regex$, and if it matches,
// assigns any capture groups to the *int64 vars.
func match(path, pattern string, vars ...interface{}) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
		case *int64:
			n, err := strconv.ParseInt(match, 10, 64)
			if err != nil {
				return false
			}
			*p = n
		default:
			panic("vars must be *int64")
		}
	}
	return true
}

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}

// allowMethod takes a HandlerFunc and wraps it in a handler that only
// responds if the request method is the given method, otherwise it
// responds with HTTP 405 Method Not Allowed.
func allowMethod(h http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.Header().Set("Allow", method)
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

// get takes a HandlerFunc and wraps it to only allow the GET method
func get(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, http.MethodGet)
}

// post takes a HandlerFunc and wraps it to only allow the POST method
func post(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, http.MethodPost)
}

// put takes a HandlerFunc and wraps it to only allow the PUT method
func put(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, http.MethodPut)
}

// delete takes a HandlerFunc and wraps it to only allow the DELETE method
func delete(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, http.MethodDelete)
}
