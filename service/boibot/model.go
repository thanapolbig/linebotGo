package boibot

import (
	"time"
)

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type reportWork struct {
	ReportWorkID int64  `json:"report_work_id"`
	DateWork     string `json:"date_work"`
	Picture      string `json:"picture"`
	Location     string `json:"location"`
	TimeCheckIn  string `json:"time_check_in"`
	TimeCheckOut string `json:"time_check_out"`
	EmpID        string `json:"emp_id"`
}

type ReportWork struct {
	DateWork     string `json:"date_work"`
	Picture      string `json:"picture"`
	Location     string `json:"location"`
	TimeCheckIn  string `json:"time_check_in"`
	TimeCheckOut string `json:"time_check_out"`
	EmpId        string `json:"emp_id"`
}

type ReportWorkinput struct {
	DateWork    string `json:"date_work"`
	Picture     string `json:"picture"`
	Location    string `json:"location"`
	TimeCheckIn string `json:"time_check_in"`
	EmpId       string `json:"emp_id"`
}

type Count struct {
	ReportWorkID int64 `json:"report_work_id"`
	SUM          int64 `json:"sum"`
}

type insertReportWork struct {
	DateWork string `json:"date_work"`
	Location string `json:"location"`
	EmpID    string `json:"emp_id"`
}

type input struct {
	Uid     string `json:"uid"`
	Message string `json:"message"`
}
type reportWork1 struct {
	Emp_id string `json:"emp_id"`
	F_name string `json:"f_name"`
	L_name string `json:"l_name"`
}

type EmpInfo struct {
	EmpID    string `json:"emp_id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Rank_emp string `json:"rank_emp"`
	UID      string `json:"uid"`
	DeptID   string `json:"dept_id"`
}
