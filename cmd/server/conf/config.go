package conf

type Defaults struct {
	Mode       string
	Format     string
	DataLength int
}

// Default stores some default request settings
var Default = Defaults{"http", "json", 1000}
