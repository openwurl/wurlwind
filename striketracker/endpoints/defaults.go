package endpoints

/*
/api/v1/
> accounts/
> barometer/
> ips/
> pops/
*/

// Base
const (
	URL = "https://striketracker.highwinds.com" // HTTPS not optional
	V1  = "/api/v1"                             // Striketracker API V1
)

// BasePaths
const (
	ACCOUNTS       = "/accounts"
	BAROMETER      = "/barometer"
	BILLINGREGIONS = "/billingRegions"
	IPS            = "/ips"
	POPS           = "/pops"
)

// BasePath wraps the various possible base paths
type BasePath string

// Service type to base path
const (
	Accounts       BasePath = ACCOUNTS
	Analytics      BasePath = ACCOUNTS
	Authentication BasePath = ACCOUNTS
	Barometer      BasePath = BAROMETER
	BillingRegions BasePath = BILLINGREGIONS
	Certificates   BasePath = ACCOUNTS
	Configuration  BasePath = ACCOUNTS
	Hcs            BasePath = ACCOUNTS
	Hosts          BasePath = ACCOUNTS
	Notification   BasePath = ACCOUNTS
	Origins        BasePath = ACCOUNTS
	Platforms      BasePath = ACCOUNTS
	Pops           BasePath = POPS
	Ips            BasePath = IPS
	Purge          BasePath = ACCOUNTS
	Search         BasePath = ACCOUNTS
	Services       BasePath = ACCOUNTS
	Sessions       BasePath = ACCOUNTS
	Users          BasePath = ACCOUNTS
)

/*
Accounts
Analytics
Authentication
Barometer - barometer
BillingRegions - billingregions
Certificates
Configuration
Hcs
Hosts
Notification
Origins
Platforms
Pops - IPS/POPS
Purge
Search
Services
Sessions
Users
*/
