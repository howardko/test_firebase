package notify

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	"github.com/gin-gonic/gin"
	"wordsmap/common/util"
)

var tmpdb PairPost
var pairID = "12345"

func SendEvent(c *gin.Context) {


}

func AddPair(pair *PairPost) (string, error){
	tmpdb = *pair
	return pairID, nil
}

func GetPair(pairID string) (*PairPost, error){
	return &tmpdb, nil
}

func sendEvent(token string, data map[string]string) map[string]string{

	ids := []string{
		token,
	}

	c := fcm.NewFcmClient(util.Config.Notify.Server_Key)
	c.NewFcmRegIdsMsg(ids, data)

	result, err := c.Send()

	if err == nil {
		return result.Results[0]
	} else {
		fmt.Println(err)
		return map[string]string{ "error": err.Error()}
	}

}
