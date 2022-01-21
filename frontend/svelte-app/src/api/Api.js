
import axios from axios

class Api {

	post(endpoint, params) {
		axios({
			method: 'post',
			baseURL: 'http://localhost:8080',
			url: endpoint,
			params: params,
		})
		.then((res) => {
			console.log(res)
		})
		.catch((e) => {
			console.log(e)
		})
	}
}


export default Api