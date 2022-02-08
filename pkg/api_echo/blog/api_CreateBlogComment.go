package blog

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/blog"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateBlogComment(c echo.Context) (*model.BlogComment, error) {
	comment := model.BlogComment{}
	if error := c.Bind(&comment); error != nil {
		return nil, error
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result, err := op.BlogCommentCreate(&comment, claims.Sub)
	return result, err
}
