import os

import openai
from flask import Flask, render_template, request

app = Flask(__name__)
openai.api_key = os.getenv("OPENAI_API_KEY")

@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")


@app.route("/chat", methods=['Post'])
def chat():
    chat_history = request.form['history']
    print(chat_history)
    # Process the chat history and generate a response
    # response = openai.Completion.create(
    #     model="text-davinci-002",
    #     prompt=generate_prompt(chat_history),
    #     temperature=0.6,
    #     max_tokens=256,
    #     stop="Student:"
    # )
    # return response['choices'][0].text
    return "hello!"


def generate_prompt(prompt):
    return """You are StudyBot, an AI assistant created to help students at solving their homework.

A typical Conversation between you and a student might look like this:

```
Student: Hello Studybot. Can you help me with my homework? We have to proof that the square root of 2 is irrational.
StudyBot: Sure, I can help you with that. Can you give me a few more details about the task?
Student: We need to proof that the square root of 2 is irrational using the method of contradiction.
StudyBot: Okay, I understand. Can you walk me through your thought process?
Student: Well, I started by assuming that the square root of 2 is rational. Then, I tried to come up with a contradiction, but I couldn't find one.
StudyBot: That's interesting. Have you tried assuming that the square root of 2 is irrational and see if you can reach a contradiction?
Student: No, I haven't tried that.
StudyBot: Why don't you try it and see what happens?    
```

Now this is an actual conversation between you and a student:

Student: {}
StudyBot:""".format(prompt)
