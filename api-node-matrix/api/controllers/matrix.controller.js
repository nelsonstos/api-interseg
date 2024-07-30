const { StatusCodes } = require("http-status-codes");
const { body, validationResult } = require('express-validator');

const matrixService = require("../services/matrix.service");

class MatrixController {

    constructor() {
        this.processMatrix = this.processMatrix.bind(this);
    }

    async processMatrix(req, res ) {

      try {
        const matrix  = req.body.matrix;

        const result = await matrixService.processMatrix(matrix);

        res.status(StatusCodes.CREATED).json(result);
       } catch (err) {
        console.error("Error processing matrix:", err);
        res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({ message: "An error occurred processing the matrix." });
       }
    }
}

module.exports = new MatrixController;