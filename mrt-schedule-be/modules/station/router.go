package station

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup){
	stationService := NewService()
	station := router.Group("/stations")
	station.GET("", func(c *gin.Context){
		GetAllStations(c, stationService)
	})
	station.GET("/:id", func(c *gin.Context){
		CheckSchedules(c, stationService)
	})


}

func GetAllStations(c *gin.Context, service Service){
	datas, err := service.GetAllStations()
	if err != nil{
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		},
	)
	return
	}
	c.JSON(
		http.StatusOK,
		APIResponse{
			Success: true,
			Message: "Successfully Get Stations",
			Data: datas,
		},
	)
}

func CheckSchedules(c *gin.Context, service Service){
	id := c.Param("id")
	datas, err := service.CheckSchedules(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		},
	)
	return
	}
	c.JSON(
		http.StatusOK,
		APIResponse{
			Success: true,
			Message: "Successfully",
			Data: datas,
		},
	)
}


type APIResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"Data"`
}