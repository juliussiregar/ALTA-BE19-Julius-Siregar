package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUserHandler
func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	var input = new(User)
	if err := c.Bind(input); err != nil {
		h.logErrorInternal("input error:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input data kurang tepat",
			"data":    nil,
		})
	}

	if err := h.userService.CreateUser(input); err != nil {
		h.logErrorInternal("failed to create user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "terdapat permasalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    input,
	})
}

// GetAllUsersHandler
func (h *UserHandler) GetAllUsersHandler(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		h.logErrorInternal("failed to get all users:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "terdapat permasalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

// GetUserByIDHandler
func (h *UserHandler) GetUserByIDHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		h.logErrorInternal("failed to get user by ID:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "terdapat permasalahan pada pengolahan data",
			"data":    nil,
		})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

// UpdateUserHandler
func (h *UserHandler) UpdateUserHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	var updatedUser = new(User)
	if err := c.Bind(updatedUser); err != nil {
		h.logErrorInternal("input error:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "input data kurang tepat",
			"data":    nil,
		})
	}

	if err := h.userService.UpdateUser(userID, updatedUser); err != nil {
		h.logErrorInternal("failed to update user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "terdapat permasalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    updatedUser,
	})
}

// DeleteUserHandler
func (h *UserHandler) DeleteUserHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		h.logErrorInternal("failed to delete user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "terdapat permasalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success delete user with ID %d", userID),
	})
}

func (h *UserHandler) logErrorInternal(message string, err error) {
	h.userService.logError(message, err)
}
