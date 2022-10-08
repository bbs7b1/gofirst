package main

import (
	"encoding/json"
	"fmt"
	"gofirst/db"
	"gofirst/rules"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	if len(params["userid"]) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	userId, err := strconv.Atoi(params["userid"][0])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	user := db.FindUser(userId)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	campaigns := db.GetAllCampaigns()

	ids := rules.GetUserCampaigns(user, campaigns, runtime.NumCPU())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)

}

func startServer(port int) {
	fmt.Printf("Server is coming up on port %d\n", port)
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("startServer failed %s", err.Error())
		os.Exit(2)
	}
}
