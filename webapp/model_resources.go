package webapp

import "github.com/jinzhu/gorm"

type QuizTemplate struct {
	gorm.Model
	Name      string
	Questions []*QuestionTemplate `gorm:"many2many:quiz_quesiton_templates;association_autoupdate:false;"`
}

type QuestionTemplate struct {
	gorm.Model
	Text                      string
	Type                      uint // choice / text / ...
	ChoiceAnswerTemplates     []*ChoiceAnswerTemplate
	FlowDiagramAnswerTemplate *FlowDiagramAnswerTemplate
}

type ChoiceAnswerTemplate struct {
	gorm.Model
	QuestionTemplateID uint
	Text               string
	IsCorrect          bool
}

type FlowDiagramAnswerTemplate struct {
	gorm.Model
	QuestionTemplateID uint
	Text               string
}
