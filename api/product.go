package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/SabariGanesh-K/prod-mgm-go/s3aws"
	"github.com/SabariGanesh-K/prod-mgm-go/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)
type CreateProductRequest struct {
	ID                 string   `form:"id"`
	UserID             string   `form:"user_id"`
	ProductName        string   `form:"product_name"`
	ProductDescription string   `form:"product_description"`
	ProductPrice       string   `form:"product_price"`
	File *multipart.FileHeader `form:"file" `
	// ProductUrls        []string `json:"product_urls"`
}
// add file to s3 codes

func (server *Server) createProduct(ctx *gin.Context){
	var req  CreateProductRequest
	if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	//add file to s3 

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	bucketName := "elasticbeanstalk-us-east-1-686995207617"
	uploader := s3manager.NewUploader(sess)
	file, err := req.File.Open()
	
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	location, err := s3aws.UploadFile(uploader, file, bucketName,util.RandomString(20))
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}
	fmt.Printf("file added in location %s",location)

	//get urls
	// ProductUrls:=[]string{"https://firebasestorage.googleapis.com/v0/b/personal-website-cc143.appspot.com/o/B612_20241011_182211_278.jpg?alt=media&token=7f8c8632-881a-4585-8996-e93927758907"}
	arg:= db.CreateProductParams{
		ID               :req.ID,
	UserID             :req.UserID ,
	ProductName        :req.ProductName,
	ProductDescription :req.ProductDescription,
	ProductPrice      :req.ProductPrice ,
	ProductUrls       :[]string{location},
	}


	
	product , err:= server.store.CreateProduct(ctx,arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//signal rabbitmq

	ctx.JSON(http.StatusOK,product)

	

}
type getProductByProductIDRequest struct {
ID             string   `uri:"id"`

}
func (server *Server) getProductByProductID(ctx *gin.Context){
	var req  getProductByProductIDRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	
	

	product , err:= server.store.GetProductByProductID(ctx,req.ID)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,product)

	

}
type AddCompressedProductImageUrlsByIDRequest struct {
	CompressedProductImagesUrls []string `json:"compressed_product_images_urls"`
	ID                          string   `json:"id"`
}
func (server *Server) addCompressedImagesByProductID (ctx *gin.Context){
	var req AddCompressedProductImageUrlsByIDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg:= db.AddCompressedProductImageUrlsByIDParams{
		CompressedProductImagesUrls: req.CompressedProductImagesUrls,
		ID: req.ID,
	}
	product,err:= server.store.AddCompressedProductImageUrlsByID(ctx,arg)
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,product)
}