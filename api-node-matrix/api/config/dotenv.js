
try {
    const dotenv = require('dotenv');
    // Loading environment variables form the .env file
    dotenv.config();
    
    //dotenv.config({ path: path.resolve(__dirname, '../../.env') });
  } catch (error) {
    console.error("Error loading .env file:", error);
  }