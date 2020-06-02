import (
	"net/http"
	"strconv"
)

// START OMIT
func (h *UserHandler) UpdateUser(c helper.GinContext) {
	var user User
	idStr = c.Param("userid") // HLxxx
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		glog.Info(err)
		return
	}
	if !h.UserSvc.Exist(id) {
		c.AbortWithStatus(http.StatusNotFound)
		glog.Info(err)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		glog.Info(err)
		return
	}
	u, err := h.UserSvc.UpdateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		glog.Info(err)
		return
	}
	c.JSON(200, u)
} // END OMIT
