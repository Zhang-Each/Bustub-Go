package buffer

type Replacer interface {

	Victim(frame_id *frame_id_t) bool

	Pin(frame_id *frame_id_t)

	Unpin(frame_id *frame_id_t)

	Size() int
}