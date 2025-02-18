package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"os"
	"simple-web-go/internal/config"
	"simple-web-go/internal/database"
	"simple-web-go/internal/handler"
	"simple-web-go/internal/routes"
	"simple-web-go/internal/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

func LoadConfig(path string) (*config.Config, error) {
	cfg := &config.Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func initDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := database.InitDB(&database.Config{
		User:            cfg.Database.User,
		Password:        cfg.Database.Password,
		Host:            cfg.Database.Host,
		Port:            cfg.Database.Port,
		DBName:          cfg.Database.DBName,
		Charset:         cfg.Database.Charset,
		ParseTime:       cfg.Database.ParseTime,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 错误处理中间件
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "服务器内部错误",
				})
			}
		}()
		c.Next()
	})

	// 创建API分组
	apiV1 := r.Group("/api")
	{
		// 初始化服务和handler
		services := service.InitServices(db)
		handlers := handler.InitHandlers(services)

		// 注册路由
		routes.RegisterUserRoutes(apiV1, handlers.SysUserHandler)
		routes.RegisterMenuRoutes(apiV1, handlers.SysMenuHandler)
	}

	return r
}

func startServer(cfg *config.Config, router *gin.Engine) error {
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("服务器启动在 %s", s.Addr)
	return s.ListenAndServe()
}

func main() {
	// 加载配置
	cfg, err := LoadConfig("configs/config.yml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	db, err := initDB(cfg)
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 设置gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化路由
	router := setupRouter(db)

	// 启动服务器
	if err := startServer(cfg, router); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
