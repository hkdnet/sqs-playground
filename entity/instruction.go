package entity

type Instruction struct {
	GroupID         string `yaml:"groupID"`
	DeduplicationID string `yaml:"deduplicationID"`
	Message         string `yaml:"message"`
}
