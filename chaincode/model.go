package main

const (
	HEALTH          = "HEALTH"
	HEALTHREPORT    = "HEALTHREPORT"
	EDUCATION       = "EDUCATION"
	QUALIFICATION   = "QUALIFICATION"
	PERSONAL        = "PERSONAL"
	IDENTITY        = "IDENTITY"
	DOB             = "DOB"
	PERSONALREQUEST = "PERSONALREQUEST"
)

type Health struct {
	DocType string           `json:"docType"`
	ID      string           `json:"id"`
	Reports map[string]string `json:"reports"` // reports UUID of Recipts pic => epoch time
}
type HealthReport struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`          // key of image in BC
	Doctor  string `json:"doctor_name"` // Doctor signed by
	Type    string `json:"type_report"` // which type of health report it is dob, injury
}
type DoB struct {
	DOBDetails HealthReport `json:"dob_details"`
	DOB        string       `json:"dob"`
	IssueTime  int64        `json:"issue_time"`
}
type Education struct {
	DocType       string            `json:"docType"`
	ID            string            `json:"id"`
	Qualification map[string]string `json:"qualification"`
}
type Qualification struct {
	DocType   string `json:"docType"`
	ID        string `json:"id"` // Qualification link
	IssueTime int64 `json:"issue_time"`
	Status    string `json:"status"` // verified = 1 not verified =0
}
type Personal struct {
	DocType  string `json:"docType"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	DOB      string `json:"dob"`
	PAddress string `json:"p_address"`
	CAddress string `json:"c_address"`
	Phone    string `json:"phone_number"`
	Parent   string `json:"parent_name"`
	Status   string `json:"status"`   // verified = 1 not verified =0
	Photo    string `json:"photo_id"` // photo link
}

type Identity struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`
	// keys of Pe,Ed,He
	PersonalDetails  string            `json:"personal_details"`
	EducationDetails string            `json:"education_etails"`
	HealthDetails    string            `json:"health_details"`
	Requests         map[string]string `json:"all_request"` // type of request => key of request
}
type Request struct {
	DocType     string `json:"docType"`
	ID          string `json:"id"`
	From        string `json:"request_maker"`
	Status      string `json:"status"` // 1 accepted -1: rejected 0: request created
	GivenData   string `json:"given_date"`
	RequestTime int64  `json:"request_time"`
}
