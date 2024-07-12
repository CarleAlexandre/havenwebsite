package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type card struct {
	Img			string
	Title		string
	Description	string
}

func main() {
	port := os.Getenv("WEBSITE_PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("www"))))

	go func() {
		http.ListenAndServe(":"+port, mux)
	}()

	func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n-----------------------\n\n")
		for {
			fmt.Print("Server Shell:> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			if strings.Compare("reload", text) == 0 {
				fmt.Println("reloading website data")
			} else if strings.Compare("exit", text) == 0 {
				fmt.Println("closing server")
				break
			} else {
				fmt.Println("Unknow Cmd")
			}
		}
	}()
}
