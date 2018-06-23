package models

import "time"

//go:generate reform

//reform:users
type User struct {
	ID       int32  `reform:"id,pk"`
	Name     string `reform:"name"`
	Email    string `reform:"email"`
	Password string `reform:"password"`

	CreatedAt time.Time `reform:"created_at"`
	UpdatedAt time.Time `reform:"updated_at"`
}

// BeforeInsert set CreatedAt and UpdatedAt.
func (u *User) BeforeInsert() error {
	u.CreatedAt = time.Now().UTC().Truncate(time.Second)
	u.UpdatedAt = u.CreatedAt
	return nil
}

// BeforeUpdate set UpdatedAt.
func (u *User) BeforeUpdate() error {
	u.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return nil
}
