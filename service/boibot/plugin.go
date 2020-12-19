package boibot

import (
	mssql "api-wecode-supplychain/database/mssql"

	log "github.com/sirupsen/logrus"
)

func logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = mssql.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}

func getEmpID(uid string) (result EmpInfo, err error) {

	//var result1 EmpInfo

	if err = mssql.DB.Raw("SELECT emp_id FROM emp_info WHERE uid = ?", uid).
		Find(&result).Error; err != nil {
		//return
	}
	log.Infof("result : %+v", result)

	return result, nil
}

func querycheck(empid string, date string) (result ReportWork, err error) {

	if err = mssql.DB.Raw("SELECT date_work ,location ,time_check_in ,emp_id ,picture FROM report_work WHERE emp_id = ? and date_work = ?", empid, date).
		Find(&result).Error; err != nil {
		//return
	}
	log.Infof("result : %+v", result)

	return result, nil
}

func checkinDB(date string, time string, EmpInfo string) (err error) {
	log.Println(EmpInfo)
	result := ReportWorkinput{
		DateWork:    date,
		Picture:     "",
		Location:    "",
		TimeCheckIn: time,
		EmpId:       EmpInfo,
	}
	log.Println(result)
	if err := mssql.DB.Table("report_work").Save(result).Error; err != nil {
		return err
	}
	log.Println(result)
	return nil
}
func checkinDBGPS(date string, location string, EmpInfo EmpInfo) (err error) {
	log.Println(EmpInfo.EmpID)
	result := ReportWorkinput{
		Location: location,
	}
	if err := mssql.DB.Table("report_work").Where("emp_id = ? and date_work = ? ", EmpInfo.EmpID, date).Update(result).Error; err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func querycheckGPS(empid string, date string) (result ReportWork, err error) {

	if err = mssql.DB.Raw("SELECT location FROM report_work WHERE emp_id = ? and date_work = ?", empid, date).
		Find(&result).Error; err != nil {
		//return
	}
	log.Infof("result : %+v", result)

	return result, nil
}

func checkinDBImage(date string, image string, EmpInfo EmpInfo) (err error) {
	log.Println(EmpInfo.EmpID)
	result := ReportWorkinput{
		Picture: image,
	}
	if err := mssql.DB.Table("report_work").Where("emp_id = ? and date_work = ? ", EmpInfo.EmpID, date).Update(result).Error; err != nil {
		return err
	}
	log.Println(result)
	return nil
}
func getReport(command string) []reportWork1 {
	var result []reportWork1
	data := mssql.DB.Raw(command).Find(&result)
	if data == nil {
		return result
	}
	return result
}
func checkEmp(input string) (EmpInfo, error) {
	var result EmpInfo
	if err := mssql.DB.Raw("SELECT * FROM Emp_info Where UID = ?", input).Find(&result).Error; err != nil {
		return EmpInfo{}, err
	}
	log.Println(result)
	return result, nil
}
func getDateCheckin(Emp_ID string, date string) (result reportWork, err error) {
	if err = mssql.DB.Raw("SELECT report_work_id,emp_id,date_work,time_check_in,location,picture FROM report_work as rw WHERE emp_id = ? AND date_work = ? ", Emp_ID, date).
		Find(&result).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return result, err

}

func getDateCheckout(Emp_ID string, date string) (result reportWork, err error) {

	if err = mssql.DB.Raw("SELECT report_work_id,emp_id,date_work,time_check_out FROM report_work WHERE emp_id = ? AND date_work = ? AND time_check_out IS NOT NULL ", Emp_ID, date).
		Find(&result).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return result, err
}

func updateTimeCheckout(Emp_ID string, date string, time string) (result1 string, err1 error) {

	result := reportWork{
		TimeCheckOut: time,
	}
	if err := mssql.DB.Table("report_work").Where("emp_id = ? AND date_work = ? AND time_check_in IS NOT NULL AND time_check_out IS NULL ", Emp_ID, date).Update(result).Error; err != nil {
		return
	}
	log.Println(result)
	return "คุณได้ทำการ Check Out สำเร็จแล้ว", err1
}

// checkout
//var command1 = `select Emp_ID from Emp_info WHERE UID = '${uid}' ;`
//var command2 = `SELECT Date FROM Report_Work WHERE Emp_ID = ${empId} AND [Date] = '${date}' AND Time_login IS NOT NULL`
//var command3 = `SELECT Date FROM Report_Work WHERE Emp_ID = ${empId} AND [Date] = '${date}' AND Time_logout IS NOT NULL`
//var command4 = `UPDATE Report_Work SET Time_logout='${time}' WHERE [Date] = '${date}' AND Time_login IS NOT NULL`

//leave
//var command1 = `SELECT Emp_ID from Emp_info WHERE UID = '${uid}' ;`
//var command2 = `SELECT Date FROM Report_Work WHERE Emp_ID = ${empId} AND [Date] = '${date}' AND Time_login is NOT NULL`
//var command3 = `SELECT Location FROM Report_Work WHERE Emp_ID = ${empId} AND [Date] = '${date}' AND Location = 'leave'`
//var command4 = `SELECT COUNT(Report_work_ID) AS sum FROM Report_Work WHERE Emp_ID = ${empId} AND Location = 'leave' `
//var command5 = `INSERT INTO Report_Work ([Date],Location,Emp_ID) VALUES ('${date}','leave','${empId}') `

func getDateLeave(Emp_ID string, date string) (result reportWork, err error) {
	if err = mssql.DB.Raw("SELECT report_work_id,emp_id,date_work,time_check_in FROM report_work as rw WHERE emp_id = ? AND date_work = ? AND time_check_in IS NOT NULL AND location IS NOT NULL", Emp_ID, date).
		Find(&result).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return result, err

}

func getLocation(Emp_ID string, date string) (result reportWork, err error) {
	if err = mssql.DB.Raw("SELECT report_work_id,location FROM report_work WHERE emp_id = ? AND date_work = ? AND location = ? ", Emp_ID, date, "leave").
		Find(&result).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return result, err
}

func getCountLeave(Emp_ID string) (result Count, err error) {

	if err = mssql.DB.Raw("SELECT COUNT(report_work_id) AS sum FROM report_work WHERE emp_id = ? AND location = ? ", Emp_ID, "leave").
		Find(&result).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return result, err
}

func insertLeave(Emp_ID string, date string) (result1 string, err error) {

	insert := insertReportWork{
		DateWork: date,
		Location: "leave",
		EmpID:    Emp_ID,
	}
	if err = mssql.DB.Table("report_work").Where("emp_id = ? AND date = ? AND location IS NULL AND time_check_in IS NULL AND time_check_out IS NULL", Emp_ID, date).Save(insert).Error; err != nil {
		//if err = mssql.DB.Raw("INSERT INTO report_work WHERE emp_id = ? AND date = ? AND location IS NULL AND time_check_in IS NULL AND time_check_out IS NULL ,", Emp_ID,date).Save(&insert).Error; err != nil {
		//log.Infof("result : %+v", err)
		return
	}
	//log.Infof("result : %+v", result)
	return "คุณได้ทำการลางานสำเร็จแล้ว", err
}

//var command5 = `INSERT INTO Report_Work ([Date],Location,Emp_ID) VALUES ('${date}','leave','${empId}') `
