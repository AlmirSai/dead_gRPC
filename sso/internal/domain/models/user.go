// internal/domain/models/user.go
package models

type User struct {
	ID 		 string
	Email	 string
	PassHash []byte
}
