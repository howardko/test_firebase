package notify

import (
	"github.com/gin-gonic/gin"
	"wordsmap/common/Response"
)

// SetupRoute for notify api
func SetupRoute(r *gin.RouterGroup) {

	// add a notification pair
	// curl -X POST -d '{"user_id":"howardko","token":"23786dgqjhbamio"}' $(ip dev):8080/v0.1/notify/pair
	r.POST("/notify/pair", AddPairHandler)

	// get notification pair
	r.GET("/notify/pair/:pair_id", GetPairHandler)

	// send event to a pair
	// curl -X POST -d '{"pair_id":"12345","title": "hello", "message": "hey", "data":{"key1": "data1"}}' $(ip dev):8080/v0.1/notify/event
	r.POST("/notify/send", SendEventHandler)

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

var SendEventHandler = func(c *gin.Context) {
	var json SendPost
	if c.BindJSON(&json) == nil {
		//switch t := json.Data.(type){
		//case map[string]interface{}:
		//	fmt.Println("here")
		//default:
		//	fmt.Println(t)
		//}
		err, results:= SendEvent(json.Pair_ID, json.Title, json.Message, json.Data)
		Response.JSON(c, results, err)
	}
}