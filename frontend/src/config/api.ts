import { Axios } from "axios";

const Api = new Axios({
    baseURL: "http://localhost:5000/api/v1",
    headers: {
        "Content-Type": "application/json"
    }
});

export default Api;