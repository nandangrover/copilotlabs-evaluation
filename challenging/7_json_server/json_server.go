import jsonServer "github.com/typicode/json-server"
server := jsonServer.New()
router := jsonServer.Router("db.json")
middlewares := jsonServer.Defaults()

// Set default middlewares (logger, static, cors and no-cache)
server.Use(middlewares)

// Add custom routes before JSON Server router
server.Get("/echo", func(req jsonServer.Request, res jsonServer.Response) {
	res.JSONP(req.Query)
})

// To handle POST, PUT and PATCH you need to use a body-parser
// You can use the one used by JSON Server
server.Use(jsonServer.BodyParser)
server.Use(func(req jsonServer.Request, res jsonServer.Response, next jsonServer.Next) {
	if req.Method == "POST" {
		req.Body["createdAt"] = time.Now()
	}
	// Continue to JSON Server router
	next()
})

// Use default router
server.Use(router)
server.Listen(3000, func() {
	fmt.Println("JSON Server is running")
})
