package controllers

import (
	"net/http"
	"strconv"

	"github.com/akkien/learngo-explorer-api/models"
	"github.com/gin-gonic/gin"
)

// GET /blocks
// Show the new thread form page
func GetBlocks(c *gin.Context) {
	// _, err := util.Session(writer, request)
	// if err != nil {
	// 	c.Redirect(http.StatusFound, "/login")
	// 	return
	// }
	page, err1 := strconv.Atoi(c.Query("page"))
	limit, err2 := strconv.Atoi(c.Query("limit"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	models := c.MustGet("db").(*models.DBModel)

	blocks, err := models.Blocks(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get data!"})
		return
	}
	// 	res := gin.H{
	// 		"title":   "Signup",
	// 		"payload": response,
	// 	}
	// 	render(c, res, "transactions.html")
	c.JSON(http.StatusOK, blocks)
}

// GET /blocks/:blockid
func GetBlock(c *gin.Context) {
	numberParam := c.Param("number")

	blockNumber, err := strconv.Atoi(numberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block number"})
		return
	}

	models := c.MustGet("db").(*models.DBModel)

	block, err := models.BlockByNumber(blockNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get data!"})
		return
	}

	c.JSON(http.StatusOK, block)
}

// GET /txs
func GetTransactions(c *gin.Context) {
	page, err1 := strconv.Atoi(c.Query("page"))
	limit, err2 := strconv.Atoi(c.Query("limit"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	models := c.MustGet("db").(*models.DBModel)

	txs, err := models.Transactions(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get data!"})
		return
	}

	c.JSON(http.StatusOK, txs)
}

// GET /blocks/:number/txs
func GetTransactionsInBlock(c *gin.Context) {
	numberParam := c.Param("number")

	blockNumber, err := strconv.Atoi(numberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block number"})
		return
	}

	models := c.MustGet("db").(*models.DBModel)

	txs, err := models.TransactionsByBlock(blockNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get data!"})
		return
	}

	c.JSON(http.StatusOK, txs)
}

// GET /txs/:hash
func GetTransaction(c *gin.Context) {
	txHash := c.Param("hash")

	models := c.MustGet("db").(*models.DBModel)

	tx, err := models.TransactionByHash(txHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get data!"})
		return
	}

	c.JSON(http.StatusOK, tx)
}
