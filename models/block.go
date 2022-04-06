package models

// Block for PostgreSQL
type Block struct {
	Difficulty       int64  `json:"difficulty" db:"difficulty"`
	ExtraData        string `json:"extraData" db:"extra_data"`
	GasLimit         int64  `json:"gasLimit" db:"gas_limit"`
	GasUsed          int64  `json:"gasUsed" db:"gas_used"`
	Hash             string `json:"hash" db:"hash"`
	LogsBloom        string `json:"logsBloom" db:"logs_bloom"`
	Miner            string `json:"miner" db:"miner"`
	MixHash          string `json:"mixHash" db:"mix_hash"`
	Nonce            string `json:"nonce" db:"nonce"`
	Number           int64  `json:"number" db:"number"`
	ParentHash       string `json:"parentHash" db:"parent_hash"`
	ReceiptsRoot     string `json:"receiptsRoot" db:"receipts_root"`
	Sha3Uncles       string `json:"sha3Uncles" db:"sha3_uncles"`
	Size             int64  `json:"size" db:"size"`
	StateRoot        string `json:"stateRoot" db:"state_root"`
	Timestamp        int64  `json:"timestamp" db:"timestamp"`
	TotalDifficulty  int64  `json:"totalDifficulty" db:"total_difficulty"`
	TransactionsRoot string `json:"transactionsRoot" db:"transactions_root"`
	TransactionCount int64  `json:"transactionCount" db:"transaction_count"`
	CreatedTimestamp string `json:"createdTimestamp" db:"created_timestamp"`
}

var blockParams = []string{
	"difficulty",
	"extra_data",
	"gas_limit",
	"gas_used",
	"hash",
	"logs_bloom",
	"miner",
	"mix_hash",
	"nonce",
	"number",
	"parent_hash",
	"receipts_root",
	"sha3_uncles",
	"size",
	"state_root",
	"timestamp",
	"total_difficulty",
	"transactions_root",
	"transaction_count",
}

// Blocks : Get all blocks in the database and returns it
func (m *DBModel) Blocks(page, limit int) (blocks []Block, err error) {
	offset := (page - 1) * limit
	m.DB.Limit(limit).Offset(offset).Find(&blocks)
	return
}

// BlockByNumber : Get all blocks in the database and returns it
func (m *DBModel) BlockByNumber(number int) (item Block, err error) {
	item = Block{}
	result := m.DB.Where("number = ?", number).First(&item)
	err = result.Error
	return
}
