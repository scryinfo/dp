// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

// url returns a url with corresponding _token that should be used to request the API provided by http://www.variflight.com/.
func url(urlWithoutQueryParams string, queryParameters map[string]string, appsecurity string) string {
	// join the query parameters with the appsecurity, and get a query string with corresponding _token.
	queryStrWithToken := queryWithToken(queryParameters, appsecurity)
	// join the query string with the front part of url to conform the final whole requested url.
	return urlWithoutQueryParams + queryStrWithToken
}

// _token returns a _token that generated from the given queryParamsMap and appsecurity.
func token(queryParamsMap map[string]string, appsecurity string) string {
	// convert query parameters map to required string, whose parameter names are made lowercase and parameters are
	// sorted in increasingly order and are concatenated with "&".
	paramsSortedStr := convertMapToStr(queryParamsMap)

	// join the appsecurity
	paramsSortedStrWithRegCode := paramsSortedStr + "&appsecurity=" + appsecurity

	// encrypt twice by MD5
	return md5md5(paramsSortedStrWithRegCode)
}

//  queryWithToken returns a query string that carries a corresponding _token.
func queryWithToken(queryParameters map[string]string, appsecurity string) string {
	// convert query parameters map to required string, whose parameter names are made lowercase and parameters are
	// sorted in increasingly order and are concatenated with "&".
	paramsSortedStr := convertMapToStr(queryParameters)

	// generate corresponding _token with the given paramsSortedStr and appsecurity
	token := _token(paramsSortedStr, appsecurity)

	// join the paramsSortedStr with _token
	return paramsSortedStr + "&_token=" + token
}

// _token returns a _token that generated from the given paramsSortedStr and appsecurity.
func _token(paramsSortedStr string, appsecurity string) string {
	// join the appsecurity
	paramsSortedStrWithRegCode := paramsSortedStr + "&appsecurity=" + appsecurity

	// encrypt twice
	return md5md5(paramsSortedStrWithRegCode)
}

// convertMapToStr returns a lowercase-keys and sorted-elements string with the given map.
func convertMapToStr(kvs map[string]string) string {
	// convert map to array with key to lower
	var strArr []string
	for key, value := range kvs {
		strArr = append(strArr, strings.ToLower(key)+"="+value)
	}
	// sort array of strings in increasing order
	sort.Strings(strArr)
	// concatenate the elements of array to create a single string
	return strings.Join(strArr, "&")
}

// md5md5 checksum of the data twice
func md5md5(data string) string {
	data = fmt.Sprint(md5.Sum([]byte(data)))
	return fmt.Sprint(md5.Sum([]byte(data)))
}
