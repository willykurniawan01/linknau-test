package tests

import (
	"testing"

	"github.com/willykurniawan01/linknau-test/app/services"
)

func TestPayOrderSuccess(t *testing.T) {
	t.Logf("== Testing: Write a unit test for a complex Go function that involves multiple dependencies and asynchronous operations. ==")
	payment := services.Payment{}
	response := payment.PayOrder("123456789")

	if response["message_action"] != "PAYMENT_SUCCESS" {
		t.Errorf("Expected PAYMENT_SUCCESS, got %s", response["message_action"])
	}
}
