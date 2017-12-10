// Package api :: params.go - http request parameters
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dockerian/dateparse"
	"github.com/gorilla/mux"
)

var (
	muxVars = mux.Vars
)

// BodyParams struct contains an JSON key/value map object
type BodyParams map[string]interface{}

// Params struct contains key/value pairs from URL path, request body, and query string
type Params struct {
	Form url.Values
	Body []byte
	Path string
	Post BodyParams
	Vars map[string]string
}

// NewParams returns pointer to an instance of Params struct with parsed http.Request
func NewParams(r *http.Request) *Params {
	r.ParseForm()

	var bodyParams = BodyParams{}
	var bytes []byte
	if r.Body != nil {
		bytes, _ = ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if len(bytes) > 0 {
			_ = json.Unmarshal(bytes, &bodyParams)
		}
	}

	pathVars := muxVars(r)

	// log.Printf("request.Body: %+v\n", r.Body)
	// log.Printf("request.Form: %+v\n", r.Form)
	// log.Printf("request.PostForm: %+v\n", r.Form)
	// log.Printf("mux.Vars: %+v\n", pathVars)
	return &Params{
		Body: bytes,
		Form: r.Form,
		Path: fmt.Sprintf("%s%s", r.Host, r.RequestURI),
		Post: bodyParams,
		Vars: pathVars,
	}
}

// GetBody method returns pointer to body param by key name
func (params *Params) GetBody(key string) interface{} {
	if key == "" || params.Body == nil {
		return params.Body
	}
	if data, okay := params.Post[key]; okay {
		return data
	}
	return params.Post
}

// GetDateRange method returns a date range by the key name
func (params *Params) GetDateRange(key string) ([]time.Time, error) {
	var dateValues []time.Time
	for _, str := range params.GetValues(key) {
		for _, strValue := range strings.Split(str, ",") {
			if strValue != "now" && len(strValue) < 4 {
				continue
			}
			// log.Println("[params] parsing date in range:", strValue)
			if dateValue, err := dateparse.ParseAny(strValue); err == nil {
				log.Printf("[params] adding '%s':'%s' as '%+v'\n", key, strValue, dateValue)
				dateValues = append(dateValues, dateValue)
			}
		}
	}
	if len(dateValues) > 0 {
		dateparse.Sort(dateValues)
		return dateValues, nil
	}
	return dateValues, errors.New("empty value")
}

// GetDateValues method returns sorted date values by the key name
func (params *Params) GetDateValues(key string) ([]time.Time, error) {
	var dateValues []time.Time
	strValues := params.GetValues(key)
	for _, str := range strValues {
		// log.Println("[params] parsing date value:", str)
		if date, err := dateparse.ParseAny(str); err == nil {
			log.Printf("[params] parsed '%s':'%s' to '%+v'\n", key, str, date)
			dateValues = append(dateValues, date)
		}
	}
	if len(dateValues) > 0 {
		dateparse.Sort(dateValues)
		return dateValues, nil
	}
	return dateValues, errors.New("empty value")
}

// GetInt method returns int value by the key name
// or the second parameter as default value
func (params *Params) GetInt(key string, defaultValues ...int) (int, error) {
	if strValue := params.GetValue(key); strValue != "" {
		return strconv.Atoi(strValue)
	}
	if len(defaultValues) > 0 {
		return defaultValues[0], nil
	}
	return 0, errors.New("empty value")
}

// GetNextPageURL returns next page URL per current page offset
func (params *Params) GetNextPageURL(pgOffsetKey string, pgOffset int) string {
	exp := fmt.Sprintf(`&?%s=[^&]+`, pgOffsetKey)
	rex := regexp.MustCompile(exp)
	origURL := rex.ReplaceAllString(params.Path, "")
	nextURL := fmt.Sprintf("%s&%s=%d", origURL, pgOffsetKey, pgOffset+1)
	return nextURL
}

// GetIntByRange method returns int value by the key name
// and within the range of rangeValues parameters
func (params *Params) GetIntByRange(key string, rangeValues ...int) int {
	var intVal int
	var minVal = math.MinInt32
	var maxVal = math.MaxInt32
	if strValue := params.GetValue(key); strValue != "" {
		intVal, _ = strconv.Atoi(strValue)
	}
	rangeLen := len(rangeValues)
	if rangeLen > 1 {
		sort.Ints(rangeValues)
		minVal, maxVal = rangeValues[0], rangeValues[rangeLen-1]
	} else if rangeLen > 0 {
		minVal = rangeValues[0]
	}
	if intVal < minVal {
		intVal = minVal
	}
	if intVal > maxVal {
		intVal = maxVal
	}
	return intVal
}

// GetValue method returns the value string by the key name
// or the second parameter as default value
func (params *Params) GetValue(key string, defaultValues ...string) string {
	if formValues, okay := params.Form[key]; okay {
		if len(formValues) > 0 {
			return formValues[0]
		}
	}
	if varValue, okay := params.Vars[key]; okay {
		return varValue
	}
	if len(defaultValues) > 0 {
		return defaultValues[0]
	}
	return ""
}

// GetValues method returns values by the key name
func (params *Params) GetValues(key string) []string {
	if formValues, okay := params.Form[key]; okay {
		if len(formValues) > 0 {
			return formValues
		}
	}
	if varValue, okay := params.Vars[key]; okay {
		return []string{varValue}
	}
	return []string{""}
}

// HasKey returns true if the params has the key; otherwise, return false
func (params *Params) HasKey(key string) bool {
	if _, hasFormKey := params.Form[key]; !hasFormKey {
		if _, hasPostKey := params.Post[key]; !hasPostKey {
			if _, hasVarsKey := params.Vars[key]; !hasVarsKey {
				return false
			}
		}
	}
	return true
}
