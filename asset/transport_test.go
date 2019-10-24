package asset

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestDecodeAddIncomeRequest(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader(`{"statement_id": 1, "business_income": 2, "business_cost": 3, "gross_profit": 4}`))
	ctx := context.Background()
	data, err := decodeAddIncomeRequest(ctx, request)
	if err != nil {
		t.Error("add income request decode error:", err)
	}
	if data.(addIncomeRequest).StatementId != 1 {
		t.Errorf("add income request statement_id = %d, expect 1", data.(addIncomeRequest).StatementId)
	}
}
