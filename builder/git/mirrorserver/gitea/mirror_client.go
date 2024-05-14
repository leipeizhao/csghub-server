package gitea

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/pulltheflower/gitea-go-sdk/gitea"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/config"
)

type MirrorClient struct {
	giteaClient *gitea.Client
	config      *config.Config
}

type TokenResponse struct {
	SHA1    string `json:"sha1"`
	Message string `json:"message"`
}

func NewMirrorClient(config *config.Config) (client *MirrorClient, err error) {
	ctx := context.Background()
	token, err := findOrCreateAccessToken(ctx, config)
	if err != nil {
		slog.Error("Failed to find or create token", slog.String("error: ", err.Error()))
		return nil, err
	}
	giteaClient, err := gitea.NewClient(
		config.MirrorServer.Host,
		gitea.SetContext(ctx),
		gitea.SetToken(token.Token),
		gitea.SetBasicAuth(config.MirrorServer.Username, config.MirrorServer.Password),
	)
	if err != nil {
		return nil, err
	}

	return &MirrorClient{giteaClient: giteaClient, config: config}, nil
}

func findOrCreateAccessToken(ctx context.Context, config *config.Config) (*database.GitServerAccessToken, error) {
	gs := database.NewGitServerAccessTokenStore()
	tokens, err := gs.FindByType(ctx, "mirror")
	if err != nil {
		slog.Error("Fail to get mirror server access token from database", slog.String("error: ", err.Error()))
		return nil, err
	}

	if len(tokens) == 0 {
		access_token, err := generateAccessTokenFromGitea(config)
		if err != nil {
			slog.Error("Fail to create mirror server access token", slog.String("error: ", err.Error()))
			return nil, err
		}
		gToken := &database.GitServerAccessToken{
			Token:      access_token,
			ServerType: database.MirrorServer,
		}

		gToken, err = gs.Create(ctx, gToken)
		if err != nil {
			slog.Error("Fail to create mirror server access token", slog.String("error: ", err.Error()))
			return nil, err
		}

		return gToken, nil
	}
	return &tokens[0], nil
}

func encodeCredentials(username, password string) string {
	credentials := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(credentials))
}

func generateAccessTokenFromGitea(config *config.Config) (string, error) {
	username := config.MirrorServer.Username
	password := config.MirrorServer.Password
	giteaUrl := fmt.Sprintf("%s/api/v1/users/%s/tokens", config.MirrorServer.Host, username)
	authHeader := encodeCredentials(username, password)
	data := map[string]any{
		"name": "access_token",
		"scopes": []string{
			"write:user",
			"write:admin",
			"write:organization",
			"write:repository",
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error encoding JSON data:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", giteaUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+authHeader)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body:", err)
		return "", err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		slog.Error("Error decoding JSON response:", err)
		return "", err
	}

	if tokenResponse.Message != "" {
		return "", errors.New(tokenResponse.Message)
	}

	return tokenResponse.SHA1, nil
}
