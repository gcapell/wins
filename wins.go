package main

import (
	"log"
	"os"
	"strings"
	
	"code.google.com/p/goplan9/plan9/acme"
)

func main() {
	w, err := acme.New()
	if err != nil {
		log.Fatal(err)
	}
	
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	chunks := strings.Split(dir, "/google3/")
	if len(chunks) != 2 {
		log.Fatal("expected exactly one google3 in", dir)
	}
	
	
	w.Name("%s/google3/+Errors", chunks[0]) 
	if err := w.Fprintf("tag", "\nacls ; blaze build --noobjfs %s:all\nacls ; blaze test --noobjfs '--test_output=streamed' '--test_arg=test.v' %s:all", chunks[1], chunks[1]); err != nil {
		log.Fatal(err)
	}
}
