from django.db import models
from mptt.models import MPTTModel, TreeForeignKey

class MainModel(MPTTModel):
    name = models.CharField(max_length=50, unique=True)
    parent = TreeForeignKey('self', null=True, blank=True, related_name='children', db_index=True)
    value = models.NullBooleanField()

    class MPTTMeta:
        order_insertion_by = ['name']