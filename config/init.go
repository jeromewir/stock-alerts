package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFromEnvironment is the initialization function for the configuration
// Call this at the beginning of the program
func ReadFromEnvironment() error {
	messengerAccessToken = os.Getenv("MESSENGER_ACCESS_TOKEN")
	messengerVerifyToken = os.Getenv("MESSENGER_VERIFY_TOKEN")
	fnacCookie = os.Getenv("COOKIE_FNAC")
	cDiscountCookie = os.Getenv("CDISCOUNT_COOKIE")
	messengerRecipientS := os.Getenv("MESSENGER_RECIPIENT_IDS")
	chromeHost = os.Getenv("CHROME_HOST")
	chromePort = os.Getenv("CHROME_PORT")
	cronIntervalSecondsString := os.Getenv("CRON_INTERVAL_SECONDS")

	messengerRecipientIDs = strings.Split(messengerRecipientS, ",")

	for i, id := range messengerRecipientIDs {
		messengerRecipientIDs[i] = strings.TrimSpace(id)
	}

	if messengerAccessToken == "" {
		return errors.New("Expected MESSENGER_ACCESS_TOKEN defined in environment")
	}

	if messengerVerifyToken == "" {
		return errors.New("Expected MESSENGER_VERIFY_TOKEN defined in environment")
	}

	if chromeHost == "" {
		return errors.New("Expected CHROME_HOST defined in environment")
	}
	
	if chromePort == "" {
		return errors.New("Expected CHROME_PORT defined in environment")
	}

	if cronIntervalSecondsString != "" {
		var err error
		cronIntervalSeconds, err = strconv.Atoi(cronIntervalSecondsString)

		if err != nil {
			fmt.Println("Warning: `CRON_INTERVAL_SECONDS` is not a number, defaulting to 300 seconds")
			cronIntervalSeconds = 300
		}
	} else {
		cronIntervalSeconds = 300
	}

	return nil
}