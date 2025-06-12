package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/GurramKarimunisa/go-user-api/internal/logger" // Corrected import path
	"github.com/GurramKarimunisa/go-user-api/internal/service" // Corrected import path
	"go.uber.org/zap" // RE-ADDED: For creating zap.Field types like zap.Error(err)
)

type UserHandler struct {
	userService *service.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

// CreateUserRequest defines the request body for creating a user
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"` //YYYY-MM-DD
}

// UpdateUserRequest defines the request body for updating a user
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		logger.Log.Error("Failed to parse request body", zap.Error(err)) // Corrected: Use zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(req); err != nil {
		logger.Log.Warn("Validation error on CreateUser", zap.Error(err)) // Corrected: Use zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dobTime, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		logger.Log.Error("Failed to parse DOB", zap.Error(err)) // Corrected: Use zap.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format for dob. Expected YYYY-MM-DD."})
	}

	user, err := h.userService.CreateUser(context.Background(), req.Name, dobTime)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Int32("userID", user.ID), zap.Error(err)) // Corrected: Use zap.Int32 and zap.Error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	logger.Log.Info("User created successfully", zap.Int32("userID", user.ID)) // Corrected: Use zap.Int32
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logger.Log.Warn("Invalid user ID format", zap.String("id", idStr)) // Corrected: Use zap.String
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetUserByID(context.Background(), int32(id))
	if err != nil {
		logger.Log.Error("Failed to get user by ID", zap.Int32("userID", int32(id)), zap.Error(err)) // Corrected: Use zap.Int32 and zap.Error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	logger.Log.Info("User retrieved successfully", zap.Int32("userID", user.ID)) // Corrected: Use zap.Int32
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers(context.Background())
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err)) // Corrected: Use zap.Error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}

	logger.Log.Info("Listed all users successfully", zap.Int("count", len(users))) // Corrected: Use zap.Int
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logger.Log.Warn("Invalid user ID format for update", zap.String("id", idStr)) // Corrected: Use zap.String
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	req := new(UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		logger.Log.Error("Failed to parse request body for update", zap.Error(err)) // Corrected: Use zap.Error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(req); err != nil {
		logger.Log.Warn("Validation error on UpdateUser", zap.Error(err)) // Corrected: Use zap.Error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dobTime, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		logger.Log.Error("Failed to parse DOB for update", zap.Error(err)) // Corrected: Use zap.Error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format for dob. Expected YYYY-MM-DD."})
	}

	user, err := h.userService.UpdateUser(context.Background(), int32(id), req.Name, dobTime)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Int32("userID", int32(id)), zap.Error(err)) // Corrected: Use zap.Int32 and zap.Error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	logger.Log.Info("User updated successfully", zap.Int32("userID", user.ID)) // Corrected: Use zap.Int32
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logger.Log.Warn("Invalid user ID format for delete", zap.String("id", idStr)) // Corrected: Use zap.String
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.userService.DeleteUser(context.Background(), int32(id))
	if err != nil {
		logger.Log.Error("Failed to delete user", zap.Int32("userID", int32(id)), zap.Error(err)) // Corrected: Use zap.Int32 and zap.Error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	logger.Log.Info("User deleted successfully", zap.Int32("userID", int32(id))) // Corrected: Use zap.Int32
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}
