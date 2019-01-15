package test_constants

import "BigBang/internal/app/feed_attributes"

const PostHash1 = "0xTestPostHash001"
const PostHash2 = "0xTestPostHash002"
const PostHash3 = "0xTestPostHash003"
const PostHash4 = "0xTestPostHash004"
const PostHash5 = "0xTestPostHash005"
const PostHash6 = "0xTestPostHash006"
const PostHash7 = "0xTestPostHash007"

const BoardId1 = "0xBoardId001"
const BoardId2 = "0xBoardId002"
const BoardId3 = "0xBoardId003"

const EmptyParentHash = "0x0000000000000000000000000000000000000000000000000000000000000000"

const PostTypeHash = "0x2fca5a5e"

var PostContent1 = feed_attributes.Content{
	Title:    "Title1",
	Text:     "Text1",
	Image:    "Image1",
	Subtitle: "Subtitle1",
	Meta:     "[{offset: 0, length: 41, type: 'url'}]",
}

var PostContent2 = feed_attributes.Content{
	Title:    "Title2",
	Text:     "Text2",
	Image:    "Image2",
	Subtitle: "Subtitle2",
	Meta:     "[{offset: 0, length: 41, type: 'url'}]",
}

var SessionContent1 = feed_attributes.Content{
	Title:    "SessionTitle1",
	Text:     "SessionText1",
	Image:    "SessionImage1",
	Subtitle: "SessionSubtitle1",
	Meta:     "[{offset: 0, length: 41, type: 'url'}]",
}

var SessionContent2 = feed_attributes.Content{
	Title:    "Sessionitle2",
	Text:     "SessionText2",
	Image:    "SessionImage2",
	Subtitle: "SessionSubtitle2",
	Meta:     "[{offset: 0, length: 41, type: 'url'}]",
}

const SessionStartTime1 = 1539108006
const SessionEndTime1 = 1539108008

const SessionStartTime2 = 1539108012
const SessionEndTime2 = 1539108018

const PostBlockTimestamp1 = 1539108001
const PostBlockTimestamp2 = 1539108002
