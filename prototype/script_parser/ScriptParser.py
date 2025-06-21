import json


class ScriptParser:
    """
        Script Parser class that will take in a text file, and parse the data from the script.
    """
    def __init__(self):
        pass

    def parse_txt_file(self,input_txt):
        parsed_txt = ""
        with open(input_txt) as f:
            parsed_txt = f.read()
        return parsed_txt


def parse_txt_script(input_txt):
    return ScriptParser().parse_txt_file(input_txt=input_txt)