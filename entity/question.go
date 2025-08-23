package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CurrentAnswers  uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	if p <= AnserPossibleD || p >= AnserPossibleA {
		return true
	}
	return false
}

const (
	AnserPossibleA PossibleAnswerChoice = iota + 1
	AnserPossibleB
	AnserPossibleC
	AnserPossibleD
)

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMeed
	QuestionDifficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	if q <= QuestionDifficultyHard && q >= QuestionDifficultyEasy {

		return true
	}
	return false
}
