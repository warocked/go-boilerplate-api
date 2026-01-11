package db

import (
	"fmt"
)

// ConnectionPoolStats represents database connection pool statistics
type ConnectionPoolStats struct {
	MaxOpenConnections int    // Maximum number of open connections
	OpenConnections    int    // Number of open connections (in use + idle)
	InUse              int    // Number of connections currently in use
	Idle               int    // Number of idle connections
	WaitCount          int64  // Total number of connections waited for
	WaitDuration       string // Total time blocked waiting for a new connection
	MaxIdleClosed      int64  // Total number of connections closed due to SetMaxIdleConns
	MaxIdleTimeClosed  int64  // Total number of connections closed due to SetConnMaxIdleTime
	MaxLifetimeClosed  int64  // Total number of connections closed due to SetConnMaxLifetime
}

// GetPoolStats returns connection pool statistics
// This is useful for monitoring and verifying connection reuse
func GetPoolStats() (*ConnectionPoolStats, error) {
	if DB == nil {
		return nil, fmt.Errorf("database is not initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	stats := sqlDB.Stats()

	return &ConnectionPoolStats{
		MaxOpenConnections: stats.MaxOpenConnections,
		OpenConnections:    stats.OpenConnections,
		InUse:              stats.InUse,
		Idle:               stats.Idle,
		WaitCount:          stats.WaitCount,
		WaitDuration:       stats.WaitDuration.String(),
		MaxIdleClosed:      stats.MaxIdleClosed,
		MaxIdleTimeClosed:  stats.MaxIdleTimeClosed,
		MaxLifetimeClosed:  stats.MaxLifetimeClosed,
	}, nil
}

func LogPoolStats() error {
	stats, err := GetPoolStats()
	if err != nil {
		return err
	}

	fmt.Printf("Database Connection Pool Stats:\n")
	fmt.Printf("  Max Open Connections: %d\n", stats.MaxOpenConnections)
	fmt.Printf("  Open Connections: %d\n", stats.OpenConnections)
	fmt.Printf("  In Use: %d\n", stats.InUse)
	fmt.Printf("  Idle: %d\n", stats.Idle)
	fmt.Printf("  Wait Count: %d\n", stats.WaitCount)
	fmt.Printf("  Wait Duration: %s\n", stats.WaitDuration)
	fmt.Printf("  Max Idle Closed: %d\n", stats.MaxIdleClosed)
	fmt.Printf("  Max Idle Time Closed: %d\n", stats.MaxIdleTimeClosed)
	fmt.Printf("  Max Lifetime Closed: %d\n", stats.MaxLifetimeClosed)

	return nil
}

// VerifyConnectionReuse verifies that connections are being reused properly
// Returns true if connections are being reused (idle connections > 0 or in use < open)
func VerifyConnectionReuse() (bool, error) {
	stats, err := GetPoolStats()
	if err != nil {
		return false, err
	}

	// Connections are being reused if:
	// 1. There are idle connections (connections available for reuse)
	// 2. Or connections are being used but pool is working (open < max)
	reusing := stats.Idle > 0 || (stats.OpenConnections < stats.MaxOpenConnections && stats.OpenConnections > 0)

	return reusing, nil
}
