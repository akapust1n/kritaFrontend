from django.conf.urls import url
from . import views
from jsonview.decorators import json_view

views.collectLargeDataWrapper()

urlpatterns = [
    url(r'^$', views.start_list, name='star_list'),
    url(r'^installInfo/$', views.show_install_info, name='show_install_info'),
    url(r'^actions/$', views.show_actions_info, name='show_actions_info'),
    url(r'^tools/$', views.show_tools_info, name='show_tools_info'),
    url(r'^images/$', views.show_images_info, name='show_images_info'),
    url(r'^assertsInfo/$', views.show_asserts_info, name='show_asserts_info'),
    url(r'^imagesGraphs/$', views.image_graphs, name='image_graphs'),
    url(r'^installGraphs/$', views.install_graphs, name='install_graphs'),
    url(r'^toolsTableUse/$', views.tools_table_use, name='tools_table_use'),
    url(r'^toolsTableActivate/$', views.tools_table_activate,
        name='tools_table_activate'),
    url(r'^actionsTable/$', views.actions_table, name='actions_table'),

]
