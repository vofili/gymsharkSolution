package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pack struct {
	Size int
}

func calculatePacks(orderQuantity int, packSizes []Pack) []int {
	packs := make([]int, len(packSizes))
	for i := len(packSizes) - 1; i >= 0; i-- {
		packs[i] = orderQuantity / packSizes[i].Size
		orderQuantity %= packSizes[i].Size
	}
	return packs
}

func packHandler(c *gin.Context) {
	orderQuantityStr := c.Query("quantity")
	orderQuantity, err := strconv.Atoi(orderQuantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order quantity"})
		return
	}

	packSizes := []Pack{
		{Size: 5000},
		{Size: 2000},
		{Size: 1000},
		{Size: 500},
		{Size: 250},
	}

	packs := calculatePacks(orderQuantity, packSizes)
	c.JSON(http.StatusOK, gin.H{"packs": packs})
}

func main() {
	r := gin.Default()
	r.GET("/packs", packHandler)
	r.Run(":8080")
}
