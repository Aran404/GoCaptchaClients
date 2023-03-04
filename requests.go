package captchaclients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (in *Instance) CreateTask(Task interface{}) error {
	Url := fmt.Sprintf("https://%v/createTask", in.Service)

	Payload := CreateTaskPayload{
		ClientKey: in.ApiKey,
		Task:      Task,
	}

	Marshalled, err := json.Marshal(&Payload)

	if err != nil {
		return err
	}

	Response, err := in.PostJson(Url, bytes.NewReader(Marshalled))

	if err != nil {
		return err
	}

	if Response["errorId"].(float64) != 0 {
		return fmt.Errorf("%v", Response["errorId"])
	}

	in.TaskId = fmt.Sprintf("%v", Response["taskId"])

	return nil
}

func (in *Instance) GetTaskResult(retries int) GetTaskResultResponse {
	start := time.Now()

	Url := fmt.Sprintf("https://%v/getTaskResult", in.Service)

	for i := 0; i < retries; i++ {
		Payload := GetResultPayload{
			ClientKey: in.ApiKey,
			TaskId: func() interface{} {
				if strings.Contains(in.TaskId, "-") {
					return fmt.Sprintf("%v", in.TaskId)
				} else {
					if strings.Contains(in.TaskId, "e") {
						parsed, _, _ := big.ParseFloat(in.TaskId, 10, 0, big.ToNearestEven)

						var descience = new(big.Int)
						descience, _ = parsed.Int(descience)

						return descience
					} else {
						parsed, _ := strconv.Atoi(in.TaskId)

						return parsed
					}
				}
			}(),
		}

		Marshalled, err := json.Marshal(&Payload)

		if err != nil {
			return GetTaskResultResponse{Error: err}
		}

		Response, err := in.PostJson(Url, bytes.NewReader(Marshalled))

		if err != nil {
			return GetTaskResultResponse{Error: err}
		}

		if Response["errorId"].(float64) != 0 {
			return GetTaskResultResponse{
				ErrorId: fmt.Sprintf("%v", Response["errorId"].(float64)),
			}
		}

		if Response["status"] == "ready" {
			return GetTaskResultResponse{
				HcaptchaToken: func(res map[string]interface{}) string {
					if val, ok := res["gRecaptchaResponse"]; ok {
						return val.(string)
					} else {
						return res["solution"].(map[string]interface{})["gRecaptchaResponse"].(string)
					}
				}(Response),
				SolveTime: time.Since(start),
			}
		}

		time.Sleep(time.Second)
	}

	return GetTaskResultResponse{}
}

func (in *Instance) PostJson(url string, body io.Reader) (map[string]interface{}, error) {
	// Captcha Services That Use Json
	resp, err := http.Post(url, "application/json", body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var Response map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&Response); err != nil {
		return nil, err
	}

	return Response, nil
}
