
func main() {
	srv := service.NewUserService()
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, srv.find())
	})
 }
 router.POST("/users", func(c *gin.Context) {
	var u service.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.String(http.StatusBadRequest, "ERROR AL CREAR EL USUARIO")
	}
	u = srv.Add(u)
	c.JSON(http.StatusCreated, u)
 })
 router.Run()
