package main

import (
	"fmt"
	"net/http"
)

func statusRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Puppet Generator is running.\n")
	fmt.Fprintf(w, "The Resources available are:\nfile,\nuser,\ngroup,\npackage,\nmount")
}

func fileRoute(w http.ResponseWriter, r *http.Request) {
	file := "file { 'name':\n\tensure => file,\n\towner  => owner,\n\tgroup  => group,\n\tmode   => mode,\n\tsource => 'puppet:///modules/class/file.txt';\n}\n"
	fmt.Fprintf(w, file)
}

func userRoute(w http.ResponseWriter, r *http.Request) {
	user := "user { 'name':\n\tcomment => 'First Last',\n\thome => '/home/name',\n\tensure => present,\n\t#shell => '/bin/bash',\n\t#uid => '501',\n\t#gid => '20',\n}\n"
	fmt.Fprintf(w, user)
}

func groupRoute(w http.ResponseWriter, r *http.Request) {
	group := "group { 'name':\n\tgid => ,\n}\n"
	fmt.Fprintf(w, group)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", statusRoute)
	mux.HandleFunc("/file", fileRoute)
	mux.HandleFunc("/user", userRoute)
	mux.HandleFunc("/group", groupRoute)
	http.ListenAndServe(":8080", mux)
}
