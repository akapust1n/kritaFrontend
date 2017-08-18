from django.db import models


class Tools(models.Model):
    name = models.CharField(max_length=100)
    time = models.FloatField()
    countUse = models.IntegerField()

class ToolsActivate(models.Model):
    name = models.CharField(max_length=100)
    time = models.FloatField()
    countUse = models.IntegerField()

class Actions(models.Model):
    name = models.CharField(max_length=100)
    countUse = models.IntegerField()
