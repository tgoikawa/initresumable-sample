package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	adminpub "google.golang.org/genproto/googleapis/iam/admin/v1"

	admin "cloud.google.com/go/iam/admin/apiv1"
	"cloud.google.com/go/storage"

	"net/http"
)

const (
	googleAccessID string = "sandbox-kosukeoikawa@appspot.gserviceaccount.com"
)

func main() {
	http.HandleFunc("/api/resumable-upload-url", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		iamClient, err := admin.NewIamClient(ctx)
		if err != nil {
			fmt.Printf("%s", err)
			w.WriteHeader(500)
			return
		}
		signer := ServiceAccountSigner{Client: iamClient}

		url, err := storage.SignedURL("resumable_test", "object1", &storage.SignedURLOptions{
			GoogleAccessID: googleAccessID,
			Method:         http.MethodPost,
			Expires:        time.Now().AddDate(0, 0, 1),
			ContentType:    "video/mp4",
			SignBytes: func(b []byte) ([]byte, error) {
				return signer.SignByte(ctx, b)
			},
			Headers: []string{"x-goog-resumable:start"}, // Resumable URL生成用
		})

		if err != nil {
			fmt.Printf("%s", err)
			w.WriteHeader(400)
			return
		}

		raw, err := json.Marshal(map[string]interface{}{
			"url": url,
		})

		if err != nil {
			fmt.Printf("%s", err)
			w.WriteHeader(500)
			return
		}

		fmt.Printf("return url")
		w.Write(raw)
		w.WriteHeader(200)

	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe("localhost:"+port, nil)
}

type ServiceAccountSigner struct {
	Client *admin.IamClient
}

func (s *ServiceAccountSigner) SignByte(ctx context.Context, b []byte) ([]byte, error) {
	resp, err := s.Client.SignBlob(ctx, &adminpub.SignBlobRequest{
		Name:        fmt.Sprintf("projects/%s/serviceAccounts/%s", "sandbox-kosukeoikawa", googleAccessID),
		BytesToSign: b,
	})

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	return resp.GetSignature(), nil
}
