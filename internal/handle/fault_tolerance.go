package handle

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/config"
	"net/http"
)

// FaultToleranceHandle toggles should fault and returns fault status.
func FaultToleranceHandle(w http.ResponseWriter, _ *http.Request) {
	fault := config.ToggleFault()
	res, _ := json.Marshal(fault)
	_, _ = w.Write(res)
}
