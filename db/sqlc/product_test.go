package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/SabariGanesh-K/prod-mgm-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Products{
	user:= createRandomUser(t)
	arg:= CreateProductParams{
		ID               : util.RandomString(20),
	UserID             : user.UserID,
	ProductName        :util.RandomString(20),
	ProductDescription :util.RandomString(20),
	ProductPrice      : "0.02",
	ProductUrls       :[]string{"https://firebasestorage.googleapis.com/v0/b/personal-website-cc143.appspot.com/o/B612_20241011_182211_278.jpg?alt=media&token=7f8c8632-881a-4585-8996-e93927758907"},
	}
	product,err:= testQueries.CreateProduct(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,product)
	require.Equal(t,user.UserID,product.UserID)
	require.Equal(t,arg.ID,product.ID)
	require.Equal(t,arg.ProductName,product.ProductName)
	require.Equal(t,arg.ProductDescription,product.ProductDescription)
	require.Equal(t,arg.ProductPrice,product.ProductPrice)
	require.NotEmpty(t,product.ProductUrls)
	return product
}
func TestCreateProduct(t *testing.T)  {

	createRandomProduct(t)


}


func TestGetProductByProductID(t *testing.T)  {
	

	product1:= createRandomProduct(t)
	product2,err:= testQueries.GetProductByProductID(context.Background(),product1.ID)
	require.NoError(t,err)
	require.NotEmpty(t,product2)
	require.Equal(t,product1.UserID,product2.UserID)
	require.Equal(t,product1.ID,product2.ID)
	require.Equal(t,product1.ProductName,product2.ProductName)
	require.Equal(t,product1.ProductDescription,product2.ProductDescription)
	require.Equal(t,product1.ProductPrice,product2.ProductPrice)
	require.NotEmpty(t,product2.ProductUrls)


}

func TestGetProductsByUserID(t *testing.T)  {
	var lastProduct Products
	for i := 0; i < 10; i++ {
		lastProduct = createRandomProduct(t)
	}

	arg:= GetProductsByUserIDParams{
		UserID: lastProduct.UserID,
		MinPrice: sql.NullString{"0",true,},
		MaxPrice: sql.NullString{"100",true,},
		ProductName: sql.NullString{"",true,},
	}
	product2,err:= testQueries.GetProductsByUserID(context.Background(),arg)

	// accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	for _, product := range product2 {
		require.NotEmpty(t, product)
		require.Equal(t, lastProduct.UserID , product.UserID)
	}




}

func TestAddCompressedProductImageUrlsByID(t *testing.T) {
	
	product1:=createRandomProduct(t)
	arg1:= AddCompressedProductImageUrlsByIDParams{
	CompressedProductImagesUrls       :[]string{"https://firebasestorage.googleapis.com/v0/b/personal-website-cc143.appspot.com/o/B612_20241011_182211_278.jpg?alt=media&token=7f8c8632-881a-4585-8996-e93927758907"},
		
		ID: product1.ID,
	}
	product2,err:= testQueries.AddCompressedProductImageUrlsByID(context.Background(),arg1)
	require.NoError(t,err)
	require.NotEmpty(t,product2.CompressedProductImagesUrls)
}