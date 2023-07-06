package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/students"
)

type CoursesClient struct {
	client DefaultClient
}

func NewCoursesClient(baseURL string) (CoursesClient, error) {
	client, err := NewDefaultClient(baseURL)
	if err != nil {
		return CoursesClient{}, err
	}
	return CoursesClient{client: client}, nil
}

func (c CoursesClient) GetCourse(ctx context.Context, id string) (students.GetCourseOutput, error) {
	path := fmt.Sprintf("/courses/%s", id)
	u, err := c.client.baseURL.Parse(path)
	if err != nil {
		return students.GetCourseOutput{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return students.GetCourseOutput{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return students.GetCourseOutput{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// TODO: handle errors
	}

	var respBody GetCourseResp
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return students.GetCourseOutput{}, err
	}

	return students.GetCourseOutput{
		ID:   respBody.ID,
		Name: respBody.Name,
	}, nil
}
