// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Resumes is the golang structure for table resumes.
type Resumes struct {
	Id                     int    `orm:"id,primary"               json:"id"`                     //
	Describe               string `orm:"describe"                 json:"describe"`               //
	FileUrl                string `orm:"file_url"                 json:"fileUrl"`                //
	CollegeMajor           string `orm:"college_major"            json:"collegeMajor"`           //
	Name                   string `orm:"name"                     json:"name"`                   //
	Gender                 int    `orm:"gender"                   json:"gender"`                 //
	Grade                  int    `orm:"grade"                    json:"grade"`                  //
	DepartmentType         int    `orm:"department_type"          json:"departmentType"`         //
	Experience             int    `orm:"experience"               json:"experience"`             //
	TelephoneNumber        string `orm:"telephone_number"         json:"telephoneNumber"`        //
	QqNumber               string `orm:"qq_number"                json:"qqNumber"`               //
	WechatNumber           string `orm:"wechat_number"            json:"wechatNumber"`           //
	InterviewId            int    `orm:"interview_id"             json:"interviewId"`            //
	InterviewerId          int    `orm:"interviewer_id"           json:"interviewerId"`          //
	SendTime               int64  `orm:"send_time"                json:"sendTime"`               //
	InitialScreeningResult int    `orm:"initial_screening_result" json:"initialScreeningResult"` //
	InitialScreeningTime   int64  `orm:"initial_screening_time"   json:"initialScreeningTime"`   //
	InterviewResult        int    `orm:"interview_result"         json:"interviewResult"`        //
	InterviewEvaluation    string `orm:"interview_evaluation"     json:"interviewEvaluation"`    //
	InterviewTime          int64  `orm:"interview_time"           json:"interviewTime"`          //
}
