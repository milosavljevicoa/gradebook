package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Name 1",
			LastName:  "Last Name 1",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 79,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 89,
				},
				Grade{
					Title: "Homework",
					Type:  GradeHomework,
					Score: 94,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Name 2",
			LastName:  "Last Name 2",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 84,
				},
				Grade{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 86,
				},
				Grade{
					Title: "Homework",
					Type:  GradeHomework,
					Score: 90,
				},
			},
		},
	}
}
