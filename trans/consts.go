package trans

// Translation of Telegram commands
const (
	COMMAND_START    = "COMMAND_START"
	COMMAND_MENU     = "COMMAND_MENU"
	COMMAND_GAVE     = "COMMAND_GAVE"
	COMMAND_GOT      = "COMMAND_GOT"
	COMMAND_RETURNED = "COMMAND_RETURNED"
	COMMAND_BALANCE  = "COMMAND_BALANCE"
	COMMAND_HISTORY  = "COMMAND_HISTORY"
	COMMAND_SETTINGS = "COMMAND_SETTINGS"
	COMMAND_HELP     = "COMMAND_HELP"
	COMMAND_CANCEL   = "COMMAND_CANCEL"
	COMMAND_CLEAR    = "COMMAND_CLEAR"
)

const (
	EMAIL_INVITE_SUBJ = "EMAIL_INVITE_SUBJ"
	EMAIL_INVITE_HTML = "EMAIL_INVITE_HTML"
	EMAIL_INVITE_TEXT = "EMAIL_INVITE_TEXT"

	SMS_INVITE_TEXT = "SMS_INVITE_TEXT"

	EMAIL_RECEIPT_SUBJ      = "EMAIL_RECEIPT_SUBJ"
	EMAIL_RECEIPT_BODY_TEXT = "EMAIL_RECEIPT_BODY_TEXT"
	EMAIL_RECEIPT_BODY_HTML = "EMAIL_RECEIPT_BODY_HTML"

	MESSENGER_RECEIPT_TEXT = "MESSENGER_RECEIPT_TEXT"

	MESSAGE_TEXT_INVITE_CREATED             = "MESSAGE_TEXT_INVITE_CREATED"
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED = "MESSAGE_TEXT_TRANSFER_CREATION_CANCELED"

	MESSAGE_TEXT_NOTHING_TO_CANCEL                 = "MESSAGE_TEXT_NOTHING_TO_CANCEL"
	MESSAGE_TEXT_SETTINGS                          = "MESSAGE_TEXT_SETTINGS"
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET               = "MESSAGE_TEXT_NOT_IMPLEMENTED_YET"
	MESSAGE_TEXT_HI                                = "MESSAGE_TEXT_HI"
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE = "MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE"
	MESSAGE_TEXT_WHATS_NEXT                        = "MESSAGE_TEXT_WHATS_NEXT"
	MESSAGE_TEXT_WHATS_NEXT_HINT                   = "MESSAGE_TEXT_WHATS_NEXT_HINT"
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE  = "MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE"

	MESSAGE_TEXT_BACK_TO_MAIN_MENU = "MESSAGE_TEXT_BACK_TO_MAIN_MENU"

	COMMAND_TEXT_DUE_RETURNS                     = "COMMAND_TEXT_DUE_RETURNS"
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER          = "MESSAGE_TEXT_OVERDUE_RETURNS_HEADER"
	MESSAGE_TEXT_DUE_RETURNS_HEADER              = "MESSAGE_TEXT_DUE_RETURNS_HEADER"
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER         = "MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER"
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY = "MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY"
	MESSAGE_TEXT_DUE_RETURNS_EMPTY               = "MESSAGE_TEXT_DUE_RETURNS_EMPTY"

	MESSAGE_TEXT_UNKNOWN_LANGUAGE               = "MESSAGE_TEXT_UNKNOWN_LANGUAGE"
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY           = "MESSAGE_TEXT_UNKNOWN_COUNTERPARTY"
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN = "MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN"
	MESSAGE_TEXT_UNKNOWN_DEBT                   = "MESSAGE_TEXT_UNKNOWN_DEBT"

	MESSAGE_TEXT_HISTORY_HEADER     = "MESSAGE_TEXT_HISTORY_HEADER"
	MESSAGE_TEXT_HISTORY_LIST       = "MESSAGE_TEXT_HISTORY_LIST"
	MESSAGE_TEXT_HISTORY_NO_RECORDS = "MESSAGE_TEXT_HISTORY_NO_RECORDS"
	MESSAGE_TEXT_WELCOME            = "MESSAGE_TEXT_WELCOME"

	MESSAGE_TEXT_INVITE_BY_EMAIL    = "MESSAGE_TEXT_INVITE_BY_EMAIL"
	MESSAGE_TEXT_INVITE_BY_SMS      = "MESSAGE_TEXT_INVITE_BY_SMS"
	MESSAGE_TEXT_INVITE_BY_TELEGRAM = "MESSAGE_TEXT_INVITE_BY_TELEGRAM"

	MESSAGE_TEXT_LETS_SEND_SMS           = "MESSAGE_TEXT_LETS_SEND_SMS"
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING = "MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING"
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING  = "MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING"

	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT = "MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT"

	MESSAGE_TEXT_INVALID_DATE            = "MESSAGE_TEXT_INVALID_DATE"
	MESSAGE_TEXT_INVALID_YEAR            = "MESSAGE_TEXT_INVALID_YEAR"
	MESSAGE_TEXT_INVALID_DAY             = "MESSAGE_TEXT_INVALID_DAY"
	MESSAGE_TEXT_INVALID_MONTH           = "MESSAGE_TEXT_INVALID_MONTH"
	MESSAGE_TEXT_INVALID_EMAIL           = "MESSAGE_TEXT_INVALID_EMAIL"
	MESSAGE_TEXT_INVALID_PHONE_NUMBER    = "MESSAGE_TEXT_INVALID_PHONE_NUMBER"
	MESSAGE_TEXT_NO_CONTACT_RECEIVED     = "MESSAGE_TEXT_NO_CONTACT_RECEIVED"
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER = "MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER"
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER  = "MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER"

	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE = "MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE"

	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME = "MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME"
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME   = "MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME"

	MESSAGE_TEXT_INVALID_FLOAT = "MESSAGE_TEXT_INVALID_FLOAT"

	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN = "MESSAGE_TEXT_OK_TRY_AGAIN"

	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES = "MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES"
	MESSAGE_TEXT_WRONG_EMAIL                   = "MESSAGE_TEXT_WRONG_EMAIL"
	MESSAGE_TEXT_WRONG_PHONE_NUMBER            = "MESSAGE_TEXT_WRONG_PHONE_NUMBER"

	MESSAGE_TEXT_ASK_INVITE_CHANNEL               = "MESSAGE_TEXT_ASK_INVITE_CHANNEL"
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL        = "MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL"
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER = "MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER"
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE         = "MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE"

	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED      = "MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED"
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED    = "MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED"
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM = "MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM"

	MESSAGE_TEXT_WRONG_INVITE_CODE         = "MESSAGE_TEXT_WRONG_INVITE_CODE"
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE        = "MESSAGE_TEXT_CHOOSE_UI_LANGUAGE"
	MESSAGE_TEXT_ABOUT_INVITES             = "MESSAGE_TEXT_ABOUT_INVITES"
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE = "MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE"

	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL = "MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL"
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS   = "MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS"

	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED = "MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED"

	MESSAGE_TEXT_LOCALE_CHANGED = "MESSAGE_TEXT_LOCALE_CHANGED"

	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY = "MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY"

	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY = "MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY"

	MESSAGE_TEXT_CHOOSE_CURRENCY = "MESSAGE_TEXT_CHOOSE_CURRENCY"

	MESSAGE_TEXT_ASK_LENDING_TYPE         = "MESSAGE_TEXT_ASK_LENDING_TYPE"
	MESSAGE_TEXT_ASK_LENDING_AMOUNT       = "MESSAGE_TEXT_ASK_LENDING_AMOUNT"
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY = "MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY"

	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT                    = "MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT"
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS = "MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS"
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM          = "MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM"

	MESSAGE_TEXT_ASK_BORROWING_TYPE         = "MESSAGE_TEXT_ASK_BORROWING_TYPE"
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT       = "MESSAGE_TEXT_ASK_BORROWING_AMOUNT"
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY = "MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY"

	MESSAGE_TEXT_YOUR_ABOUT_ADS            = "MESSAGE_TEXT_YOUR_ABOUT_ADS"
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE     = "MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE"
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME = "MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME"

	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER = "MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER"
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER   = "MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER"

	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER = "MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER"
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER   = "MESSAGE_TEXT_RECEIPT_RETURN_TO_USER"

	MESSAGE_TEXT_LOGIN_CODE = "MESSAGE_TEXT_LOGIN_CODE"

	MESSAGE_TEXT_ASK_FOR_FEEDBAK     = "MESSAGE_TEXT_ASK_FOR_FEEDBAK"
	COMMAND_TEXT_GIVE_FEEDBACK       = "COMMAND_TEXT_GIVE_FEEDBACK"
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT = "MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT"

	MESSAGE_TEXT_THANKS                                 = "MESSAGE_TEXT_THANKS"
	MESSAGE_TEXT_PLEASE_SEND_TEXT                       = "MESSAGE_TEXT_PLEASE_SEND_TEXT"
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT               = "MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT"
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER = "MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER"
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT                = "MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT"
	COMMAND_TEXT_FEEDBACK                               = "COMMAND_TEXT_FEEDBACK"
	COMMAND_TEXT_LEAVE_FEEDBACK                         = "COMMAND_TEXT_LEAVE_FEEDBACK"
	COMMAND_TEXT_ASK_FOR_FEEDBACK                       = "COMMAND_TEXT_ASK_FOR_FEEDBACK"
	COMMAND_TEXT_FEEDBACK_POSITIVE                      = "COMMAND_TEXT_FEEDBACK_POSITIVE"
	COMMAND_TEXT_FEEDBACK_NEUTRAL                       = "COMMAND_TEXT_FEEDBACK_NEUTRAL"
	COMMAND_TEXT_FEEDBACK_NEGATIVE                      = "COMMAND_TEXT_FEEDBACK_NEGATIVE"
	COMMAND_TEXT_FEEDBACK_NOT_READY                     = "COMMAND_TEXT_FEEDBACK_NOT_READY"

	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE = "MESSAGE_TEXT_ON_FEEDBACK_POSITIVE"
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL  = "MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL"
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE = "MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE"

	MESSAGE_TEXT_YOU_CAN_HELP_BY      = "MESSAGE_TEXT_YOU_CAN_HELP_BY"
	MESSAGE_TEXT_YOU_CAN_HELP_BY_HTML = "MESSAGE_TEXT_YOU_CAN_HELP_BY_HTML"
	MESSAGE_TEXT_ASK_TO_TRANSLATE     = "MESSAGE_TEXT_ASK_TO_TRANSLATE"

	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE = "MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE"

	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK = "COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK"

	MESSAGE_TEXT_TRANSFER_IS_CREATING = "MESSAGE_TEXT_TRANSFER_IS_CREATING"
	COMMAND_TEXT_PLEASE_WAIT          = "COMMAND_TEXT_PLEASE_WAIT"
	MESSAGE_TEXT_PLEASE_WAIT          = "MESSAGE_TEXT_PLEASE_WAIT"

	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT = "MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT"

	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU = "MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU"
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU = "MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU"

	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY = "MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY"
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY = "MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY"

	MESSAGE_TEXT_DUE_ON  = "MESSAGE_TEXT_DUE_ON"
	MESSAGE_TEXT_NOTE    = "MESSAGE_TEXT_NOTE"
	MESSAGE_TEXT_COMMENT = "MESSAGE_TEXT_COMMENT"

	MESSAGE_TEXT_LOGIN_TO_WEB_APP = "MESSAGE_TEXT_LOGIN_TO_WEB_APP"

	MESSAGE_TEXT_ASK_DUE            = "MESSAGE_TEXT_ASK_DUE"
	MESSAGE_TEXT_ASK_DUE_DATE       = "MESSAGE_TEXT_ASK_DUE_DATE"
	MESSAGE_TEXT_ASK_DATE_TO_REMIND = "MESSAGE_TEXT_ASK_DATE_TO_REMIND"
	MESSAGE_TEXT_WRONG_DATE         = "MESSAGE_TEXT_WRONG_DATE"

	MESSAGE_TEXT_FAILED_TO_DELETE_USER      = "MESSAGE_TEXT_FAILED_TO_DELETE_USER"
	MESSAGE_TEXT_USER_DELETED               = "MESSAGE_TEXT_USER_DELETED"
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY       = "MESSAGE_TEXT_ASK_PRIMARY_CURRENCY"
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO = "MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO"

	MESSAGE_TEXT_BALANCE_IS_ZERO                                   = "MESSAGE_TEXT_BALANCE_IS_ZERO"
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO                      = "MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO"
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER = "MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER"
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER = "MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER"

	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER = "MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER"
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER = "MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER"

	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE = "MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE"
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY         = "MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY"
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED        = "MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED"
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT         = "MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT"
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT       = "MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT"
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL                    = "MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL"

	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER = "MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER"

	MESSAGE_TEXT_RECEIPT_LINK = "MESSAGE_TEXT_RECEIPT_LINK"

	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER = "MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER"
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER   = "MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER"
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER   = "MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER"

	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY         = "MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY"
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER         = "MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER"
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT                = "MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT"
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT = "MESSAGE_TEXT_THIS__NUMBER_WILL_BE_USED_TO_SEND_RECEIPT"

	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE  = "BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE"
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT = "BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT"
	BUTTON_TEXT_DEBT_RETURNED_FULLY        = "BUTTON_TEXT_DEBT_RETURNED_FULLY"
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY    = "BUTTON_TEXT_DEBT_RETURNED_PARTIALLY"

	BUTTON_TEXT_SEE_RECEIPT_DETAILS = "BUTTON_TEXT_SEE_RECEIPT_DETAILS"

	MESSAGE_TEXT_BALANCE_HEADER                                               = "MESSAGE_TEXT_BALANCE_HEADER"
	INLINE_BUTTON_SHOW_FULL_HISTORY                                           = "INLINE_BUTTON_SHOW_FULL_HISTORY"
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED                               = "MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED"
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER                                 = "MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER"
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE                                  = "MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE"
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE = "MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE"
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE = "MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE"

	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT = "MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT"

	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY = "MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY"

	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS = "MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS"

	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM                          = "MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM"
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT = "MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT"
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL                             = "MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL"
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS                               = "MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS"

	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT = "MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT"
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT = "COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT"
	MESSAGE_TEXT_HELP                                   = "MESSAGE_TEXT_HELP"
	COMMAND_TEXT_OPEN_USER_REPORT                       = "COMMAND_TEXT_OPEN_USER_REPORT"
	COMMAND_TEXT_REPORT_A_BUG                           = "COMMAND_TEXT_REPORT_A_BUG"
	COMMAND_TEXT_SUBMIT_AN_IDEA                         = "COMMAND_TEXT_SUBMIT_AN_IDEA"
)

// Inline
const (
	INLINE_RECEIPT_TITLE           = "INLINE_RECEIPT_TITLE"
	INLINE_RECEIPT_DESCRIPTION     = "INLINE_RECEIPT_DESCRIPTION"
	INLINE_RECEIPT_MESSAGE         = "INLINE_RECEIPT_MESSAGE"
	INLINE_RECEIPT_CHOOSE_LANGUAGE = "INLINE_RECEIPT_CHOOSE_LANGUAGE"

	INLINE_INVITE_TITLE       = "INLINE_INVITE_TITLE"
	INLINE_INVITE_DESCRIPTION = "INLINE_INVITE_DESCRIPTION"
	INLINE_INVITE_MESSAGE     = "INLINE_INVITE_MESSAGE"
)

// Command texts
const (
	COMMAND_TEXT_WAIT_A_SECOND = "COMMAND_TEXT_WAIT_A_SECOND"
	COMMAND_TEXT_ACCEPT        = "COMMAND_TEXT_ACCEPT"
	COMMAND_TEXT_DECLINE       = "COMMAND_TEXT_DECLINE"

	HTML_USING_TELEGRAM = "HTML_USING_TELEGRAM"

	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM = "BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM"
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM = "BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM"

	COMMAND_TEXT_VIEW_RECEIPT_DETAILS = "COMMAND_TEXT_VIEW_RECEIPT_DETAILS"
	COMMAND_TEXT_ACCEPT_INVITE        = "COMMAND_TEXT_ACCEPT_INVITE"
	COMMAND_TEXT_MAIN_MENU_TITLE      = "Main /menu"
	COMMAND_TEXT_DISABLE_REMINDER     = "COMMAND_TEXT_DISABLE_REMINDER"
	COMMAND_TEXT_TOMORROW             = "COMMAND_TEXT_TOMORROW"
	COMMAND_TEXT_DAY_AFTER_TOMORROW   = "COMMAND_TEXT_DAY_AFTER_TOMORROW"
	COMMAND_TEXT_THIS_WEEK            = "COMMAND_TEXT_THIS_WEEK"
	COMMAND_TEXT_IN_FEW_MINUTES       = "COMMAND_TEXT_IN_FEW_MINUTES"
	COMMAND_TEXT_IN_1_WEEK            = "COMMAND_TEXT_IN_1_WEEK"
	COMMAND_TEXT_IN_1_MONTH           = "COMMAND_TEXT_IN_1_MONTH"
	COMMAND_TEXT_SET_DATE             = "COMMAND_TEXT_SET_DATE"

	COMMAND_TEXT_MONDAY    = "COMMAND_TEXT_MONDAY"
	COMMAND_TEXT_TUESDAY   = "COMMAND_TEXT_TUESDAY"
	COMMAND_TEXT_WEDNESDAY = "COMMAND_TEXT_WEDNESDAY"
	COMMAND_TEXT_THURSDAY  = "COMMAND_TEXT_THURSDAY"
	COMMAND_TEXT_FRIDAY    = "COMMAND_TEXT_FRIDAY"
	COMMAND_TEXT_SATURDAY  = "COMMAND_TEXT_SATURDAY"
	COMMAND_TEXT_SUNDAY    = "COMMAND_TEXT_SUNDAY"

	COMMAND_TEXT_YES             = "COMMAND_TEXT_YES"
	COMMAND_TEXT_YES_EXCLAMATION = "COMMAND_TEXT_YES_EXCLAMATION"
	COMMAND_TEXT_NOT_TOO_MUCH    = "COMMAND_TEXT_NOT_TOO_MUCH"
	COMMAND_TEXT_NO    = "COMMAND_TEXT_NO"

	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE    = "COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE"
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME = "COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME"
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME    = "COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME"

	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED = "MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED"

	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY = "MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY"

	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY       = "MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY"
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT    = "MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT"
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE               = "MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE"
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT            = "MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT"
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE = "MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE"
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT = "MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT"

	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER = "COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER"
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER         = "COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER"
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER            = "COMMAND_TEXT_NO_NOTE_FOR_TRANSFER"

	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER    = "COMMAND_TEXT_ADD_NOTE_TO_TRANSFER"
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER = "COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER"

	COMMAND_TEXT_CANCEL                    = "COMMAND_TEXT_CANCEL"
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION       = "COMMAND_TEXT_ADD_YOUR_OWN_OPTION"
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY = "COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY"
	COMMAND_TEXT_LANGUAGE                  = "COMMAND_TEXT_LANGUAGE"
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE = "COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE"

	COMMAND_TEXT_INVITE_ME_BY_EMAIL = "COMMAND_TEXT_INVITE_ME_BY_EMAIL"
	COMMAND_TEXT_INVITE_ME_BY_SMS   = "COMMAND_TEXT_INVITE_ME_BY_SMS"

	COMMAND_TEXT_SEND_ME_NEW_INVITE             = "COMMAND_TEXT_SEND_ME_NEW_INVITE"
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL    = "COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL"
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS      = "COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS"
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM = "COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM"

	COMMAND_TEXT_SEND_MY_PHONE_NUMBER = "COMMAND_TEXT_SEND_MY_PHONE_NUMBER"

	COMMAND_TEXT_SEND_BY_EMAIL      = "COMMAND_TEXT_SEND_BY_EMAIL"
	COMMAND_TEXT_SEND_BY_SMS        = "COMMAND_TEXT_SEND_BY_SMS"
	COMMAND_TEXT_INVITE_BY_TELEGRAM = "COMMAND_TEXT_INVITE_BY_TELEGRAM"

	COMMAND_TEXT_I_HAVE_INVITE              = "I_HAVE_INVITE"
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL       = "COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL"
	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN     = "COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN"
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES = "COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES"

	COMMAND_TEXT_GAVE        = "COMMAND_TEXT_GAVE"
	COMMAND_TEXT_GOT         = "COMMAND_TEXT_GOT"
	COMMAND_TEXT_RETURN      = "COMMAND_TEXT_RETURN"
	COMMAND_TEXT_BALANCE     = "COMMAND_TEXT_BALANCE"
	COMMAND_TEXT_SETTING     = "COMMAND_TEXT_SETTING"
	COMMAND_TEXT_HIGH_FIVE   = "COMMAND_TEXT_HIGH_FIVE"
	COMMAND_TEXT_CHANGE_LANG = "COMMAND_TEXT_CHANGE_LANG"
	COMMAND_TEXT_HELP        = "COMMAND_TEXT_HELP"
	COMMAND_TEXT_HISTORY     = "COMMAND_TEXT_HISTORY"
	//COMMAND_TEXT_DUE_RETURNS = "COMMAND_TEXT_DUE_PAYMENTS"

	COMMAND_TEXT_NEW_COUNTERPARTY = "COMAND_TEXT_NEW_COUNTERPARTY"

	COMMAND_TEXT_SUBSCRIBE_TO_APP  = "COMMAND_TEXT_SUBSCIBE_TO_APP"
	MESSAGE_TEXT_SUBSCRIBED_TO_APP = "MESSAGE_TEXT_SUBSCRIBED_TO_APP"

	COMMAND_TEXT_I_AM_FINE_WITH_BOT    = "COMMAND_TEXT_I_AM_FINE_WITH_BOT"
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP = "MESSAGE_TEXT_NOT_INTERESTED_IN_APP"

	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND       = "COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND"
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM     = "COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM"
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM = "COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM"
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS          = "COMMAND_TEXT_SEND_RECEIPT_BY_SMS"
	COMMAND_TEXT_SEND_RECEIPT_BY_VK           = "COMMAND_TEXT_SEND_RECEIPT_BY_VK"
	COMMAND_TEXT_SEND_RECEIPT_BY_OK           = "COMMAND_TEXT_SEND_RECEIPT_BY_OK"
	COMMAND_TEXT_SEND_RECEIPT_BY_FB           = "COMMAND_TEXT_SEND_RECEIPT_BY_FB"
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT          = "COMMAND_TEXT_SEND_RECEIPT_BY_TWT"

	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM = "COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM"
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT                = "COMMAND_TEXT_DO_NOT_SEND_RECEIPT"
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT           = "MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT"

	COMMAND_TEXT_SHOW_ALL_CONTACTS = "COMMAND_TEXT_SHOW_ALL_CONTACTS"

	MESSAGE_TEXT_REMINDER                     = "MESSAGE_TEXT_REMINDER"
	MESSAGE_TEXT_REMINDER_OUTSTANDING_DEBT    = "MESSAGE_TEXT_REMINDER_OUTSTANDING_DEBT"
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED     = "MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED"
	MESSAGE_TEXT_REMINDER_DISABLED            = "MESSAGE_TEXT_REMINDER_DISABLED"
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED = "MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED"
	MESSAGE_TEXT_REMINDER_SET                 = "MESSAGE_TEXT_REMINDER_SET"
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL    = "COMMAND_TEXT_REMINDER_RETURNED_IN_FULL"
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY  = "COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY"
	COMMAND_TEXT_REMINDER_NOT_RETURNED        = "COMMAND_TEXT_REMINDER_NOT_RETURNED"
	MESSAGE_TEXT_YOU_REPLIED                  = "MESSAGE_TEXT_YOU_REPLIED"
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN     = "MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN"
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY  = "MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY"

	MESSAGE_TEXT_DEBT_IS_RETURNED = "MESSAGE_TEXT_DEBT_IS_RETURNED"
	MESSAGE_TEXT_DETAILS_ARE_HERE = "MESSAGE_TEXT_DETAILS_ARE_HERE"
)

// SMS
const (
	SMS_RECEIPT_YOU_GOT             = "SMS_RECEIPT_YOU_GOT"
	SMS_RECEIPT_YOU_GAVE            = "SMS_RECEIPT_YOU_GAVE"
	SMS_CLICK_TO_CONFIRM_OR_DECLINE = "SMS_CLICK_TO_CONFIRM_OR_DECLINE"
	TELEGRAM_RECEIPT                = "TELEGRAM_RECEIPT"
)

// HTML
const (
	HTML_DATE    = "HTML_DATE"
	HTML_RECEIPT = "HTML_RECEIPT"
	HTML_AMOUNT  = "HTML_AMOUNT"
	HTML_FROM    = "HTML_FROM"
	HTML_TO      = "HTML_TO"
)

// DUE
const (
	DUE_IN_NOW       = "DUE_IN_NOW"
	DUE_IN_A_MINUTE  = "DUE_IN_A_MINUTE"
	DUE_IN_X_MINUTES = "DUE_IN_X_MINUTES"
	DUE_IN_AN_HOUR   = "DUE_IN_AN_HOUR"
	DUE_IN_X_HOURS   = "DUE_IN_X_HOURS"
	DUE_IN_X_DAYS    = "DUE_IN_X_DAYS"
)

// Website
const (
	WS_ALEX_T     = "WS_ALEX_T"
	WS_MOTTO      = "WS_MOTTO"
	WS_SHORT_DESC = "WS_SHORT_DESC"

	WS_INDEX_TITLE       = "WS_INDEX_TITLE"
	WS_INDEX_TG_BOT_H2   = "WS_INDEX_TG_BOT_H2"
	WS_INDEX_TG_BOT_P    = "WS_INDEX_TG_BOT_P"
	WS_INDEX_TG_BOT_OPEN = "WS_INDEX_TG_BOT_OPEN"

	WS_ADS_TITLE   = "WS_ADS_TITLE"
	WS_ADS_CONTENT = "WS_ADS_CONTENT"

	WS_HELP_US_TITLE   = "WS_HELP_US_TITLE"
	WS_HELP_US_CONTENT = "WS_HELP_US_CONTENT"

	WS_LIVE_DEMO = "WS_LIVE_DEMO"
)

// Facebook
const (
	FB_MAKE_RECORD_HEADER    = "HEADER_MAKE_RECORD"
	FB_MAKE_RECORD_HEADLINE  = "MAKE_RECORD_HEADLINE"
	FB_HOW_ARE_THINGS_HEADER = "FB_HOW_ARE_THINGS_HEADER"
)
