package httpserver

// HttpServer is a HTTP server for ports
type HttpServer struct {
	userService     UserService
	tokenService    TokenService
	bookService     BookService
	categoryService CategoryService
	cartService     CartService
}

// NewHttpServer creates a new HTTP server for ports
func NewHttpServer(userService UserService, tokenService TokenService, bookService BookService,
	categoryService CategoryService, cartService CartService) HttpServer {
	return HttpServer{
		userService:     userService,
		tokenService:    tokenService,
		bookService:     bookService,
		categoryService: categoryService,
		cartService:     cartService,
	}
}
