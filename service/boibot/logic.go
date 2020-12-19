package boibot

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	log "github.com/sirupsen/logrus"
)

func checkin(replyToken string, Input input) {
	log.Println(Input)
	empinfo, err := getEmpID(Input.Uid)
	if err != nil {
		return
	}
	if empinfo.EmpID == "" {
		msg := "คุณไม่อยู่ในระบบ กรุณาติดต่อ Human Resource (HR)"
		sendReplyMessage(replyToken, msg)
		return
	}
	t := time.Now()

	date := t.Format("2006-01-02")
	time := t.Format("15:04:05")
	log.Println(date)
	log.Println(time)
	var Querycheck ReportWork
	Querycheck, err = querycheck(empinfo.EmpID, date)
	if err != nil {
		log.Println(err)
		return
	}
	if Querycheck.EmpId != "" {
		msg := "คุณลงทะเบียนเข้างานไปแล้ว"
		sendReplyMessage(replyToken, msg)
		return
	}

	CheckinDB := checkinDB(date, time, empinfo.EmpID)
	if CheckinDB != nil {
		return
	}
	msg := "กรุณากรอกตำแหน่ง GPS"
	sendReplyMessage(replyToken, msg)
	return
}

func checkinGps(replyToken string, Input input) {
	log.Println(Input)
	empinfo, err := getEmpID(Input.Uid)
	if err != nil {
		return
	}
	if empinfo.EmpID == "" {
		msg := "ไม่อยู่ในระบบ"
		sendReplyMessage(replyToken, msg)
		return
	}
	t := time.Now()

	date := t.Format("2006-01-02")
	time := t.Format("15:04:05")
	log.Println(date)
	log.Println(time)

	Querycheck, err := querycheck(empinfo.EmpID, date)
	if err != nil {
		return
	}
	if Querycheck.EmpId == "" {
		msg := "กรุณา CHECK IN ก่อนทำรายการ"
		sendReplyMessage(replyToken, msg)
		return
	}
	if Querycheck.Location != ""{
		msg := "คุณได้ทำการกรอก ตำแหน่ง GPS ไปแล้ว"
		sendReplyMessage(replyToken, msg)
		return
	}

	CheckinDBGPS := checkinDBGPS(date, Input.Message, empinfo)
	if CheckinDBGPS != nil {
		return
	}
	msg := "กรุณากรอกรูปภาพ"
	sendReplyMessage(replyToken, msg)
	return
}

func checkinImage(replyToken string, Input input) {
	log.Println(Input)
	empinfo, err := getEmpID(Input.Uid)
	if err != nil {
		return
	}
	if empinfo.EmpID == "" {
		msg := "ไม่อยู่ในระบบ"
		sendReplyMessage(replyToken, msg)
		return
	}
	t := time.Now()

	date := t.Format("2006-01-02")
	time := t.Format("15:04:05")
	log.Println(date)
	log.Println(time)

	Querycheck, err := querycheck(empinfo.EmpID, date)
	if err != nil {
		return
	}

	if Querycheck.EmpId == "" {
		msg := "กรุณา CHECK IN ก่อนทำรายการ"
		sendReplyMessage(replyToken, msg)
		return
	}

	 QuerycheckGPS,err := querycheckGPS(empinfo.EmpID, date)
	 if QuerycheckGPS.Location == ""{
		 msg := "กรุณา GPS ก่อนทำรายการ"
		 sendReplyMessage(replyToken, msg)
		 return
	 }

	 if Querycheck.Picture != ""{
		 msg := "คุณได้ทำการกรอกรูปภาพไปแล้ว"
		 sendReplyMessage(replyToken, msg)
		 return
	 }



	CheckinDBGPS := checkinDBImage(date, Input.Message, empinfo)
	if CheckinDBGPS != nil {
		return
	}
	msg := "คุณได้ทำการ Check In สำเร็จ"
	sendReplyMessage(replyToken, msg)
	return
}

func sendReplyMessage(replyToken string, message string) error {
	if _, err := bot.ReplyMessage(replyToken,
		linebot.NewTextMessage(message)).Do(); err != nil {
		log.Infoln(replyToken)
		log.Infoln(message)
		return err
	}
	return nil
}

func checkout(replyToken string, uid string) {

	currentTime := time.Now()
	var date = currentTime.Format("2006-01-02")
	var time = currentTime.Format("15:04:05")
	log.Infof("date : %+v", date)
	log.Infof("time : %+v", time)

	//1
	var getEmpId EmpInfo
	getEmpId, err := getEmpID(uid)
	if err != nil {
		return
	}
	log.Infof("result : %+v", getEmpId.EmpID)
	if getEmpId.EmpID == "" {
		sendReplyMessage(replyToken, "คุณไม่อยู่ในระบบ กรุณาติดต่อ Human Resource (HR)")
		return
	} else {

		var getCheckin reportWork
		getCheckin, err2 := getDateCheckin(getEmpId.EmpID, date)
		if err2 != nil {
			log.Infof("result : %+v", err2)

		}
		log.Infof("result : %+v", (getCheckin.EmpID))
		log.Infof("result : %+v", (getCheckin))
		log.Infof("result : %+v", (getCheckin.ReportWorkID))
		if getCheckin.Location == "" && getCheckin.Picture == "" {
			sendReplyMessage(replyToken, "คุณยังไม่ได้เข้าสู่ระบบ")
			return
		}


		var getLocate reportWork
		getLocate, err5 := getLocation(getEmpId.EmpID, date)
		if err5 != nil {
			log.Infof("result : %+v", err5)

		}
		log.Infof("result : %+v", (getLocate.EmpID))
		log.Infof("result : %+v", (getLocate))
		log.Infof("result : %+v", (getLocate.ReportWorkID))
		if (getLocate.ReportWorkID) != 0 {
			sendReplyMessage(replyToken, "คุณได้ทำการลางานไปแล้ว ")
			return
		}

		var getCheckout reportWork
		getCheckout, err3 := getDateCheckout(getEmpId.EmpID, date)
		if err3 != nil {
			log.Infof("result : %+v", err3)
		}
		log.Infof("result : %+v", (getCheckout))
		log.Infof("result : %+v", (getCheckout.TimeCheckOut))
		//log.Infof("result : %+v", (getCheckout.TimeCheckout)[11:19])
		if (getCheckout.ReportWorkID) != 0 {
			sendReplyMessage(replyToken, "คุณได้ทำการออกจากระบบไปแล้ว")
			return
		}

		var result string
		result, err4 := updateTimeCheckout(getEmpId.EmpID, date, time)
		if err4 != nil {
			log.Infof("result : %+v", err4)
		}
		log.Infof("result : %+v", (result))
		sendReplyMessage(replyToken, result)
	}

	return
}

func getReportWork(command string) []reportWork1 {
	data := getReport(command)
	return data
}
func leave(replyToken string, uid string) {

	currentTime := time.Now()
	var date = currentTime.Format("2006-01-02")
	var time = currentTime.Format("15:04:05")
	log.Infof("date : %+v", date)
	log.Infof("time : %+v", time)

	var getEmpId EmpInfo
	getEmpId, err := getEmpID(uid)
	if err != nil {
		return
	}
	log.Infof("result : %+v", getEmpId.EmpID)
	if getEmpId.EmpID == "" {
		sendReplyMessage(replyToken, "คุณไม่อยู่ในระบบ กรุณาติดต่อ Human Resource (HR)")
		return
	} else {

		var getCheckinLeave reportWork
		getCheckinLeave, err2 := getDateLeave(getEmpId.EmpID, date)
		if err2 != nil {
			log.Infof("result : %+v", err2)

		}
		log.Infof("result : %+v", (getCheckinLeave.EmpID))
		log.Infof("result : %+v", (getCheckinLeave))
		log.Infof("result : %+v", (getCheckinLeave.ReportWorkID))
		if getCheckinLeave.ReportWorkID != 0  && getCheckinLeave.Location == "" && getCheckinLeave.Picture == ""{
			sendReplyMessage(replyToken, "คุณได้ทำการเข้าสู่ระบบเรียบร้อยแล้วไม่สามารถลางานได้ กรุณาติดต่อ HR ")
			return
		}

		var getLocate reportWork
		getLocate, err3 := getLocation(getEmpId.EmpID, date)
		if err3 != nil {
			log.Infof("result : %+v", err3)

		}
		log.Infof("result : %+v", (getLocate.EmpID))
		log.Infof("result : %+v", (getLocate))
		log.Infof("result : %+v", (getLocate.ReportWorkID))
		if (getLocate.ReportWorkID) != 0 {
			sendReplyMessage(replyToken, "คุณได้ทำการลางานไปแล้ว ")
			return
		}

		var getCount Count
		getCount, err4 := getCountLeave(getEmpId.EmpID)
		if err4 != nil {
			log.Infof("result : %+v", err4)

		}
		log.Infof("result : %+v", (getCount))
		if (getCount.SUM) > 15 {

			sendReplyMessage(replyToken, "คุณไม่เหลือวันลาอีกแล้ว")
			return
		}

		var result string
		result, err5 := insertLeave(getEmpId.EmpID, date)
		if err5 != nil {
			log.Infof("result : %+v", err5)
		}
		log.Infof("result : %+v", (result))
		sendReplyMessage(replyToken, result)

	}

}
