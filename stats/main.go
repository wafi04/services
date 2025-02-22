package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPU              float64 `json:"cpu"`
	MemoryPercent    float64 `json:"memory_percent"`
	TotalMemoryGB    float64 `json:"total_memory_gb"`
	UsedMemoryGB     float64 `json:"used_memory_gb"`
	FreeMemoryGB     float64 `json:"free_memory_gb"`
	TotalMemoryBytes uint64  `json:"total_memory_bytes"`
	UsedMemoryBytes  uint64  `json:"used_memory_bytes"`
	FreeMemoryBytes  uint64  `json:"free_memory_bytes"`
	Time             string  `json:"timestamp"`
}

func bytesToGB(bytes uint64) float64 {
	return float64((bytes) / (1024 * 1024 * 1024))
}
func getSystemStats() SystemStats {
	cpu, err := cpu.Percent(time.Second, false)

	if err != nil {
		log.Printf("err  : %s", err.Error())
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("mem : %s ", err.Error())
	}

	return SystemStats{
		CPU:              cpu[0],
		MemoryPercent:    memory.UsedPercent,
		TotalMemoryGB:    bytesToGB(memory.Total),
		UsedMemoryGB:     bytesToGB(memory.Used),
		FreeMemoryGB:     bytesToGB(memory.Free),
		TotalMemoryBytes: memory.Total,
		UsedMemoryBytes:  memory.Used,
		FreeMemoryBytes:  memory.Free,
		Time:             time.Now().Format(time.RFC3339),
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := getSystemStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	// Set up routes
	http.HandleFunc("/stats", statsHandler)

	// Start monitoring and logging
	go func() {
		for {
			stats := getSystemStats()
			fmt.Printf("\nSystem Statistics:\n")
			fmt.Printf("CPU Usage: %.2f%%\n", stats.CPU)
			fmt.Printf("Memory Usage: %.2f%%\n", stats.MemoryPercent)
			fmt.Printf("Total Memory: %.2f GB\n", stats.TotalMemoryGB)
			fmt.Printf("Used Memory: %.2f GB\n", stats.UsedMemoryGB)
			fmt.Printf("Free Memory: %.2f GB\n", stats.FreeMemoryGB)
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
