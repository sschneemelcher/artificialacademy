const usernameField = document.getElementById('username-input');
const passwordField = document.getElementById('password-input');
const loginButton = document.getElementById('login-button');

const login = (e) => {
  e.preventDefault()

  fetch('/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: `{"name":"${usernameField.value}","pass":"${passwordField.value}"}`
  })
    .then(response => response.status === 200 && (window.location = "/"))
}

loginButton.addEventListener('click', login)
