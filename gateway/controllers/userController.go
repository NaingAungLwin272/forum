package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	user_service "github.com/scm-dev1dev5/mtm-community-forum/gateway/services/user"
)

type UserController struct {
	UserServiceInterface user_service.UserServiceInterface
}

//Auth Process

func (controller *UserController) Login(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) ForgetPassword(ctx *gin.Context) {

	data, err := controller.UserServiceInterface.ForgetPassword(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) ResetPassword(ctx *gin.Context) {
	log.Println(ctx.Request.Body, "ctx for reset password")

	data, err := controller.UserServiceInterface.ResetPassword(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) ChangePassword(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.ChangePasswrd(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// User Processes
func (controller *UserController) CreateUser(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *UserController) GetUser(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	data := controller.UserServiceInterface.GetUsers(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, &data)
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.UpdateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.DeleteUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) CreateUserWithCsv(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.CreateUserWithCsv(ctx)
	fmt.Println(err, "error..........")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetUserSummary(ctx *gin.Context) {
	data := controller.UserServiceInterface.GetUserSummary(ctx)
	ctx.JSON(http.StatusCreated, &data)
}

func (controller *UserController) FilterUser(ctx *gin.Context) {
	data := controller.UserServiceInterface.FilterUser(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, &data)
}

func (controller *UserController) GetUserCount(ctx *gin.Context) {
	data := controller.UserServiceInterface.GetUserCount(ctx)
	ctx.JSON(http.StatusCreated, &data)
}

// Team Processes
func (controller *UserController) CreateTeam(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.CreateTeam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetTeam(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetTeam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetTeams(ctx *gin.Context) {
	data := controller.UserServiceInterface.GetTeams(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, &data)
}

func (controller *UserController) GetTeamCount(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetTeamCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) UpdateTeam(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.UpdateTeam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) DeleteTeam(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.DeleteTeam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetTeamsByDepartmentId(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetTeamsByDepartmentId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

// Department Processes
func (controller *UserController) CreateDepartment(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.CreateDepartment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *UserController) GetDepartment(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetDepartment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetDepartments(ctx *gin.Context) {
	data := controller.UserServiceInterface.GetDepartments(ctx)
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, &data)
}

func (controller *UserController) UpdateDepartment(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.UpdateDepartment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &data)
	}
}

func (controller *UserController) DeleteDepartment(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.DeleteDepartment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetDepartmentCount(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetDepartmentCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func (controller *UserController) GetUserByDisplayName(ctx *gin.Context) {
	data, err := controller.UserServiceInterface.GetUserByDisplayName(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, &data)
	}
}

func NewUserController(UserServiceInterface user_service.UserServiceInterface) *UserController {
	return &UserController{
		UserServiceInterface: UserServiceInterface,
	}
}
