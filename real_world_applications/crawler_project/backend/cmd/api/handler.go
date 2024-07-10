package main

import (
	"net/http"
	"web_crawler/models"
	"web_crawler/utils"

	"github.com/labstack/echo/v4"
)

// pingHandler handles the ping request and returns a success message.
func (app *Config) pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "The system is working fine")
}

func (app *Config) signupHandler(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		app.Logger.Error("error binding user data%s", err)
		return c.String(http.StatusBadRequest, "Invalid user data")
	}
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return c.String(http.StatusBadRequest, "Email, password and name are required")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		app.Logger.Error("error hashing password: %s", err)
		return c.String(http.StatusInternalServerError, "Error creating user")

	}

	user.Password = hashedPassword
	if err := user.Insert(); err != nil {
		app.Logger.Error("error inserting user: %v", err)
		return c.String(http.StatusInternalServerError, "Error creating user")

	}

	return c.JSON(http.StatusCreated, "user created successfully")

}

func (app *Config) loginHandler(c echo.Context) error {
	type LoginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var credentials LoginCredentials
	if err := c.Bind(&credentials); err != nil {
		app.Logger.Error("error binding login credentials: %v", err)
		return c.String(http.StatusBadRequest, "Invalid login credentials")

	}

	if credentials.Email == "" || credentials.Password == "" {
		return c.String(http.StatusBadRequest, "Email and password are required")
	}
	user, err := app.Model.User.GetUserByEmail(credentials.Email)
	if err != nil {
		app.Logger.Error("error getting user by email: %v", err)

		if err.Error() == "user not found" {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error getting user")

	}
	if err := utils.ComparePasswords(user.Password, credentials.Password); err != nil {
		app.Logger.Error("error comparing passwords: %v", err)
		return c.String(http.StatusUnauthorized, "Wrong password")

	}
	token, err := utils.GenerateJWT(user.ID, user.Username)

	if err != nil {
		app.Logger.Error("error generating JWT: %v", err)
		return c.String(http.StatusInternalServerError, "Error generating JWT")

	}
	return c.JSON(http.StatusOK, map[string]string{"token": token, "message": "Login successful"})
}

func (app *Config) getUserHandler(c echo.Context) error {
	userId := c.Get("userID").(string)
	user, err := app.Model.User.GetUserByID(userId)
	if err != nil {
		app.Logger.Error("error getting user by id: %v", err)
		if err.Error() == "user not found" {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error getting user")

	}
	return c.JSON(http.StatusOK, user)
}

func (app *Config) updateUserHandler(c echo.Context) error {
	var user models.User
	userId := c.Get("userID").(string) // Assuming userID is correctly retrieved and casted
	if err := c.Bind(&user); err != nil {
		app.Logger.Error("error binding user data: %v", err)
		return c.String(http.StatusBadRequest, "Invalid user data")
	}

	updateResult, err := user.Update(userId) // Update function now returns *mongo.UpdateResult and error
	if err != nil {
		app.Logger.Error("error updating user: %v", err)
		if err.Error() == "user not found" {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error updating user")
	}

	app.Logger.Info("update operation result: matched %v, modified %v", updateResult.MatchedCount, updateResult.ModifiedCount)

	return c.JSON(http.StatusOK, "username updated successfully")
}

func (app *Config) deleteUserHandler(c echo.Context) error {
	var user models.User
	userId := c.Get("userID").(string)
	if err := c.Bind(&user); err != nil {
		app.Logger.Error("error binding user data: %v", err)
		return c.String(http.StatusBadRequest, "Invalid user data")
	}
	if err := user.Delete(userId); err != nil {
		app.Logger.Error("error deleting user: %v", err)
		if err.Error() == "user not found" {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error deleting user")
	}
	return c.String(http.StatusOK, "User deleted successfully")
}
