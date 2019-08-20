package poker_service

//算法核心 等级不同比较等级，等级相同比较分数
type Poker struct {
	pokerAFace [4]int
	pokerBFace [4]int
	hashArrayA [15]int //这里可以优化,不一定每次进来都创建一个可以尝试放到常量那边
	hashArrayB [15]int
}
type PokerDate struct {
	fourValue   int
	treechValue [...]int
	twovalue    [...]int
	onevalue    [...]int
	three       int
	two         int
	one         int
}

func NewPokerDate() *PokerDate {
	return &PokerDate{}
}

func NewStartPoker() *Poker {
	return &Poker{}
}

func (*Poker) StartPoker(Path string) {
	pokerMap, err := NewFile().ReadFile(Path)
	if err != nil {
		panic(err)
		return
	}
	for _, pokers := range pokerMap["matches"] {
		//解析手牌
		pokerStrA := NewStartPoker().analysisStr(pokers["alice"], pokers["bob"])
		if pokerStrA[0] > pokerStrA[1] {
			//A赢
			pokers["result"] = "1"
		} else if pokerStrA[0] == pokerStrA[1] {
			pokers["result"] = "0"
		} else {
			//B赢
			pokers["result"] = "2"
		}
	}
	//写完后写入文件中
	NewFile().WriteJsonFile("./../source/", &pokerMap)
}

func (poker *Poker) analysisStr(pokerHandA, pokerHandB string) []int {
	len := len(pokerHandA)
	for i := 0; i < len; i++ {
		if pokerHandA[i] == 88 { //A处理存在赖子的情况
			poker.hashArrayA[13] = 1
		}
		if pokerHandB[i] == 88 { //B处理存在的情况
			poker.hashArrayB[13] = 1
		}
		if i&1 == 0 {
			poker.hashArrayA[Grade[string(pokerHandA[i])]]++
			poker.hashArrayB[Grade[string(pokerHandB[i])]]++ //这里应该用三目运算符
		} else {
			poker.pokerAFace[GradeFace[string(pokerHandA[i])]]++
			if i >= 9 && poker.pokerAFace[GradeFace[string(pokerHandA[i])]] >= 5 || (poker.pokerAFace[GradeFace[string(pokerHandA[i])]] >= 4 && poker.hashArrayA[13] == 1) { //有待优化
				poker.hashArrayA[14] = 1
			}
			poker.pokerBFace[GradeFace[string(pokerHandB[i])]]++
			if i >= 9 && poker.pokerBFace[GradeFace[string(pokerHandB[i])]] >= 5 || (poker.pokerBFace[GradeFace[string(pokerHandB[i])]] >= 4 && poker.hashArrayB[13] == 1) {
				poker.hashArrayB[14] = 1
			}
		}
	}
	NewPokerDate().comparativeResult(&poker.hashArrayA) //最大值,赖子.同花
	NewPokerDate().comparativeResult(&poker.hashArrayB)
	return nil
}

func (num *PokerDate) comparativeResult(poker *[15]int) {
	for i := 0; i < 13; i++ {
		switch poker[i] {
		case 1:
			num.onevalue[num.one] = i
			num.one++
		case 2:
			num.twovalue[num.two] = i
			num.two++
		case 3:
			num.treechValue[num.three] = i
			num.three++
		case 4:
			num.fourValue = i

		}
	}

	if num.fourValue != 0 {
		fourPoker(num.fourValue, num.onevalue[len(num.onevalue)-1], poker[13])
	} else if num.treechValue[0] != 0 {
		threeZones2(num, poker)
	} else if num.twovalue[0] != 0 {
		determineTheir(num, poker)
	} else {
		disorderly(num, poker)
	}
}

//这里判断三带二
func threeZones2(num *PokerDate, poker *[15]int) {
	threepok := num.treechValue[num.three-1]
	if num.twovalue[0] != 0 { //这里是三代一对
		if poker[13] == 1 { //说明这里有一个赖子
			//升级为四带一

		}
		//三带一对
	}
	//这里是三条
	//是同花顺
	//同花 顺子 三条 都有可能
	if poker[14] == 1 && poker[13] == 1 && num.one == 4 && (num.treechValue[0]-num.onevalue[0] == 3 || num.onevalue[num.one-1]-num.treechValue[0] == 3) { // 带赖子的同花顺

	} else if poker[14] == 1 && num.one == 4 && (num.treechValue[0]-num.onevalue[0] == 4 || num.onevalue[num.one-1]-num.treechValue[0] == 4) { //不带赖子同花顺

	}

	//如果存在赖子,择一定转为四张不可能转什么同花,顺子子类的,减少复杂度了
	if poker[13] == 1 { //带赖子的三带二 转四带一

	} else if poker[14] == 1 { //不带赖子,择选择同花

	} else if num.one == 4 && (num.treechValue[0]-num.onevalue[0] == 4 || num.onevalue[num.one-1]-num.treechValue[0] == 4) { //不带赖子的顺子

	} else { //普通的三带二

	}
}

//这里判断四条,不可能有同花 可能有赖子
func fourPoker(four, one, rascally int) {
	//如果是四条则赖子最大是A四带一就带一
	if rascally == 1 {
		if four != 13 { //带赖子 带赖子并且不是四个A的情况

		}

		//带赖子,是四个A的情况
	} else { //不带赖子的情况   四带一

	}

}

//这里判断是否是顺子或者散牌
func disorderly(num *PokerDate, poker *[15]int) {
	//如果 全部是散牌则说明可能为顺子,同花,还有赖子
	//先判断是否有同花顺
	if poker[13] == 1 && poker[14] == 1 && (num.onevalue[3]-num.onevalue[0] == 3 || num.onevalue[4]-num.onevalue[1] == 3 || num.onevalue[5]-num.onevalue[2] == 3) { //说明有同花也有赖子子 还应该有顺子

	} else if poker[14] == 1 && (num.onevalue[4]-num.onevalue[0] == 4 || num.onevalue[5]-num.onevalue[1] == 4 || num.onevalue[6]-num.onevalue[2] == 4) { //不带赖子的同花顺

	}

	//判断同花
	if poker[14] == 1 && poker[13] == 1 { //同花带赖子

	} else if poker[14] == 1 { //同花不带赖子

	}

	//判断顺子
	if poker[13] == 1 && num.onevalue[3]-num.onevalue[0] == 3 {

	} else if num.onevalue[4]-num.onevalue[0] == 4 || num.onevalue[5]-num.onevalue[1] == 4 {

	}

	if poker[13] == 1 { //散牌带赖子 转一对

	} else {

	}

}

//这个判断对子
func determineTheir(num *PokerDate, poker *[15]int) {
	//对子的判断很恐怖
	//先判断是否有同花顺
	if poker[13] == 1 && poker[14] == 1 && (num.onevalue[3]-num.onevalue[0] == 3 || num.twovalue[0]-num.onevalue[0] == 4 || num.onevalue[3]-num.twovalue[0] == 4) { //说明有同花也有赖子子 还应该有顺子

	} else if poker[14] == 1 && (num.onevalue[4]-num.onevalue[0] == 4 || num.twovalue[0]-num.onevalue[0] == 4 || num.onevalue[3]-num.twovalue[0] == 4) { //不带赖子的同花顺

	}

	//判断同花
	if poker[14] == 1 && poker[13] == 1 { //同花带赖子

	} else if poker[14] == 1 { //同花不带赖子

	}

	//判断顺子
	if poker[13] == 1 && (num.onevalue[3]-num.onevalue[0] == 3 || num.onevalue[4]-num.onevalue[1] == 3 || num.onevalue[5]-num.onevalue[2] == 3 || num.onevalue[6]-num.onevalue[3] == 3) {

	} else if num.onevalue[4]-num.onevalue[0] == 4 || num.onevalue[5]-num.onevalue[1] == 4 || num.onevalue[6]-num.onevalue[2] == 4 {

	}

	if poker[13] == 1 { //散牌带赖子 转一对

	} else {

	}
}
