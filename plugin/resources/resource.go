package resources

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/trace"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/common/rest"
	jwt "github.com/dgrijalva/jwt-go"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/models"

	"strings"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

// ArmadaEndpoint struct defining config info for the API
type Config struct {
	AccessToken      string
	FindingsURL      string
	NotificationsURL string
	AccountID        string
	RefreshToken     string
	Context          plugin.PluginContext
	restClient       *rest.Client
}

func (config *Config) newRequest(api string, method string, path string, body interface{}) *rest.Request {
	// Standard request header
	url := ""
	if api == "findings" {
		url = config.FindingsURL + "/" + config.AccountID + "/" + path
	} else if api == "notifications" {
		url = config.NotificationsURL + "/" + config.AccountID + "/" + path
	}
	r := rest.GetRequest(url).Method(method)
	r.Set("Authorization", fmt.Sprintf("%s", config.AccessToken))
	r.Set("X-Auth-Refresh-Token", config.RefreshToken)
	r.Set("X-Auth-Resource-Account", config.AccountID)
	r.Body(body)

	return r
}

func (config *Config) setContentType(r *rest.Request, content string) *rest.Request {
	if content == "graphql" {
		r.Set("Content-Type", "application/graphql")
	} else {
		r.Set("Content-Type", "application/json")
	}
	return r
}

func (config *Config) makeRequest(api string, method string, path string, body interface{}, successV interface{}, content string) error {
	err := config.refreshTokensIfNeeded()
	if err != nil {
		return err
	}

	// Make the request and parse results
	req := config.newRequest(api, method, path, body)
	req = config.setContentType(req, content)
	logf("IAM token is valid, preparing to make request: %v", req)

	// Armada has a standard error response struct
	resp, err := config.client().Do(req, successV, nil)
	if err != nil {
		logf("Error occurred: %s", err.Error())
		return err
	}

	log("response", resp.Body)

	logf("Request success: %v", resp)
	return nil
}

func (config *Config) refreshTokensIfNeeded() error {
	err := config.refreshIAMTokenIfNeeded()
	if err != nil {
		return err
	}

	return nil
}

// Validate the IAM token locally
func (config *Config) refreshIAMTokenIfNeeded() error {
	log("Validating IAM token locally")
	if claims, err := parseJWTToken(config.Context.IAMToken()); err == nil {
		if err := claims.Valid(); err != nil {
			log("The IAM token is expired, attempting to refresh")
			if updatedToken, err := config.Context.RefreshIAMToken(); err == nil {
				logf("Updated IAM token successfully, new token is: %s", updatedToken)
				log("Sleeping for a moment to allow token issue time to become valid")
				config.AccessToken = updatedToken
				time.Sleep(1 * time.Second)
			} else {
				log("Failed to refresh IAM token")
				return errors.New("Login token is expired. Please update tokens using 'ibmcloud login' and try again.")
			}
		}
	}

	return nil
}

func (config *Config) client() *rest.Client {
	if config.restClient == nil {
		config.restClient = &rest.Client{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					Dial: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).Dial,
					TLSHandshakeTimeout:   10 * time.Second,
					ResponseHeaderTimeout: 15 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					MaxIdleConns:          100,
					IdleConnTimeout:       90 * time.Second,
				},
			},
		}
	}

	return config.restClient
}

// GetConfig ...
func GetConfig(context plugin.PluginContext) *Config {
	return &Config{
		AccessToken:      context.IAMToken(),
		FindingsURL:      os.Getenv("FINDINGS_API_ENDPOINT"),
		NotificationsURL: os.Getenv("NOTIFICATIONS_API_ENDPOINT"),
		RefreshToken:     context.IAMRefreshToken(),
		AccountID:        context.CurrentAccount().GUID,
		Context:          context,
	}
	fmt.Errorf("Something went wrong!")
	return nil
}

// parseJWTToken parses a JWT token without verifying the signature
func parseJWTToken(token string) (*jwt.StandardClaims, error) {
	// Remove "bearer " in front if there is any
	// to lowercase as the cli sends in Bearer
	if strings.HasPrefix(token, "bearer ") {
		segments := strings.SplitAfterN(token, "bearer ", 2)
		token = segments[len(segments)-1]
	} else if strings.HasPrefix(token, "Bearer ") {
		segments := strings.SplitAfterN(token, "Bearer ", 2)
		token = segments[len(segments)-1]
	}

	// Parse the token
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, nil)
	if parsedToken != nil {
		if claims, ok := parsedToken.Claims.(*jwt.StandardClaims); ok {
			return claims, nil
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token malformed")
			}
		}
	}
	return nil, errors.New("cannot parse token")
}

// PrintJSON ...
func PrintJSON(data interface{}) (string, error) {
	contents, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func log(args ...interface{}) {
	trace.Logger.Println(args...)
}

func logf(fmt string, args ...interface{}) {
	trace.Logger.Printf(fmt, args...)
}

func (config *Config) PostGraph(query string, query_type string) (map[string]interface{}, error) {
	var contentType string
	var result map[string]interface{}
	if query_type == "graphql" {
		contentType = "graphql"
	} else if query_type == "json" {
		contentType = "json"
	}
	err := config.makeRequest("findings", "POST", "graph", query, &result, contentType)
	return result, err
}

func (config *Config) GetProviders() (models.ApiListProvidersResponse, error) {
	var result models.ApiListProvidersResponse
	err := config.makeRequest("findings", "GET", "providers", "", &result, "json")
	return result, err
}

func (config *Config) PostNote(provider string, query string) (models.ApiNote, error) {
	var result models.ApiNote
	err := config.makeRequest("findings", "POST", "providers/"+provider+"/notes", query, &result, "json")
	return result, err
}

func (config *Config) GetNotes(provider string) (models.ApiListNotesResponse, error) {
	var result models.ApiListNotesResponse
	err := config.makeRequest("findings", "GET", "providers/"+provider+"/notes", "", &result, "json")
	return result, err
}

func (config *Config) GetNoteByNoteID(provider string, id string) (models.ApiNote, error) {
	var result models.ApiNote
	err := config.makeRequest("findings", "GET", "providers/"+provider+"/notes/"+id, "", &result, "json")
	return result, err
}

func (config *Config) GetNoteByOccId(provider string, id string) (models.ApiNote, error) {
	var result models.ApiNote
	err := config.makeRequest("findings", "GET", "providers/"+provider+"/occurrences/"+id+"/note", "", &result, "json")
	return result, err
}

func (config *Config) PutNote(provider string, id string, query string) (models.ApiNote, error) {
	var result models.ApiNote
	err := config.makeRequest("findings", "PUT", "providers/"+provider+"/notes/"+id, query, &result, "json")
	return result, err
}

func (config *Config) DeleteNote(provider string, id string) (models.ApiEmpty, error) {
	var result models.ApiEmpty
	err := config.makeRequest("findings", "DELETE", "providers/"+provider+"/notes/"+id, nil, &result, "json")
	return result, err
}

func (config *Config) PostChannel(body string) (models.NotificationChannel, error) {
	var result models.NotificationChannel
	err := config.makeRequest("notifications", "POST", "notifications/channels", body, &result, "json")
	return result, err
}

func (config *Config) DeleteChannel(id string) (models.ApiEmpty, error) {
	var result models.ApiEmpty
	err := config.makeRequest("notifications", "DELETE", "notifications/channels/"+id, nil, &result, "json")
	return result, err
}

func (config *Config) DeleteChannels(body string) (models.ApiEmpty, error) {
	var result models.ApiEmpty
	err := config.makeRequest("notifications", "DELETE", "notifications/channels", body, &result, "json")
	return result, err
}

func (config *Config) GetChannels() (models.ChannelList, error) {
	var result models.ChannelList
	err := config.makeRequest("notifications", "GET", "notifications/channels", nil, &result, "json")
	return result, err
}

func (config *Config) GetChannel(id string) (models.GetNotificationChannel, error) {
	var result models.GetNotificationChannel
	err := config.makeRequest("notifications", "GET", "notifications/channels/"+id, nil, &result, "json")
	return result, err
}

func (config *Config) PutChannel(id string, body string) (models.NotificationChannel, error) {
	var result models.NotificationChannel
	err := config.makeRequest("notifications", "PUT", "notifications/channels/"+id, body, &result, "json")
	return result, err
}

func (config *Config) GetKey() (models.PublicKey, error) {
	var result models.PublicKey
	err := config.makeRequest("notifications", "GET", "notifications/public_key", nil, &result, "json")
	return result, err
}

func (config *Config) TestChannel(id string) (models.TestChannelResponse, error) {
	var result models.TestChannelResponse
	err := config.makeRequest("notifications", "GET", "notifications/channels/"+id+"/test", nil, &result, "json")
	return result, err
}

func (config *Config) ReadFile(file string) ([]byte, error) {
	JSONfile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer JSONfile.Close()
	b, err := ioutil.ReadAll(JSONfile)
	if err != nil {
		return nil, err
	}
	return b, nil
}
