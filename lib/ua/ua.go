package ua

import (
	"fmt"
	"math/rand"
)

// PCBrowserUA generates a user-agent string for a PC browser.
func PCBrowserUA() string {
	platforms := []string{
		"Windows NT 10.0; Win64; x64",
		"Windows NT 6.1; Win64; x64",
		"Macintosh; Intel Mac OS X 10_15_7",
		"Linux x86_64",
	}
	browserVersions := []string{
		"Chrome/115.0.0.0",
		"Firefox/110.0",
		"Safari/537.36",
		"Edge/93.0.0.0",
	}
	platform := platforms[rand.Intn(len(platforms))]
	browser := browserVersions[rand.Intn(len(browserVersions))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s", platform, browser)
}

// MobileBrowserUA generates a user-agent string for a mobile browser.
func MobileBrowserUA() string {
	androidVersions := []string{"9", "10", "11"}
	androidModel := []string{"ALP-AL00", "Pixel 5", "SM-G973F"}
	return fmt.Sprintf("Dalvik/2.1.0 (Linux; U; Android %s; %s Build/HUAWEIALP-AL00)", androidVersions[rand.Intn(len(androidVersions))], androidModel[rand.Intn(len(androidModel))])
}
