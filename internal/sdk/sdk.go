package sdk

import (
	_ "embed"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sailpoint-oss/sailpoint-cli/internal/util"
)

//go:embed sdkErr.md
var sdkErrBody string
var sdkErrParts = strings.Split(sdkErrBody, "====")

func HandleSDKError(resp *http.Response, sdkErr error) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(err)
	}

	return errors.New(util.RenderMarkdown(sdkErrParts[0] + util.PrettyPrint(resp.Header) + sdkErrParts[1] + util.PrettyPrint(data) + sdkErrParts[2]))
}
