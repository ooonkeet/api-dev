package main

// import "honnef.co/go/tools/config"

type application struct {
	config config
	// logger
	// db driver
}

// run
// mount

type config struct{
	addr string
	db dbConfig
}

type dbConfig struct{
	dsn string
}