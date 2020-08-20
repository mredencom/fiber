package compress

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
)

var filedata []byte

func init() {
	dat, err := ioutil.ReadFile("../.github/testdata/fs/img/fiber.png")
	if err != nil {
		panic(err)
	}
	filedata = dat
}

// go test -run Test_Compress
func Test_Compress_Gzip(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(filedata)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	utils.AssertEqual(t, "gzip", resp.Header.Get(fiber.HeaderContentEncoding))
}

func Test_Compress_Deflate(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(filedata)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Encoding", "deflate")

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	utils.AssertEqual(t, "deflate", resp.Header.Get(fiber.HeaderContentEncoding))
}

func Test_Compress_Brotli(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(filedata)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Accept-Encoding", "br")

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	utils.AssertEqual(t, "br", resp.Header.Get(fiber.HeaderContentEncoding))
}