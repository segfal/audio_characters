import requests
import os
from script_parser import parse_txt_script

txt_scripts = [f"user_input_script/{i}" for i in os.listdir("user_input_script")]


for txt in txt_scripts:
    print(
        parse_txt_script(input_txt=txt)
    )



