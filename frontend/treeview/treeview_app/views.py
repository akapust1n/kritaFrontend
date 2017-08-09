from django.core.serializers import json
import json
from django.shortcuts import render

import http.client


# Create your views here.

def show_install_info(request):
    # response = json.dumps({"timestamp1": 22, "timestamp2": 14})
    response = json.dumps({"there should be ": "error request"})  # temporary solution
    try:
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/install")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        print(response)
    except Exception as e:
        print("error")
        return render(request, "treeview_app/error_page.html")

    return render(request, "treeview_app/install_info.html", {'json_string': response})


def show_images_info(request):
    # response = json.dumps({"timestamp1": 22, "timestamp2": 14})
    response = json.dumps({"there should be ": "error request"})  # temporary solution
    try:
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/images")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        print(response)
    except Exception as e:
        print("error")
        return render(request, "treeview_app/error_page.html")

    return render(request, "treeview_app/images_info.html", {'json_string': response})


def show_tools_info(request):
    # response = json.dumps({"timestamp1": 22, "timestamp2": 14})
    response = json.dumps({"there should be ": "error request"})  # temporary solution
    try:
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/tools")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        print(response)
    except Exception as e:
        print("error")
        return render(request, "treeview_app/error_page.html")

    return render(request, "treeview_app/tools_info.html", {'json_string': response})

def show_actions_info(request):
    # response = json.dumps({"timestamp1": 22, "timestamp2": 14})
    response = json.dumps({"there should be ": "error request"})  # temporary solution
    try:
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/actions")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        print(response)
    except Exception as e:
        print("error")
        return render(request, "treeview_app/error_page.html")

    return render(request, "treeview_app/actions_info.html", {'json_string': response})


def start_list(request):
    return render(request, 'treeview_app/start_page.html', {})
