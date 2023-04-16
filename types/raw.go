package types

import (
	"strconv"

	"github.com/cyb3rko/signal-back-2/signal"
)

// StatementToStringArray formats a SqlStatement fairly literally as an array.
// Null parameters are left empty.
func StatementToStringArray(sql *signal.SqlStatement) []string {
	s := make([]string, len(sql.GetParameters()))
	for i, p := range sql.GetParameters() {
		if p.IntegerParameter != nil {
			s[i] = strconv.Itoa(int(*p.IntegerParameter))
		} else if p.StringParamter != nil {
			s[i] = *p.StringParamter
		}
	}
	return s
}

// CSV column headers.
var (
	MessageCSVHeaders = []string{
		"_ID",
		"DATE_SENT",
		"DATE_RECEIVED",
		"DATE_SERVER",
		"THREAD_ID",
		"RECIPIENT_ID",
		"RECIPIENT_DEVICE_ID",
		"TYPE",
		"BODY",
		"READ",
		"CT_L",
		"EXP",
		"M_TYPE",
		"M_SIZE",
		"ST",
		"TR_ID",
		"SUBSCRIPTION_ID",
		"RECEIPT_TIMESTAMP",
		"DELIVERY_RECEIPT_COUNT",
		"READ_RECEIPT_COUNT",
		"VIEWED_RECEIPT_COUNT",
		"MISMATCHED_IDENTITIES",
		"NETWORK_FAILURES",
		"EXPIRES_IN",
		"EXPIRE_STARTED",
		"NOTIFIED",
		"QUOTE_ID",
		"QUOTE_AUTHOR",
		"QUOTE_BODY",
		"QUOTE_MISSING",
		"QUOTE_MENTIONS",
		"QUOTE_TYPE",
		"SHARED_CONTACTS",
		"UNIDENTIFIED",
		"LINK_PREVIEWS",
		"VIEW_ONCE",
		"REACTIONS_UNREAD",
		"REACTIONS_LAST_SEEN",
		"REMOTE_DELETED",
		"MENTIONS_SELF",
		"NOTIFIED_TIMESTAMP",
		"SERVER_GUID",
		"MESSAGE_RANGES",
		"STORY_TYPE",
		"PARENT_STORY_ID",
		"EXPORT_STATE",
		"EXPORTED",
		"SCHEDULED_DATE",
	}

	RecipientCSVHeaders = []string{
		"_ID",
		"UUID",
		"USERNAME",
		"PHONE",
		"EMAIL",
		"GROUP_ID",
		"GROUP_TYPE",
		"BLOCKED",
		"MESSAGE_RINGTONE",
		"MESSAGE_VIBRATE",
		"CALL_RINGTONE",
		"CALL_VIBRATE",
		"NOTIFICATION_CHANNEL",
		"MUTE_UNTIL",
		"COLOR",
		"SEEN_INVITE_REMINDER",
		"DEFAULT_SUBSCRIPTION_ID",
		"MESSAGE_EXPIRATION_TIME",
		"REGISTERED",
		"SYSTEM_DISPLAY_NAME",
		"SYSTEM_PHOTO_URI",
		"SYSTEM_PHONE_LABEL",
		"SYSTEM_PHONE_TYPE",
		"SYSTEM_CONTACT_URI",
		"SYSTEM_INFO_PENDING",
		"PROFILE_KEY",
		"PROFILE_KEY_CREDENTIAL",
		"SIGNAL_PROFILE_NAME",
		"PROFILE_FAMILY_NAME",
		"PROFILE_JOINED_NAME",
		"SIGNAL_PROFILE_AVATAR",
		"PROFILE_SHARING",
		"LAST_PROFILE_FETCH",
		"UNIDENTIFIED_ACCESS_MODE",
		"FORCE_SMS_DELETION",
		"STORAGE_SERVICE_KEY",
		"DIRTY",
		"MENTION_SETTING",
		"STORAGE_PROTO",
		"CAPABILITIES",
		"LAST_GV1_MIGRATE_REMINDER",
		"LAST_SESSION_RESET",
		"WALLPAPER",
		"WALLPAPER_FILE",
		"ABOUT",
		"ABOUT_EMOJI",
		"SYSTEM_FAMILY_NAME",
		"SYSTEM_GIVEN_NAME",
		"EXTRAS",
		"GROUPS_IN_COMMON",
		"CHAT_COLORS",
		"CUSTOM_CHAT_COLORS",
		"BADGES",
		"PNI",
		"DISTRIBUTION_LIST_ID",
		"NEEDS_PNI_SIGNATURE",
		"UNREGISTERED_TIMESTAMP",
		"HIDDEN",
		"REPORTING_TOKEN",
		"SYSTEM_NICKNAME",
	}

	ThreadCSVHeaders = []string{
		"_ID",
		"DATE",
		"MEANINGFUL_MESSAGES",
		"RECIPIENT_ID",
		"READ",
		"TYPE",
		"ERROR",
		"SNIPPET",
		"SNIPPET_TYPE",
		"SNIPPET_URI",
		"SNIPPET_CONTENT_TYPE",
		"SNIPPET_EXTRAS",
		"UNREAD_COUNT",
		"ARCHIVED",
		"STATUS",
		"DELIVERY_RECEIPT_COUNT",
		"READ_RECEIPT_COUNT",
		"EXPIRES_IN",
		"LAST_SEEN",
		"HAS_SENT",
		"LAST_SCROLLED",
		"PINNED",
		"UNREAD_SELF_MENTION_COUNT",
	}

	CallCSVHeaders = []string{
		"_ID",
		"CALL_ID",
		"MESSAGE_ID",
		"PEER",
		"TYPE",
		"DIRECTION",
		"EVENT",
	}

	GroupsCSVHeaders = []string{
		"_ID",
		"GROUP_ID",
		"RECIPIENT_ID",
		"TITLE",
		"AVATAR_ID",
		"AVATAR_KEY",
		"AVATAR_CONTENT_TYPE",
		"AVATAR_RELAY",
		"TIMESTAMP",
		"ACTIVE",
		"AVATAR_DIGEST",
		"MMS",
		"MASTER_KEY",
		"REVISION",
		"DECRYPTED_GROUP",
		"EXPECTED_V2_ID",
		"FORMER_V1_MEMBERS",
		"DISTRIBUTION_ID",
		"DISPLAY_AS_STORY",
		"AUTH_SERVICE_ID",
		"LAST_FORCE_UPDATE_TIMESTAMP",
	}

	GroupMembershipCSVHeaders = []string{
		"_ID",
		"GROUP_ID",
		"RECIPIENT_ID",
		"UNIQUE",
	}

	GroupReceiptsCSVHeaders = []string{
		"_ID",
		"MMS_ID",
		"ADDRESS",
		"STATUS",
		"TIMESTAMP",
		"UNIDENTIFIED",
	}
)

// SQLSMS info
//
// https://github.com/signalapp/Signal-Android/blob/master/src/org/thoughtcrime/securesms/database/SmsDatabase.java#L77
type SQLSMS struct {
	ID                   uint64
	ThreadID             *uint64
	Address              *string
	AddressDeviceID      uint64 // default 1
	Person               *uint64
	DateReceived         *uint64
	DateSent             *uint64
	Protocol             uint64 // effectively default 0
	Read                 uint64 // default 0
	Status               uint64 // default -1
	Type                 *uint64
	ReplyPathPresent     *uint64
	DeliveryReceiptCount uint64 // default 0
	Subject              *string
	Body                 *string
	MismatchedIdentities *string // default null
	ServiceCenter        *string
	SubscriptionID       uint64 // default -1
	ExpiresIn            uint64 // default 0
	ExpireStarted        uint64 // default 0
	Notified             uint64 // default 0
	ReadReceiptCount     uint64 // default 0
	Unidentified         uint64 // default 0
}

// StatementToSMS converts a of SQL statement to a single SMS.
func StatementToSMS(sql *signal.SqlStatement) *SQLSMS {
	return ParametersToSMS(sql.GetParameters())
}

// ParametersToSMS converts a set of SQL parameters to a single SMS.
func ParametersToSMS(ps []*signal.SqlStatement_SqlParameter) *SQLSMS {
	if len(ps) < 22 {
		return nil
	}

	result := &SQLSMS{
		ID:                   ps[0].GetIntegerParameter(),
		ThreadID:             ps[1].IntegerParameter,
		Address:              ps[2].StringParamter,
		AddressDeviceID:      ps[3].GetIntegerParameter(),
		Person:               ps[4].IntegerParameter,
		DateReceived:         ps[5].IntegerParameter,
		DateSent:             ps[6].IntegerParameter,
		Protocol:             ps[7].GetIntegerParameter(),
		Read:                 ps[8].GetIntegerParameter(),
		Status:               ps[9].GetIntegerParameter(),
		Type:                 ps[10].IntegerParameter,
		ReplyPathPresent:     ps[11].IntegerParameter,
		DeliveryReceiptCount: ps[12].GetIntegerParameter(),
		Subject:              ps[13].StringParamter,
		Body:                 ps[14].StringParamter,
		MismatchedIdentities: ps[15].StringParamter,
		ServiceCenter:        ps[16].StringParamter,
		SubscriptionID:       ps[17].GetIntegerParameter(),
		ExpiresIn:            ps[18].GetIntegerParameter(),
		ExpireStarted:        ps[19].GetIntegerParameter(),
		Notified:             ps[20].GetIntegerParameter(),
		ReadReceiptCount:     ps[21].GetIntegerParameter(),
		//Unidentified:         ps[22].GetIntegerParameter(),
	}

	return result
}

// SQLMMS info
//
// https://github.com/signalapp/Signal-Android/blob/master/src/org/thoughtcrime/securesms/database/MmsDatabase.java#L110
type SQLMMS struct {
	ID                   uint64
	ThreadID             *uint64
	DateSent             *uint64
	DateReceived         *uint64
	MessageBox           *uint64
	Read                 uint64 // default 0
	MID                  *string
	Sub                  *string
	SubCs                *uint64
	Body                 *string
	PartCount            *uint64
	CtT                  *string
	ContentLocation      *string
	Address              *string
	AddressDeviceID      *uint64
	Expiry               *uint64
	MCls                 *string
	MessageType          *uint64
	V                    *uint64
	MessageSize          *uint64
	Pri                  *uint64
	Rr                   *uint64
	RptA                 *uint64
	RespSt               *uint64
	Status               *uint64
	TransactionID        *string
	RetrSt               *uint64
	RetrTxt              *string
	RetrTxtCs            *uint64
	ReadStatus           *uint64
	CtCls                *uint64
	RespTxt              *string
	DTm                  *uint64
	DeliveryReceiptCount uint64  // default 0
	MismatchedIdentities *string // default null
	NetworkFailure       *string // default null
	DRpt                 *uint64
	SubscriptionID       uint64 // default -1
	ExpiresIn            uint64 // default 0
	ExpireStarted        uint64 // default 0
	Notified             uint64 // default 0
	ReadReceiptCount     uint64 // default 0
	QuoteID              uint64 // default 0
	QuoteAuthor          *string
	QuoteBody            *string
	QuoteAttachment      uint64 // default -1
	QuoteMissing         uint64 // default 0
	SharedContacts       *string
	Unidentified         uint64 // default 0
}

// StatementToMMS converts a of SQL statement to a single MMS.
func StatementToMMS(sql *signal.SqlStatement) *SQLMMS {
	return ParametersToMMS(sql.GetParameters())
}

// ParametersToMMS converts a set of SQL parameters to a single MMS.
func ParametersToMMS(ps []*signal.SqlStatement_SqlParameter) *SQLMMS {
	if len(ps) < 42 {
		return nil
	}

	result := &SQLMMS{
		ID:                   ps[0].GetIntegerParameter(),
		ThreadID:             ps[1].IntegerParameter,
		DateSent:             ps[2].IntegerParameter,
		DateReceived:         ps[3].IntegerParameter,
		MessageBox:           ps[4].IntegerParameter,
		Read:                 ps[5].GetIntegerParameter(),
		MID:                  ps[6].StringParamter,
		Sub:                  ps[7].StringParamter,
		SubCs:                ps[8].IntegerParameter,
		Body:                 ps[9].StringParamter,
		PartCount:            ps[10].IntegerParameter,
		CtT:                  ps[11].StringParamter,
		ContentLocation:      ps[12].StringParamter,
		Address:              ps[13].StringParamter,
		AddressDeviceID:      ps[14].IntegerParameter,
		Expiry:               ps[15].IntegerParameter,
		MCls:                 ps[16].StringParamter,
		MessageType:          ps[17].IntegerParameter,
		V:                    ps[18].IntegerParameter,
		MessageSize:          ps[19].IntegerParameter,
		Pri:                  ps[20].IntegerParameter,
		Rr:                   ps[21].IntegerParameter,
		RptA:                 ps[22].IntegerParameter,
		RespSt:               ps[23].IntegerParameter,
		Status:               ps[24].IntegerParameter,
		TransactionID:        ps[25].StringParamter,
		RetrSt:               ps[26].IntegerParameter,
		RetrTxt:              ps[27].StringParamter,
		RetrTxtCs:            ps[28].IntegerParameter,
		ReadStatus:           ps[29].IntegerParameter,
		CtCls:                ps[30].IntegerParameter,
		RespTxt:              ps[31].StringParamter,
		DTm:                  ps[32].IntegerParameter,
		DeliveryReceiptCount: ps[33].GetIntegerParameter(),
		MismatchedIdentities: ps[34].StringParamter,
		NetworkFailure:       ps[35].StringParamter,
		DRpt:                 ps[36].IntegerParameter,
		SubscriptionID:       ps[37].GetIntegerParameter(),
		ExpiresIn:            ps[38].GetIntegerParameter(),
		ExpireStarted:        ps[39].GetIntegerParameter(),
		Notified:             ps[40].GetIntegerParameter(),
		ReadReceiptCount:     ps[41].GetIntegerParameter(),
		//QuoteID:              ps[42].GetIntegerParameter(),
		//QuoteAuthor:          ps[43].StringParamter,
		//QuoteBody:            ps[44].StringParamter,
		//QuoteAttachment:      ps[45].GetIntegerParameter(),
		//QuoteMissing:         ps[46].GetIntegerParameter(),
		//SharedContacts:       ps[47].StringParamter,
		//Unidentified:         ps[48].GetIntegerParameter(),
	}

	return result
}

// SQLPart info
//
// https://github.com/signalapp/Signal-Android/blob/master/src/org/thoughtcrime/securesms/database/AttachmentDatabase.java#L120
type SQLPart struct {
	RowID                uint64 // primary
	MmsID                *uint64
	Seq                  uint64 // default 0
	ContentType          *string
	Name                 *string
	Chset                *uint64
	ContentDisposition   *string
	Fn                   *string
	Cid                  *string
	ContentLocation      *string
	CttS                 *uint64
	CttT                 *string
	encrypted            *uint64
	TransferState        *uint64
	Data                 *string
	Size                 *uint64
	FileName             *string
	Thumbnail            *string
	ThumbnailAspectRatio *float64
	UniqueID             uint64 // not null
	Digest               []byte
	FastPreflightID      *string
	VoiceNote            uint64 // default 0
	DataRandom           []byte
	ThumbnailRandom      []byte
	Quote                uint64  // default 0
	Width                uint64  // default 0
	Height               uint64  // default 0
	Caption              *string //default null
}

// StatementToPart converts a of SQL statement to a single part.
func StatementToPart(sql *signal.SqlStatement) *SQLPart {
	return ParametersToPart(sql.GetParameters())
}

// ParametersToPart converts a set of SQL parameters to a single part.
func ParametersToPart(ps []*signal.SqlStatement_SqlParameter) *SQLPart {
	if len(ps) < 25 {
		return nil
	}
	return &SQLPart{
		RowID:                ps[0].GetIntegerParameter(),
		MmsID:                ps[1].IntegerParameter,
		Seq:                  ps[2].GetIntegerParameter(),
		ContentType:          ps[3].StringParamter,
		Name:                 ps[4].StringParamter,
		Chset:                ps[5].IntegerParameter,
		ContentDisposition:   ps[6].StringParamter,
		Fn:                   ps[7].StringParamter,
		Cid:                  ps[8].StringParamter,
		ContentLocation:      ps[9].StringParamter,
		CttS:                 ps[10].IntegerParameter,
		CttT:                 ps[11].StringParamter,
		encrypted:            ps[12].IntegerParameter,
		TransferState:        ps[13].IntegerParameter,
		Data:                 ps[14].StringParamter,
		Size:                 ps[15].IntegerParameter,
		FileName:             ps[16].StringParamter,
		Thumbnail:            ps[17].StringParamter,
		ThumbnailAspectRatio: ps[18].DoubleParameter,
		UniqueID:             ps[19].GetIntegerParameter(),
		Digest:               ps[20].GetBlobParameter(),
		FastPreflightID:      ps[21].StringParamter,
		VoiceNote:            ps[22].GetIntegerParameter(),
		DataRandom:           ps[23].GetBlobParameter(),
		ThumbnailRandom:      ps[24].GetBlobParameter(),
		Quote:                ps[25].GetIntegerParameter(),
		Width:                ps[26].GetIntegerParameter(),
		Height:               ps[27].GetIntegerParameter(),
		//Caption:              ps[28].StringParamter,
	}
}
