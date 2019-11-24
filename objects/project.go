package objects

type Project struct {
	Base
	Name string `json:"name"`
	Title string `json:"title"`
}