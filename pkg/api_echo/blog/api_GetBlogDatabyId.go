package blog

import (
	"blockmedic/pkg/model"
	op "blockmedic/pkg/operation/blog"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetBlogPostbyId(c echo.Context) (*model.Blog, error) {
	id := c.Param("id")
	blogId, _ := strconv.ParseInt(id, 10, 64)
	result, err := op.BlogPostbyId(blogId)
	return result, err
}
