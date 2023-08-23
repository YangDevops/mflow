package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mflow/apps/trigger"
	"github.com/infraboard/mflow/test/tools"
)

func TestHandleGitlabPushEvent(t *testing.T) {
	payload, err := tools.ReadContentFile("../impl/test/gitlab_push.json")
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, impl.APIPrefix()+"/gitlab", bytes.NewReader(payload))

	// 补充query参数
	qs := make(url.Values)
	qs.Add("skip_run_pipeline", "true")
	req.URL.RawQuery = qs.Encode()

	t.Logf("url: %s", req.URL)
	// 添加Header头
	req.Header.Set(trigger.GITLAB_HEADER_EVENT_NAME, "Push Hook")
	req.Header.Set(trigger.GITLAB_HEADER_EVENT_UUID, "1234")
	req.Header.Set(trigger.GITLAB_HEADER_EVENT_TOKEN, "my-secret")

	resp := httptest.NewRecorder()

	impl.HandleGitlabEvent(restful.NewRequest(req), restful.NewResponse(resp))

	if resp.Code != http.StatusOK {
		t.Log(resp.Body)
		t.Fatalf("Expected response code %d. Got %d.", http.StatusOK, resp.Code)
	}

	ins := trigger.NewDefaultRecord()
	err = json.Unmarshal(resp.Body.Bytes(), ins)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body. Error: %v", err)
	}

	t.Log(ins)
}

func TestMannulGitlabEvent(t *testing.T) {
	payload, err := tools.ReadContentFile("test/gitlab_mannul_push.json")
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, impl.APIPrefix()+"/mannul", bytes.NewReader(payload))
	t.Logf("url: %s", req.URL)

	resp := httptest.NewRecorder()
	impl.MannulGitlabEvent(restful.NewRequest(req), restful.NewResponse(resp))

	if resp.Code != http.StatusOK {
		t.Log(resp.Body)
		t.Fatalf("Expected response code %d. Got %d.", http.StatusOK, resp.Code)
	}

	ins := trigger.NewDefaultRecord()
	err = json.Unmarshal(resp.Body.Bytes(), ins)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body. Error: %v", err)
	}

	t.Log(ins)
}
