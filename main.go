package main

import "github.com/AvinFajarF/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {

}

