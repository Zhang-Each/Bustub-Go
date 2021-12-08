package buffer

import "bustub/common"

type Replacer interface {

	Victim(frameId *common.FrameId) bool

	Pin(frameId *common.FrameId)

	Unpin(frameId *common.FrameId)

	Size() int
}