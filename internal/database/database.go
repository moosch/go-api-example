package database

type LoginDetails struct {
    AuthToken string
    Username string
}

type AccountDetails struct {
    Username string
    Age      int
}

type DatabaseInterface interface {
    InsertLoginDetails(username string, authToken string) *LoginDetails
    GetLoginDetails(username string) *LoginDetails
    UpdateAuthToken(username string, authToken string)
    DeleteAuthToken(username string) error

    // InsertAccountDetails(username string, authToken string) *AccountDetails
    GetAccountDetails(username string) *AccountDetails
    UpdateAccountDetails(username string, account AccountDetails) (*AccountDetails, error)
    DeleteAccount(username string) error
}

var dbConn DatabaseInterface

func GetDatabaseConnection() (*DatabaseInterface, error) {
    if dbConn != nil {
        return &dbConn, nil
    }

    dbConn = &mockDB{}
    return &dbConn, nil
}

