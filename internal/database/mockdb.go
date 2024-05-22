package database

import (
	"errors"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
    "jim": {
        AuthToken: "123XYZ",
        Username:  "jim",
    },
    "fran": {
        AuthToken: "091IOU",
        Username:  "fran",
    },
    "lee": {
        AuthToken: "878WXW",
        Username:  "lee",
    },
}

var mockAccountDetails = map[string]AccountDetails{
    "jim": {
        Username:  "jim",
        Age:       25,
    },
    "fran": {
        Username:  "fran",
        Age:       24,
    },
    "lee": {
        Username:  "lee",
        Age:       31,
    },
}

func (d *mockDB) InsertLoginDetails(username string, authToken string) *LoginDetails {
    time.Sleep(time.Second * 1)
    var clientData = LoginDetails{
        AuthToken: authToken,
        Username: username,
    }
    clientData, ok := mockLoginDetails[username]
    if ok {
        return nil
    }
    mockLoginDetails[username] = clientData
    return &clientData
}

func (d *mockDB) GetLoginDetails(username string) *LoginDetails {
    time.Sleep(time.Second * 1)

    var clientData = LoginDetails{}
    clientData, ok := mockLoginDetails[username]
    if !ok {
        return nil
    }
    return &clientData
}

func (d *mockDB) UpdateAuthToken(username string, authToken string) {
    time.Sleep(time.Second * 1)

    clientData, ok := mockLoginDetails[username]
    if ok {
        clientData.AuthToken = authToken
        mockLoginDetails[username] = clientData
    }
}

func (d *mockDB) DeleteAuthToken(username string) error {
    time.Sleep(time.Second * 1)
    delete(mockLoginDetails, username)

    return nil
}

func (d *mockDB) GetAccountDetails(username string) *AccountDetails {
    time.Sleep(time.Second * 1)

    var clientData = AccountDetails{}
    clientData, ok := mockAccountDetails[username]
    if !ok {
        return nil
    }
    return &clientData
}

func (d *mockDB) UpdateAccountDetails(username string, account AccountDetails) (*AccountDetails, error) {
    time.Sleep(time.Second * 1)

    _, ok := mockAccountDetails[username]
    if !ok {
        return nil, errors.New("User not found.")
    }
    delete(mockAccountDetails, username)
    mockAccountDetails[account.Username] = account

    return &account, nil
}

func (d *mockDB) DeleteAccount(username string) error {
    time.Sleep(time.Second * 1)

    delete(mockAccountDetails, username)

    return nil
}

