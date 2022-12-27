const inputField = document.getElementById('chat-input');
const sendButton = document.getElementById('send-button');
const clearButton = document.getElementById('clear-button');
const chatWindow = document.getElementById('chat-window');

const messageClassList = 'fs-5 font-weight-bold mb-3 card';

const appendMessage = (message, user) => {

  // Create a new element to display the message
  const messageElement = document.createElement('div');
  const messageElementUsername = document.createElement('span');
  const messageElementMessage = document.createElement('span');
  if (user != "me") {
    messageElement.classList += ` text-start bg-dark text-light ${messageClassList}`;
  } else {
    messageElement.classList += ` text-end bg-secondary text-light ${messageClassList}`;
  }
  messageElementMessage.innerText = `${message}`;
  messageElementUsername.innerHTML = `${user}:</br>`;
  messageElementUsername.classList += ' fs-6 text-decoration-underline';
  messageElement.appendChild(messageElementUsername);
  messageElement.appendChild(messageElementMessage);

  // Add the message to the chat window
  chatWindow.appendChild(messageElement);

  // Scroll the window so the new message is in view
  window.scrollTo(0, document.body.scrollHeight);

}

const sendMessage = () => {
  // Get the message from the input field
  const message = inputField.value;

  // Clear the input field
  inputField.value = '';
  appendMessage(message, 'me');


  // Send a POST request to the server with the updated chat history
  fetch('/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: `{"content":"${message}"}`
  })
    .then(response => response.json())
    .then(result => {
      appendMessage(result.message, result.user);
      sendButton.setAttribute('disabled', '');
    })
}

const clearHistory = () => {
  fetch('/chat', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
  })
    .then(response => (response.status === 200) && (chatWindow.innerHTML = ''))
}

let shiftActive = false;
// Add an event listener to the input field that listens for the keydown event
inputField.addEventListener('keydown', event => {
  (event.key === 'Enter' && !shiftActive) && inputField.value.length && (event.preventDefault() || sendMessage());
  (event.key === 'Enter' && !shiftActive) && !inputField.value.length && event.preventDefault();
  (event.key === 'Shift') && (shiftActive = true);
});

inputField.addEventListener('keyup', event => {
  (event.key === 'Shift') && (shiftActive = false);
  if (inputField.value.length)
    sendButton.removeAttribute('disabled')
  else sendButton.setAttribute('disabled', '');
});

// Add an event listener to the send button
sendButton.addEventListener('click', sendMessage)
clearButton.addEventListener('click', clearHistory)
onload = () => window.scrollTo(0, document.body.scrollHeight);
