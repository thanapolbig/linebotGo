package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type textBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type lineMessage struct {
	Messages []textBody `json:"messages"`
}

// enable_scheduler: true
// count down : "@every 25s"
// time on day : "TZ=Asia/Bangkok 59 23 * * *" #time_zone MM:HH
func StartScheduler() {
	_enableScheduler := false
	if _enableScheduler {
		log.Infoln("Start Scheduler")
		startCheckInBroadcastScheduler()
		startCheckInBroadcastScheduler1()
		startCheckInBroadcastScheduler2()
		// startAlertScheduler()
	}
}

func startCheckInBroadcastScheduler() {
	_timeScheduler := "TZ=Asia/Bangkok 30 08 * * *" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, broadcastJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartCheckInBroadcastScheduler : ", msg)
}
func startCheckInBroadcastScheduler1() {
	_timeScheduler := "TZ=Asia/Bangkok 50 08 * * *" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, broadcastJob1)
	c.Start()
	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartCheckInBroadcastScheduler : ", msg)
}

func startCheckInBroadcastScheduler2() {
	_timeScheduler := "TZ=Asia/Bangkok 00 09 * * *" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, broadcastJob2)
	c.Start()
	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartCheckInBroadcastScheduler : ", msg)
}

func startAlertScheduler() {
	_timeScheduler := "@every 10s" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, alertCheckInJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Alert finish check in step running at: %s", c.Entries()[0].Next.String())
	log.Infoln("StartAlertScheduler : ", msg)
}

func broadcastJob() {
	log.Infoln("===============================")
	log.Infoln("Check In Broadcast Scheduler : ", time.Now())
	log.Infoln("===============================")

	var payload lineMessage
	payload.Messages = append(payload.Messages, textBody{
		Type: "text",
		Text: "อย่าลืม CHECK IN นะ !!!",
	})
	method := "POST"
	url := "https://api.line.me/v2/bot/message/broadcast"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer shWHPy2AtysDJ8IkWowJC/NTQgpxu5ddDJ3GwD3eUQ6FcDnZ+Q1VgBrF9xyusFwet9PK5cBPmmxxLthorhuZrL94biE5JuQ+aJjEboIP5dQDLsgXMur0+QiioiMCL+xJxlA3aG5aFRDGy2HhHEPpfgdB04t89/1O/w1cDnyilFU=",
	}
	apiCaller(payload, method, url, headers)

}

func broadcastJob1() {
	log.Infoln("===============================")
	log.Infoln("Check In Broadcast Scheduler : ", time.Now())
	log.Infoln("===============================")

	var payload lineMessage
	payload.Messages = append(payload.Messages, textBody{
		Type: "text",
		Text: "เหลือ 10 นาที",
	})
	method := "POST"
	url := "https://api.line.me/v2/bot/message/broadcast"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer shWHPy2AtysDJ8IkWowJC/NTQgpxu5ddDJ3GwD3eUQ6FcDnZ+Q1VgBrF9xyusFwet9PK5cBPmmxxLthorhuZrL94biE5JuQ+aJjEboIP5dQDLsgXMur0+QiioiMCL+xJxlA3aG5aFRDGy2HhHEPpfgdB04t89/1O/w1cDnyilFU=",
	}
	apiCaller(payload, method, url, headers)

}

func broadcastJob2() {
	log.Infoln("===============================")
	log.Infoln("Check In Broadcast Scheduler : ", time.Now())
	log.Infoln("===============================")

	var payload lineMessage
	payload.Messages = append(payload.Messages, textBody{
		Type: "text",
		Text: "หมดเวลาลงทะเบียน",
	})
	method := "POST"
	url := "https://api.line.me/v2/bot/message/broadcast"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer shWHPy2AtysDJ8IkWowJC/NTQgpxu5ddDJ3GwD3eUQ6FcDnZ+Q1VgBrF9xyusFwet9PK5cBPmmxxLthorhuZrL94biE5JuQ+aJjEboIP5dQDLsgXMur0+QiioiMCL+xJxlA3aG5aFRDGy2HhHEPpfgdB04t89/1O/w1cDnyilFU=",
	}
	apiCaller(payload, method, url, headers)

}

func alertCheckInJob() {
	log.Infoln("##############################")
	log.Infoln("Alert Check In Scheduler : ", time.Now())
	log.Infoln("##############################")

	var payload lineMessage
	payload.Messages = append(payload.Messages, textBody{
		Type: "text",
		Text: "อย่าลืม CHECK IN นะ !!!",
	})
	method := "POST"
	url := "https://api.line.me/v2/bot/message/broadcast"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer shWHPy2AtysDJ8IkWowJC/NTQgpxu5ddDJ3GwD3eUQ6FcDnZ+Q1VgBrF9xyusFwet9PK5cBPmmxxLthorhuZrL94biE5JuQ+aJjEboIP5dQDLsgXMur0+QiioiMCL+xJxlA3aG5aFRDGy2HhHEPpfgdB04t89/1O/w1cDnyilFU=",
	}
	apiCaller(payload, method, url, headers)

}

type apiCallError struct {
	s string
}

func apiCaller(payload interface{}, method string, url string, headers map[string]string) ([]byte, error) {

	payloadAsBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadAsBytes))
	req.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode == 200 {
		return body, nil
	} else {
		return nil, err
	}
}
