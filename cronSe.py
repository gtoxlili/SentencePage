# -*- coding: utf-8 -*-

import requests
import json


def getContent():
    result = requests.post("http://115.28.168.103:8080/yiyan/getfeeds",
                           headers={"Cookie": "JSESSIONID=181F6886466983E3ADB83ACD543F38EC"})
    settoJson(result.json())


def settoJson(values):
    version = open("version", "r")
    result = []
    values = values["textcardlist"]
    content = version.read()
    for i in values:
        if i["datetime"] == content:
            break
        if i.get("from"):
            result.append({"title": i["content"].replace(
                "\n", "<br>"), "content": i["from"]})
    version.close()
    with open("version", 'w') as f:
        f.write(values[0]["datetime"])
    saveData(result)


def saveData(result):
    with open("data.json", 'r') as f:
        datasds = json.load(f)
        result = result+datasds
        with open("data.json", 'w') as f:
            json.dump(result, f)


if __name__ == "__main__":
    getContent()
