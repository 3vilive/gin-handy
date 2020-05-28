package handy

import (
	"io"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GinCtxGetInt64 get int64 from query
func GinCtxGetInt64(c *gin.Context, name string) (int64, error) {
	strVal := c.Query(name)
	if strVal == "" {
		return 0, nil
	}

	return strconv.ParseInt(strVal, 10, 64)
}

// GinCtxGetFile is short hand function to get file reader from form
func GinCtxGetFile(c *gin.Context, name string) (*multipart.FileHeader, io.Reader, error) {
	fileForm, err := c.FormFile(name)
	if err != nil {
		return nil, nil, err
	}

	f, err := fileForm.Open()
	if err != nil {
		return nil, nil, err
	}
	return fileForm, f, nil
}

// GinCtxGetFileReader is short hand function to get file reader from form
func GinCtxGetFileReader(c *gin.Context, name string) (io.Reader, error) {
	_, reader, err := GinCtxGetFile(c, name)
	return reader, err
}

// GinCtxSetCORS write CORS headers
func GinCtxSetCORS(c *gin.Context, origin string) {
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE")
	c.Header("Access-Control-Allow-Origin", origin)
}
