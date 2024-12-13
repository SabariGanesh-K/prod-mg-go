package api

import (
	"bytes"
	"time"

	"encoding/json"
	"fmt"
	"io"

	// "mime/multipart"
	"net/http"
	"net/http/httptest"

	// "os"
	"testing"

	mockdb "github.com/SabariGanesh-K/prod-mgm-go/db/mock"
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/SabariGanesh-K/prod-mgm-go/util"

	// "github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetProductsByUserIDAPI(t *testing.T) {
	user, _ := randomUser()
	products := randomProducts(user)
	testCases := []struct {
		name          string
		body            gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{name: "OK",
			body:gin.H{
				"user_id":products[0].UserID,
				"min_price":"0",
				"max_price":"100",
				"product_name":"",
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductsByUserID(gomock.Any(), gomock.Any()).Times(1).Return(products, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			
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
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			url := "/products"


			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(data))
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
func BenchmarkGetProductByProductIDAPI(b *testing.B) {
	user, _ := randomUser()
	product := randomProduct(user)

	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	server := newBenchMarkTestServer(b, store)
	productmarshalled, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product,"marshal")
	// Pre-warm the cache (if applicable)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		DialTimeout: 100 * time.Second,

	})
	err = client.Set( product.ID, productmarshalled, 0).Err()
	require.NoError(b, err)
	fmt.Println("benchmark cache saved")

	b.Run("with cache hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/products/%s", product.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(b, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(b, http.StatusOK, recorder.Code)
		}
	})

	b.Run("without cache hit", func(b *testing.B) {
		// Clear the cache (if applicable)
		err := client.Del( product.ID).Err()
		require.NoError(b, err)

		// Stub the database call
		store.EXPECT().
			GetProductByProductID(gomock.Any(), gomock.Eq(product.ID)).
			Times(b.N).
			Return(product, nil)

		for i := 0; i < b.N; i++ {
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/products/%s", product.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(b, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(b, http.StatusOK, recorder.Code)
		}
	})
}
func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Products) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduct db.Products
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}
func randomProducts(user db.Users) []db.Products {
	prods  :=[]db.Products{}
	product1:= randomProduct(user)
	product2:= randomProduct(user)
	prods=append(prods, product1)
	prods=append(prods, product2)


	return prods
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


// func TestCreateProductAPI(t *testing.T) {
// 	user, _ := randomUser()
// 	product := randomProduct(user)
// 	file,err:= os.Open("./2.jpg")
// 	fmt.Print(err,"error")

// 	require.NoError(t,err)
// 	defer file.Close()
// 	filestate,_ := file.Stat()
// 	fileheader:=multipart.FileHeader{
// 		Filename:"2.jpg",
// 		Header:nil,
// 		Size:     filestate.Size(),

// 	}
// 	reqbody := &bytes.Buffer{}
// 	writer := multipart.NewWriter(reqbody)
// 	part, errr := writer.CreateFormFile("file", fileheader.Filename)
// 	require.NoError(t,errr)
// 	_, err = io.Copy(part, file)
// 	require.NoError(t,err)
// 	err = writer.WriteField("id", product.ID) 
// 	require.NoError(t,err)
// 	err = writer.WriteField("user_id", product.UserID) 
// 	require.NoError(t,err)
// 	err = writer.WriteField("product_name", product.ProductName) 
// 	require.NoError(t,err)
// 	err = writer.WriteField("product_description", product.ProductDescription) 
// 	require.NoError(t,err)
// 	err = writer.WriteField("product_price", product.ProductPrice) 
// 	require.NoError(t,err)
// 	err = writer.Close()
// 	require.NoError(t,err)
// 	testCases := []struct {
// 		name          string
// 		body          *bytes.Buffer
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: reqbody,
// 			buildStubs: func(store *mockdb.MockStore) {
			
// 				store.EXPECT().
// 					CreateProduct(gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(product, nil)

// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchProduct(t, recorder.Body, product)
// 			},
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)
// 			recorder := httptest.NewRecorder()
	
// 			server := newTestServer(t, store)

// 			url := "/products"
// 			request, err := http.NewRequest(http.MethodPost, url, tc.body)
// 			require.NoError(t, err)
// 			request.Header.Set("Content-Type", writer.FormDataContentType())

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}

// }
