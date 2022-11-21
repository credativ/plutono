package login

import (
	"testing"

	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginUsingPlutonoDB(t *testing.T) {
	plutonoLoginScenario(t, "When login with non-existing user", func(sc *plutonoLoginScenarioContext) {
		sc.withNonExistingUser()
		err := loginUsingPlutonoDB(sc.loginUserQuery)
		require.EqualError(t, err, models.ErrUserNotFound.Error())

		assert.False(t, sc.validatePasswordCalled)
		assert.Nil(t, sc.loginUserQuery.User)
	})

	plutonoLoginScenario(t, "When login with invalid credentials", func(sc *plutonoLoginScenarioContext) {
		sc.withInvalidPassword()
		err := loginUsingPlutonoDB(sc.loginUserQuery)

		require.EqualError(t, err, ErrInvalidCredentials.Error())

		assert.True(t, sc.validatePasswordCalled)
		assert.Nil(t, sc.loginUserQuery.User)
	})

	plutonoLoginScenario(t, "When login with valid credentials", func(sc *plutonoLoginScenarioContext) {
		sc.withValidCredentials()
		err := loginUsingPlutonoDB(sc.loginUserQuery)
		require.NoError(t, err)

		assert.True(t, sc.validatePasswordCalled)

		require.NotNil(t, sc.loginUserQuery.User)
		assert.Equal(t, sc.loginUserQuery.Username, sc.loginUserQuery.User.Login)
		assert.Equal(t, sc.loginUserQuery.Password, sc.loginUserQuery.User.Password)
	})

	plutonoLoginScenario(t, "When login with disabled user", func(sc *plutonoLoginScenarioContext) {
		sc.withDisabledUser()
		err := loginUsingPlutonoDB(sc.loginUserQuery)
		require.EqualError(t, err, ErrUserDisabled.Error())

		assert.False(t, sc.validatePasswordCalled)
		assert.Nil(t, sc.loginUserQuery.User)
	})
}

type plutonoLoginScenarioContext struct {
	loginUserQuery         *models.LoginUserQuery
	validatePasswordCalled bool
}

type plutonoLoginScenarioFunc func(c *plutonoLoginScenarioContext)

func plutonoLoginScenario(t *testing.T, desc string, fn plutonoLoginScenarioFunc) {
	t.Helper()

	t.Run(desc, func(t *testing.T) {
		origValidatePassword := validatePassword

		sc := &plutonoLoginScenarioContext{
			loginUserQuery: &models.LoginUserQuery{
				Username:  "user",
				Password:  "pwd",
				IpAddress: "192.168.1.1:56433",
			},
			validatePasswordCalled: false,
		}

		t.Cleanup(func() {
			validatePassword = origValidatePassword
		})

		fn(sc)
	})
}

func mockPasswordValidation(valid bool, sc *plutonoLoginScenarioContext) {
	validatePassword = func(providedPassword string, userPassword string, userSalt string) error {
		sc.validatePasswordCalled = true

		if !valid {
			return ErrInvalidCredentials
		}

		return nil
	}
}

func (sc *plutonoLoginScenarioContext) getUserByLoginQueryReturns(user *models.User) {
	bus.AddHandler("test", func(query *models.GetUserByLoginQuery) error {
		if user == nil {
			return models.ErrUserNotFound
		}

		query.Result = user
		return nil
	})
}

func (sc *plutonoLoginScenarioContext) withValidCredentials() {
	sc.getUserByLoginQueryReturns(&models.User{
		Id:       1,
		Login:    sc.loginUserQuery.Username,
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(true, sc)
}

func (sc *plutonoLoginScenarioContext) withNonExistingUser() {
	sc.getUserByLoginQueryReturns(nil)
}

func (sc *plutonoLoginScenarioContext) withInvalidPassword() {
	sc.getUserByLoginQueryReturns(&models.User{
		Password: sc.loginUserQuery.Password,
		Salt:     "salt",
	})
	mockPasswordValidation(false, sc)
}

func (sc *plutonoLoginScenarioContext) withDisabledUser() {
	sc.getUserByLoginQueryReturns(&models.User{
		IsDisabled: true,
	})
}
