package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/haroldpacha/customer-management/cache"
	"github.com/haroldpacha/customer-management/config"
	"github.com/haroldpacha/customer-management/controller"
	"github.com/haroldpacha/customer-management/middleware"
	"github.com/haroldpacha/customer-management/repository"
	"github.com/haroldpacha/customer-management/service"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = config.SetupDatabaseConnection()
	userRepository       repository.UserRepository       = repository.NewUserRepository(db)
	customerRepository   repository.CustomerRepository   = repository.NewCustomerRepository(db)
	departmentRepository repository.DepartmentRepository = repository.NewDepartmentRepository(db)

	rd              *redis.Client         = config.SetupRedisConnection()
	departmentCache cache.DepartmentCache = cache.NewDepartmentCache(rd, 1000)

	jwtService        service.JWTService        = service.NewJWTService()
	authService       service.AuthService       = service.NewAuthService(userRepository)
	userService       service.UserService       = service.NewUserService(userRepository)
	customerService   service.CustomerService   = service.NewCustomerService(customerRepository)
	departmentService service.DepartmentService = service.NewDepartmentService(departmentRepository)

	authController       controller.AuthController       = controller.NewAuthController(authService, jwtService)
	userController       controller.UserController       = controller.NewUserController(userService, jwtService)
	customerController   controller.CustomerController   = controller.NewCustomerController(customerService, jwtService)
	departmentController controller.DepartmentController = controller.NewDepartmentController(departmentService, departmentCache)
)

func Rest() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "PUT", "PATCH", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("login", authController.Login)
		authRoutes.POST("register", authController.Register)
	}

	userRoutes := r.Group("/api/v1/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("profile", userController.Get)
		userRoutes.PUT("", userController.Update)
	}

	customerRoutes := r.Group("/api/v1/customers", middleware.AuthorizeJWT(jwtService))
	{
		customerRoutes.GET("", customerController.All)
		customerRoutes.POST("", customerController.Insert)
		customerRoutes.GET("/:id", customerController.Get)
		customerRoutes.PUT("", customerController.Update)
		customerRoutes.DELETE("/:id", customerController.Delete)
	}

	departmentRoutes := r.Group("/api/v1/departments")
	{
		departmentRoutes.GET("/all", departmentController.All)
		departmentRoutes.POST("", departmentController.Insert)
		departmentRoutes.GET("/:id", departmentController.Show)
		departmentRoutes.PUT("", departmentController.Update)
		departmentRoutes.DELETE("/:id", departmentController.Delete)
	}

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "API v1.0",
	// 	})
	// })

	r.Run()
}
