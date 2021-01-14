package config

var (
	messengerAccessToken string
	messengerVerifyToken string
	fnacCookie string
	cDiscountCookie string
	messengerRecipientIDs []string
	chromeHost string
	chromePort string
	cronIntervalSeconds int
)

// GetMessengerAccessToken returns the access token used to interact with messenger API
func GetMessengerAccessToken() string {
	return messengerAccessToken
}

// GetMessengerVerifyToken returns the verification token for the webhook used in messenger API
func GetMessengerVerifyToken() string {
	return messengerVerifyToken
}

// GetFnacCookie returns the fnac cookie not to be triggered by datadome
func GetFnacCookie() string {
	return fnacCookie
}

// GetMessengerRecipientIDs returns the ids of the messenger alerts
func GetMessengerRecipientIDs() []string {
	return messengerRecipientIDs
}

// GetCDiscountCookie returns the cookie needed to access the page
func GetCDiscountCookie() string {
	return cDiscountCookie
}

// GetChromeHost returns the host for the headless browser
func GetChromeHost() string {
	return chromeHost
}

// GetChromePort returns the host for the headless browser
func GetChromePort() string {
	return chromePort
}

// GetCronIntervalSeconds returns the number, in minutes, when the scrapers should launch
func GetCronIntervalSeconds() int {
	return cronIntervalSeconds
}