package main

import (
	"fmt"

	"github.com/anika308142/mongoapi/routers"
)

func main() {

	fmt.Println("MongoDB Api")
	routers.MyRouter()
	//gin.Default().Run("localhost:9090")
}
