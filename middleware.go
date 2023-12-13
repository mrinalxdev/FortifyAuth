package main

import (
	"context"
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
	Roles []string
}

//Dummy Users . Will connect real db soon
var users = map[string]User{
	"admin" : User{Username: "Mrinal", Password: "adminpass", Roles: []string{"admin", "user"}},
	"user1" : User{Username: "Test1", Password: "user1", Roles: []string{"user"}},
	"user2" : User{Username: "Test2", Password: "user2", Roles: []string{"user"}},
	"user3" : User{Username: "Test3", Password: "user3", Roles: []string{"user"}},
}

func roleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			username := getUsernameFromContext(r.Context())
			
			user, found := users[username]

			if !found {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			roleMatched := false
			for _, allowedRole := range allowedRoles {
				for _, userRole := range user.Roles {
					if allowedRole == userRole {
						roleMatched = true
						break
					}
				}
				if roleMatched {
					break
				}
			}

			//Log Role Functionality

			if roleMatched {
				fmt.Printf("Role check passed for user %s with roles %v\n", username, user.Roles)

				if preCheckFunc := getPreCheckFunc(username); preCheckFunc != nil {
					preCheckFunc()
				}

				next.ServeHTTP(w, r)

				if postCheckFunc := getPostCheckFunc(username); postCheckFunc != nil {
					postCheckFunc()
				}
			} else {
				fmt.Printf("Role check failed for user %s with roles %v\n", username, user.Roles)

				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		})
	}
}

func getPreCheckFunc(username string) func() { 
	// defining custom post check logic based on the user
	if username == "admin" {
		return func() {
			fmt.Println("Execute admin post-check logic")
		} 
	}
	return nil
}

func getPostCheckFunc(username string) func() {
	if username == "admin" {
		return func() {
			fmt.Println("Execute admin post-check logic")
		} 
	}
	return nil
}

func usernameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "session-name")
		username, ok := session.Values["username"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), keyUsername, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUsernameFromContext(ctx context.Context) string {
	username, ok := ctx.Value(keyUsername).(string)
	if !ok {
		return ""
	}
	return username
}

//Key for username
type contextKey string
const keyUsername contextKey = "username"