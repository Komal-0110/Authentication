package models

import "time"

type RoleType string

const (
	UserRoleType   RoleType = "user"
	AdminRoleType  RoleType = "admin"
	VendorRoleType RoleType = "vendor"
)

type Address struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PinCode      int    `json:"pincode"`
}

type AccountStatus string

const (
	ActiveAccountStatus   AccountStatus = "active"
	DeActiveAccountStatus AccountStatus = "deactive"
)

type User struct {
	Id                int           `json:"id"`
	ExternalId        string        `json:"ext_id"`
	Username          string        `json:"userName"`
	Email             string        `json:"email"`
	Password          string        `json:"password"`
	FullName          string        `json:"fullName,omitempty"`
	PhoneNumber       string        `json:"phoneNumber,omitempty"`
	Address           Address       `json:"address,omitempty"`
	Role              RoleType      `json:"role"`
	AccountStatus     AccountStatus `json:"accountStatus"`
	ProfilePictureURL *string       `json:"profilePictureURL,omitempty"`
	WishList          []any         `json:"wishlist,omitempty"`
	Cart              []any         `json:"cart,omitempty"`
	OrderHistory      []any         `json:"orderHistort,omitempty"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         *time.Time    `json:"updatedAt,omitempty"`
}
