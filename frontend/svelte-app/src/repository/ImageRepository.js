
import Api from '/src/api/Api.js'


class ImageRepository {

	constructor() {
		this._api = new Api
	}

	post(params) {
		this._api.post('/images', params)
	}
}


export default ImageRepository