package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
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

	postStudentsReq, err := FromRegisterInput(input)
	if err != nil {
		return err
	}

	bodyBytes, err := json.Marshal(postStudentsReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var serviceErr ServiceErrorResp

		err = json.NewDecoder(resp.Body).Decode(&serviceErr)
		if err != nil {
			return err
		}

		return NewHTTPError(u.RequestURI(), serviceErr.Message, serviceErr.Status)
	}

	return nil
}
