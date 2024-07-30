const express = require('express');
const dotenv = require('dotenv');
const bodyParser = require('body-parser');
const router = require('./api/routes/routes');
const cors = require('cors');

dotenv.config();

const app = express();

// Configuración de CORS
const corsOptions = {
  origin: '*', // Permitir todos los orígenes
  methods: 'GET,HEAD,PUT,PATCH,POST,DELETE', // Métodos permitidos
  allowedHeaders: 'Content-Type,Authorization', // Encabezados permitidos
  credentials: true // Permitir el envío de cookies y encabezados de autorización
};

app.use(cors(corsOptions));

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use(express.json());

app.get('/', (req, res) => {
  res.send('Welcome to the API!');
});

//Routes

app.use("/api/v1", router);

app.listen(process.env.PORT || 3000, () => {
    console.log('Server running on port 3000');
});
