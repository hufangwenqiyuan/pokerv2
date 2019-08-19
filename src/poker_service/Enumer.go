package poker_service

//给牌进行分类，从大到小赋值
const (
	//单张大牌
	SOLA        = iota
	TWAIN       //一对
	TWOPAIRS    //两对
	THREE       //三条
	STRAIGHT    //顺子
	SAMEFLOWER  //同花
	THREE_ZONES //三带二
	QUARTIC     //四条
	SEQUENCE    //同花顺
	ROYALFLUSH  //皇家从花顺
)

var Grade = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}