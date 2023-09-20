package features_proxy

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotiFunc(ctx *gin.Context, noti_id string, title string, body string) {
	requestBodyData := map[string]interface{}{
		"to": noti_id,
		"notification": map[string]string{
			"title": title,
			"body":  body,
			"icon":  "https://res.cloudinary.com/dkjvy425b/image/upload/v1692261477/mtm_community_profile/bk0x4o6gkqlu8q53frys.png",
		},
	}

	requestBodyJSON, err := json.Marshal(requestBodyData)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(requestBodyJSON))
		if err != nil {
			return
		}

		req.Header.Set("Authorization", "key=AAAAAinRttw:APA91bH3Ub9PUSN2f_oo_RgrAPHWGhRrAD07yyrNXlFD88sQ1gTRy4OEFr1KSU63P0rQ9Ne1435EixHXBb--TSGKrB3A_Ivfv4TQ9Dt17_Hllbvsn0P-EL6lVQBOSZbleDfMZfcDO0bT") // Replace with your actual authorization token

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			// Handle error
			return
		}
		defer resp.Body.Close()
	}()
}
