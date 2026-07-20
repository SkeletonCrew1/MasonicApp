package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db *gorm.DB
type BanRecord struct {
	ID 		  uint		`gorm:"primaryKey"`
	IP        string    `gorm:"uniqueIndex;not null" json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}
type BanRequest struct {
	IP string `json:"ip" binding:"required"`
}
func initDB(dsn string){
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("waiting for postgres...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	db.AutoMigrate(&BanRecord{})
}

func banHandler(c *gin.Context) {
	var req BanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ip is required"})
		return
	}
	record := BanRecord{IP: req.IP}
	result := db.Where(BanRecord{IP: req.IP}).FirstOrCreate(&record)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not ban ip"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "banned", "ip": req.IP})
}
func checkHandler(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ip query param is required"})
		return
	}
	var count int64
	db.Model(&BanRecord{}).Where("ip = ?", ip).Count(&count)
	c.JSON(http.StatusOK, gin.H{"banned": count > 0})
}
func listHandler(c *gin.Context) {
	var records []BanRecord
	db.Order("created_at desc").Find(&records)
	c.JSON(http.StatusOK, records)
}
func main() {
	dsn := os.Getenv("DB_DSN")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	initDB(dsn)
	r := gin.Default()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	r.Use(func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	})

	r.POST("/ban", banHandler)
	r.GET("/check", checkHandler)
	r.GET("/bans", listHandler)

	log.Println("ban-service listening on:" +port)
	r.Run(":" + port)
}
