package httpclient

import (
	"context"
	"net/http"

	"github.com/tccav/student-creator-dispatcher/pkg/domain/students"
)

type StudentsClient struct {
	client DefaultClient
}

func NewStudentsClient(baseURL string) (StudentsClient, error) {
	client, err := NewDefaultClient(baseURL)
	if err != nil {
		return StudentsClient{}, err
	}
	return StudentsClient{client: client}, nil
}

func (s StudentsClient) Register(ctx context.Context, input students.RegisterInput) error {
	const path = "/students/"
	u, err := s.client.baseURL.Parse(path)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		// TODO: handle errors
	}

	return nil
}
