package boibot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

var bot *linebot.Client

func (ep *Endpoint) CallbackHandler(c *gin.Context) {
	var err error
	//var CHANNEL_SECRET = viper.GetString("boibot.channelSecret")
	//var CHANNEL_TOKEN = viper.GetString("boibot.channelToken")
	var CHANNEL_SECRET = "d53e953912578537136aae49806309b3"
	var CHANNEL_TOKEN = "shWHPy2AtysDJ8IkWowJC/NTQgpxu5ddDJ3GwD3eUQ6FcDnZ+Q1VgBrF9xyusFwet9PK5cBPmmxxLthorhuZrL94biE5JuQ+aJjEboIP5dQDLsgXMur0+QiioiMCL+xJxlA3aG5aFRDGy2HhHEPpfgdB04t89/1O/w1cDnyilFU="
	bot, err = linebot.New(
		CHANNEL_SECRET,
		CHANNEL_TOKEN,
	)
	if err != nil {
		log.Fatal(err)
	}
	events, err := bot.ParseRequest(c.Request) //message

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(400, gin.H{}) //Bad Request
		} else {
			c.JSON(500, gin.H{}) //Internet Server Error
		}
		return
	}

	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Println("===== " + message.Text + " =====")
				uid := event.Source.UserID
				text := message.Text
				Input := input{
					Uid:     uid,
					Message: text,
				}
				if text == "CHECK IN" {
					checkin(event.ReplyToken, Input)
				}

				emp, err := checkEmp(event.Source.UserID)


				if text == "CHECK IN" {
					//checkin(event.ReplyToken, message.Text)
				} else if text == "CHECK OUT" {
					checkout(event.ReplyToken, events[0].Source.UserID)
					//sendReplyMessage(event.ReplyToken,"คุณได้ทำการ Check Out สำเร็จ")
				} else if text == "LEAVE" {
					//sendReplyMessage(event.ReplyToken,"คุณได้ทำการลาหยุดไปแล้ว")
					leave(event.ReplyToken, events[0].Source.UserID)
				} else

				if emp.Rank_emp == "HR" {
					if text == "CHECK REPORT" {
						message1 := "กด 1 เพื่อเเสดงรายงานคนมาสายรายวัน\n" +
							"กด 2 เพื่อเเสดงรายงานคนลาหยุดเกิน 2 ครั้ง/เดือน\n" +
							"กด 3 เพื่อเเสดงคนมาสายเกิน 5 ครั้ง/เดือน"
						sendReplyMessage(event.ReplyToken, message1)
					}
					if text == "1" {
						command := "SELECT rw.emp_id , e.f_name , e.l_name \n" +
							"FROM Report_Work rw INNER JOIN Emp_info e ON rw.Emp_ID = e.Emp_ID\n" +
							"WHERE rw.time_check_in > '08:30:00' AND rw.date_work = '2020-12-09'\n" +
							"ORDER BY rw.emp_id ASC"
						result := getReportWork(command)
						if err != nil {
							return
						}
						message1 := fmt.Sprintf("วันนี้มีคนมาสาย %d คน \n %s", len(result), result)
						sendReplyMessage(event.ReplyToken, message1)
					} else if text == "2" {
						command := "SELECT rw.emp_id , e.f_name , e.l_name , COUNT(rw.location ) AS leave_amount\n" +
							"FROM Report_Work rw INNER JOIN Emp_info e ON rw.emp_id = e.emp_id \n" +
							"WHERE rw.location = 'leave' AND (rw.date_work BETWEEN '2020/12/01' AND '2020/12/31')\n" +
							"GROUP BY rw.emp_id , e.f_name , e.l_name \n" +
							"HAVING COUNT(rw.location ) >= 2\n" +
							"ORDER BY rw.emp_id ASC"
						result := getReportWork(command)
						if err != nil {
							return
						}
						message1 := fmt.Sprintf("เดือนนี้มีคนลาหยุดเกิน 2 ครั้ง %d คน \n %s", len(result), result)
						sendReplyMessage(event.ReplyToken, message1)
					} else if text == "3" {
						command := "SELECT rw.emp_id , e.f_name , e.l_name , COUNT(rw.time_check_in ) AS late_amount\n" +
							"FROM Report_Work rw INNER JOIN Emp_info e ON rw.emp_id = e.emp_id \n" +
							"WHERE rw.time_check_in > '09:00:00' AND (rw.date_work BETWEEN '2020/12/01' AND '2020/12/31')\n" +
							"GROUP BY rw.emp_id , e.f_name , e.l_name \n" +
							"HAVING COUNT(rw.time_check_in ) >= 5\n" +
							"ORDER BY rw.emp_id ASC"
						result := getReportWork(command)
						if err != nil {
							return
						}
						message1 := fmt.Sprintf("เดือนนี้มีคนมาสายเกิน 5 ครั้ง %d คน\n %s", len(result), result)
						sendReplyMessage(event.ReplyToken, message1)
					}
				} else if emp.Rank_emp == "Supervisor" {
					command := "SELECT emp_id,f_name,l_name\n" +
						"FROM Emp_info \n" +
						"WHERE emp_id  NOT  IN " +
						"(SELECT rw.emp_id " +
						"FROM report_work rw " +
						"WHERE rw.date_work = '2020-12-13')" +
						"AND dept_id = 'FN0001'"
					result := getReportWork(command)
					if err != nil {
						return
					}
					message1 := fmt.Sprintf("วันนี้มีคนไม่มาทำงาน %d คน \n %s", len(result), result)
					sendReplyMessage(event.ReplyToken, message1)
				} else {
					msg := "คุณไม่มีสิทธิ์เข้าถึง กรุณาติดต่อ HR"
					sendReplyMessage(event.ReplyToken, msg)
				}
			case *linebot.StickerMessage:
				log.Println("StickerMessage ================")
				log.Println("event.Source.UserID:", event.Source.UserID)
			case *linebot.LocationMessage:
				log.Println("LocationMessage ================")
				Input := input{
					Uid:     event.Source.UserID,
					Message: fmt.Sprintf("%f", message.Longitude) + "||" + fmt.Sprintf("%f", message.Latitude),
				}
				checkinGps(event.ReplyToken, Input)
			case *linebot.ImageMessage:
				log.Println("ImageMessage ================")
				Input := input{
					Uid:     event.Source.UserID,
					Message: message.ID,
				}
				checkinImage(event.ReplyToken, Input)
			default:
				//sendReplyMessage(event.ReplyToken, "Sorry, this command is not support.")
			}
		}
	}
}
