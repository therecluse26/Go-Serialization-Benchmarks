package conf

type Defaults struct {
	Mode        string
	Format      string
	DataLength  int
	Compression bool
}

// Default stores some default request settings
var Default = Defaults{"http", "json", 1000, false}
