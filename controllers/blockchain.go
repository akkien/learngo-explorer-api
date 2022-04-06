package controllers

import (
	"net/http"
	"strconv"

	"github.com/akkien/learngo-explorer-api/models"
	"github.com/akkien/learngo-explorer-api/util"
	"github.com/gin-gonic/gin"
)

// GET /blocks
// Show the new thread form page
func ReadBlocks(c *gin.Context) {
	_, err := util.Session(writer, request)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	blocks, err := models.Blocks(1, 10)
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get data")
		return
	}
	response := models.BksTxs{Blocks: blocks}
	res := gin.H{
		"title":   "Signup",
		"payload": response,
	}
	render(c, res, "blocks.html")
}

// GET /txs
// ReadTransactions Show the new thread form page
func ReadTransactions(c *gin.Context) {
	_, err := util.Session(writer, request)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	queries := request.URL.Query()
	if len(queries) > 0 {
		blockQuery := queries["block"]
		response := models.BksTxs{}

		if len(blockQuery) > 0 {
			// Query transactions in spectific block
			blockNumber, err := strconv.Atoi(blockQuery[0])
			if err != nil {
				util.ErrorMessage(writer, request, "Cannot get data")
				return
			}

			txs, err := models.TransactionsByBlock(blockNumber)
			if err != nil {
				util.ErrorMessage(writer, request, "Cannot get data")
				return
			}
			response.Txs = txs
		} else {
			// There is query but wrong format, query txs of last block
			txs, err := models.TransactionsLastBlock()
			if err != nil {
				util.ErrorMessage(writer, request, "Cannot get data")
				return
			}
			response.Txs = txs
		}

		util.GenerateHTML(writer, response, "layout", "private.navbar", "transactions")
		return
	}

	// Query last 10 txs
	txs, err := models.Transactions(1, 10)
	response := models.BksTxs{Txs: txs}
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get data")
		return
	}
	res := gin.H{
		"title":   "Signup",
		"payload": response,
	}
	render(c, res, "transactions.html")
}

// GET /blocks/:blockid
// ReadBlock Show the new thread form page
func ReadBlock(c *gin.Context) {
	_, err := util.Session(writer, request)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	path := request.URL.Path
	_, numberParam := util.ShiftPath(path)
	if numberParam == "/" {
		c.Redirect(http.StatusFound, "/txs")
		return
	}

	blockNumber, err := strconv.Atoi(numberParam)
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get data")
		return
	}

	block, err := models.BlockByNumber(blockNumber)
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get data")
		return
	}

	res := gin.H{
		"title":   "Signup",
		"payload": block,
	}
	render(c, res, "block.html")
}

// GET /tx/:txid
// ReadTransaction Show the new thread form page
func ReadTransaction(c *gin.Context) {
	_, err := util.Session(writer, request)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	path := request.URL.Path
	_, txHash := util.ShiftPath(path)
	if txHash == "/" {
		c.Redirect(http.StatusFound, "/txs")
		return
	}

	tx, err := models.TransactionByHash(txHash)
	receipt, err := models.ReceiptByHash(txHash)
	if err != nil {
		util.ErrorMessage(writer, request, "Transaction not found")
		return
	}

	response := models.TxDetail{Tx: tx, Receipt: receipt}
	res := gin.H{
		"title":   "Signup",
		"payload": response,
	}
	render(c, res, "transaction.html")
}
