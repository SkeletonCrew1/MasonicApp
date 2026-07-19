package main
import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq"
)
var db *sql.DB
type banRequest struct {
	IP string `json:"ip"`
}
type banRecord struct {
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
func banHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req banRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.IP == "" {
		http.Error(w, "ip is required", http.StatusBadRequest)
		return
	}
	_, err := db.Exec(
		`INSERT INTO banned_ips (ip) VALUES ($1) ON CONFLICT (ip) DO NOTHING`,
		req.IP,
	)
	if err != nil {
		log.Println("ban insert error:", err)
		http.Error(w, "could not ban ip", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "banned", "ip": req.IP})
}
func checkHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "ip query param is required", http.StatusBadRequest)
		return
	}
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM banned_ips WHERE ip = $1)`, ip).Scan(&exists)
	if err != nil {
		log.Println("check query error:", err)
		http.Error(w, "could not check ip", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"banned": exists})
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT ip, created_at FROM banned_ips ORDER BY created_at DESC`)
	if err != nil {
		log.Println("list query error:", err)
		http.Error(w, "could not list bans", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var records []banRecord
	for rows.Next() {
		var rec banRecord
		if err := rows.Scan(&rec.IP, &rec.CreatedAt); err != nil {
			continue
		}
		records = append(records, rec)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
func main() {
	dsn := os.Getenv("DB_DSN")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("could not open db:", err)
	}
	for i := 0; i < 10; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Println("waiting for postgres...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS banned_ips (
			id SERIAL PRIMARY KEY,
			ip TEXT UNIQUE NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal("could not create table:", err)
	}
	http.HandleFunc("/ban", withCORS(banHandler))
	http.HandleFunc("/check", withCORS(checkHandler))
	http.HandleFunc("/bans", withCORS(listHandler))
	log.Println("ban-service listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
