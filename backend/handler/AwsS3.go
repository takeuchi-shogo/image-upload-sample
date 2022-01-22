package handler

import (
	"log"
	"os"

	"golang-image-upload/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Upload() {
	c := config.NewConfig()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     c.AwsS3.AccessKeyID,
				SecretAccessKey: c.AwsS3.SecretAccessKey,
			}),
			Region: aws.String(c.AwsS3.Region),
		},
		Profile: "default",
	}))

	targetFilePath := "./sample.txt"
	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal("err: ", err)
	}
	defer file.Close()

	bucketName := c.AwsS3.Bucket // 保存するバケットの名前
	objectKey := "sample-key"    // ここはS3に保存される名前になる

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}

// func S3Download() {
// 	c := config.NewConfig()

// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		Config: aws.Config{
// 			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
// 				AccessKeyID:     c.AwsS3.AccessKeyID,
// 				SecretAccessKey: c.AwsS3.SecretAccessKey,
// 			}),
// 			Region: aws.String(c.AwsS3.Region),
// 		},
// 		Profile: "default",
// 	}))

// 	d := s3manager.NewDownloader(sess)

// 	// ダウンロードしたデータを入れるものが必要
// 	buf := aws.NewWriteAtBuffer([]byte{})

// 	n, err := d.Download(buf, &s3.GetObjectInput{
// 		Bucket: aws.String(c.AwsS3.Bucket),
// 		Key:    aws.String("BotIcon.jpeg"),
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var img image.Image
// 	fmt.Println(img)

// 	img, err = jpeg.Decode((bytes.NewReader(buf.Bytes())))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	w, _ := os.Create("./img/BotIconDownload.jpeg")
// 	jpeg.Encode(w, img, &jpeg.Options{Quality: 100})

// 	// data := bytes.NewBuffer(buf.Bytes()).String()
// 	// fmt.Println(data)
// 	fmt.Println("size: ", n)

// 	// c.JSON(200, "success!!")

// }
