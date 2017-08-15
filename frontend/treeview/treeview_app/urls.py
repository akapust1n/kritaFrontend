from django.conf.urls import url
from . import views
from jsonview.decorators import json_view

urlpatterns = [
    url(r'^$', views.start_list, name='star_list'),
    url(r'^installInfo/$', views.show_install_info, name='show_install_info'),
    url(r'^actions/$', views.show_actions_info, name='show_actions_info'),
    url(r'^tools/$', views.show_tools_info, name='show_tools_info'),
    url(r'^images/$', views.show_images_info, name='show_images_info'),
    url(r'^imagesGraphs/$', views.image_graphs, name='image_graphs'), 
    url(r'^installGraphs/$', views.install_graphs , name='install_graphs'),  
]

