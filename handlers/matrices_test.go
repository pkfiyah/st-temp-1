package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name          string
		handler       func(w http.ResponseWriter, r *http.Request)
		endpoint      string
		testDataPath  string
		expected      string
		expectSuccess bool
	}{
		{
			name:         "Echo Sunny Test",
			handler:      Echo,
			endpoint:     "http://localhost:8080/echo",
			testDataPath: "../testfiles/matrix.csv",
			expected:     "1,2,3\n4,5,6\n7,8,9\n",
		},
		{
			name:         "Invert Sunny Test",
			handler:      Invert,
			endpoint:     "http://localhost:8080/invert",
			testDataPath: "../testfiles/matrix.csv",
			expected:     "1,4,7\n2,5,8\n3,6,9\n",
		},
		{
			name:         "Flatten Sunny Test",
			handler:      Flatten,
			endpoint:     "http://localhost:8080/flatten",
			testDataPath: "../testfiles/matrix.csv",
			expected:     "1,2,3,4,5,6,7,8,9\n",
		},
		{
			name:         "Sum Sunny Test",
			handler:      Sum,
			endpoint:     "http://localhost:8080/sum",
			testDataPath: "../testfiles/matrix.csv",
			expected:     "45\n",
		},
		{
			name:         "Sum Sunny Test",
			handler:      Multiply,
			endpoint:     "http://localhost:8080/multiply",
			testDataPath: "../testfiles/matrix.csv",
			expected:     "362880\n",
		},
		{
			name:         "Echo Heavy Test",
			handler:      Echo,
			endpoint:     "http://localhost:8080/echo",
			testDataPath: "../testfiles/matrix_hardmode.csv",
			expected:     "1,-3,3,5,3\n4,-5.5,charm,100,-94\n7,shoryuken,9,-9,42\na,test,1,08,hadoken\nbulb,squirt,1,-8,2\n",
		},
		{
			name:         "Invert Heavy Test",
			handler:      Invert,
			endpoint:     "http://localhost:8080/invert",
			testDataPath: "../testfiles/matrix_hardmode.csv",
			expected:     "1,4,7,a,bulb\n-3,-5.5,shoryuken,test,squirt\n3,charm,9,1,1\n5,100,-9,08,-8\n3,-94,42,hadoken,2\n",
		},
		{
			name:         "Flatten Heavy Test",
			handler:      Flatten,
			endpoint:     "http://localhost:8080/flatten",
			testDataPath: "../testfiles/matrix_hardmode.csv",
			expected:     "1,-3,3,5,3,4,-5.5,charm,100,-94,7,shoryuken,9,-9,42,a,test,1,08,hadoken,bulb,squirt,1,-8,2\n",
		},
		{
			name:         "Sum Heavy Test",
			handler:      Sum,
			endpoint:     "http://localhost:8080/sum",
			testDataPath: "../testfiles/matrix_hardmode.csv",
			expected:     "72\n",
		},
		{
			name:         "Multiply Heavy Test",
			handler:      Multiply,
			endpoint:     "http://localhost:8080/multiply",
			testDataPath: "../testfiles/matrix_hardmode.csv",
			expected:     "15472622592000\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup data to make Http call to server
			file, err := os.Open(tt.testDataPath)
			if err != nil {
				t.Fatal("Test file not found")
			}

			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", "matrix_test.csv")
			if err != nil {
				t.Fatalf("error: %+v\n", err)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				t.Fatalf("error: %+v\n", err)
			}
			writer.Close()
			req := httptest.NewRequest(http.MethodPost, tt.endpoint, body)
			w := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Handle the call
			tt.handler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error: %+v\n", err)
			}

			if string(data) != tt.expected {
				t.Errorf("Expected: \n%s\nbut got: \n%s", tt.expected, string(data))
			}
		})
	}
}
