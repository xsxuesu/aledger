package common


const ADMIN_ORG = "mmOrg"
const ADMIN_Name = "mmadmin"
const ADMIN_ADDRESS = "0xdce75625e4267CE1d15ab8EBa69811cDA23101F4"

const VERSION = "Ledger-Code v1.0"

const LEDGER = "ledger"
const ACCOUNT = "account"
const TOKEN = "token"

const ACCOUNT_PRE = "ACCOUNT_"
const TOKEN_PRE = "TOKEN_"
const Ledger_PRE = "LEDGER"
const SIGN_PRE = "SIGN"

const CompositeIndexName = "pre~tkn~name"
const CompositeRequestIndexName = "pre~tkn~name~txid"
const Evt_payment uint8 = 255
const TOPIC = "LEDGER_TX_%s"

const GETSTAT_ERR = 500
const PUTSTAT_ERR = 600
const MARSH_ERR   = 400
const COMPOSTEKEY_ERR = 700
const EVENT_ERR =800
const (
	TKNERR_EXIST   = 501
	TKNERR_LOCKED  = 502
	TKNERR_PREMISSON	 = 503
)

const (
	ACCOUNT_EXIST 	  = 504
	ACCOUNT_NOT_EXIST = 505
	ACCOUNT_PREMISSION = 506
	ACCOUNT_LOCK = 507
	ACCOUNT_COMMONNAME = 508
)

const (
	Right_ERR = 504
	Param_ERR = 505
	FORMAT_ERR = 508
	Balance_NOT_ENOUGH = 506
	STATUS_ERR = 507
)

const(
	PENDING_SIGN = "Pending"
	SENT_SIGN 	= "Sent"
	Failed_SIGN = "Refused"

	TOKEN_INIT  = "Init"
	TOKEN_ISSUE = "Issue"
	TOKEN_BURN	= "Burn"
	TOKEN_BREAK = "Break"
	TOKEN_MERGE = "Merge"
)