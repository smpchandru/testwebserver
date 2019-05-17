package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handle(t *testing.T) {
	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handle(tt.args.w, tt.args.req)
			bt, _ := ioutil.ReadAll(tt.args.w.Body)
			if string(bt) != "Hello" {
				t.Fatal("Not got actual resp")
			}
		})
	}
}

func Test_main(t *testing.T) {
    fmt.Printf(formafunc name(varName, type){
    
    t, a)
}
