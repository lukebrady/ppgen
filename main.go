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
	group := "group { 'name':\n\tgid => 1,\n}\n"
	fmt.Fprintf(w, group)
}

func packageRoute(w http.ResponseWriter, r *http.Request) {
	pack := "package { 'name':\n\tensure => installed,\n}\n"
	fmt.Fprintf(w, pack)
}

func execRoute(w http.ResponseWriter, r *http.Request) {
	exec := "exec { 'name':\n\tcommand => '/bin/echo',\n\t# path => '/usr/bin:/usr/sbin:/bin:/usr/local/bin',\n\t# refreshonly => true,\n}\n"
	fmt.Fprintf(w, exec)
}

func cronRoute(w http.ResponseWriter, r *http.Request) {
	cron := "cron { 'name':\n\tcommand => '/path/to/executable',\n\t# user => 'root',\n\t# hour => 1,\n\t# minute  => 0,\n\t}\n"
	fmt.Fprintf(w, cron)
}

func serviceRoute(w http.ResponseWriter, r *http.Request) {
	service := "service { 'name':\n\tensure => running,\n\tenable => true,\n\thasrestart => true,\n\thasstatus  => true,\n\t# pattern => 'name',\n}\n"
	fmt.Fprintf(w, service)
}

func mountRoute(w http.ResponseWriter, r *http.Request) {
	mount := "mount { 'name':\n\tensure  => present,\n\tdevice  => 'device',\n\tfstype  => 'fstype',\n\toptions => 'opts';\n}\n"
	fmt.Fprintf(w, mount)
}

func yumRepoRoute(w http.ResponseWriter, r *http.Request) {
	yum := "yumrepo { 'name':\n\tbaseurl => '',\n\tdescr => 'The  repository',\n\tenabled => '1',\n\tgpgcheck => '1',\n\tgpgkey => 'file:///etc/pki/rpm-gpg/RPM-GPG-KEY-',\n\tmirrorlist => '',\n}\n"
	fmt.Fprintf(w, yum)
}

func ifRoute(w http.ResponseWriter, r *http.Request) {
	ifv := "if test {\n\t# enter puppet code\n}\n"
	fmt.Fprintf(w, ifv)
}

func elseifRoute(w http.ResponseWriter, r *http.Request) {
	elseifv := "elseif test {\n\t# enter puppet code\n}\n"
	fmt.Fprintf(w, elseifv)
}

func elseRoute(w http.ResponseWriter, r *http.Request) {
	elsev := "else {\n\t# enter puppet code\n}\n"
	fmt.Fprintf(w, elsev)
}

func caseRoute(w http.ResponseWriter, r *http.Request) {
	elsev := "case $variable {\n\t'value': {\n\t# code\n}\n\tdefault: {\n\t# code\n\t}\n}\n"
	fmt.Fprintf(w, elsev)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", statusRoute)
	mux.HandleFunc("/file", fileRoute)
	mux.HandleFunc("/user", userRoute)
	mux.HandleFunc("/group", groupRoute)
	mux.HandleFunc("/package", packageRoute)
	mux.HandleFunc("/exec", execRoute)
	mux.HandleFunc("/cron", cronRoute)
	mux.HandleFunc("/service", serviceRoute)
	mux.HandleFunc("/mount", mountRoute)
	mux.HandleFunc("/yumrepo", yumRepoRoute)
	mux.HandleFunc("/if", ifRoute)
	mux.HandleFunc("/elseif", elseifRoute)
	mux.HandleFunc("/else", elseRoute)

	http.ListenAndServe(":8080", mux)
}
