import django_tables2 as tables
from .models import Tools
from .models import ToolsActivate
from .models import Actions

class ToolsTable(tables.Table):
    class Meta:
        model = Tools
        # add class="paleblue" to <table> tag
        order_by = '-countUse'
        attrs = {'class': 'paleblue'}

class ToolsActivateTable(tables.Table):
    class Meta:
        model = ToolsActivate
        order_by = '-countUse'
        # add class="paleblue" to <table> tag
        attrs = {'class': 'paleblue'}

class ActionsTable(tables.Table):
    class Meta:
        model = Actions
        order_by = '-countUse'
        attrs = {'class': 'paleblue'}