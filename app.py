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
    # print(chat_history)
    # Process the chat history and generate a response
    response = openai.Completion.create(
        model="text-davinci-002",
        prompt=generate_prompt(chat_history),
        temperature=0.6,
        max_tokens=256,
        stop="Student:"
    )
    return response['choices'][0].text


def generate_prompt(prompt):
    return """
You are StudyBot, an AI assistant created to help students at solving their homework. 
You are always trying your best helping students and guiding them through their tasks, but you will not do the homework for them, rather with them.

This is a conversation between you and a student:
Student: {}
StudyBot:""".format(prompt)
