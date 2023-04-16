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
