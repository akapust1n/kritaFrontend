from django.core.serializers import json
import json
from django.shortcuts import render
import http.client
from graphos.renderers import gchart
from graphos.sources.simple import SimpleDataSource
from graphos.renderers.gchart import BarChart


# Create your views here.

def show_install_info(request):
    # response = json.dumps({"timestamp1": 22, "timestamp2": 14})
    # temporary solution
    response = json.dumps({"there should be ": "error request"})
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
    # temporary solution
    response = json.dumps({"there should be ": "error request"})
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
    # temporary solution
    response = json.dumps({"there should be ": "error request"})
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
    # temporary solution
    response = json.dumps({"there should be ": "error request"})
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


def getImageInfo(type, distribution, subsection):
    conn = http.client.HTTPConnection("localhost:8080")
    conn.request("GET", "/get/images?type=" + type)
    r1 = conn.getresponse()
    response = r1.read()  # what will happen if response code will be not 200
    conn.close()
    response = response.decode("utf-8")
    decoded = json.loads(response)

    resultList = [['Type', subsection], ]
    for x in distribution:
        print(decoded[x][subsection])
        resultList.append([x, decoded[x][subsection]])

    return resultList


def image_graphs(request):

    try:
        response = getImageInfo(
            "width", ["L500", "L1000", "L2000", "L4000", "L8000", "M8000", "Unknown"], "Count")
        response1 = []
        response2 = getImageInfo("height", [
                                 "L500", "L1000", "L2000", "L4000", "L8000", "M8000", "Unknown"], "Count")
        response3 = []
        response4 = getImageInfo("numlayers", [
                                 "L1", "L2",  "L4", "L8", "L16", "L32", "L64", "M64", "Unknown"], "Count")
        response5 = []
        response6 = getImageInfo("filesize", [
            "Mb1", "Mb5", "Mb10", "Mb25", "Mb50", "Mb100", "Mb200", "Mb400", "Mb800", "More800", "Unknown"], "Count")
    except Exception as e:
        print(e)
        return render(request, "treeview_app/error_page.html")

# DataSource object
    data_source = SimpleDataSource(data=response)
   # data_source1 = SimpleDataSource(data=response1)
    data_source2 = SimpleDataSource(data=response2)
   # data_source3 = SimpleDataSource(data=response3)
    data_source4 = SimpleDataSource(data=response4)
    #data_source5 = SimpleDataSource(data=response5)
    data_source6 = SimpleDataSource(data=response6)


# Chart object
    chart = BarChart(data_source)
   # chart1 = BarChart(data_source1)
    chart2 = BarChart(data_source2)
   # chart3 = BarChart(data_source3)
    chart4 = BarChart(data_source4)
   # chart5 = BarChart(data_source5)
    chart6 = BarChart(data_source6)

    context = {'chart': chart, 'chart2': chart2,
               "chart4": chart4,   "chart6": chart6
               }
    return render(request, 'treeview_app/image_graphs.html', context)
