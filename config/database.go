package config
import (
"fmt"
"log"
"os"
"github.com/KicaiBadai/gin-firebase-backend/models"
"gorm.io/driver/mysql"
"gorm.io/gorm"
"gorm.io/gorm/logger"
)
// DB adalah instance GORM global yang dipakai di seluruh aplikasi
var DB *gorm.DB
func InitDatabase() {
// Ambil konfigurasi dari environment variables
host := os.Getenv("DB_HOST")
port := os.Getenv("DB_PORT")
user := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
dbname := os.Getenv("DB_NAME")
// Format DSN (Data Source Name) untuk MySQL
// Format: user:pass@tcp(host:port)/dbname?params
dsn := fmt.Sprintf(
"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
user, password, host, port, dbname,
)
// Konfigurasi GORM
gormConfig := &gorm.Config{
Logger: logger.Default.LogMode(logger.Info), // Log semua query SQL
}
// Buka koneksi
var err error
DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
if err != nil {
log.Fatalf("Gagal koneksi ke database: %v", err)
}
// Setup connection pool
sqlDB, err := DB.DB()
if err != nil {log.Fatalf("Gagal mendapatkan sql.DB: %v", err)
}
sqlDB.SetMaxOpenConns(25) // Maksimal 25 koneksi terbuka
sqlDB.SetMaxIdleConns(10) // Maksimal 10 koneksi idle
// AutoMigrate: buat/update tabel sesuai struct model
// GORM akan buat tabel jika belum ada
err = DB.AutoMigrate(
&models.User{},
&models.Product{},
)
if err != nil {
log.Fatalf("AutoMigrate gagal: %v", err)
}
log.Println("Database terhubung dan tabel sudah di-migrate")}