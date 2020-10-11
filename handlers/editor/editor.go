package editor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ExampleRequest struct {
	FirstName string `json:"first_name" form:"first_name" query:"first_name"`
	LastName  string `json:"last_name" form:"last_name" query:"last_name"`
}

func Editor(c echo.Context) error {
	// Bind the input data to ExampleRequest
	exampleRequest := new(ExampleRequest)
	if err := c.Bind(exampleRequest); err != nil {
		return err
	}

	// Manipulate the input data
	greeting := exampleRequest.FirstName + " " + exampleRequest.LastName

	// Pretty print the json []byte
	var resp bytes.Buffer
	var b = []byte(
		fmt.Sprintf(`{
		 "first_name": %q,
		 "last_name": %q,
		 "msg": "Hello %s"
	   }`, exampleRequest.FirstName, exampleRequest.LastName, greeting),
	)
	err := json.Indent(&resp, b, "", "  ")
	if err != nil {
		return err
	}

	// Return the json to the client
	return c.JSONBlob(
		http.StatusOK,
		[]byte(
			fmt.Sprintf("%s", resp.Bytes()),
		),
	)

}
