package order

import (
	"fmt"
	"net/http"

	"github.com/joseluis8906/pocone/pkg/log"
)

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var task CreateTask
	task.GenerateID()
	task.PersistOrder(r.Context())
	task.MarshalOrder()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s creating order: %s", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func removeItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func changeItemQty(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func pay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func resume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var task GetTask
	task.ExtractID(r)
	task.SearchTheDB(r.Context())
	task.Encode()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s showing order: %v", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}
