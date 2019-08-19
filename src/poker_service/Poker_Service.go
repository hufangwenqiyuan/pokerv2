package poker_service

//算法核心 等级不同比较等级，等级相同比较分数
type Poker struct {}

func NewStartPoker() *Poker {
	return &Poker{}
}

func (*Poker) StartPoker(Path string) {
	pokerMap, err := NewFile().ReadFile(Path)
	if err != nil {
		panic(err)
		return
	}
	for _,pokers:= range pokerMap["matches"] {
		//解析手牌
		pokerStrA := analysisStr(pokers["alice"])
		pokerStrB :=analysisStr(pokers["bob"])
		if pokerStrA[0] > pokerStrB[0]{
			//A赢
			pokers["result"] = "1"
		}else if pokerStrA[0] == pokerStrB[0]{ //打分？？？？
			//应该在这里给它们打分
			if  pokerStrA[1] > pokerStrB[1]{
				//A赢
				pokers["result"] = "1"
			}else if pokerStrA[1] == pokerStrB[1]{
				//平局
				pokers["result"] = "O"
			}else {
				//B赢
				pokers["result"] = "2"
			}
		}else {
			//B赢
			pokers["result"] = "2"
		}
	}
	//写完后写入文件中
	 NewFile().WriteJsonFile("./../source/result.json",&pokerMap)
}

func analysisStr(player string)(int,[]int) {
	var hashArray [...]int                //定义一个长度为13的数组 做一个实验，清空快还是创建快
	hashArray[Grade[string(player[0])]]++ //这里还必须转换一下为string入股使用rune则不需要
	pokerMap := make(map[uint8]int)
	PokerFace := player[1] //定义一个字符串来接收字母，这样就能判定是否是顺子了,
	length := len(player)
	pokerMap[player[0]]++
	var result  int
	var tierce int //默认为零 记录最后的结果 标记是否是同花
	for i := 2; i < length; i++ {
		if i&1 == 0 { //与运算，这里说明是奇数 也就是数字，牌面
			pokerMap[player[i]]++ //map中对应的值++ 这里不转换是为了加快速度，毕竟转换也需要时间
			hashArray[Grade[string(player[i])]]++   //这里必须优化
			//相同的算等级，不同的算分数
		//	result[1] =+ player[i]  //其实不需要每次都算分数把，毕竟费时  其中有一种特殊的 A2345这是最小的孙子
		} else {
			if tierce == 0 && PokerFace != player[i] {
				tierce = 1 //判断出不是同花，可以标记一下 只要标记不是1则表示是同花
			}
		}
	}
	mapLength := len(pokerMap)
	switch mapLength {
	case 2: //这里有两种可能，第一种四个，三代二 而且必须是一对
		if pokerMap[player[0]] == 2 || pokerMap[player[0]] == 3 { //这里是满长彩  因为现比较前三个所以三个值的权值比二两个高      重13开始
			result = THREE_ZONES
			return result
		}
		result = QUARTIC
		return  result//这里是四个
	case 3: //两对，三带二，而且是不是一对的
		//上面可以不管同花，但这里需要判断，上面判断没意义，因为比同花大
		if tierce == 0 {
			result = SAMEFLOWER
			return result//这里是同花
		}
		if pokerMap[player[0]] == 2 || pokerMap[player[2]] == 2 {
			//这里是两对
			result = TWOPAIRS
			return result
		}
		result = THREE
		return result//三带二
	case 5: //这里可能是顺子或者散牌
	var haaryLenth = len(hashArray)
		if tierce == 0 { //这里是同花顺   因为同花比顺子大
			if 	hashArray[haaryLenth]&hashArray[haaryLenth-9]&hashArray[haaryLenth-10]&hashArray[haaryLenth-11]&hashArray[haaryLenth-12] == 1{ //这里一定需要修改  前往后判断都比这个快
				result = STRAIGHT
				result[1] = 0
				return result
			} else if hashArray[haaryLenth]&hashArray[haaryLenth-1]&hashArray[haaryLenth-2]&hashArray[haaryLenth-3]&hashArray[haaryLenth-4] == 1 { //顺子
				//这里是同花顺
				result = SEQUENCE
				return result
			}
			result = SAMEFLOWER
			return result//这里是同花
		}
	 if	hashArray[haaryLenth]&hashArray[haaryLenth-9]&hashArray[haaryLenth-10]&hashArray[haaryLenth-11]&hashArray[haaryLenth-12] == 1{
		 result = STRAIGHT
		 return result
	} else if  hashArray[haaryLenth]&hashArray[haaryLenth-1]&hashArray[haaryLenth-2]&hashArray[haaryLenth-3]&hashArray[haaryLenth-4] == 1 {//顺子
			result = STRAIGHT
		}
		result = SOLA
		return  result//散牌


	default: //这里只有一种可能是一对
		if tierce == 0 {
			result = SAMEFLOWER
			return result//这里是同花
		}
		result = TWAIN
		return result

	}
}
