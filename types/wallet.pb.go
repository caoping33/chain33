// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WalletTxDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Height     int64        `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,6,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,7,opt,name=fromaddr" json:"fromaddr,omitempty"`
	Txhash     []byte       `protobuf:"bytes,8,opt,name=txhash,proto3" json:"txhash,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *WalletTxDetail) Reset()                    { *m = WalletTxDetail{} }
func (m *WalletTxDetail) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetail) ProtoMessage()               {}
func (*WalletTxDetail) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *WalletTxDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WalletTxDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *WalletTxDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WalletTxDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WalletTxDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *WalletTxDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WalletTxDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *WalletTxDetail) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *WalletTxDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `protobuf:"bytes,1,rep,name=txDetails" json:"txDetails,omitempty"`
}

func (m *WalletTxDetails) Reset()                    { *m = WalletTxDetails{} }
func (m *WalletTxDetails) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetails) ProtoMessage()               {}
func (*WalletTxDetails) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *WalletTxDetails) GetTxDetails() []*WalletTxDetail {
	if m != nil {
		return m.TxDetails
	}
	return nil
}

type WalletAccountStore struct {
	Privkey   string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label     string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *WalletAccountStore) Reset()                    { *m = WalletAccountStore{} }
func (m *WalletAccountStore) String() string            { return proto.CompactTextString(m) }
func (*WalletAccountStore) ProtoMessage()               {}
func (*WalletAccountStore) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *WalletAccountStore) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *WalletAccountStore) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WalletAccountStore) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *WalletAccountStore) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

type WalletPwHash struct {
	PwHash  []byte `protobuf:"bytes,1,opt,name=pwHash,proto3" json:"pwHash,omitempty"`
	Randstr string `protobuf:"bytes,2,opt,name=randstr" json:"randstr,omitempty"`
}

func (m *WalletPwHash) Reset()                    { *m = WalletPwHash{} }
func (m *WalletPwHash) String() string            { return proto.CompactTextString(m) }
func (*WalletPwHash) ProtoMessage()               {}
func (*WalletPwHash) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *WalletPwHash) GetPwHash() []byte {
	if m != nil {
		return m.PwHash
	}
	return nil
}

func (m *WalletPwHash) GetRandstr() string {
	if m != nil {
		return m.Randstr
	}
	return ""
}

type WalletStatus struct {
	IsLock       bool `protobuf:"varint,1,opt,name=isLock" json:"isLock,omitempty"`
	IsAutoMining bool `protobuf:"varint,2,opt,name=isAutoMining" json:"isAutoMining,omitempty"`
	HasSeed      bool `protobuf:"varint,3,opt,name=hasSeed" json:"hasSeed,omitempty"`
}

func (m *WalletStatus) Reset()                    { *m = WalletStatus{} }
func (m *WalletStatus) String() string            { return proto.CompactTextString(m) }
func (*WalletStatus) ProtoMessage()               {}
func (*WalletStatus) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *WalletStatus) GetIsLock() bool {
	if m != nil {
		return m.IsLock
	}
	return false
}

func (m *WalletStatus) GetIsAutoMining() bool {
	if m != nil {
		return m.IsAutoMining
	}
	return false
}

func (m *WalletStatus) GetHasSeed() bool {
	if m != nil {
		return m.HasSeed
	}
	return false
}

type WalletAccounts struct {
	Wallets []*WalletAccount `protobuf:"bytes,1,rep,name=wallets" json:"wallets,omitempty"`
}

func (m *WalletAccounts) Reset()                    { *m = WalletAccounts{} }
func (m *WalletAccounts) String() string            { return proto.CompactTextString(m) }
func (*WalletAccounts) ProtoMessage()               {}
func (*WalletAccounts) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *WalletAccounts) GetWallets() []*WalletAccount {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletAccount struct {
	Acc   *Account `protobuf:"bytes,1,opt,name=acc" json:"acc,omitempty"`
	Label string   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *WalletAccount) Reset()                    { *m = WalletAccount{} }
func (m *WalletAccount) String() string            { return proto.CompactTextString(m) }
func (*WalletAccount) ProtoMessage()               {}
func (*WalletAccount) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *WalletAccount) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WalletAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type WalletUnLock struct {
	Passwd  string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
	Timeout int64  `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *WalletUnLock) Reset()                    { *m = WalletUnLock{} }
func (m *WalletUnLock) String() string            { return proto.CompactTextString(m) }
func (*WalletUnLock) ProtoMessage()               {}
func (*WalletUnLock) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *WalletUnLock) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *WalletUnLock) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type GenSeedLang struct {
	Lang int32 `protobuf:"varint,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *GenSeedLang) Reset()                    { *m = GenSeedLang{} }
func (m *GenSeedLang) String() string            { return proto.CompactTextString(m) }
func (*GenSeedLang) ProtoMessage()               {}
func (*GenSeedLang) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *GenSeedLang) GetLang() int32 {
	if m != nil {
		return m.Lang
	}
	return 0
}

type GetSeedByPw struct {
	Passwd string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSeedByPw) Reset()                    { *m = GetSeedByPw{} }
func (m *GetSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*GetSeedByPw) ProtoMessage()               {}
func (*GetSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *GetSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type SaveSeedByPw struct {
	Seed   string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *SaveSeedByPw) Reset()                    { *m = SaveSeedByPw{} }
func (m *SaveSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*SaveSeedByPw) ProtoMessage()               {}
func (*SaveSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *SaveSeedByPw) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *SaveSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ReplySeed struct {
	Seed string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
}

func (m *ReplySeed) Reset()                    { *m = ReplySeed{} }
func (m *ReplySeed) String() string            { return proto.CompactTextString(m) }
func (*ReplySeed) ProtoMessage()               {}
func (*ReplySeed) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *ReplySeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type ReqWalletSetPasswd struct {
	Oldpass string `protobuf:"bytes,1,opt,name=oldpass" json:"oldpass,omitempty"`
	Newpass string `protobuf:"bytes,2,opt,name=newpass" json:"newpass,omitempty"`
}

func (m *ReqWalletSetPasswd) Reset()                    { *m = ReqWalletSetPasswd{} }
func (m *ReqWalletSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetPasswd) ProtoMessage()               {}
func (*ReqWalletSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *ReqWalletSetPasswd) GetOldpass() string {
	if m != nil {
		return m.Oldpass
	}
	return ""
}

func (m *ReqWalletSetPasswd) GetNewpass() string {
	if m != nil {
		return m.Newpass
	}
	return ""
}

type ReqNewAccount struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
}

func (m *ReqNewAccount) Reset()                    { *m = ReqNewAccount{} }
func (m *ReqNewAccount) String() string            { return proto.CompactTextString(m) }
func (*ReqNewAccount) ProtoMessage()               {}
func (*ReqNewAccount) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *ReqNewAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type MinerFlag struct {
	Flag int32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
}

func (m *MinerFlag) Reset()                    { *m = MinerFlag{} }
func (m *MinerFlag) String() string            { return proto.CompactTextString(m) }
func (*MinerFlag) ProtoMessage()               {}
func (*MinerFlag) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{14} }

func (m *MinerFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

type ReqWalletTransactionList struct {
	FromTx    []byte `protobuf:"bytes,1,opt,name=fromTx,proto3" json:"fromTx,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqWalletTransactionList) Reset()                    { *m = ReqWalletTransactionList{} }
func (m *ReqWalletTransactionList) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletTransactionList) ProtoMessage()               {}
func (*ReqWalletTransactionList) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{15} }

func (m *ReqWalletTransactionList) GetFromTx() []byte {
	if m != nil {
		return m.FromTx
	}
	return nil
}

func (m *ReqWalletTransactionList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqWalletTransactionList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqWalletImportPrivKey struct {
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletImportPrivKey) Reset()                    { *m = ReqWalletImportPrivKey{} }
func (m *ReqWalletImportPrivKey) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletImportPrivKey) ProtoMessage()               {}
func (*ReqWalletImportPrivKey) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{16} }

func (m *ReqWalletImportPrivKey) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqWalletImportPrivKey) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletSendToAddress struct {
	From        string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To          string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note        string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
	Istoken     bool   `protobuf:"varint,5,opt,name=istoken" json:"istoken,omitempty"`
	TokenSymbol string `protobuf:"bytes,6,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
}

func (m *ReqWalletSendToAddress) Reset()                    { *m = ReqWalletSendToAddress{} }
func (m *ReqWalletSendToAddress) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSendToAddress) ProtoMessage()               {}
func (*ReqWalletSendToAddress) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{17} }

func (m *ReqWalletSendToAddress) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqWalletSendToAddress) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetIstoken() bool {
	if m != nil {
		return m.Istoken
	}
	return false
}

func (m *ReqWalletSendToAddress) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

type ReqWalletSetFee struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReqWalletSetFee) Reset()                    { *m = ReqWalletSetFee{} }
func (m *ReqWalletSetFee) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetFee) ProtoMessage()               {}
func (*ReqWalletSetFee) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{18} }

func (m *ReqWalletSetFee) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReqWalletSetLabel struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletSetLabel) Reset()                    { *m = ReqWalletSetLabel{} }
func (m *ReqWalletSetLabel) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetLabel) ProtoMessage()               {}
func (*ReqWalletSetLabel) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{19} }

func (m *ReqWalletSetLabel) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqWalletSetLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletMergeBalance struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *ReqWalletMergeBalance) Reset()                    { *m = ReqWalletMergeBalance{} }
func (m *ReqWalletMergeBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletMergeBalance) ProtoMessage()               {}
func (*ReqWalletMergeBalance) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{20} }

func (m *ReqWalletMergeBalance) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReqStr struct {
	Reqstr string `protobuf:"bytes,1,opt,name=reqstr" json:"reqstr,omitempty"`
}

func (m *ReqStr) Reset()                    { *m = ReqStr{} }
func (m *ReqStr) String() string            { return proto.CompactTextString(m) }
func (*ReqStr) ProtoMessage()               {}
func (*ReqStr) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{21} }

func (m *ReqStr) GetReqstr() string {
	if m != nil {
		return m.Reqstr
	}
	return ""
}

type ReplyStr struct {
	Replystr string `protobuf:"bytes,1,opt,name=replystr" json:"replystr,omitempty"`
}

func (m *ReplyStr) Reset()                    { *m = ReplyStr{} }
func (m *ReplyStr) String() string            { return proto.CompactTextString(m) }
func (*ReplyStr) ProtoMessage()               {}
func (*ReplyStr) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{22} }

func (m *ReplyStr) GetReplystr() string {
	if m != nil {
		return m.Replystr
	}
	return ""
}

type ReqTokenPreCreate struct {
	CreatorAddr  string `protobuf:"bytes,1,opt,name=creator_addr,json=creatorAddr" json:"creator_addr,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Symbol       string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
	Introduction string `protobuf:"bytes,4,opt,name=introduction" json:"introduction,omitempty"`
	OwnerAddr    string `protobuf:"bytes,5,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
	Total        int64  `protobuf:"varint,6,opt,name=total" json:"total,omitempty"`
	Price        int64  `protobuf:"varint,7,opt,name=price" json:"price,omitempty"`
}

func (m *ReqTokenPreCreate) Reset()                    { *m = ReqTokenPreCreate{} }
func (m *ReqTokenPreCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenPreCreate) ProtoMessage()               {}
func (*ReqTokenPreCreate) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{23} }

func (m *ReqTokenPreCreate) GetCreatorAddr() string {
	if m != nil {
		return m.CreatorAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTokenPreCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenPreCreate) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *ReqTokenPreCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReqTokenPreCreate) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ReqSellToken struct {
	Sell  *TradeForSell `protobuf:"bytes,1,opt,name=sell" json:"sell,omitempty"`
	Owner string        `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
}

func (m *ReqSellToken) Reset()                    { *m = ReqSellToken{} }
func (m *ReqSellToken) String() string            { return proto.CompactTextString(m) }
func (*ReqSellToken) ProtoMessage()               {}
func (*ReqSellToken) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{24} }

func (m *ReqSellToken) GetSell() *TradeForSell {
	if m != nil {
		return m.Sell
	}
	return nil
}

func (m *ReqSellToken) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ReqBuyToken struct {
	Buy   *TradeForBuy `protobuf:"bytes,1,opt,name=buy" json:"buy,omitempty"`
	Buyer string       `protobuf:"bytes,2,opt,name=buyer" json:"buyer,omitempty"`
}

func (m *ReqBuyToken) Reset()                    { *m = ReqBuyToken{} }
func (m *ReqBuyToken) String() string            { return proto.CompactTextString(m) }
func (*ReqBuyToken) ProtoMessage()               {}
func (*ReqBuyToken) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{25} }

func (m *ReqBuyToken) GetBuy() *TradeForBuy {
	if m != nil {
		return m.Buy
	}
	return nil
}

func (m *ReqBuyToken) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

type ReqRevokeSell struct {
	Revoke *TradeForRevokeSell `protobuf:"bytes,1,opt,name=revoke" json:"revoke,omitempty"`
	Owner  string              `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
}

func (m *ReqRevokeSell) Reset()                    { *m = ReqRevokeSell{} }
func (m *ReqRevokeSell) String() string            { return proto.CompactTextString(m) }
func (*ReqRevokeSell) ProtoMessage()               {}
func (*ReqRevokeSell) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{26} }

func (m *ReqRevokeSell) GetRevoke() *TradeForRevokeSell {
	if m != nil {
		return m.Revoke
	}
	return nil
}

func (m *ReqRevokeSell) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*WalletTxDetail)(nil), "types.WalletTxDetail")
	proto.RegisterType((*WalletTxDetails)(nil), "types.WalletTxDetails")
	proto.RegisterType((*WalletAccountStore)(nil), "types.WalletAccountStore")
	proto.RegisterType((*WalletPwHash)(nil), "types.WalletPwHash")
	proto.RegisterType((*WalletStatus)(nil), "types.WalletStatus")
	proto.RegisterType((*WalletAccounts)(nil), "types.WalletAccounts")
	proto.RegisterType((*WalletAccount)(nil), "types.WalletAccount")
	proto.RegisterType((*WalletUnLock)(nil), "types.WalletUnLock")
	proto.RegisterType((*GenSeedLang)(nil), "types.GenSeedLang")
	proto.RegisterType((*GetSeedByPw)(nil), "types.GetSeedByPw")
	proto.RegisterType((*SaveSeedByPw)(nil), "types.SaveSeedByPw")
	proto.RegisterType((*ReplySeed)(nil), "types.ReplySeed")
	proto.RegisterType((*ReqWalletSetPasswd)(nil), "types.ReqWalletSetPasswd")
	proto.RegisterType((*ReqNewAccount)(nil), "types.ReqNewAccount")
	proto.RegisterType((*MinerFlag)(nil), "types.MinerFlag")
	proto.RegisterType((*ReqWalletTransactionList)(nil), "types.ReqWalletTransactionList")
	proto.RegisterType((*ReqWalletImportPrivKey)(nil), "types.ReqWalletImportPrivKey")
	proto.RegisterType((*ReqWalletSendToAddress)(nil), "types.ReqWalletSendToAddress")
	proto.RegisterType((*ReqWalletSetFee)(nil), "types.ReqWalletSetFee")
	proto.RegisterType((*ReqWalletSetLabel)(nil), "types.ReqWalletSetLabel")
	proto.RegisterType((*ReqWalletMergeBalance)(nil), "types.ReqWalletMergeBalance")
	proto.RegisterType((*ReqStr)(nil), "types.ReqStr")
	proto.RegisterType((*ReplyStr)(nil), "types.ReplyStr")
	proto.RegisterType((*ReqTokenPreCreate)(nil), "types.ReqTokenPreCreate")
	proto.RegisterType((*ReqSellToken)(nil), "types.ReqSellToken")
	proto.RegisterType((*ReqBuyToken)(nil), "types.ReqBuyToken")
	proto.RegisterType((*ReqRevokeSell)(nil), "types.ReqRevokeSell")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 1024 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x06, 0x25, 0xcb, 0x96, 0x46, 0xb2, 0xf3, 0x66, 0xdf, 0xc4, 0x60, 0x8d, 0xb6, 0x51, 0x16,
	0x4d, 0xe3, 0x02, 0x85, 0x81, 0x26, 0xb7, 0x02, 0x05, 0x62, 0x37, 0x70, 0x1c, 0xd4, 0x0e, 0x8c,
	0x95, 0x8b, 0xf6, 0x56, 0xac, 0xc8, 0xb1, 0x45, 0x98, 0xe2, 0xca, 0xcb, 0xd5, 0xd7, 0x5f, 0xe9,
	0xb9, 0x7f, 0xa8, 0xff, 0xa8, 0x98, 0xd9, 0xa5, 0x44, 0xb6, 0xee, 0xa1, 0xb7, 0x79, 0x66, 0xe7,
	0x63, 0xe7, 0x99, 0xd9, 0x21, 0x61, 0xb0, 0xd4, 0x79, 0x8e, 0xee, 0x64, 0x66, 0x8d, 0x33, 0xa2,
	0xe3, 0xd6, 0x33, 0x2c, 0x8f, 0x9e, 0x3a, 0xab, 0x8b, 0x52, 0x27, 0x2e, 0x33, 0x85, 0x3f, 0x39,
	0xfa, 0xdf, 0x38, 0x37, 0xc9, 0x7d, 0x32, 0xd1, 0x59, 0xa5, 0xd9, 0xd7, 0x49, 0x62, 0xe6, 0x45,
	0x70, 0x3d, 0x3a, 0xc0, 0x15, 0x26, 0x73, 0x67, 0xac, 0xc7, 0xf2, 0xf7, 0x16, 0x1c, 0xfc, 0xc2,
	0xb1, 0x6f, 0x56, 0xef, 0xd1, 0xe9, 0x2c, 0x17, 0x12, 0x5a, 0x6e, 0x15, 0x47, 0xc3, 0xe8, 0xb8,
	0xff, 0x46, 0x9c, 0x70, 0xaa, 0x93, 0x9b, 0x6d, 0x26, 0xd5, 0x72, 0x2b, 0xf1, 0x2d, 0xec, 0x59,
	0x4c, 0x30, 0x9b, 0xb9, 0xb8, 0xd5, 0x30, 0x54, 0x5e, 0xfb, 0x5e, 0x3b, 0xad, 0x2a, 0x13, 0x71,
	0x08, 0xbb, 0x13, 0xcc, 0xee, 0x26, 0x2e, 0x6e, 0x0f, 0xa3, 0xe3, 0xb6, 0x0a, 0x48, 0x3c, 0x83,
	0x4e, 0x56, 0xa4, 0xb8, 0x8a, 0x77, 0x58, 0xed, 0x81, 0xf8, 0x1c, 0x7a, 0x5c, 0x85, 0xcb, 0xa6,
	0x18, 0x77, 0xf8, 0x64, 0xab, 0xa0, 0x58, 0x7a, 0x4a, 0x05, 0xc5, 0xbb, 0x3e, 0x96, 0x47, 0xe2,
	0x08, 0xba, 0xb7, 0xd6, 0x4c, 0x75, 0x9a, 0xda, 0x78, 0x6f, 0x18, 0x1d, 0xf7, 0xd4, 0x06, 0x93,
	0x8f, 0x5b, 0x4d, 0x74, 0x39, 0x89, 0xbb, 0xc3, 0xe8, 0x78, 0xa0, 0x02, 0x12, 0x5f, 0x02, 0xf8,
	0x9a, 0x3e, 0xe9, 0x29, 0xc6, 0x3d, 0xf6, 0xaa, 0x69, 0xe4, 0x39, 0x3c, 0x69, 0x72, 0x53, 0x8a,
	0xb7, 0xd0, 0x73, 0x15, 0x88, 0xa3, 0x61, 0xfb, 0xb8, 0xff, 0xe6, 0x79, 0x28, 0xbd, 0x69, 0xaa,
	0xb6, 0x76, 0x72, 0x01, 0xc2, 0x1f, 0x9e, 0xfa, 0x5e, 0x8c, 0x9c, 0xb1, 0x28, 0x62, 0xd8, 0x9b,
	0xd9, 0x6c, 0x71, 0x8f, 0x6b, 0x26, 0xbb, 0xa7, 0x2a, 0x48, 0xbc, 0xe4, 0x7a, 0x8c, 0x39, 0x73,
	0xdb, 0x53, 0x1e, 0x08, 0x01, 0x3b, 0x5c, 0x5d, 0x9b, 0x95, 0x2c, 0x13, 0x57, 0xc4, 0xca, 0xc8,
	0xe9, 0xe9, 0x8c, 0x59, 0xec, 0xa9, 0xad, 0x42, 0xbe, 0x83, 0x81, 0xcf, 0x7b, 0xbd, 0xbc, 0xa0,
	0x7a, 0x0f, 0x61, 0x77, 0xc6, 0x12, 0x27, 0x1c, 0xa8, 0x80, 0xe8, 0x26, 0x56, 0x17, 0x69, 0xe9,
	0x6c, 0xc8, 0x58, 0x41, 0x99, 0x56, 0x11, 0x46, 0x4e, 0xbb, 0x79, 0x49, 0x11, 0xb2, 0xf2, 0xd2,
	0x24, 0xf7, 0x1c, 0xa1, 0xab, 0x02, 0x12, 0x12, 0x06, 0x59, 0x79, 0x3a, 0x77, 0xe6, 0x2a, 0x2b,
	0xb2, 0xe2, 0x8e, 0xc3, 0x74, 0x55, 0x43, 0x47, 0x59, 0x26, 0xba, 0x1c, 0x21, 0xa6, 0x5c, 0x42,
	0x57, 0x55, 0x50, 0xbe, 0xab, 0x66, 0x30, 0xf0, 0x53, 0x8a, 0x13, 0xd8, 0xf3, 0x13, 0x5f, 0x91,
	0xfc, 0xac, 0x41, 0x72, 0xb0, 0x53, 0x95, 0x91, 0xfc, 0x00, 0xfb, 0x8d, 0x13, 0x31, 0x84, 0xb6,
	0x4e, 0x92, 0x30, 0xc5, 0x07, 0xc1, 0xb9, 0x72, 0xa3, 0xa3, 0xc7, 0x49, 0xde, 0x52, 0xf6, 0x73,
	0xc1, 0x85, 0x11, 0x65, 0xba, 0x2c, 0x97, 0x69, 0xe8, 0x51, 0x40, 0x54, 0x0c, 0xf1, 0x6c, 0xe6,
	0xfe, 0x01, 0xb4, 0x55, 0x05, 0xe5, 0x4b, 0xe8, 0x7f, 0xc0, 0x82, 0xea, 0xba, 0xd4, 0xc5, 0x1d,
	0x75, 0x2d, 0xd7, 0xc5, 0x1d, 0xbb, 0x77, 0x14, 0xcb, 0xf2, 0x15, 0x99, 0x38, 0x32, 0x39, 0x5b,
	0x5f, 0x2f, 0xff, 0x2d, 0x87, 0xfc, 0x1e, 0x06, 0x23, 0xbd, 0xc0, 0x8d, 0x9d, 0x80, 0x9d, 0x92,
	0xd8, 0xf3, 0x56, 0x2c, 0xd7, 0x7c, 0x5b, 0x0d, 0xdf, 0x17, 0xd0, 0x53, 0x38, 0xcb, 0xd7, 0xe4,
	0xfc, 0x98, 0xa3, 0xbc, 0x00, 0xa1, 0xf0, 0x21, 0x34, 0x17, 0xdd, 0xf5, 0xa6, 0x2c, 0x93, 0xa7,
	0x14, 0xa3, 0x9a, 0xc9, 0x00, 0xe9, 0xa4, 0xc0, 0x25, 0x9f, 0x84, 0x19, 0x09, 0x50, 0xbe, 0x82,
	0x7d, 0x85, 0x0f, 0x9f, 0x70, 0x59, 0x71, 0xbf, 0x61, 0x36, 0xaa, 0x33, 0xfb, 0x02, 0x7a, 0x57,
	0x59, 0x81, 0xf6, 0x3c, 0xd7, 0xcc, 0xca, 0x6d, 0xae, 0x37, 0xac, 0x90, 0x2c, 0x6f, 0x21, 0xde,
	0xdc, 0xa8, 0xb6, 0x6f, 0x2e, 0xb3, 0x92, 0x37, 0x08, 0xbd, 0xe6, 0x9b, 0x55, 0x35, 0xb9, 0x1e,
	0x51, 0x2a, 0xce, 0xc9, 0x77, 0xea, 0x28, 0x0f, 0xe8, 0x55, 0xa4, 0x99, 0x45, 0x76, 0xe7, 0x59,
	0xeb, 0xa8, 0xad, 0x42, 0x5e, 0xc0, 0xe1, 0x26, 0xcf, 0xc7, 0xe9, 0xcc, 0x58, 0x77, 0x6d, 0xb3,
	0xc5, 0x4f, 0xb8, 0xfe, 0xaf, 0x2f, 0x52, 0xfe, 0x11, 0xd5, 0x42, 0x8d, 0xb0, 0x48, 0x6f, 0xcc,
	0x69, 0x9a, 0x5a, 0x2c, 0x4b, 0x2e, 0xd0, 0x9a, 0x69, 0x45, 0x39, 0xc9, 0xe2, 0x00, 0x5a, 0xce,
	0x84, 0x08, 0x2d, 0x67, 0x6a, 0xab, 0xac, 0xdd, 0x58, 0x65, 0x02, 0x76, 0x0a, 0xe3, 0x30, 0xbc,
	0x67, 0x96, 0xe9, 0x6a, 0x59, 0xe9, 0xcc, 0x3d, 0x16, 0xbc, 0x12, 0xbb, 0xaa, 0x82, 0x62, 0x08,
	0x7d, 0x16, 0x46, 0xeb, 0xe9, 0xd8, 0xe4, 0xbc, 0x15, 0x7b, 0xaa, 0xae, 0x92, 0xdf, 0xc0, 0x93,
	0x7a, 0xab, 0xcf, 0xb1, 0xbe, 0x45, 0xa3, 0x7a, 0x6a, 0xf9, 0x03, 0x3c, 0xad, 0x9b, 0x5e, 0x36,
	0x16, 0x4f, 0x54, 0x5b, 0x3c, 0x8f, 0x13, 0xf2, 0x1a, 0x9e, 0x6f, 0xdc, 0xaf, 0xd0, 0xde, 0xe1,
	0x99, 0xce, 0x75, 0x91, 0x60, 0x28, 0x3d, 0xaa, 0x4a, 0x97, 0x43, 0xd8, 0x55, 0xf8, 0x30, 0x72,
	0xbc, 0x9b, 0x2d, 0x3e, 0xd0, 0xea, 0x09, 0xc3, 0xef, 0x91, 0xfc, 0x1a, 0xba, 0x7e, 0x80, 0x9d,
	0xa5, 0xdd, 0x6e, 0x49, 0xde, 0x5a, 0x6d, 0xb0, 0xfc, 0x33, 0xe2, 0x2b, 0xdf, 0x50, 0xbd, 0xd7,
	0x16, 0x7f, 0xb4, 0xa8, 0x1d, 0x8a, 0x97, 0x30, 0x48, 0x48, 0x32, 0xf6, 0xb7, 0xda, 0xd5, 0xfb,
	0x41, 0x47, 0x4d, 0x62, 0x96, 0x69, 0xed, 0xb7, 0x02, 0xcb, 0xda, 0x7f, 0x5c, 0x4a, 0x4f, 0xa3,
	0x5f, 0xb2, 0x01, 0xf1, 0x7a, 0x2b, 0x9c, 0x35, 0xe9, 0xdc, 0xcf, 0x94, 0xef, 0x4c, 0x43, 0x27,
	0xbe, 0x00, 0x30, 0xcb, 0x02, 0x43, 0xc2, 0x8e, 0xdf, 0xc5, 0xac, 0x39, 0x0d, 0x84, 0x39, 0xe3,
	0x74, 0x1e, 0x3e, 0x5b, 0x1e, 0x90, 0x76, 0x66, 0xb3, 0x04, 0xf9, 0x93, 0xd5, 0x56, 0x1e, 0xc8,
	0x2b, 0x18, 0x10, 0x3b, 0x98, 0xe7, 0x5c, 0x96, 0x78, 0x4d, 0xef, 0x37, 0xcf, 0xc3, 0x36, 0xfb,
	0xff, 0xf6, 0x9b, 0x9c, 0xe2, 0xb9, 0xb1, 0x64, 0xa7, 0xd8, 0x80, 0xc2, 0x71, 0xc6, 0xaa, 0x2b,
	0x0c, 0xe4, 0x47, 0xe8, 0x2b, 0x7c, 0x38, 0x9b, 0xaf, 0x7d, 0xb4, 0xaf, 0xa0, 0x3d, 0x9e, 0xaf,
	0xff, 0xf9, 0x81, 0xe7, 0x60, 0x67, 0xf3, 0xb5, 0xa2, 0x63, 0x0a, 0x35, 0x9e, 0xaf, 0xb7, 0xa1,
	0x18, 0xc8, 0x5f, 0xf9, 0xad, 0x2b, 0x5c, 0x98, 0x7b, 0xa4, 0xbc, 0xe2, 0x3b, 0x6a, 0x1f, 0xa1,
	0x10, 0xef, 0xb3, 0xbf, 0xc5, 0xdb, 0x9a, 0xaa, 0x60, 0xf8, 0xf8, 0x25, 0xc7, 0xbb, 0xfc, 0x3f,
	0xf2, 0xf6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x1a, 0xa6, 0x84, 0xea, 0x08, 0x00, 0x00,
}
