package main

import (
	"fmt"
	"github.com/labi-le/tempest"
)

type News struct {
	Title string `template:"title"`
	Body  string `template:"body"`
}

func main() {
	news := News{
		Title: "{title} {product} disappeared",
		Body:  "{body} It is not yet known when they will appear again...",
	}

	replacements := []tempest.Replacement{
		{
			Tag:   "title",
			Value: "Something terrible happened today...",
		},

		{
			Tag:   "body",
			Value: "Today, {product} disappeared in stores...",
		},

		{
			Tag:   "product",
			Value: "Green onions",
		},
	}

	ShowExampleReplaceStruct(news, replacements)
	ShowExampleReplaceStructByTag(news, replacements)
	ShowExampleReplaceString(news.Title, replacements)
}

func ShowExampleReplaceStruct(n News, replacements []tempest.Replacement) {
	tempest.ReplaceStruct(&n, replacements)
	fmt.Printf("1: Title: %s\nBody: %s\n", n.Title, n.Body)
	fmt.Println()
	fmt.Println("In the first case, we have replaced all the tags in the structure")
	fmt.Println()
}

func ShowExampleReplaceStructByTag(n News, replacements []tempest.Replacement) {
	tempest.ReplaceStructByTag(&n, replacements)
	fmt.Printf("2: Title: %s\nBody: %s\n", n.Title, n.Body)
	fmt.Println()
	fmt.Println("In the second only those that were marked with the template tag")
	fmt.Println()
}

func ShowExampleReplaceString(title string, replacements []tempest.Replacement) {
	fmt.Println(tempest.ReplaceString(title, replacements))
	fmt.Println()
	fmt.Println("3: In the third case, we have replaced all the tags in the string")
}
