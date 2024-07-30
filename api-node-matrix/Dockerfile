# Usa la imagen base oficial de Node.js 20
FROM node:20

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de configuración de dependencias
COPY package*.json ./

# Crea un archivo .npmrc para autenticación (si es necesario)
# Nota: Asegúrate de agregar un archivo .npmrc adecuado para la autenticación si estás utilizando un registro npm privado.
# RUN echo "//registry.npmjs.org/:_authToken=${NPM_TOKEN}" > .npmrc

# Instala las dependencias
RUN npm install

# Copia el resto del código fuente al contenedor
COPY . .

# Expone el puerto que la aplicación escuchará
EXPOSE 3000

# Define el argumento de construcción para el entorno
ARG NODE_ENV=production

# Establece la variable de entorno para el entorno
ENV NODE_ENV=${NODE_ENV}

# Define el comando por defecto para ejecutar la aplicación basado en el entorno
CMD ["npm", "run", "${NODE_ENV}"]
