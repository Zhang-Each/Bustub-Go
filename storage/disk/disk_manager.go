package disk

import (
	"bustub/common"
	"os"
	"sync"
)

type DiskManager struct {
	logIO      *os.File
	dbIO       os.File
	logName    string
	fileName   string
	numFlushes int
	numWrites  int
	flushLog   bool
	dbLatch    sync.Mutex
}

func NewDiskManager(file string) *DiskManager {
	return &DiskManager{fileName: file}
}

func (dm *DiskManager) ShutDown()  {

}

func (dm *DiskManager) WritePage(id common.PageId, data []byte)  {

}

func (dm *DiskManager) ReadPage(id common.PageId, data []byte)  {

}

func (dm *DiskManager) WriteLog(log []byte, size int) {

}

func (dm *DiskManager) ReadLog(log []byte, size int, offset int) bool {
	return false
}

func (dm *DiskManager) GetNumFlushes() int {
	return dm.numFlushes
}

func (dm *DiskManager) GetNumWrites() int {
	return dm.numWrites
}

func (dm *DiskManager) GetFlushState() bool {
	return dm.flushLog
}

