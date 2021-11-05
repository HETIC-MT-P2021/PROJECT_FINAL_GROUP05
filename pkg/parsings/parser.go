package parsings

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/pkg/models"
)

var (
	inputUriParam string
)

const (
	RegexFindType         = `^(\/[\w\-]+)`
	RegexFindParams       = `(?m)-.\s*([^\s]+)`
	RegexFindParamsValue  = `(?m)(?P<value> (.*?)$)`
	RegexFindOptionsValue = `(?:-options)(?:\s)+(?P<Filter>[a-z ]+)`
)

func getTypeFromString(input string) string {
	typeRegex := regexp.MustCompile(RegexFindType)
	mediaType := typeRegex.FindString(input)
	return mediaType
}

func getParamsFromString(input string) []string {
	var regexParams = regexp.MustCompile(RegexFindParams)
	return regexParams.FindAllString(input, -1)
}

func getOptionsFromString(input string) models.Options {
	var startParamInSeconds int
	var durationParamInSeconds int
	var params = getParamsFromString(input)
	var options models.Options
	var filters []string

	reValue := regexp.MustCompile(RegexFindParamsValue)

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

	filters = getFilterFromString(input)

	options.DurationInSeconds = durationParamInSeconds
	options.StartInSeconds = startParamInSeconds
	options.Filters = filters
	return options
}

func getFilterFromString(input string) []string {
	var regexFilter = regexp.MustCompile(RegexFindOptionsValue)
	match := regexFilter.FindStringSubmatch(input)
	filter := regexFilter.SubexpIndex("Filter")
	return strings.Fields(match[filter])
}

func ParseCommand(inputString string) models.DownloadQueueMessage {
	var message models.DownloadQueueMessage

	if inputUriParam != "" {
		inputUriParam = ""
	}

	getOptionsFromString(inputString)
	message.Type = getTypeFromString(inputString)

	message.Options = getOptionsFromString(inputString)
	message.MediaLink = inputUriParam

	return message
}
