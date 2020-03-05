// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import "fmt"

const (
	// status codes defined by VariFlight API
	// 0	用户不存在
	// 1	待审核
	// 2	当前访问ip尚未加入白名单
	// 3	缺少参数或参数验证失败
	// 4	暂无数据权限
	// 10	暂无数据
	// 11	未知错误
	UserNotExists          VariFlightStatusCode = 0
	CheckPending           VariFlightStatusCode = 1
	IPIsNotInWhiteList     VariFlightStatusCode = 2
	ParamValidationFailure VariFlightStatusCode = 3
	NoDataPermission       VariFlightStatusCode = 4
	NoData                 VariFlightStatusCode = 10
	UnknownError           VariFlightStatusCode = 11
)

// GetUrlError wraps error returned by http.Get() during calling VariFlightCaller
type GetUrlError struct {
	Method       APIMethod
	Url          string
	WrappedError error
}

func newGetUrlError(method APIMethod, url string, wrappedErr error) *GetUrlError {
	return &GetUrlError{method, url, wrappedErr}
}

func (r *GetUrlError) Error() string {
	return fmt.Sprintf("%v failed, url: %v, error: %v", r.Method, r.Url, r.WrappedError)
}

func (r *GetUrlError) Unwrap() error {
	return r.WrappedError
}

//  VariFlightDataQueryError represents error message responded by VariFlightCaller
type VariFlightDataQueryError struct {
	Method     APIMethod
	Url        string
	StatusCode VariFlightStatusCode
	Status     VariFlightStatus

	ErrCode int
	ErrMsg  string
}

type VariFlightStatusCode int

type VariFlightStatus string

func newDataQueryError(method APIMethod, url string, statusCode VariFlightStatusCode, status VariFlightStatus, errCode int, errMsg string) *VariFlightDataQueryError {
	return &VariFlightDataQueryError{method, url, statusCode, status, errCode, errMsg}
}

func (e *VariFlightDataQueryError) Error() string {
	return fmt.Sprintf("%v failed, url: %v, status code: %v, status: %v, error code: %v, error: %v.",
		e.Method, e.Url, e.StatusCode, e.Status, e.ErrCode, e.ErrMsg)
}

// DecodeJsonError wraps error of decoding message from VariFlightCaller
type DecodeJsonError struct {
	Method     APIMethod
	Url        string
	StatusCode int
	Status     string

	WrappedError error
}

func newDecodeJsonError(method APIMethod, url string, statusCode int, status string, wrappedError error) *DecodeJsonError {
	return &DecodeJsonError{method, url, statusCode, status, wrappedError}
}

func (e *DecodeJsonError) Error() string {
	return fmt.Sprintf("%v failed, url: %v, status code: %v, status: %v, error: %v.",
		e.Method, e.Url, e.StatusCode, e.Status, e.WrappedError)
}

func (e *DecodeJsonError) Unwrap() error {
	return e.WrappedError
}

// DBAccessError wraps error from accessing database.
type DBAccessError struct {
	Query        string
	WrappedError error
}

func newDBAccessError(query string, wrappedError error) *DBAccessError {
	return &DBAccessError{query, wrappedError}
}

func (e *DBAccessError) Error() string {
	return fmt.Sprintf("failed to access database, query: %v, error: %v", e.Query, e.WrappedError)
}

func (e DBAccessError) Unwrap() error {
	return e.WrappedError
}
