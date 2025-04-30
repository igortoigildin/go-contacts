package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	authclient "github.com/igortoigildin/go-contacts/gen-auth-client"
)

func main() {
	cfg := authclient.NewConfiguration()
	ctx := context.Background()
	cfg.Servers = authclient.ServerConfigurations{
		authclient.ServerConfiguration{
			URL: "http://localhost:8080",
		},
	}
	client := authclient.NewAPIClient(cfg)

	// Данные для регистрации (можно заменить на реальные)
	registerReq := authclient.RegisterRequest{
		Username: "newuser",
		Password: "newpassword",
	}

	// Отправляем запрос на регистрацию
	resp, httpResp, err := client.AuthAPI.AuthRegisterPost(ctx).Body(registerReq).Execute()

	// Обработка ошибок
	check("register", err, httpResp)

	// Обработка ошибок
	check("register", err, httpResp)

	// Успешный ответ
	fmt.Println("🎉 User registered successfully:")
	fmt.Println("Response: ", resp.GetMessage())

	// --- 1. Login ---
	loginReq := authclient.LoginRequest{
		Username: "newuser",
		Password: "newpassword",
	}

	loginResp, httpResp, err := client.AuthAPI.AuthLoginPost(ctx).Body(loginReq).Execute()
	check("login", err, httpResp)

	fmt.Println("Token:", loginResp.Token)

	// // --- Verify ---
	verifyResp, httpResp, err := client.AuthAPI.AuthVerifyGet(ctx).Execute()
	check("verify", err, httpResp)

	fmt.Println(" Valid:", verifyResp.Valid)
	fmt.Println(" Username:", verifyResp.Username)
}

func check(stage string, err error, resp *http.Response) {
	if err != nil {
		if resp != nil {
			log.Fatalf("[%s] HTTP %d: %v", stage, resp.StatusCode, err)
		}
		log.Fatalf("[%s] error: %v", stage, err)
	}
}
