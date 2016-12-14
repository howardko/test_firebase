package notify

import (
	"github.com/gin-gonic/gin"
	"wordsmap/common/Response"
)

// SetupRoute for notify api
func SetupRoute(r *gin.RouterGroup) {

	// update_tunnel_agent
	// curl -X POST -d '{"user_id":"howardko","token":"23786dgqjhbamio' $(ip dev):8080/v0.1/notify/pair
	r.POST("/notify/pair", AddPairHandler)

	// get_tunnels
	r.GET("/notify/pair/:pair_id", GetPairHandler)

}

var GetPairHandler = func(c *gin.Context) {
	pairID := c.Param("pair_id")
	pair, err := GetPair(pairID)
	Response.JSON(c, pair, err)
}

var AddPairHandler = func(c *gin.Context) {
	var json PairPost
	if c.BindJSON(&json) == nil {
		pairID, err := AddPair(&json)
		Response.JSON(c, pairID, err)
	}
}