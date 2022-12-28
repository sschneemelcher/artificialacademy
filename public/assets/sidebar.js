const newchatButton = document.getElementById('newchat-button');

const newChat = () => {
  fetch('/new', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response => response.status === 200 && (window.location = "/"))
}

newchatButton.addEventListener('click', newChat)
