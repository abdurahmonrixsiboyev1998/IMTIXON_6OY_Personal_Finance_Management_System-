
package models

import "mime/multipart"

type Register struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type File struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}
