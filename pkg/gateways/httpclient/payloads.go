package httpclient

type GetCourseResp struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MinPeriods int    `json:"min_periods"`
	MaxPeriods int    `json:"max_periods"`
}
