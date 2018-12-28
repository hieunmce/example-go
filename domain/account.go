package domain

type Account struct {
	Model
	UserName       string `json:"user_name"`
	DigestPassword string `json:"digest_password"`
	Type           string `json:"type"`
}
