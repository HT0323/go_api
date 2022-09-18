package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/HT0323/go_api/api/middlewares"
)

// エラー発生した場合のレスポンス処理を行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError

	// errをMyAppError型のappErrに変換
	if !errors.As(err, &appErr) {
		// 変換できない場合はunKnownエラーを作成
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData:
		statusCode = http.StatusBadRequest
	case BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
