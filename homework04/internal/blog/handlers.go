package blog

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type registerRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type postRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type commentRequest struct {
	Content string `json:"content" binding:"required"`
}

func (a *App) handleRegister(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.abortError(c, http.StatusBadRequest, "invalid register payload")
		return
	}

	hash, err := hashPassword(req.Password)
	if err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to hash password")
		return
	}

	u := User{Username: req.Username, PasswordHash: hash, Email: req.Email}
	if err := a.DB.Create(&u).Error; err != nil {
		a.abortError(c, http.StatusConflict, "username or email already exists")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": u.ID, "username": u.Username, "email": u.Email})
}

func (a *App) handleLogin(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.abortError(c, http.StatusBadRequest, "invalid login payload")
		return
	}

	var u User
	if err := a.DB.Where("username = ?", req.Username).First(&u).Error; err != nil {
		a.abortError(c, http.StatusUnauthorized, "invalid username or password")
		return
	}
	if err := comparePassword(u.PasswordHash, req.Password); err != nil {
		a.abortError(c, http.StatusUnauthorized, "invalid username or password")
		return
	}

	token, err := a.newToken(u)
	if err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *App) handleCreatePost(c *gin.Context) {
	var req postRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.abortError(c, http.StatusBadRequest, "invalid post payload")
		return
	}

	post := Post{Title: req.Title, Content: req.Content, UserID: userIDFromContext(c)}
	if err := a.DB.Create(&post).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to create post")
		return
	}

	if err := a.DB.Preload("User").First(&post, post.ID).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to load created post")
		return
	}

	c.JSON(http.StatusCreated, toPostResponse(post))
}

func (a *App) handleListPosts(c *gin.Context) {
	var posts []Post
	if err := a.DB.Preload("User").Order("id asc").Find(&posts).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to list posts")
		return
	}

	resp := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		resp = append(resp, toPostResponse(p))
	}
	c.JSON(http.StatusOK, resp)
}

func (a *App) handleGetPost(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		a.abortError(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post Post
	if err := a.DB.Preload("User").First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			a.abortError(c, http.StatusNotFound, "post not found")
			return
		}
		a.abortError(c, http.StatusInternalServerError, "failed to get post")
		return
	}

	c.JSON(http.StatusOK, toPostResponse(post))
}

func (a *App) handleUpdatePost(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		a.abortError(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post Post
	if err := a.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			a.abortError(c, http.StatusNotFound, "post not found")
			return
		}
		a.abortError(c, http.StatusInternalServerError, "failed to get post")
		return
	}

	if post.UserID != userIDFromContext(c) {
		a.abortError(c, http.StatusForbidden, "no permission to update this post")
		return
	}

	var req postRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.abortError(c, http.StatusBadRequest, "invalid post payload")
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	if err := a.DB.Save(&post).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to update post")
		return
	}

	if err := a.DB.Preload("User").First(&post, post.ID).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to load updated post")
		return
	}

	c.JSON(http.StatusOK, toPostResponse(post))
}

func (a *App) handleDeletePost(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		a.abortError(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post Post
	if err := a.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			a.abortError(c, http.StatusNotFound, "post not found")
			return
		}
		a.abortError(c, http.StatusInternalServerError, "failed to get post")
		return
	}

	if post.UserID != userIDFromContext(c) {
		a.abortError(c, http.StatusForbidden, "no permission to delete this post")
		return
	}

	if err := a.DB.Delete(&post).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to delete post")
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *App) handleCreateComment(c *gin.Context) {
	postID, ok := parseIDParam(c, "id")
	if !ok {
		a.abortError(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post Post
	if err := a.DB.First(&post, postID).Error; err != nil {
		a.abortError(c, http.StatusNotFound, "post not found")
		return
	}

	var req commentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.abortError(c, http.StatusBadRequest, "invalid comment payload")
		return
	}

	comment := Comment{Content: req.Content, UserID: userIDFromContext(c), PostID: postID}
	if err := a.DB.Create(&comment).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to create comment")
		return
	}

	if err := a.DB.Preload("User").First(&comment, comment.ID).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to load created comment")
		return
	}

	c.JSON(http.StatusCreated, toCommentResponse(comment))
}

func (a *App) handleListComments(c *gin.Context) {
	postID, ok := parseIDParam(c, "id")
	if !ok {
		a.abortError(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post Post
	if err := a.DB.First(&post, postID).Error; err != nil {
		a.abortError(c, http.StatusNotFound, "post not found")
		return
	}

	var comments []Comment
	if err := a.DB.Preload("User").Where("post_id = ?", postID).Order("id asc").Find(&comments).Error; err != nil {
		a.abortError(c, http.StatusInternalServerError, "failed to list comments")
		return
	}

	resp := make([]CommentResponse, 0, len(comments))
	for _, cm := range comments {
		resp = append(resp, toCommentResponse(cm))
	}
	c.JSON(http.StatusOK, resp)
}
