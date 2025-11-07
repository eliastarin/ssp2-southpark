package app

import (
	"encoding/json"
	"net/http"

	"github.com/eliastarin/ssp2-southpark/go-api/domain"
	"github.com/eliastarin/ssp2-southpark/go-api/ports"
)

type Handlers struct {
	pub ports.MessagePublisher
}

func NewHandlers(pub ports.MessagePublisher) *Handlers {
	return &Handlers{pub: pub}
}

func (h *Handlers) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handlers) PostMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var m domain.Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := m.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.pub.Publish(m); err != nil {
		http.Error(w, "failed to publish", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
