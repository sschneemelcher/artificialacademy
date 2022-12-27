const logoutButton = document.getElementById('logout-button');

const logout = () => {
  fetch('/logout', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(response => response.status === 200 && (window.location = "/login"))
}

logoutButton.addEventListener('click', logout)
