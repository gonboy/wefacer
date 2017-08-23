package models
import(
	"wefacer/core"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"log"
)


type FaceAuto struct {
	ResultNum int `json:"result_num"`
	Result    []struct {
		Location struct {
			Left   int `json:"left"`
			Top    int `json:"top"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"location"`
		FaceProbability int     `json:"face_probability"`
		RotationAngle   int     `json:"rotation_angle"`
		Yaw             float64 `json:"yaw"`
		Pitch           float64 `json:"pitch"`
		Roll            float64 `json:"roll"`
		Age				float64 `json:"age"`
		Landmark        []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"landmark"`
		Landmark72 []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"landmark72"`
		Gender             string  `json:"gender"`
		Expression		   int 	   `json:"expression"`
		GenderProbability  float64 `json:"gender_probability"`
		Glasses            int     `json:"glasses"`
		GlassesProbability float64 `json:"glasses_probability"`
		Race               string  `json:"race"`
		RaceProbability    float64 `json:"race_probability"`
		Qualities          struct {
			Occlusion struct {
				LeftEye    float64 `json:"left_eye"`
				RightEye   float64 `json:"right_eye"`
				Nose       float64 `json:"nose"`
				Mouth      float64 `json:"mouth"`
				LeftCheek  float64 `json:"left_cheek"`
				RightCheek float64 `json:"right_cheek"`
				Chin       float64 `json:"chin"`
			} `json:"occlusion"`
			Blur         float64 `json:"blur"`
			Illumination int     `json:"illumination"`
			Completeness int     `json:"completeness"`
			Type         struct {
				Human   float64 `json:"human"`
				Cartoon float64 `json:"cartoon"`
			} `json:"type"`
		} `json:"qualities"`
	} `json:"result"`
	LogID int64 `json:"log_id"`
}


type IdentifyFace interface{
	DentifyFace(request IRequest,requesthead RequestHead,requesthandlechan chan FaceAuto,faceAutoerrchan chan bool)
}
//百度++
type BaiduDentifyFace struct{

}
func (baiduDentifyFace BaiduDentifyFace) DentifyFace(request IRequest,requesthead RequestHead,requesthandlechan chan FaceAuto,faceAutoerrchan chan bool){
	resp, err1 := http.Get(request.(ImageRequest).PicUrl)

	content,err2 := ioutil.ReadAll(resp.Body)
	if err1!=nil||err2!=nil{
		faceAutoerrchan<-true
		return
	}
	defer resp.Body.Close()
	basestring:=core.GetImageBase64(content)
	res,err:=http.PostForm(RecognitionUrl,url.Values{"access_token":{AutoGeneratedValue.AccessToken},"image":{basestring},"max_face_num":{"5"},"face_fields":{"beauty,age,expression,faceshape,gender,glasses,landmark,race,qualities"}})
	if err != nil {
		log.Println(err.Error())
	}
	defer res.Body.Close()
	body,_ := ioutil.ReadAll(res.Body)
	var faceAuto FaceAuto
	errUnmarshal:=json.Unmarshal(body,&faceAuto)
	if errUnmarshal != nil {
		log.Println(errUnmarshal.Error())
	}
	requesthandlechan<-faceAuto
}

//Face++
type FaceAddDentifyFace struct{

}
func (faceAddDentifyFace FaceAddDentifyFace) DentifyFace(request IRequest,requesthead RequestHead,requesthandlechan chan FaceAuto,faceAutoerrchan chan bool){

}


