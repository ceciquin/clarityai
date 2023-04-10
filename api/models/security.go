package main


// define a struct to represent a security
type Security struct {
    ID    int     `json:"id"`
    Score float64 `json:"score"`
}

// GET /securities endpoint
r.GET("/securities", func(c *gin.Context) {
	// query database to get list of securities
	securities := []Security{
		{ID: "ABC", Name: "ABC Corporation", Symbol: "ABC", Exchange: "NYSE"},
		{ID: "DEF", Name: "DEF Corporation", Symbol: "DEF", Exchange: "NASDAQ"},
		// add more securities as needed
	}

	// return list of securities as a JSON response
	c.JSON(200, gin.H{"securities": securities})
})
