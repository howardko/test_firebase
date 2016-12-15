package notify

import (
	"github.com/NaySoftware/go-fcm"
	"wordsmap/common/util"
	"wordsmap/common/Error"
)

var tmpdb PairPost
var pairID = "12345"

func AddPair(pair *PairPost) (string, error){
	tmpdb = *pair
	return pairID, nil
}

func GetPair(pairID string) (*PairPost, error){
	return &tmpdb, nil
}

func SendEvent(pairID string, title string, message string, data map[string]string) (error, map[string]string){

	// get token from pair
	token := tmpdb.Token
	ids := []string{
		token,
	}

	c := fcm.NewFcmClient(util.Config.Notify.Server_Key)
	payload := fcm.NotificationPayload{Title: title, Body: message}
	c.SetNotificationPayload(&payload)
	c.NewFcmRegIdsMsg(ids, data)

	result, err := c.Send()

	if err == nil{
		if result.Ok == false{
			return Error.SendEventError, map[string]string{"error": string(result.StatusCode)}
		}

		if result.Success == 1 {
			return nil, result.Results[0]
		}else{
			return Error.SendEventError, map[string]string{"error": string(result.StatusCode)}
		}
	} else {
		return err, map[string]string{"error": err.Error()}
	}

}
