package controller



import (

	"github.com/gin-gonic/gin"

	"net/http"

	"sync/atomic"

)





var usersLoginInfo = map[string]User{}



var userIdSequence = int64(1)



type UserLoginResponse struct {

	Response

	UserId int64  `json:"user_id,omitempty"`

	Token  string `json:"token"`

}



type UserResponse struct {

	Response

	User User `json:"user"`

}



func Register(c *gin.Context) {

	username := c.Query("username")

	password := c.Query("password")



	token := username + password



	if _, exist := usersLoginInfo[token]; exist {

		c.JSON(http.StatusOK, UserLoginResponse{

			Response: Response{StatusCode: 1, StatusMsg: "用户已存在"},

		})

	} else {

		atomic.AddInt64(&userIdSequence, 1)

		newUser := User{

			Id:   userIdSequence,

			Name: username,

		}

		usersLoginInfo[token] = newUser

		c.JSON(http.StatusOK, UserLoginResponse{

			Response: Response{StatusCode: 0},

			UserId:   userIdSequence,

			Token:    username + password,

		})

	}

}



func Login(c *gin.Context) {

	username := c.Query("username")

	password := c.Query("password")



	token := username + password



	if user, exist := usersLoginInfo[token]; exist {

		c.JSON(http.StatusOK, UserLoginResponse{

			Response: Response{StatusCode: 0},

			UserId:   user.Id,

			Token:    token,

		})

	} else {

		c.JSON(http.StatusOK, UserLoginResponse{

			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},

		})

	}

}



func UserInfo(c *gin.Context) {

	token := c.Query("token")



	if user, exist := usersLoginInfo[token]; exist {

		c.JSON(http.StatusOK, UserResponse{

			Response: Response{StatusCode: 0},

			User:     user,

		})

	} else {

		c.JSON(http.StatusOK, UserResponse{

			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})

	}

}

