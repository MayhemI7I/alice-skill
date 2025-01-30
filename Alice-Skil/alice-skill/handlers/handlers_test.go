package handlers

import (
    "bytes"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "local/alice-skill/internal/urlstorage"
    

)


// Тесты для GET обработчика
func TestHandleGet(t *testing.T) {
    tests := []struct {
        name         string
        setupStorage func() *urlstorage.URLStorage
        urlPath      string
        wantStatus   int
        wantLocation string
        wantBody     string
    }{
        {
            name: "successful redirect",
            setupStorage: func() *urlstorage.URLStorage {
                s := urlstorage.NewURLStorage()
                s.SaveURL("abc123", "/long-url.com")
                return s
            },
            urlPath:      "/abc123",
            wantStatus:   http.StatusTemporaryRedirect,
            wantLocation: "/long-url.com",
        },
        {
            name: "not found",
            setupStorage: func() *urlstorage.URLStorage {
                return urlstorage.NewURLStorage()
            },
            urlPath:    "/missing",
            wantStatus: http.StatusNotFound,
            wantBody:   "URL not found\n",
        },
        {
            name: "invalid short URL",
            setupStorage: func() *urlstorage.URLStorage {
                s := urlstorage.NewURLStorage()
                s.SaveURL("invalid", " ")
                return s
            },
            urlPath:    "/invalid",
            wantStatus: http.StatusBadRequest,
            wantBody:   "Long URL does not exist\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            storage := tt.setupStorage()
            req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
            res := httptest.NewRecorder()
            
            HandleURL(storage).ServeHTTP(res, req)

            assert.Equal(t, tt.wantStatus, res.Code, "status code mismatch")
            assert.Equal(t, tt.wantLocation, res.Header().Get("Location"), "location header mismatch")
            assert.Contains(t, res.Body.String(), tt.wantBody, "response body mismatch")
        })
    }
}

// Тесты для POST обработчика
func TestHandlePost(t *testing.T) {
    tests := []struct {
        name        string
        body        string
        wantStatus  int
        wantBody    string
        wantContain bool // Проверять ли содержимое тела
    }{
        {
            name:        "successful creation",
            body:        "/valid-url.com",
            wantStatus:  http.StatusCreated,
            wantContain: true,
        },
        {
            name:       "empty body",
            body:       "",
            wantStatus: http.StatusBadRequest,
            wantBody:   "Invalid request body",
        },
        {
            name:       "invalid URL",
            body:       " ",
            wantStatus: http.StatusBadRequest,
            wantBody:   "Invalid URL\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            storage := urlstorage.NewURLStorage()
            req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.body))
            res := httptest.NewRecorder()

            HandleURL(storage).ServeHTTP(res, req)

            assert.Equal(t, tt.wantStatus, res.Code, "status code mismatch")
            assert.Equal(t, "text/plain; charset=utf-8", res.Header().Get("Content-Type"), "content type mismatch")

            if tt.wantContain {
                // Проверяем что вернулся корректный короткий URL
                shortURL := res.Body.String()
                assert.NotEmpty(t, shortURL, "short URL should not be empty")
                
                // Проверяем что URL сохранился в storage
                longURL, err := storage.GetURL(shortURL)
                require.NoError(t, err)
                assert.Equal(t, tt.body, longURL, "storage value mismatch")
            } else {
                assert.Contains(t, res.Body.String(), tt.wantBody, "response body mismatch")
            }
        })
    }
}


// Тесты для обработки методов
func TestHandleMethodNotAllowed(t *testing.T) {
    methods := []string{
        http.MethodPut,
        http.MethodDelete,
        http.MethodPatch,
    }

    storage := urlstorage.NewURLStorage()
    for _, method := range methods {
        t.Run(fmt.Sprintf("method %s", method), func(t *testing.T) {
            req := httptest.NewRequest(method, "/", nil)
            res := httptest.NewRecorder()

            HandleURL(storage).ServeHTTP(res, req)

            assert.Equal(t, http.StatusMethodNotAllowed, res.Code, "status code mismatch")
            assert.Contains(t, res.Body.String(), "Method not allowed", "response body mismatch")
        })
    }
}

// // Тест парсинга URL
// func TestParseURL(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		path       string
// 		storage    map[string]string
// 		wantLong   string
// 		wantStatus int
// 	}{
// 		{
// 			name:       "valid path",
// 			path:       "/valid",
// 			storage:    map[string]string{"valid": "https://ok.com"},
// 			wantLong:   "https://ok.com",
// 			wantStatus: http.StatusOK,
// 		},
// 		{
// 			name:       "empty path",
// 			path:       "/",
// 			wantStatus: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			storage := urlstorage.NewURLStorage()
// 			for k, v := range tt.storage {
// 				storage.SaveURL(k, v)
// 			}

// 			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
// 			res := httptest.NewRecorder()

// 			result := ParseURL(storage, res, req)

// 			if tt.wantStatus != http.StatusOK {
// 				require.Equal(t, tt.wantStatus, res.Code)
// 				return
// 			}

// 			assert.Equal(t, tt.wantLong, result)
// 		})
// 	}
// }
