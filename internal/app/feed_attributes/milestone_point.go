package feed_attributes

type MilestonePoint int64

func (mp MilestonePoint) Value() int64 {
	return int64(mp)
}
