package rereadhttpbody

import (
	"bytes"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func HttpBodyReadAll(c echo.Context) (bodyBytes []byte, err error) {
	if c.Request().Body == nil {
		return nil, nil
	}
	bodyBytes, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return
	}
	c.Request().Body.Close()
	return
}

func HttpBodyReWrite(c echo.Context, bodyBytes []byte) {
	if bodyBytes == nil {
		c.Request().Body = nil
	} else {
		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
}

func run() {
	e := echo.New()
	e.Any("/", func(c echo.Context) error {
		bodyBytes, err := HttpBodyReadAll(c)
		if err != nil {
			fmt.Println(err.Error())
		}

		if bodyBytes == nil {
			fmt.Println("bodybytes nil")
		}
		fmt.Println("first body bytes: ", string(bodyBytes))

		HttpBodyReWrite(c, bodyBytes)
		if c.Request().Body == nil {
			fmt.Println("context body nil")
		}

		bodyBytes, _ = io.ReadAll(c.Request().Body)
		fmt.Println("second body bytes: ", string(bodyBytes))
		return nil
	})
	e.Logger.Fatal(e.Start(":8081"))
}
