package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willykurniawan01/linknau-test/app/constant"
)

func GenerateAPIcallID() string {
	nowTime := time.Now().Unix()
	currentDate := time.Now()

	hourTm := currentDate.Hour()
	minuteTm := currentDate.Minute()
	secondTm := currentDate.Second()

	startID := strconv.FormatInt(nowTime, 10) + strconv.Itoa(hourTm) + strconv.Itoa(minuteTm) + strconv.Itoa(secondTm)
	randomNum := rand.Intn(10000000) + 1
	apiCallID := fmt.Sprintf("API_CALL_%s_%d", startID, randomNum)
	return apiCallID
}

func ResponseApi(c *gin.Context, statustype string, data map[string]interface{}) {
	var status constant.StatusCode
	for _, s := range constant.Response {
		if s.Code == statustype {
			status = s
			break
		}
	}
	currentTime := time.Now()
	resp := gin.H{
		"message_id":               GenerateAPIcallID(),
		"message_action":           status.Code,
		"message_desc":             status.Message,
		"message_data":             data,
		"message_request_datetime": currentTime.Format("2006-01-02 15:04:05"),
	}

	c.JSON(status.Status, resp)
}
