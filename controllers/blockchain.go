package controllers

import (
	"net/http"
	"strconv"

	"github.com/akkien/explorer-modern/models"
	"github.com/akkien/explorer-modern/util"
)

// GET /blocks
// Show the new thread form page
func ReadBlocks(writer http.ResponseWriter, request *http.Request) {
	_, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	blocks, err := models.Blocks(1, 10)
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get data")
		return
	}
	response := models.BksTxs{Blocks: blocks}
	util.GenerateHTML(writer, response, "layout", "private.navbar", "blocks")
}

// GET /txs
// ReadTransactions Show the new thread form page
func ReadTransactions(writer http.ResponseWriter, request *http.Request) {
	_, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
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
	util.GenerateHTML(writer, response, "layout", "private.navbar", "transactions")
}

// GET /blocks/:blockid
// ReadBlock Show the new thread form page
func ReadBlock(writer http.ResponseWriter, request *http.Request) {
	_, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	path := request.URL.Path
	_, numberParam := util.ShiftPath(path)
	if numberParam == "/" {
		http.Redirect(writer, request, "/txs", 302)
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

	util.GenerateHTML(writer, block, "layout", "private.navbar", "block")
}

// GET /tx/:txid
// ReadTransaction Show the new thread form page
func ReadTransaction(writer http.ResponseWriter, request *http.Request) {
	_, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}

	path := request.URL.Path
	_, txHash := util.ShiftPath(path)
	if txHash == "/" {
		http.Redirect(writer, request, "/txs", 302)
		return
	}

	tx, err := models.TransactionByHash(txHash)
	receipt, err := models.ReceiptByHash(txHash)
	if err != nil {
		util.ErrorMessage(writer, request, "Transaction not found")
		return
	}

	response := models.TxDetail{Tx: tx, Receipt: receipt}
	util.GenerateHTML(writer, response, "layout", "private.navbar", "transaction")
}
