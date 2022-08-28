package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createReq() (*http.Request, context.Context) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	return req.WithContext(ctx), ctx
}

func TestConfig_AddDefaultData(t *testing.T) {
	req, ctx := createReq()

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get error data")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, ctx := createReq()
	auth := testApp.IsAuthenticated(req)

	if auth {
		t.Error("returns true for authenticated, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)
	auth = testApp.IsAuthenticated(req)

	if !auth {
		t.Error("returns false for authenticated, when it should be true")
	}
}

func TestConfig_render(t *testing.T) {
	pathToTemplates = "./templates"
	rr := httptest.NewRecorder()
	req, _ := createReq()

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Error("failed to render page")
	}
}
