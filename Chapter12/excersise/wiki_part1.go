package main

import (
	"fmt"
	"io/ioutil"
)

// Page 页
type Page struct {
	Title string
	Body  []byte
}

// Save 保存
func (p *Page) Save() (err error) {
	/* 	file, err := os.OpenFile(p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	   	if err != nil {
	   		fmt.Printf("An error occurred while openning or creating\n")
	   		return err
	   	}
	   	defer file.Close()
	   	outputWriter := bufio.NewWriter(file)
	   	outputWriter.Write(p.Body)
	   	outputWriter.Flush()
	   	return nil */
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

// Load 加载
func (p *Page) Load(title string) (err error) {
	/* 	buf, err := ioutil.ReadFile(title)
	   	if err != nil {
	   		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	   		return err
	   	}
	   	p.Body = buf
	   	return nil */
	p.Title = title
	p.Body, err = ioutil.ReadFile(p.Title)
	return err
}

func wikiPart1() {
	page := new(Page)
	page.Title = "Chapter.md"
	page.Body = []byte("hello world")
	page.Save()
	page.Load("Chapter.md")
	fmt.Println(string(page.Body))
}
