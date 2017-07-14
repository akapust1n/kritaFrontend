from django.shortcuts import render
from jsonview.decorators import json_view
import json
from django.shortcuts import render_to_response
from treeview_app.models import MainModel

# Create your views here.

def show_install_info(request):
    json_string = json.dumps({"timestamp1": 22, "timestamp2": 14})
    return render(request,"treeview_app/install_info.html", {'json_string': json_string})
# def show_genres(request):
#     return render("treeview_app/genres.html",
#                          {'nodes':Genre.objects.all()})                          
def post_list(request):
    return render(request, 'treeview_app/post_list.html', {})
