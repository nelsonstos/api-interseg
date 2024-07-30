const express = require('express');

const MatrixController = require('../controllers/matrix.controller');

const router = express.Router();
//matrix
router.post('/matrix/statistics', MatrixController.processMatrix);

module.exports = router;