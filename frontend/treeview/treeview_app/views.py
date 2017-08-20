from django.core.serializers import json
import json
from django.shortcuts import render
import http.client
from graphos.renderers import gchart
from graphos.sources.simple import SimpleDataSource
from graphos.renderers.gchart import BarChart
from .models import Tools, ToolsActivate, Actions
from django_tables2 import RequestConfig
from .tables import ToolsTable, ToolsActivateTable, ActionsTable
import threading
import sched, time



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

def show_asserts_info(request):
    response = json.dumps({"there should be ": "error request"})
    try:
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/asserts")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        print(response)
    except Exception as e:
        print("error")
        return render(request, "treeview_app/error_page.html")

    return render(request, "treeview_app/asserts_info.html", {'asserts': response})
    
def start_list(request):
    return render(request, 'treeview_app/start_page.html', {})


def getImageInfo(type, distribution, subsection, title):
    conn = http.client.HTTPConnection("localhost:8080")
    conn.request("GET", "/get/images?type=" + type)
    r1 = conn.getresponse()
    response = r1.read()  # what will happen if response code will be not 200
    conn.close()
    print("test1")
    response = response.decode("utf-8")
    print(type)
    decoded = json.loads(response)

    resultList = [[title, subsection], ]
    for x in distribution:
        # print(decoded[x][subsection])
        resultList.append([x, decoded[x][subsection]])
    print(type + "TYPE")
    return resultList


def image_graphs(request):
    WIDTH_COUNT = 0
    HEIGHT_COUNT = 1
    NUMLAYERS_COUNT = 2
    FILESIZE_COUNT = 3
    COLORPROFILE_COUNT = 4
    response = [[], [], [], [], []]
    try:
        response[WIDTH_COUNT] = getImageInfo(
            "width", ["L500", "L1000", "L2000", "L4000", "L8000", "M8000", "Unknown"], "Count", "Width")

        response[HEIGHT_COUNT] = getImageInfo("height", [
            "L500", "L1000", "L2000", "L4000", "L8000", "M8000", "Unknown"], "Count", "Height")
        response[NUMLAYERS_COUNT] = getImageInfo("numlayers", [
            "L1", "L2",  "L4", "L8", "L16", "L32", "L64", "M64", "Unknown"], "Count", "Num layers")
        response[FILESIZE_COUNT] = getImageInfo("filesize", [
            "Mb1", "Mb5", "Mb10", "Mb25", "Mb50", "Mb100", "Mb200", "Mb400", "Mb800", "More800", "Unknown"], "Count", "File size")
        response[COLORPROFILE_COUNT] = getImageInfo("colorprofile", [
            "RGBA", "CMYK", "Grayscale", "XYZ", "YCbCr", "Lab", "Unknown"], "Count", "Color Profile")
    except Exception as e:
        print(e)
        return render(request, "treeview_app/error_page.html")

# DataSource object
    data_source = [0, 0, 0, 0, 0]
    data_source[WIDTH_COUNT] = SimpleDataSource(data=response[WIDTH_COUNT])
    data_source[HEIGHT_COUNT] = SimpleDataSource(data=response[HEIGHT_COUNT])
    data_source[NUMLAYERS_COUNT] = SimpleDataSource(
        data=response[NUMLAYERS_COUNT])
    data_source[FILESIZE_COUNT] = SimpleDataSource(
        data=response[FILESIZE_COUNT])
    data_source[COLORPROFILE_COUNT] = SimpleDataSource(
        data=response[COLORPROFILE_COUNT])


# Chart objects
    chart = [0, 0, 0, 0, 0]
    chart[WIDTH_COUNT] = BarChart(data_source[WIDTH_COUNT])
    chart[HEIGHT_COUNT] = BarChart(data_source[HEIGHT_COUNT])
    chart[NUMLAYERS_COUNT] = BarChart(data_source[NUMLAYERS_COUNT])
    chart[FILESIZE_COUNT] = BarChart(data_source[FILESIZE_COUNT])
    chart[COLORPROFILE_COUNT] = BarChart(data_source[COLORPROFILE_COUNT])

    context = {'chart': chart[WIDTH_COUNT], 'chart2': chart[HEIGHT_COUNT],
               "chart4": chart[NUMLAYERS_COUNT],   "chart6": chart[FILESIZE_COUNT],
               "chart8": chart[COLORPROFILE_COUNT]
               }
    return render(request, 'treeview_app/image_graphs.html', context)


def getInstallInfo(type, distribution, subsection, title):
    print("get installInfo")
    conn = http.client.HTTPConnection("localhost:8080")
    conn.request("GET", "/get/install?type=" + type)
    r1 = conn.getresponse()
    response = r1.read()  # what will happen if response code will be not 200
    conn.close()
    response = response.decode("utf-8")
    decoded = json.loads(response)
    print(decoded)

    resultList = [[title, subsection], ]
    for x in distribution:
        # print(decoded[x][subsection])
        resultList.append([x, decoded[x][subsection]])
    print(type + "TYPE")
    return resultList


def install_graphs(request):
    OS_COUNT = 0
    WINDOWS_COUNT = 1
    LINUX_COUNT = 2
    MAC_COUNT = 3
    CPU_ARCHITECTURE_COUNT = 4
    CPU_CORES_COUNT = 5
    COMPILER_COUNT = 6
    LOCALE_COUNT = 7
    ISINTEL_COUNT = 8

    response = []
    for i in range(8):
        response.append([])
    print("DD")
    try:
        response[OS_COUNT] = getInstallInfo(
            "os", ["Windows", "Linux", "Mac", "Unknown"], "Count", "Os")
        response[WINDOWS_COUNT] = getInstallInfo("windows", [
            "V7", "V8", "V81", "V10", "Other"], "Count", "Versions of Windows")
        response[LINUX_COUNT] = getInstallInfo("linux", [
            "Ubuntu1404", "Ubuntu1410",  "Ubuntu1504", "Ubuntu1510", "Ubuntu1604", "Ubuntu1610", "Ubuntu1704", "Other"], "Count", "Version of Linux")
        response[MAC_COUNT] = getInstallInfo("mac", [
            "V1012", "Other"], "Count", "Versions of Macs")
        response[CPU_ARCHITECTURE_COUNT] = getInstallInfo("architecture", [
            "X86_64", "X86", "Other", "Unknown"], "Count", "Kinds of architecture")
        response[CPU_CORES_COUNT] = getInstallInfo("cores", [
            "C1", "C2", "C3", "C4", "C6", "C8", "Other", "Unknown"], "Count", "Number of cores")
        response[COMPILER_COUNT] = getInstallInfo("compiler", [
            "GCC", "Clang", "MSVC", "Other", "Unknown"], "Count", "Compilers")
        response[LOCALE_COUNT] = getInstallInfo("locale", [
            "English", "Russian", "Other", "Unknown"], "Count", "Locales")
    except Exception as e:
        print(e)
        return render(request, "treeview_app/error_page.html")

# DataSource object
    data_source = []
    for i in range(8):
        data_source.append([])
    data_source[OS_COUNT] = SimpleDataSource(data=response[OS_COUNT])
    data_source[WINDOWS_COUNT] = SimpleDataSource(data=response[WINDOWS_COUNT])
    data_source[LINUX_COUNT] = SimpleDataSource(
        data=response[LINUX_COUNT])
    data_source[MAC_COUNT] = SimpleDataSource(
        data=response[MAC_COUNT])
    data_source[CPU_ARCHITECTURE_COUNT] = SimpleDataSource(
        data=response[CPU_ARCHITECTURE_COUNT])
    data_source[CPU_CORES_COUNT] = SimpleDataSource(
        data=response[CPU_CORES_COUNT])
    data_source[COMPILER_COUNT] = SimpleDataSource(
        data=response[COMPILER_COUNT])
    data_source[LOCALE_COUNT] = SimpleDataSource(
        data=response[LOCALE_COUNT])

# Chart objects
    chart = []
    for i in range(8):
        chart.append([])
    chart[OS_COUNT] = BarChart(data_source[OS_COUNT])
    chart[WINDOWS_COUNT] = BarChart(data_source[WINDOWS_COUNT])
    chart[LINUX_COUNT] = BarChart(data_source[LINUX_COUNT])
    chart[MAC_COUNT] = BarChart(data_source[MAC_COUNT])
    chart[CPU_ARCHITECTURE_COUNT] = BarChart(
        data_source[CPU_ARCHITECTURE_COUNT])
    chart[CPU_CORES_COUNT] = BarChart(
        data_source[CPU_CORES_COUNT])
    chart[COMPILER_COUNT] = BarChart(
        data_source[COMPILER_COUNT])
    chart[LOCALE_COUNT] = BarChart(
        data_source[LOCALE_COUNT])

    context = {'chart': chart[OS_COUNT], 'chart1': chart[WINDOWS_COUNT],
               "chart2": chart[LINUX_COUNT],   "chart3": chart[MAC_COUNT],
               "chart4": chart[CPU_ARCHITECTURE_COUNT],
               "chart5": chart[CPU_CORES_COUNT],
               "chart6": chart[COMPILER_COUNT],
               "chart7": chart[LOCALE_COUNT],
               }
    return render(request, 'treeview_app/install_graphs.html', context)


def tools_table_use(request):

    table = ToolsTable(Tools.objects.all())
    RequestConfig(request).configure(table)

    context = {'tools': table}
    return render(request, 'treeview_app/tools_table.html', context)

def tools_table_activate(request):
   
    table = ToolsActivateTable(ToolsActivate.objects.all())
    RequestConfig(request).configure(table)
    context = {'tools': table}
    return render(request, 'treeview_app/tools_table.html', context)

def actions_table(request):
    table = ActionsTable(Actions.objects.all())
    RequestConfig(request).configure(table)
    context = {'actions': table}
    return render(request, 'treeview_app/actions_table.html', context)
    
def collectLargeData():
    s = sched.scheduler(time.time, time.sleep)

    def collect(sc):
        print("collect data...")
        ToolsActivate.objects.all().delete()
        Tools.objects.all().delete()
        Actions.objects.all().delete()

        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/tools")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        response = response.decode("utf-8")
        decoded = json.loads(response)
        for x in decoded["ToolsActivate"]:
                # print(decoded[x][subsection])
            dd = ToolsActivate(name=x["Name"], countUse=x["CountUse"], time = x["Time"])
            dd.save()
        for x in decoded["ToolsUse"]:
            # print(decoded[x][subsection])
            dd = Tools(name=x["Name"], countUse=x["CountUse"], time = x["Time"])
            dd.save()
        #actions
        conn = http.client.HTTPConnection("localhost:8080")
        conn.request("GET", "/get/actions")
        r1 = conn.getresponse()
        response = r1.read()  # what will happen if response code will be not 200
        conn.close()
        response = response.decode("utf-8")
        decoded = json.loads(response)
        for x in decoded["Actions"]:
                # print(decoded[x][subsection])
            dd = Actions(name=x["Name"], countUse=x["CountUse"])
            dd.save()
        s.enter(3600, 1, collect, (sc,)) #updates every hour

    s.enter(0.1, 1, collect, (s,))
    s.run()

def collectLargeDataWrapper():
    t = threading.Timer(10.0, collectLargeData)
    t.start() 
