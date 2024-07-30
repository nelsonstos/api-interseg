const apigomatrix = require('../clients/apigomatrix.client');

class MatrixService {
    async processMatrix(data) {

        const matrix = await apigomatrix.factorizationMatrix(data);
        console.log("result: ", matrix)

        const {q_matrix, r_matrix } = matrix.data

        // ConvertiNG JSON string to matrixes
        const qMatrix = JSON.parse(q_matrix);
        const rMatrix = JSON.parse(r_matrix);

        const combinedStats = this.calculateCombinedStats([qMatrix, rMatrix]);

        const isQMatrixDiagonal = this.isDiagonal(qMatrix);
        const isRMatrixDiagonal = this.isDiagonal(rMatrix);
        const stats = {
            factorization_id: matrix.data.id,
            ...combinedStats,
            isQMatrixDiagonal,
            isRMatrixDiagonal,
        }
        const statsData = await apigomatrix.calculateMatrixStats(stats);

        if(statsData.status !== "success") {
            throw new Error("Failed to calculate matrix stats");
        }

        return statsData;
    }
    // Calulating the combin matrix statisc
    calculateCombinedStats(matrices) {
        let combinedArray = matrices.flat(2); 
        let max = Math.max(...combinedArray);
        let min = Math.min(...combinedArray);
        let sum = combinedArray.reduce((acc, val) => acc + val, 0);
        let avg = sum / combinedArray.length;

        return { max, min, avg, sum };
    }

    // Validating if the matrix is Diagonal
    isDiagonal(matrix) {
        for (let i = 0; i < matrix.length; i++) {
            for (let j = 0; j < matrix[i].length; j++) {
                if (i !== j && matrix[i][j] !== 0) {
                    return false;
                }
            }
        }
        return true;
    }
}

module.exports = new MatrixService;