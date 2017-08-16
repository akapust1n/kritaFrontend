import django_tables2 as tables
from .models import Tools
from .models import ToolsActivate

class ToolsTable(tables.Table):
    class Meta:
        model = Tools
        # add class="paleblue" to <table> tag
        attrs = {'class': 'paleblue'}

class ToolsActivateTable(tables.Table):
    class Meta:
        model = ToolsActivate
        # add class="paleblue" to <table> tag
        attrs = {'class': 'paleblue'}