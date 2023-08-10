package temperature

import "github.com/joho/godotenv"

func Outdoor() {
	godotenv.Load(".env")

	// fmt.Print("It's freaking cold outside!")
}
