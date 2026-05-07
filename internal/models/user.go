package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"sirekap/internal/database"
)

// User represents a user in the system
type User struct {
	ID       int
	Nama     string
	NIP      string
	Email    string
	Username string
	Password string
	Level    string
}

// GetUserByUsernameOrEmail finds a user by username or email
func GetUserByUsernameOrEmail(identifier string) (*User, error) {
	query := `SELECT id, nama, nip, email, username, password, level 
			  FROM user 
			  WHERE username = ? OR email = ?`

	user := &User{}
	var nip sql.NullString

	err := database.DB.QueryRow(query, identifier, identifier).Scan(
		&user.ID,
		&user.Nama,
		&nip,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Level,
	)

	if err != nil {
		return nil, err
	}

	if nip.Valid {
		user.NIP = nip.String
	}

	return user, nil
}

// CheckPassword verifies the password against the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GetInitials returns the user initials for display
func (u *User) GetInitials() string {
	if len(u.Nama) == 0 {
		return "?"
	}
	
	// Simple initials: first letter of first and last word
	runes := []rune(u.Nama)
	initials := string(runes[0])
	
	// Find last space and get first letter after it
	for i := len(runes) - 1; i > 0; i-- {
		if runes[i] == ' ' && i+1 < len(runes) {
			initials += string(runes[i+1])
			break
		}
	}
	
	if len(initials) == 1 && len(runes) > 1 {
		initials += string(runes[1])
	}
	
	return initials
}
