package news

import "fmt"

type MyResponse struct {
	News     []Data `json:"data"`
	Category string `json:"category"`
}

type Data struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Title   string `json:"title"`
}

func (data Data) PrintInfo() string {
	return fmt.Sprintf(
		"author : %s || title: %s\n Content: %s",
		data.Author,
		data.Title,
		data.Content)
}
