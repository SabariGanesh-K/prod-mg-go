package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/SabariGanesh-K/prod-mgm-go/db/mock"
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/SabariGanesh-K/prod-mgm-go/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateProductAPI(t *testing.T) {
	user, _ := randomUser()
	product := randomProduct(user)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"id":                  product.ID,
				"user_id":             product.UserID,
				"product_name":        product.ProductName,
				"product_description": product.ProductDescription,
				"product_price":       product.ProductPrice,
				"product_urls":        product.ProductUrls,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateProductParams{
					ID:                 product.ID,
					UserID:             product.UserID,
					ProductName:        product.ProductName,
					ProductDescription: product.ProductDescription,
					ProductPrice:       "0.02",
					ProductUrls:        product.ProductUrls,
				}
				store.EXPECT().
					CreateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(product, nil)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/products"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func TestGetProductByProductIDAPI(t *testing.T) {
	user, _ := randomUser()
	product := randomProduct(user)
	testCases := []struct {
		name          string
		ID            string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{name: "OK",
			ID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductByProductID(gomock.Any(), gomock.Any()).Times(1).Return(product, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/products/%s", tc.ID)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}
func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Products) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduct db.Products
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}
func randomProduct(user db.Users) db.Products {
	return db.Products{
		ID:                          util.RandomString(20),
		UserID:                      user.UserID,
		ProductName:                 util.RandomString(20),
		ProductDescription:          util.RandomString(20),
		ProductPrice:                "0.02",
		ProductUrls:                 []string{"https://firebasestorage.googleapis.com/v0/b/personal-website-cc143.appspot.com/o/B612_20241011_182211_278.jpg?alt=media&token=7f8c8632-881a-4585-8996-e93927758907"},
		CompressedProductImagesUrls: []string{},
	}
}
