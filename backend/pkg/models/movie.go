package models

type Movie struct {
 ID       string `json:"ID" omitempty`
 Title    string `json:"title"`
 Director string `json:"director"`
}
