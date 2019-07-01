package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	handlers "github.com/kimberly.luna/proxy-app/api/handlers"
	server "github.com/kimberly.luna/proxy-app/api/server"
	utils "github.com/kimberly.luna/proxy-app/api/utils"
	assert "github.com/stretchr/testify/assert"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		//función anónima
		utils.LoadEnv()
		app := server.SetUp()
		handlers.HandlerRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("Server Running... ")
}

type Response struct {
	Status       string         `json:"status,omitempty"`
	Response     string         `json:"result,omitempty"`
	ResponseText []ResponseText `json:res,omitempty`
}

type ResponseText struct {
	Domain string
}

func TestAlgorithm(t *testing.T) {

	cases := []struct {
		// Attr
		Domain string
		Output string
	}{
		//structs
		{Domain: "alpha", Output: `["alpha"]`},
		{Domain: "beta", Output: `["beta, alpha"]`},
		{Domain: "", Output: "domain error"},
	}

	valuesToCompare := &Response{}
	client := http.Client{}

	for _, singleCase := range cases {
		fmt.Println("domain: ", singleCase.Domain)
		req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
		req.Header.Add("domain", singleCase.Domain)

		response, err := client.Do(req)
		bytes, err := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(bytes, valuesToCompare)

		fmt.Println("HEY ", valuesToCompare)
		assert.Nil(t, err)
		assert.Equal(t, singleCase.Output, valuesToCompare.Response)
	}

}
