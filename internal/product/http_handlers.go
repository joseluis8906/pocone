package product

import (
	"fmt"
	"net/http"

	"github.com/joseluis8906/pocone/pkg/log"
)

func all(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var task addProductTask
	task.Unmarshal(r.Body)
	task.PersistOnDB(r.Context())
	task.Marshal()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s adding new product: %v", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}

func get(w http.ResponseWriter, r *http.Request) {
	var task getProductTask
	task.ExtractParams(r.URL.Query())
	task.SearchOnDB(r.Context())
	task.Marshal()

	data, err := task.Result()
	if err != nil {
		log.Printf("%s getting products: %v", log.Error, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", data)
}
