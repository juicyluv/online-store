package domain

import "time"

type User struct {
	Id         *int64     `json:"id,omitempty"`
	Telephone  *string    `json:"telephone,omitempty"`
	Password   *string    `json:"password,omitempty"`
	Email      *string    `json:"email,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

type GetUserRequest struct {
	UserId int64 `json:"userId"`
}
