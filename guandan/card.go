package guandan

import (
	"github.com/ajenpan/poker_algorithm/poker"
)

type GDCards = poker.Cards

type GDDeck = poker.Cards

func NewDeck() *GDDeck {
	deck := poker.NewDeck()
	deck.Append(poker.NewDeck())
	return deck
}

type DeckType int

const (
	DeckPass           DeckType = 0  // 过
	DeckSingle         DeckType = 1  // 单张
	DeckPair           DeckType = 2  // 对子
	DeckThree          DeckType = 3  // 3张
	DeckThree_with_two DeckType = 4  // 3带2
	DeckStraight       DeckType = 5  // 顺子
	DeckFlush_straight DeckType = 6  // 同花顺
	DeckStraight_pair  DeckType = 7  // 连对
	DeckStraight_three DeckType = 8  // 钢板 (2个连续的三张牌)
	DeckBomb4          DeckType = 9  // 4炸
	DeckBomb5          DeckType = 10 // 5炸
	DeckBomb6          DeckType = 11 // 6炸
	DeckBomb7          DeckType = 12 // 7炸
	DeckBomb8          DeckType = 13 // 8炸
	DeckBomb9          DeckType = 14 // 9炸
	DeckBomb10         DeckType = 15 // 10炸
	DeckBomb_joker     DeckType = 16 // 王炸
	DeckWindflow       DeckType = 17 // 接风
)

type DeckPower struct {
	DeckType  DeckType
	DeckValue int
}

// result 1: dp > other, 0: dp == other, -1: dp < other, -2: cannot compare, -3: error
func (dp *DeckPower) Compare(other *DeckPower) int {
	if dp.DeckType == DeckWindflow || other.DeckType == DeckWindflow {
		return -2
	}
	if other.DeckType == DeckPass {
		return 1
	}
	if dp.DeckType == DeckPass {
		return -1
	}

	// same type
	if dp.DeckType == other.DeckType {
		if dp.DeckValue > other.DeckValue {
			return 1
		} else if dp.DeckValue < other.DeckValue {
			return -1
		}
		return 0
	}

	dpIsBomb := false
	if dp.DeckType >= DeckBomb4 && dp.DeckType <= DeckBomb_joker {
		dpIsBomb = true
	}
	otherIsBomb := false
	if other.DeckType >= DeckBomb4 && other.DeckType <= DeckBomb_joker {
		otherIsBomb = true
	}

	if !dpIsBomb && !otherIsBomb {
		// 牌型不同且非炸弹牌型无法比较
		return -2
	}

	// dp是炸弹，other不是炸弹
	if dpIsBomb && !otherIsBomb {
		return 1
	}

	// dp不是炸弹，other是炸弹
	if !dpIsBomb && otherIsBomb {
		return -1
	}

	if dp.DeckType > other.DeckType {
		return 1
	}

	if dp.DeckType < other.DeckType {
		return -1
	}
	return -2
}

func GetDeckType(wildcard poker.CardRank, cards *poker.Cards) DeckType {
	if cards.Size() == 0 {
		return DeckPass
	} else if cards.Size() == 1 {
		return DeckSingle
	}

	normalCards := poker.NewEmptyCards()
	wildCards := poker.NewEmptyCards()

	cardSize := cards.Size()
	suitCnt := make(map[poker.CardSuit]int)
	rankCnt := make(map[poker.CardRank]int)

	for _, card := range cards.Inner {
		if card.Rank() == wildcard {
			wildCards.Push(card)
		} else {
			normalCards.Push(card)
			rankCnt[card.Rank()]++
			suitCnt[card.Suit()]++
		}
	}

	switch cardSize {
	case 2:
		// DeckPair
		if len(rankCnt) == 1 {
			return DeckPair
		} else {
			return DeckPass
		}
	case 3:
		// DeckThree
		if len(rankCnt) == 1 {
			return DeckThree
		} else {
			return DeckPass
		}
	case 4:
		// DeckBomb_joker
		// DeckBomb4
		if suitCnt[poker.JOKER] == 2 {
			return DeckBomb_joker
		}
		if len(rankCnt) == 1 {
			return DeckBomb4
		} else {
			return DeckPass
		}
	case 5:
		// 三带二
		// 顺子
		// 同花顺
		// 5炸
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	default:
	}
	return DeckPass
}