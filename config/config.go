package config

var (
	messengerAccessToken string
	messengerVerifyToken string
	fnacCookie string
	messengerRecipientIDs []string
	chromeHost string
	chromePort string
	cronIntervalSeconds int
	serverPort int
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

// GetServerPort returns the HTTP port the server will be binded to
func GetServerPort() int {
	return serverPort
}