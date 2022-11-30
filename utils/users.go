package utils

type User struct {
	UserID       string `json:"userID" bigtable:"userID"`
	PasswordHash string `json:"passwordHash" bigtable:"passwordHash"`
	PasswordSalt string `json:"passwordSalt" bigtable:"passwordSalt"`
}

const tableUsers = "team2_User"
const userColumnFamilyName = "UserString"
const userIDColumnName = "userID"
const passwordHashColumnName = "passwordHash"
const passwordSaltColumnName = "passwordSalt"
