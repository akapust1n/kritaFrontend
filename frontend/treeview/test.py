from jsonview.decorators import json_view

@json_view
def my_view(request):
    return {
        'foo': 'bar',
    }
