package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}

// Custom handler added by me
func CustomHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"response": "This API is powered by GoBuffalo"}))
}

// second custom handler added by me
func SecondCustomHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"resonse": "There is always plan B."}))
}
