package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 72, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 62, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeBoundaries(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{100, "A"}, {90, "A"}, {89, "B"},
		{80, "B"}, {79, "C"},
		{70, "C"}, {69, "D"},
		{60, "D"}, {59, "F"}, {0, "F"},
	}

	for _, tc := range tests {
		gradeCalculator := NewGradeCalculator()

		gradeCalculator.AddGrade("assignment", tc.score, Assignment)
		gradeCalculator.AddGrade("exam", tc.score, Exam)
		gradeCalculator.AddGrade("essay", tc.score, Essay)

		actual_value := gradeCalculator.GetFinalGrade()

		if tc.expected != actual_value {
			t.Errorf("Score %d: expected '%s'; got '%s' instead", tc.score, tc.expected, actual_value)
		}
	}
}

// Assignments avg 80, exams avg 90, essays avg 100 => 40 + 31.5 + 15 = 86.5 => B
func TestGetGradeAveragesMultipleGradesPerCategory(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 70, Assignment)
	gradeCalculator.AddGrade("assignment 2", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("exam 2", 100, Exam)
	gradeCalculator.AddGrade("essay 1", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeWithNoGrades(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := map[GradeType]string{
		Assignment: "assignment",
		Exam:       "exam",
		Essay:      "essay",
	}

	for gradeType, expected_value := range tests {
		if actual_value := gradeType.String(); expected_value != actual_value {
			t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
		}
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
