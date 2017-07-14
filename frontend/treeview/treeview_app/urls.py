from django.conf.urls import url
from . import views
from jsonview.decorators import json_view


urlpatterns = [
    url(r'^$', views.post_list, name='post_list'),
    url(r'^installInfo/$', views.show_install_info,name='show_install_info'),

]

