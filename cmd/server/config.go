package main

type Defaults struct {
	Mode string
	Format string
	DataLength int
}

var Default = Defaults{"http", "json", 1000}
