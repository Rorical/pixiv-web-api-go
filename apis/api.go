package apis

import "github.com/dghubble/sling"

type PixivWebApi struct {
	sling *sling.Sling
}

const (
	apiBase = "https://www.pixiv.net/"
)

func GetPixivWebApi() *PixivWebApi {
	return &PixivWebApi{
		sling: sling.New().Base(apiBase)
	}
}

func (api *PixivWebApi) request(path string, data interface{}, params ...interface{}) error {
	path = fmt.Sprintf(path, ...params)
	_, err = api.sling.New().Get(path).ReceiveSuccess(data)
	return err
}

func (api *PixivWebApi) IllustDetail(id int) {
	response := &IllustDetailResponse{}
	err := api.request("ajax/illust/%s", response, id)
}