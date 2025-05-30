package utility

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GenerateAccNumber() string {
	const accPrefix = "ACN"

	currentTime := time.Now().Unix()

	timeStr := strconv.FormatInt(currentTime, 10)

	return accPrefix + timeStr
}

func GetRequestID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	return strconv.Atoi(idStr)
}
