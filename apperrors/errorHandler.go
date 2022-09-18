package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
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
