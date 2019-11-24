package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Prabandham/larks_go/db"
	"github.com/Prabandham/larks_go/interfaces"
)

func loadRoutes(api *gin.Engine) {
	file, err := os.Open("Routes")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	int := interfaces.Interface{}
	intValue := reflect.ValueOf(int)
	for scanner.Scan() {
		route := scanner.Text()
		if !strings.HasPrefix(route, "#") && len(route) > 0 {
			strippedRoute := strings.Replace(route, "  ", " ", 0)
			array := strings.Split(strippedRoute, " ")
			if len(array) > 0 {
				var pathString [3]string
				index := 0
				for _, a := range array {
					if len(a) > 0 {
						pathString[index] = a
						index += 1
					}
				}
				handler := intValue.MethodByName(pathString[2]).Interface().(func(*gin.Context))
				api.Handle(pathString[1], pathString[0], handler)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := db.GetConnection()
	db.SetLogger()
	db.MigrateModels()

	server := gin.Default()

	loadRoutes(server)
	// description := docs.FuncDescription(ListProjects)

	// fmt.Println(description)

	server.Run(":3001")
}
