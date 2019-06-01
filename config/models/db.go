package models

type DB struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Url		 string `json:"url"`
	Username string `json:"username"`
	Password string `json:"passwd"`
	Port     int    `json:"port"`
}
