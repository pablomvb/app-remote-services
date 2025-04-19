package bin

import (
	"encoding/json"
	"net/http"
	"os"
)

func ls(w http.ResponseWriter, r *http.Request) {
	// get path from query request
	strPath := r.URL.Query().Get("path")

	dirContent, err := os.ReadDir(strPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a list of file names
	var fileNames []string
	for _, entry := range dirContent {
		if entry.IsDir() {
			fileNames = append(fileNames, entry.Name())
		} else {
			fileNames = append(fileNames, entry.Name())
		}
	}

	// create a json response
	jsonBuffer := map[string]interface{}{
		"files": fileNames,
	}

	jsonResponse, err := json.Marshal(jsonBuffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set the header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
