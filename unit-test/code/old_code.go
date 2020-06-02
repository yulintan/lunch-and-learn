// START OMIT
func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("userid")

	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		glog.Info(err)
	}
	c.BindJSON(&user)

	db.Save(&user)
	c.JSON(200, user)
}

// END OMIT
