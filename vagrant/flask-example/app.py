from flask import Flask, render_template
import random

app = Flask(__name__)

images = [
    "https://icatcare.org/app/uploads/2018/07/Thinking-of-getting-a-cat.png",
    "https://static01.nyt.com/images/2020/04/22/science/22VIRUS-PETCATS1/merlin_150476541_233fface-f503-41af-9eae-d90a95eb6051-superJumbo.jpg?quality=90&auto=webp",
    "https://cdn.mos.cms.futurecdn.net/VSy6kJDNq2pSXsCzb6cvYF-1024-80.jpg.webp",
]


@app.route("/")
def index():
    url = random.choice(images)
    return render_template("index.html", url=url)


if __name__ == "__main__":
    app.run(host="0.0.0.0")
