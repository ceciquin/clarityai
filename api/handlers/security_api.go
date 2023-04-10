package main

import (
	"clarityai/api/cache"
	"clarityai/db"
	"context"
	"go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
)

func main() {

	//start the implementaton of the observability strategy using Jaeger and opentelemetry
    jaegerExporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
    if err != nil {
        log.Fatal(err)
    }

    tracerProvider := trace.NewTracerProvider(
        trace.WithSampler(trace.AlwaysSample()),
        trace.WithResource(resource.NewWithAttributes(
            semconv.ServiceNameKey.String("my-api"),
        )),
        trace.WithBatcher(jaegerExporter),
    )

    otel.SetTracerProvider(tracerProvider)

    tracer := otel.Tracer("my-api")

	// Initialize a new Redis client
	redisClient, err := cache.NewCache("localhost:6379", "", 0)
	if err != nil {
		return
	}

	//begin gin routes
	r := gin.Default()

	// GET /securities endpoint
	r.GET("/securities", func(c *gin.Context) {
		// Start a new span
        ctx, span := tracer.Start(r.Context(), "securities.request")
        defer span.End()

		// Check if the list of securities IDs exists in Redis cache
		securities, err := redisClient.Get("security")
		if err != nil {
			// If the score does not exist in cache, retrieve it from the database
			[]securities, err = getSecurityFromDatabase()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to retrieve securities"})
				return
			}

			// Store the score in Redis cache for future requests
			redisClient.Set(context.Background(), "securities", securities, 0)
		}

		// Set attributes on the span
		span.SetAttributes(semconv.HTTPMethodKey.String(r.Method))
        span.SetAttributes(semconv.HTTPURLKey.String(r.URL.String()))

		// Return the score to the client
		c.JSON(200, gin.H{"score": securities})

	})

	// GET /securities/{security_id}/scores endpoint
	r.GET("/securities/:security_id/scores", func(c *gin.Context) {
		// return score details for a given security ID
		securityID := c.Param("id")

		// Check if the security score exists in Redis cache
		score, err := redisClient.Get(securityID)
		if err != nil {
			// If the score does not exist in cache, retrieve it from the database
			score, err = getScoreFromDatabase(securityID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to retrieve score"})
				return
			}

			// Store the score in Redis cache for future requests
			redisClient.Set(context.Background(), securityID, score, 0)
		}

		// Return the score to the client
		c.JSON(200, gin.H{"score": score})
	})

	// start the server
	r.Run(":8080")

	// Close the Redis connection when the program exits
	defer 

}

// A helper function that retrieves a security score from the database
func getScoreFromDatabase(securityID string) (string, error) {
	// Connect to the database
	db.Connect()
	myDB := db.GetDB()
	defer db.Close()

	// Query the database for the security score
	var score string
	err := myDB.QueryRow("SELECT score FROM security WHERE securityid = $1", securityID).Scan(&score)
	if err != nil {
		return "", err
	}

	return score, nil
}

// A helper function that retrieves a security score from the database
func getSecurityFromDatabase() ([]string, error) {
	// Connect to the database
	db.Connect()
	myDB := db.GetDB()
	defer db.Close()

	// Query the database for the security List
	rows, err := myDB.Query("SELECT securityname FROM security")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var securityList []string
	for rows.Next() {
		var security string
		err := rows.Scan(&security)
		if err != nil {
			return nil, err
		}
		securityList = append(securityList, security)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return securityList, nil
}
