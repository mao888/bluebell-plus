
import axios from 'axios';
axios.defaults.baseURL = "/api/v1/";
axios.interceptors.request.use((config) => {
  let loginResult = JSON.parse(localStorage.getItem("loginResult"));	// 解析从localStorage里拿出的loginResult
  if (loginResult) { 
	const token = loginResult.token		// 取出accessToken
	config.headers.Authorization = `Bearer ${token}`;	// 将accessToken放入到请求头里
  }
  return config;
}, (error) => {
	return Promise.reject(error);
});

axios.interceptors.response.use(
	response => {
		if (response.status === 200) {
			return Promise.resolve(response.data);
		} else {
			return Promise.reject(response.data);
		}
	},
	(error) => {
		console.log('error', error);
	}
);

export default axios;