package restapi

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func createGin() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		// res http.ResponseWriter, req *http.Request
		resp := make(map[string]string)
		resp["status"] = "ok"
		resp["message"] = "11-https-server/health"
		resp["route"] = "/health"

		c.JSON(200, resp)
	})

	r.PUT("/api/v1/uplado", func(c *gin.Context) {
		// res http.ResponseWriter, req *http.Request
		var create struct {
			Pid string `json:"pid"`
			Tid string `json:"tid"`
		}
		json.NewDecoder(c.Request.Body).Decode(&create)

		fmt.Println("[/api/create] req.body:", create)

		// jsonResp := makeResponse(name, "/api/v1/create")
		resp := make(map[string]string)
		resp["status"] = "ok"
		resp["message"] = "create message is:" + create.Pid + "/" + create.Tid
		resp["route"] = "/api/v1/create"
		c.JSON(200, resp)
	})

	r.POST("/api/v1/create", func(c *gin.Context) {
		// res http.ResponseWriter, req *http.Request
		var create struct {
			Pid string `json:"pid"`
			Tid string `json:"tid"`
		}
		json.NewDecoder(c.Request.Body).Decode(&create)

		fmt.Println("[/api/create] req.body:", create)

		// jsonResp := makeResponse(name, "/api/v1/create")
		resp := make(map[string]string)
		resp["status"] = "ok"
		resp["message"] = "create message is:" + create.Pid + "/" + create.Tid
		resp["route"] = "/api/v1/create"
		c.JSON(200, resp)
	})

	// r.HandleFunc("/api/v1/install", func(res http.ResponseWriter, req *http.Request) {
	// 	var install InstallBodyStruct

	// 	json.NewDecoder(req.Body).Decode(&install)

	// 	fmt.Println("[/api/install] req.body:")
	// 	JsonPrettyPrint(install)

	// 	// jsonResp := makeResponse(name, "/api/v1/create")
	// 	resp := make(map[string]string)
	// 	resp["status"] = "ok"
	// 	resp["message"] = "install message is:" + install.Pid + "/" + install.Tid
	// 	resp["route"] = "/api/v1/install"
	// 	jsonResp, _ := json.Marshal(resp)

	// 	res.Write(jsonResp)
	// })

	// r.HandleFunc("/api/v1/assign", func(res http.ResponseWriter, req *http.Request) {
	// 	var assign AssignBodyStruct

	// 	json.NewDecoder(req.Body).Decode(&assign)

	// 	fmt.Println("[/api/assign] req.body:")
	// 	JsonPrettyPrint(assign)

	// 	// jsonResp := makeResponse(name, "/api/v1/create")
	// 	resp := make(map[string]string)
	// 	resp["status"] = "ok"
	// 	resp["message"] = "assign message is:" + assign.Pid
	// 	resp["route"] = "/api/v1/assign"
	// 	jsonResp, _ := json.Marshal(resp)

	// 	res.Write(jsonResp)
	// })

	return r
}
