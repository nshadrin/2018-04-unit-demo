import os
import sys
import datetime

def application(environ, start_response):
    output = "<html><head><meta http-equiv=\"refresh\" content=\"1\"><body bgcolor='#dfd'>I'm a python app! Current time: "
    output += datetime.datetime.now().strftime("%Y-%m-%d %I:%M:%S %p")
    start_response('200 OK', [('Content-type', 'text/html')])
    return [output.encode('utf-8')]
