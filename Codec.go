package gotextenc

type Codec[SourceT CharLike, TargetT CharLike] interface {
	Transcode([]SourceT, []TargetT, bool) (int, int, error)
	Reset(uint64)
}
