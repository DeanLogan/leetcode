package main

type questionResponse struct {
	Data struct {
		Question *question `json:"question"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type question struct {
	QuestionID       string   `json:"questionId"`
	Title            string   `json:"title"`
	TitleSlug        string   `json:"titleSlug"`
	Content          string   `json:"content"`
	Difficulty       string   `json:"difficulty"`
	ExampleTestcases string   `json:"exampleTestcases"`
	Hints            []string `json:"hints"`
	TopicTags        []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"topicTags"`
}
