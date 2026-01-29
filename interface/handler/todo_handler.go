package handler

import (
	"go-practice/usecase/todo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TodoHandler はTodoのHTTPハンドラ
type TodoHandler struct {
	createUseCase   *todo.CreateTodoUseCase
	listUseCase     *todo.ListTodosUseCase
	completeUseCase *todo.CompleteTodoUseCase
	deleteUseCase   *todo.DeleteTodoUseCase
}

// NewTodoHandler はTodoHandlerのコンストラクタ
func NewTodoHandler(
	createUseCase *todo.CreateTodoUseCase,
	listUseCase *todo.ListTodosUseCase,
	completeUseCase *todo.CompleteTodoUseCase,
	deleteUseCase *todo.DeleteTodoUseCase,
) *TodoHandler {
	return &TodoHandler{
		createUseCase:   createUseCase,
		listUseCase:     listUseCase,
		completeUseCase: completeUseCase,
		deleteUseCase:   deleteUseCase,
	}
}

// CreateTodoRequest はTodo作成リクエスト
type CreateTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

// CreateTodoResponse はTodo作成レスポンス
type CreateTodoResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// CreateTodo はTodo作成エンドポイント
// POST /api/todos
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest

	// リクエストボディをバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Usecaseを実行
	output, err := h.createUseCase.Execute(c.Request.Context(), todo.CreateTodoInput{
		Title:     req.Title,
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// レスポンスを返す
	c.JSON(http.StatusCreated, CreateTodoResponse{
		ID:    output.ID,
		Title: output.Title,
	})
}

// TodoResponse はTodo表示用レスポンス
type TodoResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Completed   bool    `json:"completed"`
	CreatedAt   string  `json:"created_at"`
	CompletedAt *string `json:"completed_at"`
}

// ListTodos はTodo一覧取得エンドポイント
// GET /api/todos
func (h *TodoHandler) ListTodos(c *gin.Context) {
	// Usecaseを実行
	dtos, err := h.listUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch todos",
		})
		return
	}

	// DTOをレスポンスに変換
	responses := make([]TodoResponse, 0, len(dtos))
	for _, dto := range dtos {
		var completedAt *string
		if dto.CompletedAt != nil {
			formatted := dto.CompletedAt.Format("2006-01-02T15:04:05Z07:00")
			completedAt = &formatted
		}

		responses = append(responses, TodoResponse{
			ID:          dto.ID,
			Title:       dto.Title,
			Completed:   dto.Completed,
			CreatedAt:   dto.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			CompletedAt: completedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": responses,
	})
}

// CompleteTodo はTodo完了エンドポイント
// PUT /api/todos/:id/complete
func (h *TodoHandler) CompleteTodo(c *gin.Context) {
	id := c.Param("id")

	// Usecaseを実行
	err := h.completeUseCase.Execute(c.Request.Context(), todo.CompleteTodoInput{
		ID:          id,
		CompletedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo completed successfully",
	})
}

// DeleteTodo はTodo削除エンドポイント
// DELETE /api/todos/:id
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// Usecaseを実行
	err := h.deleteUseCase.Execute(c.Request.Context(), todo.DeleteTodoInput{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}
