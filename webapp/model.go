package webapp

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

const questionsPerQuiz = 10

type Quiz struct {
	gorm.Model
	UserID    uint
	User      *User
	Questions []*Question
	Active    bool
}

type Question struct {
	gorm.Model
	QuizID            uint
	Text              string
	Type              uint
	IsAnswered        bool
	ChoiceAnswers     []*ChoiceAnswer
	TextAnswer        *TextAnswer
	FlowDiagramAnswer *FlowDiagramAnswer
}

type ChoiceAnswer struct {
	gorm.Model
	QuestionID uint
	Text       string
	IsCorrect  bool
	IsSelected bool
}

type TextAnswer struct {
	gorm.Model
	QuestionID uint
	Text       string `sql:"size:999999"`
}

type FlowDiagramAnswer struct {
	gorm.Model
	QuestionID uint
	Text       string `sql:"size:999999"`
}

func newQuiz(u *User, noQ int) *Quiz {
	q := &Quiz{
		UserID: u.ID,
		Active: true,
	}

	var qts []QuestionTemplate

	h.db.
		Model(&QuestionTemplate{}).
		Preload("ChoiceAnswerTemplates").
		Order(gorm.Expr("rand()")).
		Limit(noQ).
		Find(&qts)

	for _, qt := range qts {
		qt.addToQuiz(q)
	}

	h.db.Save(&q)

	return q
}

func findQuiz(u *User) (Quiz, error) {
	q := Quiz{
		UserID: u.ID,
		Active: true,
	}
	result := h.db.Where(&q).
		Preload("Questions").
		Preload("Questions.ChoiceAnswers").
		Preload("Questions.TextAnswer").
		Preload("Questions.FlowDiagramAnswer").
		First(&q)

	return q, result.Error
}

func findAllFinishedForUser(u *User) []Quiz {
	var qs []Quiz

	h.db.Model(&Quiz{}).
		Preload("Questions").
		Preload("Questions.ChoiceAnswers").
		Preload("Questions.TextAnswer").
		Preload("Questions.FlowDiagramAnswer").
		Where("user_id = ?", u.ID).
		Where("active = 0").
		Find(&qs)

	return qs
}

func find(id int) Quiz {
	var q Quiz
	h.db.
		Preload("Questions").
		Preload("Questions.ChoiceAnswers").
		Preload("Questions.TextAnswer").
		Preload("Questions.FlowDiagramAnswer").
		First(&q, id)

	return q
}

func (q *Question) markChoicesAsSelected(ids []string) error {
	for _, aID := range ids {
		for _, a := range q.ChoiceAnswers {
			id, err := strconv.Atoi(aID)
			if err != nil {
				return fmt.Errorf("invalid answer id given %v: %v", aID, err)
			}

			if a.ID == uint(id) {
				a.IsSelected = true
			}
		}
	}

	return nil
}

func (q *Question) saveChoices(answerIds []string, quiz *Quiz) error {
	err := q.markChoicesAsSelected(answerIds)
	if err != nil {
		return fmt.Errorf("invalid answers: %v", err)
	}

	q.IsAnswered = true
	h.db.Save(q)

	return nil
}

func (q *Question) saveText(text string, quiz *Quiz) error {
	q.TextAnswer.Text = text
	q.IsAnswered = true
	h.db.Save(q)

	return nil
}

func (q *Question) saveFlowDiagram(text string, quiz *Quiz) error {
	q.FlowDiagramAnswer.Text = text
	q.IsAnswered = true
	h.db.Save(q)

	return nil
}

func (q *Quiz) getNextQuestion() (*Question, error) {
	for _, question := range q.Questions {
		if !question.IsAnswered {
			return question, nil
		}
	}

	return nil, fmt.Errorf("all questions have been answered")
}

func (q *Quiz) close() {
	q.Active = false
	h.db.Save(&q)
}

func (qt QuestionTemplate) addToQuiz(quiz *Quiz) {
	q := &Question{
		IsAnswered: false,
		QuizID:     quiz.ID,
		Text:       qt.Text,
		Type:       qt.Type,
	}

	h.db.Save(&q)

	q.TextAnswer = &TextAnswer{Text: "", QuestionID: q.ID}
	q.FlowDiagramAnswer = &FlowDiagramAnswer{Text: "", QuestionID: q.ID}
	for _, cat := range qt.ChoiceAnswerTemplates {
		ca := &ChoiceAnswer{
			QuestionID: q.ID,
			Text:       cat.Text,
			IsCorrect:  cat.IsCorrect,
			IsSelected: false,
		}
		h.db.Save(&ca)
		q.ChoiceAnswers = append(q.ChoiceAnswers, ca)
	}

	quiz.Questions = append(quiz.Questions, q)
}
