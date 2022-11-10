package login

import (
	"errors"
	"testing"

	"github.com/credativ/plutono/pkg/models"
	"github.com/credativ/plutono/pkg/services/ldap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticateUser(t *testing.T) {
	authScenario(t, "When a user authenticates without setting a password", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(nil, sc)
		mockLoginUsingLDAP(false, nil, sc)

		loginQuery := models.LoginUserQuery{
			Username: "user",
			Password: "",
		}
		err := authenticateUser(&loginQuery)

		require.EqualError(t, err, ErrPasswordEmpty.Error())
		assert.False(t, sc.plutonoLoginWasCalled)
		assert.False(t, sc.ldapLoginWasCalled)
		assert.Empty(t, sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When a user authenticates having too many login attempts", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(ErrTooManyLoginAttempts, sc)
		mockLoginUsingPlutonoDB(nil, sc)
		mockLoginUsingLDAP(true, nil, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, ErrTooManyLoginAttempts.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.False(t, sc.plutonoLoginWasCalled)
		assert.False(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Empty(t, sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When plutono user authenticate with valid credentials", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(nil, sc)
		mockLoginUsingLDAP(true, ErrInvalidCredentials, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.NoError(t, err)
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.False(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Equal(t, "plutono", sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When plutono user authenticate and unexpected error occurs", func(sc *authScenarioContext) {
		customErr := errors.New("custom")
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(customErr, sc)
		mockLoginUsingLDAP(true, ErrInvalidCredentials, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, customErr.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.False(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Equal(t, "plutono", sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When a non-existing plutono user authenticate and ldap disabled", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(models.ErrUserNotFound, sc)
		mockLoginUsingLDAP(false, nil, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, models.ErrUserNotFound.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.True(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Empty(t, sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When a non-existing plutono user authenticate and invalid ldap credentials", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(models.ErrUserNotFound, sc)
		mockLoginUsingLDAP(true, ldap.ErrInvalidCredentials, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, ErrInvalidCredentials.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.True(t, sc.ldapLoginWasCalled)
		assert.True(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Equal(t, "ldap", sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When a non-existing plutono user authenticate and valid ldap credentials", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(models.ErrUserNotFound, sc)
		mockLoginUsingLDAP(true, nil, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.NoError(t, err)
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.True(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Equal(t, "ldap", sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When a non-existing plutono user authenticate and ldap returns unexpected error", func(sc *authScenarioContext) {
		customErr := errors.New("custom")
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(models.ErrUserNotFound, sc)
		mockLoginUsingLDAP(true, customErr, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, customErr.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.True(t, sc.ldapLoginWasCalled)
		assert.False(t, sc.saveInvalidLoginAttemptWasCalled)
		assert.Equal(t, "ldap", sc.loginUserQuery.AuthModule)
	})

	authScenario(t, "When plutono user authenticate with invalid credentials and invalid ldap credentials", func(sc *authScenarioContext) {
		mockLoginAttemptValidation(nil, sc)
		mockLoginUsingPlutonoDB(ErrInvalidCredentials, sc)
		mockLoginUsingLDAP(true, ldap.ErrInvalidCredentials, sc)
		mockSaveInvalidLoginAttempt(sc)

		err := authenticateUser(sc.loginUserQuery)

		require.EqualError(t, err, ErrInvalidCredentials.Error())
		assert.True(t, sc.loginAttemptValidationWasCalled)
		assert.True(t, sc.plutonoLoginWasCalled)
		assert.True(t, sc.ldapLoginWasCalled)
		assert.True(t, sc.saveInvalidLoginAttemptWasCalled)
	})
}

type authScenarioContext struct {
	loginUserQuery                   *models.LoginUserQuery
	plutonoLoginWasCalled            bool
	ldapLoginWasCalled               bool
	loginAttemptValidationWasCalled  bool
	saveInvalidLoginAttemptWasCalled bool
}

type authScenarioFunc func(sc *authScenarioContext)

func mockLoginUsingPlutonoDB(err error, sc *authScenarioContext) {
	loginUsingPlutonoDB = func(query *models.LoginUserQuery) error {
		sc.plutonoLoginWasCalled = true
		return err
	}
}

func mockLoginUsingLDAP(enabled bool, err error, sc *authScenarioContext) {
	loginUsingLDAP = func(query *models.LoginUserQuery) (bool, error) {
		sc.ldapLoginWasCalled = true
		return enabled, err
	}
}

func mockLoginAttemptValidation(err error, sc *authScenarioContext) {
	validateLoginAttempts = func(*models.LoginUserQuery) error {
		sc.loginAttemptValidationWasCalled = true
		return err
	}
}

func mockSaveInvalidLoginAttempt(sc *authScenarioContext) {
	saveInvalidLoginAttempt = func(query *models.LoginUserQuery) error {
		sc.saveInvalidLoginAttemptWasCalled = true
		return nil
	}
}

func authScenario(t *testing.T, desc string, fn authScenarioFunc) {
	t.Helper()

	t.Run(desc, func(t *testing.T) {
		origLoginUsingPlutonoDB := loginUsingPlutonoDB
		origLoginUsingLDAP := loginUsingLDAP
		origValidateLoginAttempts := validateLoginAttempts
		origSaveInvalidLoginAttempt := saveInvalidLoginAttempt

		sc := &authScenarioContext{
			loginUserQuery: &models.LoginUserQuery{
				Username:  "user",
				Password:  "pwd",
				IpAddress: "192.168.1.1:56433",
			},
		}

		t.Cleanup(func() {
			loginUsingPlutonoDB = origLoginUsingPlutonoDB
			loginUsingLDAP = origLoginUsingLDAP
			validateLoginAttempts = origValidateLoginAttempts
			saveInvalidLoginAttempt = origSaveInvalidLoginAttempt
		})

		fn(sc)
	})
}
