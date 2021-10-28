package parsings

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
	"regexp"
	"strconv"
	"strings"
)

var (
	mediaType              string
	message                models.DownloadQueueMessage
	startParamInSeconds    int
	durationParamInSeconds int
	inputUriParam          string
)

func getTypeFromString(input string) string {
	typeRegex := regexp.MustCompile(`^([\w\-]+)`)
	mediaType = typeRegex.FindString(input)
	return mediaType
}

func getParamsFromString(input string) []string {
	var re = regexp.MustCompile(`(?m)-. ([^\s]+)`)
	return re.FindAllString(input, -1)
}

func getOptionsFromString(input string) models.Options {
	var params = getParamsFromString(input)
	var options models.Options

	reValue := regexp.MustCompile(`(?m)(?P<value> (.*?)$)`)

	for _, match := range params {
		var value = strings.ReplaceAll(reValue.FindString(match), " ", "")

		switch string(match[1]) {
		case "s":
			startParamInSeconds, _ = strconv.Atoi(value)
			break
		case "d":
			durationParamInSeconds, _ = strconv.Atoi(value)
			break
		case "i":
			inputUriParam = value
			break
		default:
			break
		}
	}

	options.DurationInSeconds = startParamInSeconds
	options.StartInSeconds = durationParamInSeconds
	return options
}

func ParseCommand(inputString string) models.DownloadQueueMessage {
	getOptionsFromString(inputString)
	message.Type = getTypeFromString(inputString)

	message.Options = getOptionsFromString(inputString)
	message.MediaLink = inputUriParam

	return message
}
