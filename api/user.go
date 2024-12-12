package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"

	"github.com/SabariGanesh-K/prod-mgm-go/util"
)

type createUserRequest struct {
	UserID         string `json:"user_id"`
	Password string `json:"password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}
	
type userResponse struct {
	UserID         string `json:"user_id"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	FilesOwned        []string  `json:"files_owned"`
}

func newUserResponse(user db.Users) userResponse {
	return userResponse{
		UserID:          user.UserID,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
		FilesOwned:        user.FilesOwned,

	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	fmt.Print(req.Password)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Printf(req.Password)
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		UserID:          req.UserID,

		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,

	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	UserID string `json:"user_id" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	

	rsp := loginUserResponse{
	
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
