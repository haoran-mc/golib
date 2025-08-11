package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/haoran-mc/golib/pkg/log"
	"github.com/labstack/echo/v4"
)

func SingleFileUpload(c echo.Context) error {
	file, err := c.FormFile("f1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	// Note: in a real world app, you'd want to sanitize the filename.
	dst, err := os.Create(fmt.Sprintf("/tmp/%s", file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

func MultiFilesUpload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]

	for index, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(fmt.Sprintf("/tmp/%s_%d", file.Filename, index))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		log.Info("file uploaded", "filename", file.Filename)
	}

	return c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
