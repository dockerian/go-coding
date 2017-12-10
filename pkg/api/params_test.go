// +build all common pkg api params

// Package api :: params_test.go
package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/dockerian/dateparse"
	"github.com/stretchr/testify/assert"
)

// TestNewParams tests NewParams constructor
func TestNewParams(t *testing.T) {
	params := getNewParamsFromGetRequest(t, "/test/path", "key=value")
	assert.NotNil(t, params)

	num, err := params.GetInt("foo", 12345)
	assert.Equal(t, 12345, num)
	assert.Nil(t, err)

	assert.True(t, params.HasKey("key"))
	assert.False(t, params.HasKey("foo"))
	val := params.GetValue("foo", "bar")
	assert.Equal(t, "bar", val)
}

// TestParamsGetDateRange tests NewParams GetDateRange method
func TestParamsGetDateRange(t *testing.T) {
	var times []time.Time
	var slice = []string{
		"2017-11-11 11:01:01",
		"2009-11-22T11:22:02",
		"2009-11-23T03:33:03 UTC",
		"2009-11-29T04:44:04-04:00",
		"",
	}
	query := "q=name1&q=name2&dt=2017-11-30,xx&range=2017-08-10,now&time="
	for _, str := range slice {
		query += "," + str
		log.Printf("[test] parsing '%s'\n", str)
		if t, err := dateparse.ParseAny(str); err == nil {
			log.Printf("[test] +adding '%+v'\n", t)
			times = append(times, t)
		}
	}
	dateparse.Sort(times)

	dt, _ := dateparse.ParseAny("2017-11-30")
	datesRangeBegin := []time.Time{dt}

	param := getNewParamsFromGetRequest(t, "test/path", query)
	assert.NotNil(t, param)
	t.Logf("param: %+v\n", param)
	dateValues, err := param.GetDateRange("time")
	msg := fmt.Sprintf("expect dates: %+v\nresult dates: %+v\n", times, dateValues)
	assert.Equal(t, times, dateValues, msg)
	assert.Nil(t, err)

	dtValues, err := param.GetDateRange("dt")
	assert.Equal(t, datesRangeBegin, dtValues)

	rangeValues, err := param.GetDateRange("range")
	// t.Logf("parsed range: %v\n", rangeValues)
	assert.NotNil(t, rangeValues)
	assert.NotEqual(t, []time.Time{}, rangeValues)
	assert.Nil(t, err)

	nilValues, err := param.GetDateRange("date")
	assert.Equal(t, "empty value", err.Error())
	assert.Nil(t, nilValues)
}

// TestParamsGetDateValues tests NewParams GetDateValues method
func TestParamsGetDateValues(t *testing.T) {
	var times []time.Time
	var slice = []string{
		"2017-11-11 11:01:01",
		"2009-11-22T11:22:02",
		"2009-11-23T03:33:03 UTC",
		"2009-11-29T04:44:04-04:00",
		"",
	}
	query := "q=name1&q=name2&dt=2017-11-30&range=2017-08-10,now"
	for _, str := range slice {
		query += "&time=" + str
		log.Printf("[test] parsing '%s'\n", str)
		if t, err := dateparse.ParseAny(str); err == nil {
			log.Printf("[test] +adding '%+v'\n", t)
			times = append(times, t)
		}
	}
	dateparse.Sort(times)

	dt, _ := dateparse.ParseAny("2017-11-30")
	datesOneItem := []time.Time{dt}

	param := getNewParamsFromGetRequest(t, "test/path", query)
	assert.NotNil(t, param)
	t.Logf("param: %+v\n", param)
	dateValues, err := param.GetDateValues("time")
	msg := fmt.Sprintf("expect dates: %+v\nresult dates: %+v\n", times, dateValues)
	assert.Equal(t, times, dateValues, msg)
	assert.Nil(t, err)

	dtValues, err := param.GetDateValues("dt")
	assert.Equal(t, datesOneItem, dtValues)

	emptyValue, err := param.GetDateValues("range")
	assert.Nil(t, emptyValue)
	assert.NotNil(t, err)

	nilValues, err := param.GetDateValues("date")
	assert.Equal(t, "empty value", err.Error())
	assert.Nil(t, nilValues)
}

// TestParamsGetIntByRange tests NewParams GetIntByRange method
func TestParamsGetIntByRange(t *testing.T) {
	params := getNewParamsFromGetRequest(t, "test/path", "size=10&q=test")
	assert.NotNil(t, params)
	t.Logf("param: %+v\n", params)
	var integers = []int{234, 5, 67, 8, 9, 0, -101}
	v0 := params.GetIntByRange("number")
	assert.Equal(t, 0, v0)
	v1 := params.GetIntByRange("size")
	assert.Equal(t, 10, v1)
	v2 := params.GetIntByRange("size", 5, 100)
	assert.Equal(t, 10, v2)
	v3 := params.GetIntByRange("size", 3, 5)
	assert.Equal(t, 5, v3)
	v4 := params.GetIntByRange("size", 20, 50)
	assert.Equal(t, 20, v4)
	v5 := params.GetIntByRange("size", integers...)
	assert.Equal(t, 10, v5)
	v6 := params.GetIntByRange("size", 25)
	assert.Equal(t, 25, v6)
}

// TestParamsGetNextPageURL tests NewParams GetNextPageURL method
func TestParamsGetNextPageURL(t *testing.T) {
	rqURL := "/test/path"
	tests := []struct {
		query      string
		pageOffset int
		nextURL    string
	}{
		{
			"pgsz=20", 0, "pgsz=20&pg=1",
		},
		{
			"pgsz=10&pg=3", 3, "pgsz=10&pg=4",
		},
		{
			"pgsz=10&pg=-1", 0, "pgsz=10&pg=1",
		},
	}
	for idx, test := range tests {
		params := getNewParamsFromGetRequest(t, rqURL, test.query)
		params.Path = fmt.Sprintf("%s/%s", rqURL, test.query)
		assert.NotNil(t, params)
		msg := fmt.Sprintf("%s ==> %s [%s]\n", test.query, test.nextURL, params.Path)
		log.Printf("Test %2d: %s\n", idx+1, msg)
		result := params.GetNextPageURL("pg", test.pageOffset)
		expected := fmt.Sprintf("%s/%s", rqURL, test.nextURL)
		assert.Equal(t, expected, result, msg)
	}
}

// TestParamsGetValue tests NewParams GetValue method
func TestParamsGetValue(t *testing.T) {
	p1 := getNewParamsFromGetRequest(t, "/test/path", "key=v1&key=v2&num1=100")
	t.Logf("p1: %+v\n", p1)
	assert.NotNil(t, p1)
	assert.Equal(t, "", p1.GetValue("foo"))
	assert.Equal(t, "contextValue", p1.GetValue("contextKey"))
	assert.Equal(t, "v1", p1.GetValue("key"))
	assert.Equal(t, []uint8([]byte(nil)), p1.GetBody(""))

	num1, _ := p1.GetInt("num1")
	assert.Equal(t, 100, num1)
	num2, err := p1.GetInt("num2")
	assert.Equal(t, "empty value", err.Error())
	assert.Equal(t, 0, num2)

	p2 := getNewParamsFromPostRequest(t,
		"/test/path?user=name1&q=name2&somekey=somevalue&num=xxx",
		[]byte(`{"msg":"post messages"}`))
	t.Logf("p2: %+v\n", p2)
	assert.NotNil(t, p2)
	assert.Equal(t, "", p2.GetValue("key"))
	assert.Equal(t, "contextValue", p2.GetValue("contextKey"))
	assert.Equal(t, "name1", p2.GetValue("user"))
	assert.Equal(t, p2.Body, p2.GetBody(""))
	assert.Equal(t, "post messages", p2.GetBody("msg"))
	assert.Equal(t, p2.Post, p2.GetBody("foo"))

	num, err := p2.GetInt("num")
	assert.Equal(t, "strconv.Atoi: parsing \"xxx\": invalid syntax", err.Error())
	assert.Equal(t, 0, num)
}

// TestParamsGetValues tests NewParams GetValues method
func TestParamsGetValues(t *testing.T) {
	params := getNewParamsFromPostRequest(t,
		"/test/path?user=name1&user=name2&somekey=somevalue",
		[]byte(`{"msg":"post messages"}`))
	assert.NotNil(t, params)
	assert.Equal(t, []string{""}, params.GetValues("key"))
	assert.Equal(t, []string{"contextValue"}, params.GetValues("contextKey"))
	assert.Equal(t, []string{"name1", "name2"}, params.GetValues("user"))

	p2 := getNewParamsFromPostRequest(t,
		"/test/path?user=name1&q=name2&somekey=somevalue",
		[]byte(`post message`))
	t.Logf("p2: %+v\n", p2)
	assert.NotNil(t, p2)
	assert.Equal(t, BodyParams{}, p2.GetBody("body"))
	assert.Equal(t, BodyParams{}, p2.GetBody("msg"))
}

func getNewParamsFromGetRequest(t *testing.T, reqPath, queryString string) *Params {
	reqUrl := fmt.Sprintf("%s?%s", reqPath, queryString)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	muxVars = muxVarsMock
	return NewParams(req)
}

func getNewParamsFromPostRequest(t *testing.T, reqPath string, body []byte) *Params {
	req, err := http.NewRequest("POST", reqPath, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	muxVars = muxVarsMock
	return NewParams(req)
}

func muxVarsMock(r *http.Request) map[string]string {
	return map[string]string{"contextKey": "contextValue"}
}
