package v2

import (
	"encoding/json"
	"strings"
	"testing"
)

// Amounts above 131,072 can't be held to the cent in a float32, so a
// Transaction must carry Amount as float64 to send the exact value to Decta.
func TestTransactionMarshalPreservesCents(t *testing.T) {
	tests := []struct {
		cents int64
		want  string
	}{
		{cents: 13107201, want: `"amount":131072.01`},
		{cents: 20000013, want: `"amount":200000.13`},
		{cents: 123456789, want: `"amount":1234567.89`},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			tx := NewTransaction(float64(tt.cents)/100, "EUR", "balance adjustment", "TOP_UP")

			body, err := json.Marshal(tx)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if !strings.Contains(string(body), tt.want) {
				t.Errorf("body %s does not contain %s", body, tt.want)
			}
		})
	}
}
