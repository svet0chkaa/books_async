package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)


type ExecutionTimeRequest struct {
	AccessKey int64  `json:"access_key"`
	ExecutionTime int `json:"execution_time"`
}

type Request struct {
	OrderId int64 `json:"order_id"`
}


func (h *Handler) issueExecutionTime(c *gin.Context) {
	var input Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueExecutionTime:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(4 * time.Second)
		sendExecutionTimeRequest(input)
	}()
}

func sendExecutionTimeRequest(request Request) {

	var executionTime = -1
	if rand.Intn(10) % 10 >= 2 {
	 executionTime = rand.Intn(15)
	}

	answer := ExecutionTimeRequest{
		AccessKey: 123,
		ExecutionTime: executionTime,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/orders/%d/update_execution_time/", request.OrderId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
