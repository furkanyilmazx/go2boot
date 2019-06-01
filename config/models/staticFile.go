package models

type StaticFile struct {
	Root      string            `json:"root"`
	Locations map[string]string `json:"locations"`
}
