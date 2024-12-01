package jack

import (
	"context"
	"fmt"
	jack "server/api"
)

func (s *jackServer) GetStudent(ctx context.Context, request *jack.GetStudentRequest) (*jack.GetStudentResponse, error) {
	testData := map[uint64]*jack.Student{
		1: &jack.Student{
			Id:   1,
			Name: "Nguyen Van A",
			Gpa:  3.4,
		},
		2: &jack.Student{
			Id:   2,
			Name: "Nguyen Van B",
			Gpa:  3.1,
		},
		3: &jack.Student{
			Id:   3,
			Name: "Nguyen Van C",
			Gpa:  1.3,
		},
	}

	student, exist := testData[request.GetStudentId()]
	if !exist {
		return nil, fmt.Errorf("student not found")
	}
	return &jack.GetStudentResponse{Student: student}, nil
}
