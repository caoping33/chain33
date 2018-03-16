package trade

import (
	"code.aliyun.com/chain33/chain33/account"
	"code.aliyun.com/chain33/chain33/types"
	dbm "code.aliyun.com/chain33/chain33/common/db"
	"code.aliyun.com/chain33/chain33/common"
)

type SellDB struct {
	types.SellOrder
}

//TODO STATUS没有什么作用，可以考虑删除
func NewSellDB(sellorder types.SellOrder) (selldb *SellDB) {
	selldb = &SellDB{sellorder}
	if types.InvalidStartTime != selldb.Starttime {
		selldb.Status = types.NotStart
	}
	return
}

func (selldb *SellDB) Save(db dbm.KVDB) {
	set := selldb.GetKVSet()
	for i := 0; i < len(set); i++ {
		db.Set(set[i].GetKey(), set[i].Value)
	}
}

func (selldb *SellDB) GetSellLogs(tradeType int32) *types.ReceiptLog {
	log := &types.ReceiptLog{}
	log.Ty = tradeType
	key := tokenSellKey(selldb.Hash)
	receiptTrade := &types.ReceiptTradeSell{key}
	log.Log = types.Encode(receiptTrade)

	return log
}

func (selldb *SellDB) GetBuyLogs(buyerAddr string, sellid string, boardlotcnt int64, sellorder *types.SellOrder, txhash string) *types.ReceiptLog {
	log := &types.ReceiptLog{}
	log.Ty = types.TyLogTradeBuy
	receiptBuy := &types.ReceiptTradeBuy{
		buyerAddr,
		sellid,
		selldb.Tokensymbol,
		txhash,
		boardlotcnt,
		sellorder.Amountperboardlot,
		sellorder.Priceperboardlot}
	log.Log = types.Encode(receiptBuy)

	return log
}

func getSellOrderFromID(sellid []byte, db dbm.KVDB) (*types.SellOrder, error) {
	value, err := db.Get(sellid)
	if err != nil {
		tradelog.Error("getSellOrderFromID", "Failed to get value frim db wiht sellid", string(sellid))
		return nil, err
	}

	var sellorder types.SellOrder
	if err = types.Decode(value, &sellorder); err != nil {
		tradelog.Error("getSellOrderFromID", "Failed to decode sell order", string(sellid))
		return nil, err
	}
	return &sellorder, nil
}

//key:mavl-token-xxx, value:token
func (selldb *SellDB) GetKVSet() (kvset []*types.KeyValue) {
	value := types.Encode(&selldb.SellOrder)
	key := tokenSellKey(selldb.Hash)
	kvset = append(kvset, &types.KeyValue{key, value})
	return kvset
}

type TradeAction struct {
	coinsAccount *account.AccountDB
	db        dbm.KVDB
	txhash    string
	fromaddr  string
	blocktime int64
	height    int64
	execaddr  string
}

//func NewTradeAction(db dbm.KVDB, tx *types.Transaction, toaddr, execaddr string, blocktime, height int64) *TradeAction {
func NewTradeAction(t *Trade, tx *types.Transaction) *TradeAction {
	hash := common.Bytes2Hex(tx.Hash())
	fromaddr := account.PubKeyToAddress(tx.GetSignature().GetPubkey()).String()
	return &TradeAction{t.GetCoinsAccount(), t.GetDB(), hash, fromaddr,
	t.GetBlockTime(), t.GetHeight(), t.GetAddr()}
}

func (action *TradeAction) TradeSell(sell *types.TradeForSell) (*types.Receipt, error) {
	tokenAccDB := account.NewTokenAccount(sell.Tokensymbol, action.db)
	//tokenAcc := tokenAccDB.LoadAccount(action.fromaddr)

	//确认发起此次出售或者众筹的余额是否足够
	totalAmount := sell.GetTotalboardlot() * sell.GetAmountperboardlot()
	receipt, err := tokenAccDB.ExecFrozen(action.fromaddr, action.execaddr, totalAmount)
	if err != nil {
		tradelog.Error("trade sell ", "addr", action.fromaddr, "execaddr", action.execaddr, "amount", totalAmount)
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	sellorder := types.SellOrder{
		sell.GetTokensymbol(),
		action.fromaddr,
		sell.GetAmountperboardlot(),
		sell.GetPriceperboardlot(),
		sell.GetTotalboardlot(),
		0,
		sell.GetStarttime(),
		sell.GetStoptime(),
		sell.GetCrowdfund(),
		action.txhash,
		types.OnSale,
	}
	tokendb := NewSellDB(sellorder)
	tokendb.Save(action.db)
	logs = append(logs, receipt.Logs...)
	logs = append(logs, tokendb.GetSellLogs(types.TyLogTradeSell))
	kv = append(kv, receipt.KV...)
	kv = append(kv, tokendb.GetKVSet()...)

	receipt = &types.Receipt{types.ExecOk, kv, logs}
	return receipt, nil
}

func (action *TradeAction) TradeBuy(buyorder *types.TradeForBuy) (*types.Receipt, error) {
	sellidByte := []byte(buyorder.Sellid)
	sellorder, err := getSellOrderFromID(sellidByte, action.db)
	if err != nil {
		return nil, types.ErrTSellOrderNotExist
	}
	//TODO 使用blocktime，避免不同主机的时间差异
	if sellorder.Status == types.NotStart && sellorder.Starttime > action.blocktime {
		return nil, types.ErrTSellOrderNotStart
	} else if sellorder.Status == types.OnSale && sellorder.Totalboardlot-sellorder.Soldboardlot < buyorder.Boardlotcnt {
		return nil, types.ErrTSellOrderNotEnough
	} else if sellorder.Status == types.SoldOut {
		return nil, types.ErrTSellOrderSoldout
	} else if sellorder.Status == types.Revoked {
		return nil, types.ErrTSellOrderRevoked
	} else if sellorder.Status == types.Expired {
		return nil, types.ErrTSellOrderExpired
	}
	sellorder.Status = types.OnSale

	//首先购买费用的划转
	receiptFromAcc, err := action.coinsAccount.ExecTransfer(action.fromaddr, sellorder.Address, action.execaddr, buyorder.Boardlotcnt * sellorder.Priceperboardlot)
	if err != nil {
		tradelog.Error("account.Transfer ", "addrFrom", action.fromaddr, "addrTo", sellorder.Address,
			"amount", buyorder.Boardlotcnt*sellorder.Priceperboardlot)
		return nil, err
	}
	//然后实现购买token的转移,因为这部分token在之前的卖单生成时已经进行冻结
	//TODO: 创建一个LRU用来保存token对应的子合约账户的地址
	tokenAccDB := account.NewTokenAccount(sellorder.Tokensymbol, action.db)
	receiptFromExecAcc, err := tokenAccDB.ExecTransferFrozen(sellorder.Address, action.fromaddr, action.execaddr, buyorder.Boardlotcnt * sellorder.Amountperboardlot)
	if err != nil {
		tradelog.Error("account.ExecTransfer token ", "addrFrom", sellorder.Address, "addrTo", action.fromaddr,
			"execaddr", action.execaddr, "amount", buyorder.Boardlotcnt*sellorder.Amountperboardlot)
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	sellorder.Soldboardlot += buyorder.Boardlotcnt
	if sellorder.Soldboardlot == sellorder.Totalboardlot {
		sellorder.Status = types.SoldOut
	}
	tokendb := NewSellDB(*sellorder)
	tokendb.Save(action.db)

	logs = append(logs, receiptFromAcc.Logs...)
	logs = append(logs, receiptFromExecAcc.Logs...)
	logs = append(logs, tokendb.GetBuyLogs(action.fromaddr, buyorder.Sellid, buyorder.Boardlotcnt, sellorder, action.txhash))
	kv = append(kv, receiptFromAcc.KV...)
	kv = append(kv, receiptFromExecAcc.KV...)
	return &types.Receipt{types.ExecOk, kv, logs}, nil
}

func (action *TradeAction) TradeRevokeSell(revoke *types.TradeForRevokeSell) (*types.Receipt, error) {
	sellidByte := []byte(revoke.Sellid)
	sellorder, err := getSellOrderFromID(sellidByte, action.db)
	if err != nil {
		return nil, types.ErrTSellOrderNotExist
	}

	if sellorder.Status == types.SoldOut {
		return nil, types.ErrTSellOrderSoldout
	} else if sellorder.Status == types.Revoked {
		return nil, types.ErrTSellOrderRevoked
	} else if sellorder.Status == types.Expired {
		return nil, types.ErrTSellOrderExpired
	}

	if action.fromaddr != sellorder.Address {
		return nil, types.ErrTSellOrderRevoke
	}
	//然后实现购买token的转移,因为这部分token在之前的卖单生成时已经进行冻结
	//TODO: 创建一个LRU用来保存token对应的子合约账户的地址
	tokenAccDB := account.NewTokenAccount(sellorder.Tokensymbol, action.db)
	tradeRest := (sellorder.Totalboardlot - sellorder.Soldboardlot) * sellorder.Amountperboardlot
	receiptFromExecAcc, err := tokenAccDB.ExecActive(sellorder.Address, action.execaddr, tradeRest)
	if err != nil {
		tradelog.Error("account.ExecActive token ", "addrFrom", sellorder.Address, "execaddr", action.execaddr, "amount", tradeRest)
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	sellorder.Status = types.Revoked
	tokendb := NewSellDB(*sellorder)
	tokendb.Save(action.db)

	logs = append(logs, receiptFromExecAcc.Logs...)
	logs = append(logs, tokendb.GetSellLogs(types.TyLogTradeRevoke))
	kv = append(kv, receiptFromExecAcc.KV...)
	return &types.Receipt{types.ExecOk, kv, logs}, nil
}

func tokenSellKey(hash string) (key []byte) {
	key = append(key, []byte("mavl-trade-sell-")...)
	key = append(key, []byte(hash)...)

	return key
}
