package blog

import (
	"blockmedic/pkg/env"
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/blog"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
)

func DeleteBlogComment(c echo.Context) (*model.BlogComment, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)
	blogCommentId, _ := strconv.ParseInt(id, 10, 64)
	result, err := op.BlogCommentDelete(blogCommentId, claims.Sub, claims.Auth)
	return result, err
}
