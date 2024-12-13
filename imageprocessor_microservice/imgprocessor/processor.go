package imgprocessor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/SabariGanesh-K/prod-img-proc-service-go/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"github.com/rs/zerolog/log"
)

func ProcessImage(message []byte) {
    config, err := util.LoadConfig(".")
    // 1. Parse the JSON payload
    var data map[string]string
    err = json.Unmarshal(message, &data)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to parse JSON message" )
        return
    }

    imageURL := data["url"]
    productID := data["product_id"]

    // 2. Download the image
    log.Info().Msg("Downloading image" )
    resp, err := http.Get(imageURL)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to download image"  )
        return
    }
    defer resp.Body.Close()

    // 3. Decode the image
    img, _, err := image.Decode(resp.Body)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to decode image"  )
        return
    }

    // 4. Resize the image
    resizedImg := resize.Resize(300, 0, img, resize.Lanczos3)

    // 5. Compress the image
    buf := new(bytes.Buffer)
    err = jpeg.Encode(buf, resizedImg, nil)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to encode image"  )
        return
    }

	

   
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

	bucketName := config.AwsBucketName
	uploader := s3manager.NewUploader(sess)
    log.Info().Msg("Compressed image. Uploading to S3...." )
    // 6. Upload to S3
    key := fmt.Sprintf("compressed-%s.jpg", util.RandomString(20))
    result, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
        Body:   buf,
    })
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to upload image to S3"  )
        return
    }
    compressedImageURL := result.Location
    log.Info().Msg("Uploaded to S3. Updating Database..." )

    // 7. Update the database (call the API endpoint)
    apiURL := fmt.Sprintf("http://0.0.0.0:8083/products/addcompressed")
    requestBody, err := json.Marshal(map[string]interface{}{
        "compressed_product_images_urls": []string{compressedImageURL},
        "id":                              productID,
    })
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to marshal request body" )
        return
    }


    resp, err = http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to update database"  )
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal().Err(err).Msg("Database update failed")
        return
    }
    log.Info().Msg("Database with compressed Image updated. Message cleared." )


}