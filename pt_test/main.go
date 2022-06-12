package main

type Request struct {
	RequestID string
	Member    *Member
}
type Member struct {
	Name string
	Age  int
}
