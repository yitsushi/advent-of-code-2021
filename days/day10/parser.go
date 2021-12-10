package day10

import "github.com/yitsushi/go-aoc/generic"

type (
	blockKind string
	blockType string
)

const (
	squareBracket    blockKind = "square"      // []
	curlyBracket     blockKind = "curly"       // {}
	angleBracket     blockKind = "angle"       // <>
	parentheses      blockKind = "parentheses" // ()
	unknownBlockType blockKind = "unknown"

	opening              blockType = "opening"
	closing              blockType = "closing"
	unknownBlockPosition blockType = "unknown"
)

// Item is one item XD.
type Item struct {
	Kind blockKind
	Type blockType
}

func typeof(character rune) (Item, error) {
	switch character {
	case '[':
		return Item{Kind: squareBracket, Type: opening}, nil
	case ']':
		return Item{Kind: squareBracket, Type: closing}, nil
	case '{':
		return Item{Kind: curlyBracket, Type: opening}, nil
	case '}':
		return Item{Kind: curlyBracket, Type: closing}, nil
	case '<':
		return Item{Kind: angleBracket, Type: opening}, nil
	case '>':
		return Item{Kind: angleBracket, Type: closing}, nil
	case '(':
		return Item{Kind: parentheses, Type: opening}, nil
	case ')':
		return Item{Kind: parentheses, Type: closing}, nil
	default:
		return Item{
			Kind: unknownBlockType,
			Type: unknownBlockPosition,
		}, UnknownCharacterError{Character: character}
	}
}

func parse(sequence string) (bool, Item, generic.Stack) {
	stack := generic.NewStack()

	for _, character := range sequence {
		item, err := typeof(character)
		if err != nil {
			return false, item, stack
		}

		if item.Type == closing {
			last, ok := stack.Pop().(Item)
			if !ok {
				continue
			}

			if last.Kind != item.Kind {
				return false, item, stack
			}

			continue
		}

		stack.Push(item)
	}

	return true, Item{Kind: "", Type: ""}, stack
}

//nolint:gomnd // I see no reasons to put them in a constant.
func scoreError(item Item) int {
	switch item.Kind {
	case parentheses:
		return 3
	case squareBracket:
		return 57
	case curlyBracket:
		return 1197
	case angleBracket:
		return 25137
	case unknownBlockType:
		return 0
	default:
		return 0
	}
}

//nolint:gomnd // I see no reasons to put them in a constant.
func scoreMissing(item Item) int {
	switch item.Kind {
	case parentheses:
		return 1
	case squareBracket:
		return 2
	case curlyBracket:
		return 3
	case angleBracket:
		return 4
	case unknownBlockType:
		return 0
	default:
		return 0
	}
}
