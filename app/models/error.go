package models

type Error struct {
	HttpCode  int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
