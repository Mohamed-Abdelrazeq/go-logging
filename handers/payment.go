package handlers

import "net/http"

func HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
		return
	}
	w.Write([]byte("Create Payment Intent"))
}
