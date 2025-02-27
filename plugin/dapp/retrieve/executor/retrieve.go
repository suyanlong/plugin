// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"fmt"
	"github.com/33cn/chain33/common/address"

	dbm "github.com/33cn/chain33/common/db"
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	rt "github.com/33cn/plugin/plugin/dapp/retrieve/types"
)

var (
	minPeriod int64 = 60
	rlog            = log.New("module", "execs.retrieve")
)

var (
	zeroDelay       int64
	zeroPrepareTime int64
	zeroRemainTime  int64
)

var driverName = "retrieve"

//Init retrieve
func Init(name string, cfg *types.Chain33Config, sub []byte) {
	drivers.Register(cfg, GetName(), newRetrieve, cfg.GetDappFork(driverName, "Enable"))
	InitExecType()
}

//InitExecType ...
func InitExecType() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Retrieve{}))
}

// GetName method
func GetName() string {
	return newRetrieve().GetName()
}

// Retrieve def
type Retrieve struct {
	drivers.DriverBase
}

func newRetrieve() drivers.Driver {
	r := &Retrieve{}
	r.SetChild(r)
	r.SetExecutorType(types.LoadExecutorType(driverName))
	return r
}

// GetDriverName method
func (r *Retrieve) GetDriverName() string {
	return driverName
}

// CheckTx nil
func (r *Retrieve) CheckTx(tx *types.Transaction, index int) error {
	return nil
}

func calcRetrieveKey(backupAddr string, defaultAddr string) []byte {
	key := fmt.Sprintf("LODB-retrieve-backup:%s:%s", address.FormatAddrKey(backupAddr),
		address.FormatAddrKey(defaultAddr))
	return []byte(key)
}

func calcRetrieveAssetKey(backupAddr, defaultAddr, assetExec, assetSymbol string) []byte {
	key := fmt.Sprintf("LODB-retrieve-backup-asset:%s:%s:%s:%s", address.FormatAddrKey(backupAddr),
		address.FormatAddrKey(defaultAddr), assetExec, assetSymbol)
	return []byte(key)
}

func getRetrieveAsset(db dbm.KVDB, backupAddr, defaultAddr, assetExec, assetSymbol string) (*rt.RetrieveQuery, error) {
	return getRetrieve(db, calcRetrieveAssetKey(backupAddr, defaultAddr, assetExec, assetSymbol))
}

func getRetrieveInfo(db dbm.KVDB, backupAddr string, defaultAddr string) (*rt.RetrieveQuery, error) {
	return getRetrieve(db, calcRetrieveKey(backupAddr, defaultAddr))
}

func getRetrieve(db dbm.KVDB, key []byte) (*rt.RetrieveQuery, error) {
	info := rt.RetrieveQuery{}
	retInfo, err := db.Get(key)
	if err != nil {
		return nil, err
	}

	err = types.Decode(retInfo, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// CheckReceiptExecOk return true to check if receipt ty is ok
func (r *Retrieve) CheckReceiptExecOk() bool {
	return true
}
