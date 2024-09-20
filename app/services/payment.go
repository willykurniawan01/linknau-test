package services

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"time"

	"github.com/willykurniawan01/linknau-test/app/helpers"
)

type Result struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type Payment struct{}

// Simulates hitting an external wallet service
func (p *Payment) ProcessEmoney(results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	results <- Result{Type: "e-money", Status: "Success"}
}

// Simulates hitting an external point service
func (p *Payment) ProcessPoints(results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	results <- Result{Type: "points", Status: "Success"}
}

// PayOrder processes both payments concurrently and returns a structured response
func (p *Payment) PayOrder(orderId string) (response map[string]interface{}) {
	// Handling Panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	apiCall := helpers.GenerateAPIcallID()

	var wg sync.WaitGroup
	results := make(chan Result, 2)

	wg.Add(2)

	go p.ProcessEmoney(results, &wg)

	go p.ProcessPoints(results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	successCount := 0

	for result := range results {
		if result.Status == "Success" {
			successCount++
		}
	}

	log.Println(successCount)

	currentTime := time.Now()
	if successCount == 2 {
		response = map[string]interface{}{
			"message_id":     apiCall,
			"message_action": "PAYMENT_SUCCESS",
			"message_desc":   "Pembayaran sudah berhasil di proses.",
			"message_data": map[string]interface{}{
				"order_id": orderId,
			},
			"message_request_datetime": currentTime.Format("2006-01-02 15:04:05"),
		}
		return response
	}

	response = map[string]interface{}{
		"message_id":     apiCall,
		"message_action": "PAYMENT_FAILED",
		"message_desc":   "Maaf pembayaran gagal di proses.",
		"message_data": map[string]interface{}{
			"order_id": orderId,
		},
		"message_request_datetime": currentTime.Format("2006-01-02 15:04:05"),
	}

	return response
}
