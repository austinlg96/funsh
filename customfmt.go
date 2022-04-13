package main

import (
	"fmt"
	"os"
)

var LOGFILE = "/logs/log"

type customfmt struct {
	Print   func(...any) (int, error)
	Printf  func(string, ...any) (int, error)
	Println func(...any) (int, error)
	Scanln  func(*string) (int, error)
	Sprint  func(...any) string
	Sprintf func(string, ...any) string
}

func Print(a ...any) (int, error) {
	f, _ := OpenLog()
	defer f.Close()
	f.WriteString(fmt.Sprintln(a...))
	i, e := fmt.Print(a...)
	return i, e
}

func Printf(s string, a ...any) (int, error) {
	f, _ := OpenLog()
	defer f.Close()
	f.WriteString(fmt.Sprintf(s, a...))
	i, e := fmt.Printf(s, a...)
	return i, e
}

func Println(a ...any) (int, error) {
	f, _ := OpenLog()
	defer f.Close()
	f.WriteString(fmt.Sprintln(a...))
	i, e := fmt.Println(a...)
	return i, e
}

func Scanln(p *string) (int, error) {
	i, e := fmt.Scanln(p)
	f, _ := OpenLog()
	defer f.Close()
	f.WriteString(fmt.Sprintln(*p))
	return i, e
}

func Sprint(a ...any) string {
	s := fmt.Sprint(a...)
	return s
}

func Sprintf(s string, a ...any) string {
	r := fmt.Sprintf(s, a...)
	return r
}

func OpenLog() (*os.File, error) {
	filename := LOGFILE
	_ = os.Chown(filename, 0, 0)
	f, e := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0222)
	f.WriteString(os.Getenv("SSH_CONNECTION") + ":     ")
	return f, e
}

var myfmt = customfmt{Print, Printf, Println, Scanln, Sprint, Sprintf}
