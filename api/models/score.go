package main


// define a struct to represent score details for a security
type Scores struct {
    Growth   float64 `json:"growth"`
    Value    float64 `json:"value"`
    Quality  float64 `json:"quality"`
    Momentum float64 `json:"momentum"`
}

// GET /securities/{security_id}/scores endpoint
r.GET("/securities/:security_id/scores", func(c *gin.Context) {
    securityID := c.Param("security_id")

    // query database to get score details for the given security ID
    scores := Scores{
        Growth:   8.5,
        Value:    7.9,
        Quality:  9.1,
        Momentum: 6.8,
    }

    // return score details as a JSON response
    c.JSON(200, gin.H{"id": securityID, "scores": scores})
})
