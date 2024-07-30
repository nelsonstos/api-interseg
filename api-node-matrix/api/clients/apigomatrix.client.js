const axios  = require("axios");
require('../config/dotenv');

class ApiGOMatrixClient {

    constructor() {
        this.instance = axios.create({
            baseURL: process.env.API_GO_MATRIX_URL,
            timeout: 10000,
            headers: {
                'cache-control': 'no-cache',
                'Content-Type': 'application/json'
                
              }
        });
    }

    async factorizationMatrix(matrix) {
        const response = await this.instance.post( `/api/v1/matrix`, { matrix }, {   
            });
        return response.data;
    }

    async calculateMatrixStats(matrix) {
        console.log("datadata : ", matrix);
        const response = await this.instance.post( `/api/v1/matrix/stats`, matrix , {   
            });
        return response.data;
    }

}

module.exports = new ApiGOMatrixClient;
