package sdk

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
)

/* 图片上传接口,
支持格式有：jpg/jpeg、png等图片格式，注意入参图片必须转码为base64编码
*/
//type GoodsImageUploadResponse struct {
//	ImageUploadResponse ImageUploadResponse `json:"goods_image_upload_response"`
//}

type ImageUploadResponse struct {
	ImageUrl string `json:"image_url"`
}

func (g *GoodsAPI) GoodsImageUpload(imagePath string) (resp ImageUploadResponse, err error) {
	srcByte, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image := base64.StdEncoding.EncodeToString(srcByte)
	params := NewParamsWithType("pdd.goods.image.upload")
	params.Set("image", base64Image)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "goods_image_upload_response")
	json.Unmarshal(bytes, &resp)
	return
}