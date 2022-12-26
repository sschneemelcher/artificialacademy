// Get references to the input field and send button
const inputField = document.getElementById('chat-input');
const sendButton = document.getElementById('send-button');

const messageClassList = 'fs-5 font-weight-bold mb-3 card';

// Get a reference to the chat window
const chatWindow = document.getElementById('chat-window');

const sendMessage = () => {
  // Get the message from the input field
  const message = inputField.value;

  // Clear the input field
  inputField.value = '';

  // Create a new element to display the message
  const messageElement = document.createElement('div');
  messageElement.classList += ` text-end bg-secondary text-light ${messageClassList}`;
  messageElement.innerText = `${message}`;

  // Add the message to the chat window
  chatWindow.appendChild(messageElement);

  // Scroll the window so the new message is in view
  window.scrollTo(0, document.body.scrollHeight);

  // Send a POST request to the server with the updated chat history
  fetch('/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: `{"text":"${message}"}`
  })
    .then(response => response.text())
    .then(result => {
      // Display the server's response in the chat window
      const responseElement = document.createElement('div');
      responseElement.classList += ` bg-dark text-light ${messageClassList}`;
      responseElement.innerText = `${result}`;
      chatWindow.appendChild(responseElement);
      window.scrollTo(0, document.body.scrollHeight);
    })

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
