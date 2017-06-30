package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gernest/front"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/spf13/viper"
	"github.com/writeas/go-strip-markdown"
)

type Post struct {
	URI     string   `json:"location"`
	Title   string   `json:"title"`
	Content string   `json:"text"`
	Tags    []string `json:"tags"`
}

func ParsePost(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading file: ", path, err)
		return
	}

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	f, body, err := m.Parse(strings.NewReader(string(buf)))

	post := Post{}
	if title, ok := f["title"]; ok {
		post.Title = stripmd.Strip(title.(string))
	}
	text := stripmd.Strip(body)
	text = strip.StripTags(text)

	post.Content = text
	post.URI = strings.Replace(strings.TrimSuffix(strings.TrimPrefix(path, "content"), ".md"), "index", "", -1)

	fmt.Printf("Parsed %s\n", post.URI)
	// Add the baseurl in front of the post location
	post.URI = viper.GetString("baseurl") + post.URI
	mtx.Lock()
	posts = append(posts, post)
	mtx.Unlock()

	wg.Done()

}
