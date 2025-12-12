function showLogin() {
  loginDialog = document.getElementById("authenticate");
  loginDialog.showModal();
}

function hideLogin() {
  loginDialog = document.getElementById("authenticate");
  loginDialog.close();
}

const bcrypt = require("bcrypt");
const saltRounds = 10; // Recommended value, adjust based on security needs and performance

async function hashPassword(password) {
  const hashedPassword = await bcrypt.hash(password, saltRounds);
  return hashedPassword;
}

// Example usage
// hashPassword('mySecretPassword123')
//     .then(hash => console.log('Hashed password:', hash))
//     .catch(err => console.error(err));
