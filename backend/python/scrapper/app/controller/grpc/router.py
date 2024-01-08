from ...service.service import Service

from .file import FileHandler


def init_routes(service: Service):
    return [FileHandler(service.file)]
