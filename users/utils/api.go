package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/micro/micro/v3/service/errors"
	"github.com/nyaruka/phonenumbers"

	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"
)

var (
	dateRegex = regexp.MustCompile(`date\((.+)\)`)
)

type listRequestBody struct {
	SpecId  string `json:"specId"`
	Filters struct {
		AccountStatusId []string `json:"account_status_id"`
		Deleted         string   `json:"deleted"`
	} `json:"filters"`
	Start  int `json:"start"`
	Count  int `json:"count"`
	Orders []struct {
		Name     string `json:"name"`
		Asc      bool   `json:"asc"`
		NullSide bool   `json:"nullSide"`
	} `json:"orders"`
	ExcludedFieldNames []string `json:"excludedFieldNames"`
}

type Account struct {
	AccountNotes            string  `json:"account_notes"`
	AccountStatusId         int     `json:"account_status_id"`
	Address                 string  `json:"address"`
	Address2                string  `json:"address2"`
	AdminType               int     `json:"admin_type"`
	Balance                 float64 `json:"balance"`
	BouncedId               int     `json:"bounced_id"`
	ByAccountId             int     `json:"by_account_id"`
	CheckBox                bool    `json:"checkBox"`
	City                    string  `json:"city"`
	CoachType               int     `json:"coachType"`
	ColumnNameWarning       bool    `json:"columnNameWarning"`
	ConvertedFromLessonOnly bool    `json:"converted_from_lesson_only"`
	CustFld                 string  `json:"cust_fld"`
	Deleted                 int     `json:"deleted"`
	DtDropped               string  `json:"dt_dropped"`
	DtJoined                string  `json:"dt_joined"`
	DtLastsignedon          string  `json:"dt_lastsignedon"`
	DtLregfee               string  `json:"dt_lregfee"`
	DtModified              string  `json:"dt_modified"`
	DtPayMethodModified     string  `json:"dt_pay_method_modified"`
	DtPswdSent              string  `json:"dt_pswd_sent"`
	DuesAccountVaultTokenId int     `json:"dues_account_vault_token_id"`
	Email                   string  `json:"email"`
	EmailOpt1               string  `json:"email_opt1"`
	EmailOpt1Valid          bool    `json:"email_opt1_valid"`
	EmailOpt1ValidationCode int     `json:"email_opt1_validation_code"`
	EmailOpt2               string  `json:"email_opt2"`
	EmailOpt2Valid          bool    `json:"email_opt2_valid"`
	EmailOpt2ValidationCode int     `json:"email_opt2_validation_code"`
	EmailOpt3               string  `json:"email_opt3"`
	EmailOpt3Valid          bool    `json:"email_opt3_valid"`
	EmailOpt3ValidationCode int     `json:"email_opt3_validation_code"`
	EmailValid              bool    `json:"email_valid"`
	EmailValidationCode     int     `json:"email_validation_code"`
	EmergencyContact        string  `json:"emergency_contact"`
	EmergencyContact2       string  `json:"emergency_contact2"`
	EmergencyPhone          string  `json:"emergency_phone"`
	EmergencyPhone2         string  `json:"emergency_phone2"`
	EncryptedCredentials    struct {
		EncryptedPassword string `json:"encryptedPassword"`
	} `json:"encryptedCredentials"`
	EncryptedPassword           string `json:"encrypted_password"`
	FinancialRequirements       int    `json:"financial_requirements"`
	FirstName                   string `json:"first_name"`
	Guard1Firstname             string `json:"guard1_firstname"`
	Guard1Lastname              string `json:"guard1_lastname"`
	Guard1PhoneH                string `json:"guard1_phone_h"`
	Guard1PhoneM                string `json:"guard1_phone_m"`
	Guard1PhoneW                string `json:"guard1_phone_w"`
	Guard2Firstname             string `json:"guard2_firstname"`
	Guard2Lastname              string `json:"guard2_lastname"`
	Guard2PhoneH                string `json:"guard2_phone_h"`
	Guard2PhoneM                string `json:"guard2_phone_m"`
	Guard2PhoneW                string `json:"guard2_phone_w"`
	HasValidEmail               bool   `json:"hasValidEmail"`
	HasValidSms                 bool   `json:"hasValidSms"`
	Id                          int    `json:"id"`
	LastError                   string `json:"lastError"`
	LastName                    string `json:"last_name"`
	LessonAccountStatus         int    `json:"lesson_account_status"`
	LessonOnly                  int    `json:"lesson_only"`
	LessonsAccountVaultTokenId  int    `json:"lessons_account_vault_token_id"`
	LscClubId                   int    `json:"lsc_club_id"`
	MedCarrier                  string `json:"med_carrier"`
	MedCarrierPhone             string `json:"med_carrier_phone"`
	MemberCount                 int    `json:"memberCount"`
	Mi                          string `json:"mi"`
	OdAndroid                   bool   `json:"od_android"`
	OdIos                       bool   `json:"od_ios"`
	Oid                         int    `json:"oid"`
	OndeckInstalled             bool   `json:"ondeckInstalled"`
	OndemandAccountVaultTokenId int    `json:"ondemand_account_vault_token_id"`
	OptionBits                  int    `json:"option_bits"`
	Password                    string `json:"password"`
	PasswordResetToken          string `json:"password_reset_token"`
	PayMethod                   string `json:"pay_method"`
	PhoneH                      string `json:"phone_h"`
	PhoneW                      string `json:"phone_w"`
	PushNotificationEnabled     bool   `json:"pushNotificationEnabled"`
	SeUuid                      string `json:"se_uuid"`
	ShowDbDetails               bool   `json:"showDbDetails"`
	ShowSearchCount             bool   `json:"showSearchCount"`
	Sms1                        string `json:"sms1"`
	Sms1Carrier                 int    `json:"sms1_carrier"`
	Sms1Valid                   bool   `json:"sms1_valid"`
	Sms1ValidationCode          int    `json:"sms1_validation_code"`
	Sms2                        string `json:"sms2"`
	Sms2Carrier                 int    `json:"sms2_carrier"`
	Sms2Valid                   bool   `json:"sms2_valid"`
	Sms2ValidationCode          int    `json:"sms2_validation_code"`
	SpnosorCategoryId           int    `json:"spnosor_category_id"`
	SponsorFullName             string `json:"sponsor_full_name"`
	SponsorPhoneM               string `json:"sponsor_phone_m"`
	SponsorPhoneO               string `json:"sponsor_phone_o"`
	State                       string `json:"state"`
	TableName                   string `json:"tableName"`
	TeamId                      int    `json:"team_id"`
	Title                       string `json:"title"`
	VerifiedEmail               string `json:"verified_email"`
	YmcaAccountVaultTokenId     int    `json:"ymca_account_vault_token_id"`
	Zip                         string `json:"zip"`
}

type ListAccountsResponse struct {
	Count  int       `json:"count"`
	Result []Account `json:"result"`
}

type FullAccount struct {
	Id                        int     `json:"id"`
	Email                     string  `json:"email"`
	EmailValid                bool    `json:"emailValid"`
	EmailOpt1                 string  `json:"emailOpt1"`
	EmailOpt1Valid            bool    `json:"emailOpt1Valid"`
	EmailOpt2                 string  `json:"emailOpt2"`
	EmailOpt2Valid            bool    `json:"emailOpt2Valid"`
	EmailOpt3                 string  `json:"emailOpt3"`
	EmailOpt3Valid            bool    `json:"emailOpt3Valid"`
	Sms1                      string  `json:"sms1"`
	Sms1Carrier               int     `json:"sms1Carrier"`
	Sms1VerifyCode            int     `json:"sms1VerifyCode"`
	Sms1Valid                 bool    `json:"sms1Valid"`
	Sms2                      string  `json:"sms2"`
	Sms2Carrier               int     `json:"sms2Carrier"`
	Sms2VerifyCode            int     `json:"sms2VerifyCode"`
	Sms2Valid                 bool    `json:"sms2Valid"`
	HomePhone                 string  `json:"homePhone"`
	WorkPhone                 string  `json:"workPhone"`
	FirstName                 string  `json:"firstName"`
	Mi                        string  `json:"mi"`
	LastName                  string  `json:"lastName"`
	Address                   string  `json:"address"`
	Address2                  string  `json:"address2"`
	City                      string  `json:"city"`
	State                     string  `json:"state"`
	Zip                       string  `json:"zip"`
	Title                     string  `json:"title"`
	Balance                   float64 `json:"balance"`
	Notes                     string  `json:"notes"`
	Deleted                   int     `json:"deleted"`
	AccountStatusId           int     `json:"accountStatusId"`
	AdminType                 int     `json:"adminType"`
	BouncedId                 int     `json:"bouncedId"`
	ByAccountId               int     `json:"byAccountId"`
	CustomField               string  `json:"customField"`
	DroppedDate               string  `json:"droppedDate"`
	DroppedDate1              string  `json:"dropped_date"`
	JoinedDate                string  `json:"joinedDate"`
	JoinedDate1               string  `json:"joined_date"`
	LastSignedOnDate          string  `json:"lastSignedOnDate"`
	LastSignedonDate          string  `json:"last_signedon_date"`
	LatestRegistrationFeeDate string  `json:"latestRegistrationFeeDate"`
	LatestRegistrationfeeDate string  `json:"latest_registrationfee_date"`
	ModifiedDate              string  `json:"modifiedDate"`
	ModifiedDate1             string  `json:"modified_date"`
	SentPasswordDate          string  `json:"sentPasswordDate"`
	SentPasswordDate1         string  `json:"sent_password_date"`
	EmergencyContact          string  `json:"emergencyContact"`
	EmergencyPhone            string  `json:"emergencyPhone"`
	EmergencyContact2         string  `json:"emergencyContact2"`
	EmergencyPhone2           string  `json:"emergencyPhone2"`
	GearAdmin                 int     `json:"gearAdmin"`
	Guard1FirstName           string  `json:"guard1FirstName"`
	Guard1LastName            string  `json:"guard1LastName"`
	Guard1HomePhone           string  `json:"guard1HomePhone"`
	Guard1WorkPhone           string  `json:"guard1WorkPhone"`
	Guard1MedicalPhone        string  `json:"guard1MedicalPhone"`
	Guard2FirstName           string  `json:"guard2FirstName"`
	Guard2LastName            string  `json:"guard2LastName"`
	Guard2HomePhone           string  `json:"guard2HomePhone"`
	Guard2WorkPhone           string  `json:"guard2WorkPhone"`
	Guard2MedicalPhone        string  `json:"guard2MedicalPhone"`
	LessonAdmin               int     `json:"lessonAdmin"`
	LessonOnly                int     `json:"lessonOnly"`
	LscClubId                 int     `json:"lscClubId"`
	MedicalCarrier            string  `json:"medicalCarrier"`
	MedicalCarrierPhone       string  `json:"medicalCarrierPhone"`
	MemberSearch              int     `json:"memberSearch"`
	PasswordResetExpiration   string  `json:"passwordResetExpiration"`
	PayMethod                 string  `json:"payMethod"`
	PayMethodLastUpdated      string  `json:"payMethodLastUpdated"`
	PaymethodLastupdatedDate  string  `json:"paymethod_lastupdated_date"`
	TeamId                    int     `json:"teamId"`
	TeamName                  string  `json:"teamName"`
	TouchpadAdmin             int     `json:"touchpadAdmin"`
	FinancialAdmin            int     `json:"financialAdmin"`
	TumAdmin                  int     `json:"tumAdmin"`
	LessonAccountStatus       int     `json:"lessonAccountStatus"`
	Metadata                  struct {
		Invisible struct {
			Items []interface{} `json:"items"`
		} `json:"invisible"`
		Readonly struct {
			Items []interface{} `json:"items"`
		} `json:"readonly"`
	} `json:"metadata"`
	Errors struct {
		Items []interface{} `json:"items"`
	} `json:"errors"`
	FinancialRequirements []struct {
		Name    string `json:"name"`
		Id      int    `json:"id"`
		Checked bool   `json:"checked"`
	} `json:"financialRequirements"`
	PictureFile            string   `json:"pictureFile"`
	CoachType              int      `json:"coachType"`
	TeamAlias              string   `json:"teamAlias"`
	MobileLastSignedOn     string   `json:"mobileLastSignedOn"`
	MobileLastsignedonDate string   `json:"mobile_lastsignedon_date"`
	WorkforceAdmin         int      `json:"workforceAdmin"`
	Pin                    string   `json:"pin"`
	HasOndeck              bool     `json:"hasOndeck"`
	DisplayFields          []string `json:"displayFields"`
	DisplayGroupFields     []int    `json:"displayGroupFields"`
	SeUuid                 string   `json:"se_uuid"`
	SeFirstName            string   `json:"se_first_name"`
	SeLastName             string   `json:"se_last_name"`
}

type Member struct {
	AccountId        int    `json:"account_id"`
	Bdonation        bool   `json:"bdonation"`
	BfinaOther       bool   `json:"bfina_other"`
	BfinaRepresented bool   `json:"bfina_represented"`
	Bio              string `json:"bio"`
	Bnewsletter      bool   `json:"bnewsletter"`
	ByAccountId      int    `json:"by_account_id"`
	Citizen          string `json:"citizen"`
	CustFld          string `json:"cust_fld"`
	Deleted          int    `json:"deleted"`
	Disabilities     []struct {
		Name    string `json:"name"`
		Id      int    `json:"id"`
		Checked bool   `json:"checked"`
	} `json:"disabilities"`
	Disability         int    `json:"disability"`
	DtAttach           string `json:"dt_attach"`
	AttachDate         string `json:"attach_date"`
	DtDivingCert       string `json:"dt_diving_cert"`
	DivingCertDate     string `json:"diving_cert_date"`
	DtDob              string `json:"dt_dob"`
	Dob                string `json:"dob"`
	DtDropped          string `json:"dt_dropped"`
	DroppedDate        string `json:"dropped_date"`
	DtInactive         string `json:"dt_inactive"`
	InactiveDate       string `json:"inactive_date"`
	DtJoined           string `json:"dt_joined"`
	JoinedDate         string `json:"joined_date"`
	DtLastRegGen       string `json:"dt_last_reg_gen"`
	LastRegGenDate     string `json:"last_reg_gen_date"`
	DtLastRegistered   string `json:"dt_last_registered"`
	LastRegisteredDate string `json:"last_registered_date"`
	DtLregfee          string `json:"dt_lregfee"`
	LregfeeDate        string `json:"lregfee_date"`
	DtModified         string `json:"dt_modified"`
	ModifiedDate       string `json:"modified_date"`
	Email              string `json:"email"`
	EmailDel           string `json:"email_del"`
	EmailValid         bool   `json:"email_valid"`
	Ethnic             int    `json:"ethnic"`
	Ethnicity          []struct {
		Name    string `json:"name"`
		Id      int    `json:"id"`
		Checked bool   `json:"checked"`
	} `json:"ethnicity"`
	FirstName             string      `json:"first_name"`
	HideFromSearch        bool        `json:"hide_from_search"`
	HsGraduation          interface{} `json:"hs_graduation"`
	Id                    int         `json:"id"`
	IsCertified           int         `json:"is_certified"`
	IsCoach               int         `json:"is_coach"`
	IsUnsanctionedAthlete bool        `json:"is_unsanctioned_athlete"`
	IsVolunteer           int         `json:"is_volunteer"`
	LastName              string      `json:"last_name"`
	LessonOnly            int         `json:"lesson_only"`
	LocationId            int         `json:"location_id"`
	MedDrPhone            string      `json:"med_dr_phone"`
	MedDrName             string      `json:"med_dr_name"`
	MedNotes              string      `json:"med_notes"`
	Medication            string      `json:"medication"`
	Metadata              struct {
		Invisible struct {
			Items []struct {
				Name  string `json:"name"`
				Value bool   `json:"value"`
			} `json:"items"`
		} `json:"invisible"`
		Readonly struct {
			Items []struct {
				Name  string `json:"name"`
				Value bool   `json:"value"`
			} `json:"items"`
		} `json:"readonly"`
	} `json:"metadata"`
	MemberStatusId              int    `json:"member_status_id"`
	Mi                          string `json:"mi"`
	ObsoleteUsmsNo              string `json:"obsolete_usms_no"`
	Options                     int    `json:"options"`
	Phone                       string `json:"phone"`
	PictureFile                 string `json:"picture_file"`
	Prefer                      string `json:"prefer"`
	RegStatus                   int    `json:"reg_status"`
	RosterGroupId               int    `json:"roster_group_id"`
	Season                      int    `json:"season"`
	Sex                         int    `json:"sex"`
	ShirtSize                   string `json:"shirt_size"`
	Sms                         string `json:"sms"`
	SmsCarrier                  int    `json:"sms_carrier"`
	SmsValid                    bool   `json:"sms_valid"`
	SmsVerifyCode               int    `json:"sms_verify_code"`
	SwimmerIdentity             string `json:"swimmer_identity"`
	SwimmerIdentityWasGenerated bool   `json:"swimmer_identity_was_generated"`
	SwimsuitSize                string `json:"swimsuit_size"`
	TeamGroupId                 int    `json:"team_group_id"`
	TeamId                      int    `json:"team_id"`
	TeamSubGroupId              int    `json:"team_sub_group_id"`
	WarmupJacketSize            string `json:"warmup_jacket_size"`
	WarmupPantSize              string `json:"warmup_pant_size"`
	Errors                      struct {
		Items []interface{} `json:"items"`
	} `json:"errors"`
	AsaRegistration struct {
		Id                 int         `json:"id"`
		DtCreated          interface{} `json:"dt_created"`
		DtModified         interface{} `json:"dt_modified"`
		AsaAmateur         bool        `json:"asa_amateur"`
		AsaLevel           int         `json:"asa_level"`
		AsaNo              interface{} `json:"asa_no"`
		RepresentedCountry int         `json:"represented_country"`
		DisciplineBits     interface{} `json:"discipline_bits"`
		DisabilityBits     interface{} `json:"disability_bits"`
		OpportunityBits    interface{} `json:"opportunity_bits"`
		AsaHideDetails     bool        `json:"asa_hide_details"`
		AsaEthnicity       interface{} `json:"asa_ethnicity"`
		AsaEthnicities     interface{} `json:"asa_ethnicities"`
	} `json:"asa_registration"`
	AsaLevels                 []interface{} `json:"asa_levels"`
	AsaCountries              []interface{} `json:"asa_countries"`
	Disciplines               []interface{} `json:"disciplines"`
	Opportunities             []interface{} `json:"opportunities"`
	AsaLevel                  int           `json:"asa_level"`
	RepresentedCountry        int           `json:"represented_country"`
	DisciplineBits            int           `json:"discipline_bits"`
	OpportunityBits           int           `json:"opportunity_bits"`
	AsaHideDetails            bool          `json:"asa_hide_details"`
	DualClubCode              string        `json:"dual_club_code"`
	DtDualJoined              string        `json:"dt_dual_joined"`
	DualJoinedDate            interface{}   `json:"dual_joined_date"`
	RankDual                  bool          `json:"rank_dual"`
	DualRec                   bool          `json:"dual_rec"`
	AsaAmateur                bool          `json:"asa_amateur"`
	GenerateSwimmerId         interface{}   `json:"generateSwimmerId"`
	IgnorePictureFile         bool          `json:"ignorePictureFile"`
	AccountFirstname          string        `json:"account_firstname"`
	AccountLastname           string        `json:"account_lastname"`
	HasValidAccountEmail      bool          `json:"hasValidAccountEmail"`
	HasValidAccountSMS        bool          `json:"hasValidAccountSMS"`
	AccountStatus             int           `json:"account_status"`
	DtDbsCheck                string        `json:"dt_dbs_check"`
	DbsCheckDate              interface{}   `json:"dbs_check_date"`
	Note                      string        `json:"note"`
	AccountStatus1            int           `json:"accountStatus"`
	DtBackStartCert           string        `json:"dt_back_start_cert"`
	BackStartCertDate         string        `json:"back_start_cert_date"`
	DtBackStartNoLedgeCert    string        `json:"dt_back_start_no_ledge_cert"`
	BackStartNoLedgeCertDate  string        `json:"back_start_no_ledge_cert_date"`
	ForwardStartCompletedStep int           `json:"forward_start_completed_step"`
	DisplayNcsaBanner         bool          `json:"display_ncsa_banner"`
	IsWorkforceMember         bool          `json:"isWorkforceMember"`
	DisplayFields             []string      `json:"displayFields"`
	DisplayGroupFields        []int         `json:"displayGroupFields"`
	ModifiedTimestamp         int64         `json:"modifiedTimestamp"`
}

func ConvertPhoneNumbers(phoneNumberString string) (string, error) {
	if phoneNumberString == "" {
		return "", nil
	}

	phoneNumber, err := phonenumbers.Parse(phoneNumberString, "US")
	if err != nil {
		return "", err
	}

	return phonenumbers.Format(phoneNumber, phonenumbers.E164), nil
}

func ConvertTeamUnifyDate(date string) (time.Time, error) {
	subMatches := dateRegex.FindStringSubmatch(date)
	if subMatches == nil {
		return time.Time{}, errors.InternalServerError("users.utils.ConvertTeamUnifyDate", "error converting %s into timestamp", date)
	}

	timestamp, err := strconv.ParseInt(subMatches[1], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.UnixMilli(timestamp), nil
}

func GetCookies(ctx context.Context, service cookiestealer.CookieStealerService) (string, error) {
	cookieRes, err := service.StealTeamUnifyCookies(ctx, &cookiestealer.StealTeamUnifyCookiesRequest{})
	if err != nil {
		return "", err
	}
	if cookieRes.Unready {
		return "", errors.InternalServerError("users.utils.GetCookies", "Unable to access teamunify at the moment")
	}
	return cookieRes.Cookies, nil
}

func sendTeamUnifyRequest(r *http.Request, cookies string) (*http.Response, error) {
	r.Header.Set("Cookie", cookies)

	client := http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func objToReader(obj interface{}) (*bytes.Reader, error) {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(objBytes), nil
}

func ListAccounts(ctx context.Context, cookies string) (*ListAccountsResponse, error) {
	url := "https://www.teamunify.com/api/dataViewService/list"

	body := listRequestBody{
		SpecId: "ama-accounts",
		Filters: struct {
			AccountStatusId []string `json:"account_status_id"`
			Deleted         string   `json:"deleted"`
		}{
			AccountStatusId: []string{"20"},
			Deleted:         "0",
		},
		Start: 0,
		Count: 130,
		Orders: []struct {
			Name     string `json:"name"`
			Asc      bool   `json:"asc"`
			NullSide bool   `json:"nullSide"`
		}{
			{
				Name:     "email",
				Asc:      true,
				NullSide: false,
			},
		},
		ExcludedFieldNames: []string{
			"zip",
			"pay_method",
			"lessons_account_vault_token_id",
			"ondemand_account_vault_token_id",
			"dt_pay_method_modified",
			"memberCount",
			"deleted",
			"city",
			"balance",
			"cust_fld",
			"guard1",
			"guard2",
			"guard1_phone_w",
			"guard1_phone_h",
			"guard1_phone_m",
			"guard2_phone_w",
			"guard2_phone_h",
			"guard2_phone_m",
			"med_carrier",
			"med_carrier_phone",
			"sms1",
			"sms2",
			"sms_valid",
			"state",
			"touchpad_admin",
			"emergency_contact",
			"emergency_phone",
			"emergency_contact2",
			"emergency_phone2",
			"address",
			"ondeckLastIn",
			"ondeckInstalled",
			"se_uuid",
			"tum_admin",
			"email_valid",
			"coachType",
			"address2",
			"accountPin",
			"dt_lregfee",
		},
	}

	reader, err := objToReader(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, reader)
	if err != nil {
		return nil, err
	}

	response, err := sendTeamUnifyRequest(request, cookies)
	if err != nil {
		return nil, err
	}

	var res ListAccountsResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetAccount(ctx context.Context, cookies string, id int) (*FullAccount, error) {
	url := "https://www.teamunify.com/api/amaService/getAccountById"

	bodyObj := struct {
		Id int `json:"id"`
	}{
		Id: id,
	}
	body, err := objToReader(bodyObj)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	response, err := sendTeamUnifyRequest(request, cookies)
	if err != nil {
		return nil, err
	}

	var account FullAccount
	err = json.NewDecoder(response.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetMembers(ctx context.Context, cookies string, id int) ([]Member, error) {
	url := "https://www.teamunify.com/api/amaService/getMembersByAccountId"

	bodyObj := struct {
		Id int `json:"accountId"`
	}{
		Id: id,
	}
	body, err := objToReader(bodyObj)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	response, err := sendTeamUnifyRequest(request, cookies)
	if err != nil {
		return nil, err
	}

	var members []Member
	err = json.NewDecoder(response.Body).Decode(&members)
	if err != nil {
		return nil, err
	}

	return members, nil
}
