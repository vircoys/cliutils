package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	nhttp "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	tu "gitlab.jiagouyun.com/cloudcare-tools/cliutils/testutil"
)

func TestBytesBody(t *testing.T) {
	errOK := NewNamespaceErr(nil, nhttp.StatusOK, "")
	bytesBody := "this is bytes response body"

	router := gin.New()
	g := router.Group("")
	g.GET("/bytes-body", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.Writer.Header().Set("X-Latest-Time", fmt.Sprintf("%s", time.Now()))
		errOK.HttpBody(c, []byte(bytesBody))
	})

	ts := httptest.NewServer(router)

	defer ts.Close()

	time.Sleep(time.Second)

	resp, err := http.Get(fmt.Sprintf("%s%s", ts.URL, "/bytes-body"))
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	for k, v := range resp.Header {
		t.Logf("%s: %v", k, v)
	}

	tu.Equals(t, bytesBody, string(b))
}

func TestHTTPErr(t *testing.T) {
	errTest := NewNamespaceErr(errors.New("test error"), nhttp.StatusForbidden, "testing")
	errOK := NewNamespaceErr(nil, nhttp.StatusOK, "")

	DefaultNamespace = "testing2"
	errTest2 := NewErr(errors.New("test error2"), nhttp.StatusForbidden)

	router := gin.New()
	g := router.Group("")

	okbody := map[string]interface{}{
		"data1": 1,
		"data2": "abc",
	}

	g.GET("/err", func(c *gin.Context) { HttpErr(c, errTest) })
	g.GET("/err2", func(c *gin.Context) { HttpErr(c, errTest2) })
	g.GET("/err3", func(c *gin.Context) { HttpErr(c, fmt.Errorf("500 error")) })
	g.GET("/errf", func(c *gin.Context) { HttpErrf(c, errTest, "%s: %s", "this is a test error", "ignore me") })
	g.GET("/ok", func(c *gin.Context) { errOK.HttpBody(c, okbody) })
	g.GET("/oknilbody", func(c *gin.Context) { errOK.HttpBody(c, nil) })
	g.GET("/errmsg", func(c *gin.Context) { err := Error(errTest, "this is a error with specific message"); HttpErr(c, err) })
	g.GET("/errfmsg", func(c *gin.Context) {
		err := Errorf(errTest, "%s: %v", "this is a message with fmt", map[string]int{"abc": 123})
		HttpErr(c, err)
	})

	srv := nhttp.Server{
		Addr:    ":8090",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != nhttp.ErrServerClosed {
			t.Log(err)
		}
	}()

	time.Sleep(time.Second)
	defer srv.Close()

	cases := []struct {
		u      string
		expect string
	}{
		{
			u:      "http://localhost:8090/errmsg",
			expect: `{"error_code":"testing.testError","message":"this is a error with specific message"}`,
		},

		{
			u:      "http://localhost:8090/errfmsg",
			expect: `{"error_code":"testing.testError","message":"this is a message with fmt: map[abc:123]"}`,
		},

		{
			u: "http://localhost:8090/err",
			expect: func() string {
				j, err := json.Marshal(errTest)
				if err != nil {
					t.Fatal(err)
				}
				return string(j)
			}(),
		},
		{
			u: "http://localhost:8090/err2",
			expect: func() string {
				j, err := json.Marshal(errTest2)
				if err != nil {
					t.Fatal(err)
				}
				return string(j)
			}(),
		},
		{
			u:      "http://localhost:8090/err3",
			expect: `{"error_code":"testing2.500Error"}`,
		},
		{
			u:      "http://localhost:8090/errf",
			expect: `{"error_code":"testing.testError","message":"this is a test error: ignore me"}`,
		},
		{
			u: "http://localhost:8090/ok",
			expect: func() string {
				x := struct {
					Content interface{} `json:"content"`
				}{
					Content: okbody,
				}
				j, err := json.Marshal(x)
				if err != nil {
					t.Fatal(err)
				}
				return string(j)
			}(),
		},

		{
			u:      "http://localhost:8090/oknilbody",
			expect: "",
		},
	}

	for _, tc := range cases {
		t.Run("", func(t *testing.T) {

			resp, err := http.Get(tc.u)
			if err != nil {
				t.Logf("get error: %s, ignored", err)
				return
			}

			if resp.Body != nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Error(err)
					return
				}

				tu.Equals(t, tc.expect, string(body))
				resp.Body.Close()
			} else {
				t.Error("body should not be nil")
			}
		})
	}
}
