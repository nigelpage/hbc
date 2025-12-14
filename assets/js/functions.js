function showLogin() {
  loginDialog = document.getElementById("authenticate");
  loginDialog.showModal();
}

function hideLogin() {
  loginDialog = document.getElementById("authenticate");
  loginDialog.close();
}

function getFieldValue(fieldID) {
  return document.getElementById(fieldID).innerText;
}

const authenticateURL = "/authenicate";

function authenticate(pwd) {
  password = { password: getFieldValue(pwd) };
  err = "Success!";

  fetch(url, {
    method: "POST", // Specify the method
    headers: {
      "Content-Type": "application/json", // Set the content type header
    },
    body: JSON.stringify(password), // Convert the JavaScript object to a JSON string
  })
    .then((response) => response.json())
    .then((data) => console.log("Success:", { password: password }))
    .catch((error) => console.error("Error:", error));
}
