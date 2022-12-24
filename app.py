import os
import time


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
    # response = openai.Completion.create(
    #     model="text-davinci-003",
    #     prompt=generate_prompt(chat_history),
    #     temperature=0.6,
    #     max_tokens=256,
    #     stop="Student:"
    # )
    # return response['choices'][0].text
    time.sleep(2)
    return 'Sure thing! Neural networks are composed of artificial neurons, which are connected to each other in layers. Each neuron is responsible for taking in a certain input, processing it, and then outputting a result. The neurons in the first layer take in the inputs, and the neurons in the last layer output the results. In between, the neurons in the hidden layers process the data and pass it on to the next layer. The connections between the neurons are weighted, meaning that some values are given more importance than others. This is how the neural network is able to form complex patterns and make predictions. Does this answer your question?'


def generate_prompt(prompt):
    return """
You are StudyBot, an AI assistant created to help students at solving their homework. 
You are always trying your best helping students and guiding them through their tasks,
as well as providing encouragement. The following is a conversation between you and a student.

Student: {}
StudyBot:""".format(prompt)
