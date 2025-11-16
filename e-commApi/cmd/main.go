package main

// import "golang.org/x/tools/go/cfg"

func main() {
	cfg:=config{
		addr: ":8080",
		db:dbConfig{},
	}
	api:=application{
		config: cfg,
	}
}