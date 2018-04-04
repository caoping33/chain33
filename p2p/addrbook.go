package p2p

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"gitlab.33.cn/chain33/chain33/common/db"
	//pb "gitlab.33.cn/chain33/chain33/types"
)

const (
	Addrkey = "addrs"
	PrivKey = "privkey"
)

func (a *AddrBook) Start() error {
	log.Debug("addrbook start")
	a.loadDb()
	go a.saveRoutine()
	return nil
}

func (a *AddrBook) Close() {
	a.Quit <- struct{}{}
	a.bookDb.Close()

}

//peer address manager
type AddrBook struct {
	mtx      sync.Mutex
	ourAddrs map[string]*NetAddress
	addrPeer map[string]*knownAddress
	filePath string
	key      string
	bookDb   db.DB
	Quit     chan struct{}
}

type knownAddress struct {
	kmtx        sync.Mutex
	Addr        *NetAddress `json:"addr"`
	Attempts    uint        `json:"attempts"`
	LastAttempt time.Time   `json:"lastattempt"`
	LastSuccess time.Time   `json:"lastsuccess"`
}

func (a *AddrBook) GetPeerStat(addr string) *knownAddress {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if peer, ok := a.addrPeer[addr]; ok {
		return peer
	}
	return nil

}

func (a *AddrBook) SetAddrStat(addr string, run bool) (*knownAddress, bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if peer, ok := a.addrPeer[addr]; ok {
		if run {
			peer.markGood()

		} else {
			peer.markAttempt()
		}
		return peer, true
	}
	return nil, false
}

func NewAddrBook(filePath string) *AddrBook {
	peers := make(map[string]*knownAddress, 0)
	a := &AddrBook{
		ourAddrs: make(map[string]*NetAddress),
		addrPeer: peers,
		filePath: filePath,
		Quit:     make(chan struct{}),
	}

	a.Start()
	return a
}

func (a *AddrBook) setKey(key string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	a.key = key
}

func (a *AddrBook) GetKey() string {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	return a.key
}

func (a *AddrBook) initKey() {
	var maxRetry = 10
	key, err := P2pComm.GenPrivkey()
	if err != nil {
		for i := 0; i < maxRetry; i++ {
			key, err = P2pComm.GenPrivkey()
			if err == nil {
				break
			}
			if i == maxRetry-1 && err != nil {
				panic(err.Error())
			}
		}

	}

	a.setKey(hex.EncodeToString(key))
}

func newKnownAddress(addr *NetAddress) *knownAddress {
	return &knownAddress{
		Addr:        addr,
		Attempts:    0,
		LastAttempt: time.Now(),
	}
}

func (ka *knownAddress) markGood() {
	ka.kmtx.Lock()
	defer ka.kmtx.Unlock()
	now := time.Now()
	ka.LastAttempt = now
	ka.Attempts = 0
	ka.LastSuccess = now
}

func (ka *knownAddress) Copy() *knownAddress {
	ka.kmtx.Lock()
	ret := knownAddress{
		kmtx:        sync.Mutex{},
		Addr:        ka.Addr.Copy(),
		Attempts:    ka.Attempts,
		LastAttempt: ka.LastAttempt,
		LastSuccess: ka.LastSuccess,
	}
	ka.kmtx.Unlock()
	return &ret
}

func (ka *knownAddress) markAttempt() {
	ka.kmtx.Lock()
	defer ka.kmtx.Unlock()
	now := time.Now()
	ka.LastAttempt = now
	ka.Attempts += 1

}

func (ka *knownAddress) GetLastOk() time.Time {
	ka.kmtx.Lock()
	defer ka.kmtx.Unlock()
	return ka.LastSuccess
}

func (ka *knownAddress) GetAttempts() uint {
	ka.kmtx.Lock()
	defer ka.kmtx.Unlock()
	return ka.Attempts
}

func (a *AddrBook) ISOurAddress(addr *NetAddress) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if _, ok := a.ourAddrs[addr.String()]; ok {
		return true
	}
	return false
}

func (a *AddrBook) IsOurStringAddress(addr string) bool {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if _, ok := a.ourAddrs[addr]; ok {
		return true
	}
	return false
}

func (a *AddrBook) AddOurAddress(addr *NetAddress) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	log.Debug("Add our address to book", "addr", addr)
	a.ourAddrs[addr.String()] = addr
}

func (a *AddrBook) Size() int {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	return len(a.addrPeer)
}

type addrBookJSON struct {
	Addrs []*knownAddress `json:"addrs"`
}

func (a *AddrBook) saveToDb() {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	addrs := []*knownAddress{}
	for _, ka := range a.addrPeer {

		if len(P2pComm.AddrRouteble([]string{ka.Addr.String()})) == 0 {
			continue
		}
		addrs = append(addrs, ka.Copy())
	}
	if len(addrs) == 0 {
		return
	}
	aJSON := &addrBookJSON{
		Addrs: addrs,
	}

	jsonBytes, err := json.Marshal(aJSON)
	if err != nil {
		log.Error("Failed to save AddrBook to file", "err", err)
		return
	}
	log.Debug("saveToDb", "addrs", string(jsonBytes))
	a.bookDb.Set([]byte(Addrkey), jsonBytes)

}

func (a *AddrBook) Pubkey() (string, error) {
	pubkey, err := P2pComm.Pubkey(a.GetKey())
	if err != nil {
		return "", err
	}

	return pubkey, nil
}

// Returns false if file does not exist.
// cmn.Panics if file is corrupt.

func (a *AddrBook) loadDb() bool {
	a.bookDb = db.NewDB("addrbook", "leveldb", a.filePath, 128)

	privkey := a.bookDb.Get([]byte(PrivKey))

	if len(privkey) == 0 {
		a.initKey()
		a.bookDb.Set([]byte(PrivKey), []byte(a.GetKey()))
		return false
	}
	a.setKey(string(privkey))
	iteror := a.bookDb.Iterator(nil, false)
	for iteror.Next() {
		if string(iteror.Key()) == Addrkey {
			//读取存入的其他节点地址信息
			aJSON := &addrBookJSON{}
			dec := json.NewDecoder(strings.NewReader(string(iteror.Value())))
			err := dec.Decode(aJSON)
			if err != nil {
				log.Crit("Error reading file %s: %v", a.filePath, err)
			}

			for _, ka := range aJSON.Addrs {
				log.Debug("loadDb", "peer", ka)
				a.addrPeer[ka.Addr.String()] = ka
			}
		}
	}
	return true

}

// Save saves the book.
func (a *AddrBook) Save() {
	a.saveToDb()
}

func (a *AddrBook) saveRoutine() {
	dumpAddressTicker := time.NewTicker(120 * time.Second)
	defer dumpAddressTicker.Stop()
out:
	for {
		select {
		case <-dumpAddressTicker.C:
			a.Save()
		case <-a.Quit:
			break out
		}
	}

	log.Info("Address handler done")
}

// NOTE: addr must not be nil
func (a *AddrBook) AddAddress(addr *NetAddress) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	log.Debug("Add address to book", "addr", addr)
	if addr == nil {
		return
	}

	if _, ok := a.ourAddrs[addr.String()]; ok {
		// Ignore our own listener address.
		return
	}
	//已经添加的不重复添加
	if _, ok := a.addrPeer[addr.String()]; ok {
		return
	}

	ka := newKnownAddress(addr)
	a.addrPeer[ka.Addr.String()] = ka
	return
}

func (a *AddrBook) RemoveAddr(peeraddr string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if _, ok := a.addrPeer[peeraddr]; ok {
		delete(a.addrPeer, peeraddr)
	}
}

func (a *AddrBook) GetPeers() []*NetAddress {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	var peerlist []*NetAddress
	for _, peer := range a.addrPeer {
		peerlist = append(peerlist, peer.Addr)
	}
	return peerlist
}

func (a *AddrBook) GetAddrs() []string {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	var addrlist []string
	for _, peer := range a.addrPeer {

		if peer.GetAttempts() == 0 {
			addrlist = append(addrlist, peer.Addr.String())
		}

	}
	return addrlist
}
