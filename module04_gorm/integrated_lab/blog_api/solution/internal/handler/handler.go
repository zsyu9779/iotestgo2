package handler

import (
	"errors"
	"net/http"
	"strconv"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct{ service service.PostService }

func New(postService service.PostService) *Handler { return &Handler{service: postService} }

type createRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Comment string   `json:"comment"`
	Tags    []string `json:"tags"`
}

func (h *Handler) List(c *gin.Context) {
	posts, err := h.service.ListPosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) Create(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := h.service.CreatePost(c.Request.Context(), request.Title, request.Content, request.Tags)
	h.writeCreate(c, post, err)
}

func (h *Handler) CreateWithComment(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := h.service.CreatePostWithComment(c.Request.Context(), request.Title, request.Content, request.Comment, request.Tags)
	h.writeCreate(c, post, err)
}

func (h *Handler) writeCreate(c *gin.Context, post *service.PostResponse, err error) {
	if errors.Is(err, service.ErrInvalidInput) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (h *Handler) Delete(c *gin.Context) {
	value, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	err = h.service.DeletePost(c.Request.Context(), uint(value))
	if errors.Is(err, service.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": service.ErrNotFound.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func Router(postService service.PostService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	h := New(postService)
	router.GET("/posts", h.List)
	router.POST("/posts", h.Create)
	router.POST("/posts/with-comment", h.CreateWithComment)
	router.DELETE("/posts/:id", h.Delete)
	return router
}
