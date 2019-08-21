package poker_service

import "sort"

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
	for i := 0; i < len; i+=2 {
		if pokerHandA[i] == 88 { //A处理存在赖子的情况
			poker.hashArrayA[13] = 1
		}
		if pokerHandB[i] == 88 { //B处理存在的情况
			poker.hashArrayB[13] = 1
		}

			poker.hashArrayA[Grade[string(pokerHandA[i])]]++
			poker.hashArrayB[Grade[string(pokerHandB[i])]]++ //这里应该用三目运算符
			poker.pokerAFace[GradeFace[string(pokerHandA[i+1])]]++
			if i >= 9 && (poker.pokerAFace[GradeFace[string(pokerHandA[i+1])]] == 5 || (poker.pokerAFace[GradeFace[string(pokerHandA[i+1])]] == 4 && poker.hashArrayA[13] == 1)){ //有待优化
				poker.hashArrayA[14] = 1
			}
			poker.pokerBFace[GradeFace[string(pokerHandB[i+1])]]++
			if i >= 9 && (poker.pokerBFace[GradeFace[string(pokerHandB[i+1])]] == 5 || (poker.pokerBFace[GradeFace[string(pokerHandB[i+1])]] == 4 && poker.hashArrayB[13] == 1)) {
				poker.hashArrayB[14] = 1
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
		fourPoker(num.fourValue, num.onevalue[num.one-1], poker[13])
	} else if num.treechValue[0] != 0 {
		threeZones2(num, poker)
	} else if num.twovalue[0] != 0 {
		determineTheir(num, poker)
	} else {
		disorderly(num, poker)
	}
}

//这里判断三带二
func threeZones2(num *PokerDate, poker *[15]int) (result [2]int ){
	threepok := num.treechValue[num.three-1]
	if num.twovalue[0] != 0 { //这里是三代一对
		if poker[13] == 1 { //说明这里有一个赖子
			//升级为四带一
			result[0] = QUARTIC
			if num.onevalue[0] > num.twovalue[0] {
				result[1] = threepok*10 +num.onevalue[0]
			}else {
				result[1] = threepok*10 +num.twovalue[0]
			}
			return
		}
		result[0] = THREE_ZONES    //普通的满堂彩
		result[0] = threepok*10 + num.twovalue[0]
		return
	}else if num.three == 2{
		if poker[13] == 1 {
			result[0] = QUARTIC
			result[1] = threepok*10 +num.treechValue[0]
			return
		}else {
			result[0] = THREE_ZONES
			result[1] = threepok*10 +num.treechValue[0]
			return
		}
	}

	//这里是三条
	//是同花顺
	//同花 顺子 三条 都有可能
	if poker[14] == 1 && poker[13] == 1 && num.one == 3 &&  compareNumber(num.treechValue,num.onevalue,2) { // 带赖子的同花顺
		result[0] = SEQUENCE
		res := compareNumber2(num.treechValue,num.onevalue,2)
			result[1] = res[3]
		return
	} else if poker[14] == 1 && num.one == 4 &&  compareNumber(num.treechValue,num.onevalue,3){ //不带赖子同花顺
		result[0] = SEQUENCE
         // res := compareNumber2(num.treechValue,num.onevalue,2)
		res := compareNumber2(num.treechValue,num.onevalue,3)
		result[1] = res[4]
		return
	}

	//如果存在赖子,择一定转为四张不可能转什么同花,顺子子类的,减少复杂度了
	if poker[13] == 1 { //带赖子的三带二 转四带一
	result[0] = QUARTIC
	result[1] = num.treechValue[0] * 10 + 13
	return
	} else if poker[14] == 1 { //不带赖子,择选择同花
	result [0] = SAMEFLOWER
	res := compareNumber2(num.treechValue,num.onevalue,2)
	result [1] =  res[3]*1000+res[2]*100+res[1]*10+res[0]
	return
	} else if num.one == 4 && compareNumber(num.treechValue,num.onevalue,3) { //不带赖子的顺子
	res := compareNumber2(num.treechValue,num.onevalue,2)
		result[0]= STRAIGHT
	result[1] = res[3]
	return
	} else { //普通的三带二
		result[0]= THREE
		result[1] = num.treechValue[0]*100+num.treechValue[3]*10+num.treechValue[2]
	return
	}
}


func compareNumber2(three [...]int,one [...]int,leng int) (result [...]int) {
	one [leng+1] = three[0]
	//one = one.sort()
	result = sort.Ints(one)
return one
}

func compareNumber(three [...]int,one [...]int,leng int) bool {
	if three[0] >one[leng]{
		if three[0] - one[0] == 3  {
			return true
		}else if three[0] - one[0] == 4  {
			return true
		}
		return false
	}else if three[0] < one[0] {
		if one[leng] - three[0] == 3 {
			return true
		}else if one[leng] - three[0] == 4  {
			return true
		}
		return false
	}else {
		if one[leng] - one[0] == 3{
			return true
		}else if one[leng] - one[0] == 4  {
			return true
		}
		return false
	}

}


//这里判断四条,不可能有同花 可能有赖子
func fourPoker(four, one, rascally int)(result [2]int ){
	//如果是四条则赖子最大是A四带一就带一
	if rascally == 1 {
		if four != 13 { //带赖子 带赖子并且不是四个A的情况
		  result[0] = QUARTIC
		  result[1] = four *10 + 13
		  return
		}
		result[1] = four *10 + 12
		result[0] = QUARTIC
		return
		//带赖子,是四个A的情况
	} else { //不带赖子的情况   四带一
		result[0] = QUARTIC
		result[1] = four + one
		return
	}

}

//这里判断是否是顺子或者散牌
func disorderly(num *PokerDate, poker *[15]int) (result [2]int ){
	//如果 全部是散牌则说明可能为顺子,同花,还有赖子
	//先判断是否有同花顺
	if poker[13] == 1 && poker[14] == 1 && compart3(num.onevalue) { //说明有同花也有赖子子 还应该有顺子
		result[0]=SEQUENCE
		return
	} else if poker[14] == 1 && (num.onevalue[4]-num.onevalue[0] == 4 || num.onevalue[5]-num.onevalue[1] == 4 || num.onevalue[6]-num.onevalue[2] == 4) { //不带赖子的同花顺
	result[0] = SEQUENCE
	return
	}

	//判断同花
	if poker[14] == 1 && poker[13] == 1 { //同花带赖子
	result[0] = SAMEFLOWER
	return
	} else if poker[14] == 1 { //同花不带赖子
		result[0] = SAMEFLOWER
		return
	}

	//判断顺子
	if poker[13] == 1 && num.onevalue[3]-num.onevalue[0] == 3 {
		result[0] = STRAIGHT
		return
	} else if num.onevalue[4]-num.onevalue[0] == 4 || num.onevalue[5]-num.onevalue[1] == 4 {
		result[0] = STRAIGHT
		return
	}

	if poker[13] == 1 { //散牌带赖子 转一对
	    result[0] = TWAIN
	    return
	} else {
		result[0] = SOLA
		return
	}

}

func compart3(input [...]int) bool  {
	if  input[3]-input[0] == 3 || input[4]-input[1] == 3 || input[5]-input[2] == 3{
		return true
	}else  if  input[3]-input[0] == 4 || input[4]-input[1] == 4 || input[5]-input[2] == 4{
		return true
	}
	return false
}

//这个判断对子
func determineTheir(num *PokerDate, poker *[15]int)(result [2]int ) {
	//对子的判断很恐怖
	//先判断是否有同花顺
	switch num.two {
	case 3:
		ThreeTwoPoker(num , poker)
	case 1:
		oneTwoPoker(num , poker)
	default:   //如果能进来说明一定是带有两个的可能了，因为这里一定是有对子的
		twoTwoPoker(num , poker)

	}
}


func oneTwoPoker(num *PokerDate, poker *[15]int)(result [2]int ){
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

func twoTwoPoker(num *PokerDate, poker *[15]int)(result [2]int ) {
	//有两对二，说明还有三张牌，可能存在顺子，同花之类的   //有一种可能性没想到，顺子中间断了，是可以用癞子补上的，凑成同花顺，同花，顺子       该死的癞子很麻烦
	//y有可能存在同花顺
	if poker[13] == 1 && poker[14] == 1 && (num.twovalue[1]-num.twovalue[0] == 1) { //有大问题

	} else if poker[14] == 1 && && (num.twovalue[1]-num.onevalue[0] == 4) { //没有癞子的同花顺

	}

	//同花的可能  癞子凑同花
	if poker[13] == 1 && poker[14] == 1{  //可惜的是前面已经判断了才能得出同花，上面不能这样得出同花，有问题

	}else {

	}
}

func ThreeTwoPoker(num *PokerDate, poker *[15]int) (result [2]int ) {
//如果是散个二不肯能出现顺子，但是可能出现同花 //这里不考虑同花，三代二比同花大
if poker[13] == 1 {  //升级为三带二

}else {  //没有癞子，则说明可能存在同花
if poker[14] == 1{  //同花

}
	//这里是普通的三带二
}
}