package dao

type Student struct {
	StudentId string      `json:"studentId" bson:"studentId"`
	ExamId    string      `json:"examId" bson:"examId"`
	CourseId  string      `json:"course" bson:"course"`
	Settings  map[int]int `json:"settings" bson:"settings"`
}

type AutoGenerated struct {
	Id         string     `yaml:"id" json:"id" bson:"id"`
	Course     string     `yaml:"course" json:"course" bson:"course"`
	General    General    `yaml:"general" json:"general"`
	Settings   []Settings `yaml:"settings" json:"settings"`
	Problems   []Problems `yaml:"problems" json:"problems"`
	Autograder Autograder `yaml:"autograder" json:"autograder"`
}

type General struct {
	Description     string `yaml:"description" json:"description" bson:"description"`
	HandinDirectory string `yaml:"handin_directory" json:"handin_directory" bson:"handin_directory"`
	MaxSubmissions  int    `yaml:"max_submissions" json:"max_submissions" bson:"max_submissions"`
	DisableHandins  bool   `yaml:"disable_handins" json:"disable_handins" bson:"disable_handins"`
	DueAt           string `yaml:"due_at" json:"due_at" bson:"due_at"`
	EndAt           string `yaml:"end_at" json:"end_at" bson:"end_at"`
	GradingDeadline string `yaml:"grading_deadline" json:"grading_deadline" bson:"grading_deadline"`
	HandinFilename  string `yaml:"handin_filename" json:"handin_filename" bson:"handin_filename"`
	Writeup         string `yaml:"writeup" json:"writeup" bson:"writeup"`
	MaxSize         int    `yaml:"max_size" json:"max_size" bson:"max_size"`
	CategoryName    string `yaml:"category_name" json:"category_name" bson:"category_name"`
	Name            string `yaml:"name" json:"name" bson:"name"`
	DisplayName     string `yaml:"display_name" json:"display_name" bson:"display_name"`
	MaxGraceDays    int    `yaml:"max_grace_days" json:"max_grace_days" bson:"max_grace_days"`
	Handout         string `yaml:"handout" json:"handout" bson:"handout"`
	HasSvn          bool   `yaml:"has_svn" json:"has_svn" bson:"has_svn"`
	StartAt         string `yaml:"start_at" json:"start_at" bson:"start_at"`
}

type Settings struct {
	Id     int    `yaml:"id" json:"id" bson:"id"`
	Tag    string `yaml:"tag" json:"tag" bson:"tag"`
	Number int    `yaml:"number" json:"number" bson:"number"`
}

type Problems struct {
	Name        string        `yaml:"name" json:"name" bson:"name"`
	Type        string        `yaml:"type" json:"type" bson:"type"`
	Description []Description `yaml:"description" json:"description" bson:"description"`
	MaxScore    int           `yaml:"max_score" json:"max_score" bson:"max_score"`
	Optional    bool          `yaml:"optional" json:"optional" bson:"optional"`
}

type Description struct {
	Name string `yaml:"name" json:"name" bson:"name"`
	//Answer string `yaml:"answer" json:"answer" bson:"answer"`
	Score int `yaml:"score" json:"score" bson:"score"`
}

type Autograder struct {
	AutogradeTimeout int    `yaml:"autograde_timeout" json:"autograde_timeout" bson:"autograde_timeout"`
	AutogradeImage   string `yaml:"autograde_image" json:"autograde_image" bson:"autograde_image"`
	ReleaseScore     bool   `yaml:"release_score" json:"release_score" bson:"release_score"`
}
