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

			}
		})
	}
}

func getPreCheckFunc(username string) func() { return nil}

func getPostCheckFun(username string) func() {return nil}

func usernameMiddleware(next http.Handler) http.Handler {return nil}

func getUsernameFromContext(ctx context.Context) string{return nil}
