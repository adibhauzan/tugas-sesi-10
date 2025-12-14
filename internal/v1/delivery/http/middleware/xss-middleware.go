package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func XSSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := bluemonday.StrictPolicy()

		query := c.Request.URL.Query()
		for key, values := range query {
			for i, v := range values {
				values[i] = p.Sanitize(v)
			}
			query[key] = values
		}
		c.Request.URL.RawQuery = query.Encode()

		_ = c.Request.ParseForm()
		for key, values := range c.Request.PostForm {
			for i, v := range values {
				values[i] = p.Sanitize(v)
			}
			c.Request.PostForm[key] = values
		}

		if c.Request.Body != nil && c.ContentType() == "application/json" {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil && len(bodyBytes) > 0 {
				var jsonData map[string]interface{}
				if err := json.Unmarshal(bodyBytes, &jsonData); err == nil {
					sanitizeJSONMap(jsonData, p)
					safeBody, _ := json.Marshal(jsonData)
					c.Request.Body = io.NopCloser(bytes.NewBuffer(safeBody))
					c.Request.ContentLength = int64(len(safeBody))
				} else {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}
		}

		c.Next()
	}
}

func sanitizeJSONMap(m map[string]interface{}, p *bluemonday.Policy) {
	for k, v := range m {
		switch val := v.(type) {
		case string:
			m[k] = p.Sanitize(val)
		case map[string]interface{}:
			sanitizeJSONMap(val, p)
		case []interface{}:
			sanitizeJSONArray(val, p)
		}
	}
}

func sanitizeJSONArray(a []interface{}, p *bluemonday.Policy) {
	for i, v := range a {
		switch val := v.(type) {
		case string:
			a[i] = p.Sanitize(val)
		case map[string]interface{}:
			sanitizeJSONMap(val, p)
		case []interface{}:
			sanitizeJSONArray(val, p)
		}
	}
}
