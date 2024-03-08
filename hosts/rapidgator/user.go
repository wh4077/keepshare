// https://rapidgator.net/article/api/user

package rapidgator

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// https://rapidgator.net/article/api/user#login
const userLoginRequired2FAAPIFormat = "/api/v2/user/login?login=%s&password=%s&code=%s"
const userloginAPIFormat = "/api/v2/user/login?login=%s&password=%s"

// https://rapidgator.net/article/api/user#info
const userInfoAPIFormat = "/api/v2/user/info?token=%s"

type LoginInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

const (
	testLogin    = "40779007@qq.com"
	testPassword = "4077rapid"
)

type CurrentAccount struct {
	Account         string
	Password        string
	Token           string
	LastUpdatedTime time.Time
	Expiration      time.Time
	Mux             sync.Mutex
}

// var currentAccount = CurrentAccount{Account: testLogin, Password: testPassword}
var currentAccount = CurrentAccount{
	Account:    "",
	Password:   "",
	Token:      "",
	Expiration: time.Now(),
}

// 4 hours;
// it is an empirical value.
const tokenExpiresIn time.Duration = 4 * 60 * 60 * time.Second

func getAccount() (account string, err error) {
	currentAccount.Mux.Lock()
	defer currentAccount.Mux.Unlock()

	return currentAccount.Account, nil
}

func setAccount(account string, password string) (err error) {
	currentAccount.Mux.Lock()
	defer currentAccount.Mux.Unlock()

	if (account == "") || (password == "") {
		return fmt.Errorf("RapidGator setAccount err: empty account info")
	}

	currentAccount.Account = account
	currentAccount.Password = password
	currentAccount.Token = ""

	newToken, err := updateToken(account, password)
	if err != nil {
		return fmt.Errorf("RapidGator updateToken err: %w", err)
	}

	currentAccount.Token = newToken
	currentAccount.Expiration = time.Now().Add(tokenExpiresIn)

	return nil
}

func getToken() (token string, err error) {
	currentAccount.Mux.Lock()
	defer currentAccount.Mux.Unlock()

	if currentAccount.Expiration.After(time.Now()) && currentAccount.Token != "" {
		return currentAccount.Token, nil
	}

	// maybe don't set account info yet.
	if (currentAccount.Account == "") || (currentAccount.Password == "") {
		return "", fmt.Errorf("RapidGator getToken err: empty account info.")
	}

	newToken, err := updateToken(currentAccount.Account, currentAccount.Password)
	if err != nil {
		return "", fmt.Errorf("RapidGator updateToken err: %w", err)
	}

	currentAccount.Token = newToken
	currentAccount.Expiration = time.Now().Add(tokenExpiresIn)

	return currentAccount.Token, nil
}

// if has one existed token, which operation is more efficient, "login" or "info"?
func updateToken(account string, password string) (token string, err error) {
	var loginResponse struct {
		Status   int `json:"status"`
		Response struct {
			Token string `json:"token"`
		} `json:"response"`
	}

	userLoginURL := baseURL + fmt.Sprintf(userloginAPIFormat, account, password)

	// resp, err := client.R().
	resp, respErr := restyClient.R().
		SetResult(&loginResponse).
		Post(userLoginURL)

	if respErr != nil {
		return "", fmt.Errorf("RapidGator login err: %w", err)
	}

	if loginResponse.Status != http.StatusOK {
		return "", fmt.Errorf("RapidGator updateToken loginResponse.Status %d != http.StatusOK", loginResponse.Status)
	}

	token = loginResponse.Response.Token
	if token == "" {
		return "", fmt.Errorf("RapidGator updateToken token is empty: %s", resp)
	}

	return token, nil
}

func getUserInfo(token string) (userInfo string, err error) {
	userInfoURL := baseURL + fmt.Sprintf(userInfoAPIFormat, token)

	resp, err := restyClient.R().
		Post(userInfoURL)

	if err != nil {
		return "", fmt.Errorf("RapidGator getUserInfo err: %s", resp)
	}

	userInfo = resp.String()

	return userInfo, nil
}
