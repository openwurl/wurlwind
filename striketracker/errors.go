package striketracker

// Possible errors for convenient matching
//
//  if err = striketracker.ErrUnhandledFatal {}
const (
	ErrUnhandledFatal            = "0: An unhandled fatal error has occurred"
	ErrGeneric                   = "1: Generic error"
	ErrMissingRequiredParameter  = "2: A required parameter is missing from a request"
	ErrInvalidHashCode           = "3: A hash code given is not of a valid format"
	ErrAccountContextNotFound    = "4: The account request in the account context was not found"
	ErrInvalidRequestJSON        = "5: The request data is not valid JSON"
	ErrDisabledFeature           = "6: This feature is temporarily disabled"
	ErrMissingConfigValue        = "7: An expected configuration value was not set"
	ErrMissingGroupByParam       = "100: The groupBy parameter on an analytics request is missing or invalid"
	ErrInvalidGranularity        = "101: An invalid granularity parameter was provided for an analytics request"
	ErrMissingOrInvalidFilter    = "102: A required filter is missing or a filter contains a bad value"
	ErrInvalidDateRange          = "103: The date range provided for an analytics request is not valid"
	ErrInvalidEndDate            = "104: The end date provided is too far in the future (past the end of the current month)"
	ErrExpiredQuery              = "105: The query results for the requested job ID have expired"
	ErrUnavailableQuery          = "106: The query results are not yet available (probably because the query is still running). @var int"
	ErrBulkAnalyticsDisabled     = "107: The account sending the request does not have the Bulk Analytics service enabled."
	ErrMissingGrantType          = "200: The grant_type parameter for authentication is missing or an invalid value"
	ErrInvalidAccountContext     = "201: The specified account context is invalid (account is suspended or deleted)"
	ErrUnauthenticated           = "203: Resource requires authentication but user is not authenticated"
	ErrPasswordExpired           = "204: Resource requires authentication but the user's password is expired"
	ErrInvalidConfirmationCode   = "205: Confirmation code is invalid or expired"
	ErrNotWhitelisted            = "206: The Client IP is not in the Users IP Whitelist"
	ErrPermissionDenied          = "300: The authenticated user does not have permissions to perform the requested action"
	ErrUserSuspended             = "301: The authenticated user is suspended"
	ErrAccountSuspended          = "302: The account is suspended"
	ErrUnassociatedUser          = "303: The authenticated user does not have an associated account"
	ErrValidationFailure         = "400: A resource has failed validation"
	ErrDuplicateOrigin           = "401: A duplicate origin exists with the same hostname, port, and path"
	ErrNotFound                  = "404: The requested resource was not found"
	ErrEndpointNotFound          = "405: The requested endpoint was not found, please check your url"
	ErrResourceExists            = "409: This resource already exists"
	ErrResourceDeleted           = "410: This resource has been deleted or is expired"
	ErrWildcardConflict          = "411: There is a conflict with a wildcard hostname"
	ErrLockedResource            = "423: The requested resource is locked"
	ErrLimitExceeded             = "429: Your use of this resource exceeds specified rate limit"
	ErrDatabaseDown              = "503: Unable to reach the database. Please try again later."
	ErrCDNUnresponsive           = "504: Unable to send configuration to the CDN. Please try again later."
	ErrApplicationMaintenance    = "505: Application is currently down for maintenance. Please try again later."
	ErrHCSDown                   = "506: Unable to reach HCS. Please try again later."
	ErrSOLRDown                  = "507: Unable to reach the SOLR API. Please try again later."
	ErrAnalyticsDown             = "508: Unable to retrieve current analytics data"
	ErrAnalyticsTimeout          = "509: Analytics request timed out. Please try again later."
	ErrAnalyticsResourceNotFound = "510: Unable to find analytics resource. Please try again later."
	ErrAnalyticsDBDown           = "511: Unable to reach analytics database. Please try again later."
	ErrHCSValidationFailure      = "600: HCS validation failed"
	ErrInvalidHCSAuthToken       = "601: HCS Invalid Auth Token"
	ErrHCSNotEnabled             = "602: HCS Service Not Enabled"
	ErrMaxTenantsHCS             = "603: HCS Max tenants exceeded"
	ErrEveryStreamDisabled       = "700: EveryStream Service Not Enabled"
	ErrEveryStreamNotFound       = "701: EveryStream Account Not Found"
	ErrEveryStreamSuspended      = "702: EveryStream Account Suspended"
	ErrEncodingJobStarted        = "703: Encoding Job already in progress"
	ErrEveryStreamQuota          = "704: EveryStream encoded quota exceeded"
	ErrTransmuxDisabled          = "705: Transmux Service is not enabled on this account"
	ErrTransmuxUnprovisioned     = "706: Transmux Service has not been provisioned"
	ErrTransmuxSuspended         = "707: Transmux Service is suspended"
)
